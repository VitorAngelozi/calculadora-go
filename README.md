# calculadora-go
Calculadora interativa em Go feita para estudo
# calculadora-go
Calculadora interativa em Go feita para estudo
# Calculadora Interativa em Go

Projeto desenvolvido em **Go (Golang)** com o objetivo de praticar conceitos de modularização, entrada e saída, controle de fluxo e pacotes. 

O programa consiste em uma **calculadora de terminal** capaz de realizar operações básicas e avançadas:

- Soma
- Subtração
- Multiplicação
- Divisão
- Potência
- Raiz quadrada
- Resolução de equações do segundo grau

---

- `main.go`: Contém o menu interativo, leitura de entradas do usuário e chamadas para o pacote `matematica`.  
- `matematica/matematica.go`: Contém todas as funções matemáticas implementadas como um pacote separado, seguindo boas práticas de modularização.

---

## Funcionalidades

1. **Soma, Subtração, Multiplicação, Divisão**  
   - Operações básicas com tratamento de tipos (`int` e `float64`).  

2. **Potência e Raiz Quadrada**  
   - Utiliza funções do pacote `math` (`math.Pow`, `math.Sqrt`) para cálculos precisos.  

3. **Equação do Segundo Grau**  
   - Calcula raízes reais de equações do segundo grau (`ax² + bx + c = 0`).  
   - Retorna `NaN` caso não haja soluções reais.  

---

## Como Executar

### Rodando diretamente com Go
```bash
go run .

### Rodando com o terminal

go build
./calculadora-go   # No Linux/Mac
calculadora-go.exe # No Windows

##Exemplo de uso
===================================
         CALCULADORA GO
===================================
1 - Somar
2 - Subtrair
3 - Multiplicar
4 - Dividir
5 - Potência
6 - Raiz Quadrada
7 - Equação do 2º Grau
0 - Sair
Escolha uma opção: 1
Digite o primeiro número: 5
Digite o segundo número: 7
Resultado: 12


