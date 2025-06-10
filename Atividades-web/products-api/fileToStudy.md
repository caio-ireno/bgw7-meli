# Arquitetura

A ideia é compreender como podemos organizar nossos arquivos e diretórios dentro de um projeto.

Por que a estrutura de pastas é importante?

Ter uma estrutura de pastas bem organizada é muito importante por alguns motivos principais:

- Separação de Camadas;
- Melhor organização'
- Reutilização;
- Manutenção facilitada;
- Modularidade;

### Estrutura de pasta em camadas

```plaintext
project 
├── cmd                       # Arquivos relacionados a comandos
 │ └── app                   # Ponto de entrada do aplicativo
 │ └── main.go           # Lógica principal do aplicativo
 ├── internal                  # Base de código interna
 │ ├── handlers              # Manipuladores de solicitação HTTP (controladores)
 │ │ └── user_handler.go   # Manipulador específico do usuário
 │ ├── services              # Lógica de negócios (camada de serviço)
 │ │ └── user_service.go   # Serviço específico do usuário
 │ ├── repositories          # Acesso a dados (camada de repositório)
 │ │ └── user_repo.go      # Repositório específico do usuário
 │ └── modelos                # Modelos de dados (entidades)
 │ └── user.go           # Modelo de usuário
 ├── pkg                       # Utilitários ou auxiliares compartilhados
 ├── configs                   # Arquivos de configuração
 ├── go.mod                    # Definição do módulo Go
 └── go.sum                    # Arquivo de checksum do módulo Go
```

```plaintext
project-root/
|-- cmd/
|-- pkg/
|-- internal/
|-- go.mod
|-- main.go (às vezes dentro de cmd/)
```

### CMD

Principais aplicações para este projeto.

Gerenciar como a aplicação é iniciada. Em projetos maiores pode haver vários aplicativos, por isso pode haver diferentes subpastas em cmd/ (ex: cmd/app, cmd/worker, etc).


### INTERNAL

Aplicativo privado e código de biblioteca. Este é o código que você não quer que outros importem em seus aplicativos ou bibliotecas. Observe que esse padrão de layout é imposto pelo próprio compilador Go.

#### Handler

Responsabilidade: Recebem as requisições HTTP, validam dados, acionam serviços e retornam respostas HTTP.

Exemplo: user_handler.go define endpoints como POST /users, GET /users/{id}, etc.

Conhecidos também como controllers.

#### Services

Responsabilidade: Onde reside a lógica de negócio principal da aplicação. Orquestra chamadas a repositories e executa regras de negócio.

Exemplo: user_service.go pode implementar funções como “CriarUsuário”, “ValidarSenha”, coordena diferentes processos.


#### Repository

Responsabilidade: Comunicação com a fonte de dados/armazenamento (banco de dados, API externa, arquivos).

Exemplo: user_repo.go implementa funções como “BuscarUsuárioPorID”, “SalvarUsuárioNoDB”, abstraindo detalhes do acesso ao dado.


# Parâmetro por ID

Temos duas abordagens para pegar o ID em uma requisicao do tipo GET.

Primeiro usando a biblioteca do chi por path params.

Segunda usando busca por query Params.

``` go
    pathId := chi.URLParam(r, "id")
    queryParansId := r.URL.Query().Get("id")
```

Para a primeira opcao a url seria: http://localhosto:8080/products/1

para a segunda opcao a url seria: http://localhosto:8080/products?id=1

# ToDomain()

O método ToDomain() serve justamente para transformar o que chega no request (DTO) em uma entidade do domínio, aplicando regras de negócio e cálculos necessários.

No seu exemplo, não faz sentido pedir a idade se você já tem a data de nascimento, pois a idade pode ser calculada. O ToDomain() é o lugar ideal para fazer essa transformação.

```go
type Studant struct {
    Name     string
    DataNasc time.Time
    Idade    int
}

// Mapper - DTO
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

Devemos usar o método ToDomain() principalmente em métodos que recebem dados do cliente e precisam transformar esses dados (DTO/request) em entidades do seu domínio, como:

- POST (criação de recurso)
- PUT (atualização total do recurso)
- PATCH (atualização parcial do recurso)
