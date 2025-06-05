diferenca entre: %d, %s, %f e outros

Esses são **formatadores** usados em funções como `fmt.Printf` em Go (e em outras linguagens também):

- **%d**: Formata um número inteiro (ex: `int`)
  - Exemplo: `fmt.Printf("%d", 10)` → `10`
- **%s**: Formata uma string
  - Exemplo: `fmt.Printf("%s", "texto")` → `texto`
- **%f**: Formata um número de ponto flutuante (float)
  - Exemplo: `fmt.Printf("%f", 3.14)` → `3.140000`
- **%v**: Formata qualquer valor (valor "padrão" para o tipo)
  - Exemplo: `fmt.Printf("%v", true)` → `true`
- **%T**: Mostra o tipo do valor
  - Exemplo: `fmt.Printf("%T", 10)` → `int`

o que é [][]string, uma matriz de strings?

Sim! `[][]string` é uma **matriz de strings** (ou slice de slices de string).

- Cada elemento do slice externo é um slice de strings.
- Usado, por exemplo, para representar linhas e colunas de um CSV:
  - Cada linha do CSV é um `[]string` (cada coluna é um campo).
  - O arquivo inteiro é um `[][]string` (todas as linhas).

**Exemplo:**
```go
dados := [][]string{
    {"1", "Joao", "Brasil"},
    {"2", "Maria", "Argentina"},
}
```
Aqui, `dados[0][1]` seria `"Joao"`.