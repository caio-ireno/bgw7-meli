# Minha dúvidas e estudos organizados

# Sumário

- [Arquitetura](#arquitetura)
- [Patch (Atualização Parcial)](#patch-atualizacao-parcial)
- [Obtendo Parâmetros via ID (Path vs Query)](#obtendo-parametros-via-id-path-vs-query)
- [ToDomain() – Mapeando DTO para Domínio](#todomain-mapeando-dto-para-dominio)
- [Uso de Memória: Valor vs Ponteiro em Structs](#uso-de-memoria-valor-vs-ponteiro-em-structs)

---

# Arquitetura

## Por que a estrutura de pastas é importante?

Organizar bem os arquivos e diretórios traz muitos benefícios, como:

- **Separação de Camadas:** Facilita entender responsabilidades.
- **Melhor organização:** Arquivos agrupados por função.
- **Reutilização:** Código modular é mais fácil de reaproveitar.
- **Manutenção facilitada:** Fácil localizar e alterar funções.
- **Modularidade:** Menos impacto colateral em mudanças.

---

## Exemplo de Estrutura em Camadas

```
project-root/
├── cmd/                # Pontos de entrada das aplicações
│   └── app/          
│       └── main.go 
├── internal/           # Código privado do projeto
│   ├── handlers/       # Controllers: lógica HTTP, endpoints
│   │   └── user_handler.go
│   ├── services/       # Lógica de negócio
│   │   └── user_service.go
│   ├── repositories/   # Acesso a dados (DB, APIs)
│   │   └── user_repo.go
│   └── models/         # Entidades do domínio
│       └── user.go
├── pkg/                # Utilitários e helpers públicos
├── configs/            # Arquivos de configuração
├── go.mod              # Definição de dependências Go
└── go.sum              # Checksum das dependências
```

---

## CMD

Gerencia como a aplicação é iniciada. Em projetos grandes, pode haver vários binários diferentes (ex: `cmd/app`, `cmd/worker`).

---

## INTERNAL

Onde vive o código privado (não deve ser importado por outros projetos).  
Organização sugerida:

### Handlers (Controllers)

- **Responsabilidade:** Recebem requisições HTTP, validam dados, acionam serviços e retornam respostas.
- **Exemplo:**  
  `user_handler.go` define endpoints como POST/GET `/users`.

### Services

- **Responsabilidade:** Implementam regras de negócio, coordenam repositórios, processos etc.
- **Exemplo:**  
  `user_service.go` com funções como “CriarUsuário”, “ValidarSenha”.

### Repositories

- **Responsabilidade:** Lógica de acesso a dados (DB, arquivos, APIs externas).
- **Exemplo:**  
  `user_repo.go` com funções como “BuscarUsuárioPorID”, “SalvarUsuárioNoDB”.

---

# Patch (Atualização Parcial)

## Fluxo do PATCH

1. **Buscar produto existente:**
   ```go
   product, ok := c.storage[productId]
   ```

2. **Preencher reqBody com os dados atuais:**
   ```go
   reqBody := dto.UpdateBodyProduct{
     Name:     product.Name,
     Type:     product.Type,
     Quantity: product.Quantity,
     Price:    product.Price,
   }
   ```
   Assim, campos não enviados serão mantidos.

3. **Sobrescrever apenas os campos enviados:**
   ```go
   err = json.NewDecoder(r.Body).Decode(&reqBody)
   ```

4. **Converter para domínio e atualizar storage:**
   ```go
   pr := reqBody.ToDomain()
   pr.Id = productId
   c.storage[productId] = &pr
   ```

5. **Retornar produto atualizado**
   ```go
   Data: &pr
   ```

---

## Resumo do PATCH

- Você **sobrescreve** o produto antigo, atualizando só os campos enviados.
- O "truque" de preencher `reqBody` com o estado anterior garante que campos não enviados não sejam resetados.
- Esse é o funcionamento típico de PATCH.

---

# Obtendo Parâmetros via ID (Path vs Query)

Duas formas de obter um parâmetro `id` em uma requisição GET:

```go
pathId := chi.URLParam(r, "id")
queryParamsId := r.URL.Query().Get("id")
```

- **Path Param:**  
  Exemplo: `http://localhost:8080/products/1`
- **Query Param:**  
  Exemplo: `http://localhost:8080/products?id=1`

---

# ToDomain() – Mapeando DTO para Domínio

O método `ToDomain()` transforma o DTO do request numa entidade do domínio, aplicando regras de negócio.

**Exemplo:**  
Se chega data de nascimento no DTO, calcula a idade na entidade de domínio.

```go
type Studant struct {
    Name     string
    DataNasc time.Time
    Idade    int
}

// Mapper - DTO → Domínio
func (r RequestBodyStudant) ToDomain() (Studant, error) {
    dataNasc, err := time.Parse("2006-01-02", r.DataNasc)
    if err != nil {
        return Studant{}, err
    }
    idade := calculateAge(dataNasc)
    return Studant{
        Name:     r.Name,
        DataNasc: dataNasc,
        Idade:    idade,
    }, nil
}

func calculateAge(birth time.Time) int {
    now := time.Now()
    years := now.Year() - birth.Year()
    if now.YearDay() < birth.YearDay() {
        years--
    }
    return years
}

type RequestBodyStudant struct {
    Name     string `json:"name"`
    DataNasc string `json:"data_nasc"` // formato: "2006-01-02"
}
```

**Use ToDomain sempre que:**

- Receber dados do cliente (POST/PUT/PATCH)
- Precisar transformar/validar antes de salvar/trabalhar

---

# Uso de memoria no &output.ResponseBodyProduct{...}

A diferença entre passar o endereço (&output.ResponseBodyProduct{...}) ou o valor (output.ResponseBodyProduct{...}) é principalmente sobre como o Go lida com structs e ponteiros:


## Passando o endereço (`&struct{...}`)

- Cria um **ponteiro** para a struct (não faz cópia).
- Mais eficiente para structs grandes.
- Permite modificar o objeto original.
- Convencional em Go, especialmente para retornos de handlers.

## Passando o valor (`struct{...}`)

- Faz uma **cópia** do struct.
- Menos eficiente se a struct crescer.
- Imutável do lado do receptor.

**Ambas funcionam com `json.NewEncoder(w).Encode()`.  
Ponteiro é preferido por eficiência e consistência.**

