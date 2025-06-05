**Exercício Integrado – Impostos sobre o Salário**

**Objetivo:**  
Desenvolver um conjunto de funções e validações para calcular e validar salários a partir do valor mensal e do número de horas trabalhadas, utilizando tratamento de erros personalizado em diferentes abordagens. O objetivo é consolidar conceitos de uso de erros customizados, funções que retornam erros e múltiplos valores, além de práticas com `Is()`, `errors.New()`, e `fmt.Errorf()`.

---

**Enunciado:**  

1. Em sua função principal (`main`), declare uma variável chamada `salary` do tipo `int` e atribua um valor à sua escolha.

2. Implemente validações sequenciais para a variável `salary`, utilizando diferentes técnicas de tratamento de erro:
   
   - Se o valor de `salary` for **menor que 150000**, crie e retorne um erro customizado com uma estrutura que implemente o método `Error()`, contendo a mensagem:  
     `"Error: the salary entered does not reach the taxable minimum"`  
     Caso contrário, imprima no console:  
     `"Must pay tax"`
     
   - Se o valor de `salary` for **menor ou igual a 10000**, crie um erro customizado com uma estrutura que implemente o método `Error()` e a mensagem:  
     `"Error: salary is less than 10000"`  
     Utilize a função `Is()` para validar e tratar este erro dentro do `main`.
     
   - Refatore a checagem anterior para que, em vez de implementar `Error()`, utilize `errors.New()` para gerar o erro e trate-o apropriadamente no `main`.

   - Por último, implemente uma validação usando `fmt.Errorf()`, onde a mensagem de erro informe o valor de `salary` inserido, no formato:  
     `"Error: the minimum taxable amount is 150,000 and the salary entered is: [salary]"`  
     Troque `[salary]` pelo valor efetivo informado.
  
3. Desenvolva uma função para permitir o cálculo do salário mensal de um trabalhador, a partir do número de horas trabalhadas e do valor da hora:

   - A função deve receber como parâmetros o número de horas trabalhadas no mês e o valor da hora.
   - A função deve retornar dois valores: o salário calculado e um erro.
   - Se o salário mensal calculado for igual ou superior a **150000**, deduza 10% de imposto sobre o valor.
   - Caso o número de horas seja **menor que 80** ou negativo, retorne um erro com a seguinte mensagem:  
     `"Error: the worker cannot have worked less than 80 hours per month"`
   - Trate a chamada dessa função no `main`, mostrando o salário líquido ou a mensagem de erro correspondente.

---

**O que se espera:**  
- Utilização correta de erros customizados em Go, com implementações usando estrutura própria, `errors.New()`, e `fmt.Errorf()`.
- Aplicação de validações sequenciais sobre as condições propostas.
- Implementação de função com múltiplos retornos, incluindo tratamento adequado para valores e erros.
- Correta exibição de mensagens no console, de acordo com os diferentes cenários e validações.

**Retornos esperados:**  
- Mensagens de erro apropriadas para cada cenário inválido
- Mensagem `"Must pay tax"` quando aplicável
- Resultado do salário mensal com desconto do imposto, quando pertinente

