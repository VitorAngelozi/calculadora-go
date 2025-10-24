# Calculadora Interativa em Go

Projeto desenvolvido em **Go (Golang)** com o objetivo de praticar conceitos de **modularização, entrada e saída, controle de fluxo e pacotes**.

A aplicação é uma **calculadora de terminal** capaz de realizar operações básicas e avançadas de forma interativa.

---

## Funcionalidades

1. **Operações básicas**
   - Soma
   - Subtração
   - Multiplicação
   - Divisão
   - Suporte a `int` e `float64`

2. **Operações avançadas**
   - Potência (`math.Pow`)
   - Raiz quadrada (`math.Sqrt`)

3. **Equação do segundo grau**
   - Calcula raízes reais da equação `ax² + bx + c = 0`
   - Retorna `NaN` caso não haja soluções reais

---

## Estrutura do Projeto

calculadora-go/
│
├─ main.go # Menu interativo e chamadas para o pacote matematica
└─ matematica/
└─ matematica.go # Funções matemáticas como pacote separado


- `main.go`: Contém o menu, leitura de entradas e chamadas para o pacote.
- `matematica/matematica.go`: Contém todas as funções matemáticas implementadas de forma modular.

---

## Pré-requisitos

Antes de executar o projeto, certifique-se de ter o **Go instalado**:

```bash
go version
# Deve exibir algo como: go version go1.21.1 windows/amd64

##Como executar
###1- Clonando o projeto
git clone https://github.com/seu-usuario/calculadora-go.git
cd calculadora-go

###2- rodando diretamente do go
go run .

###3. Compilando para gerar executável
go build -o calculadora-go

# Executando:
./calculadora-go   # Linux/Mac
calculadora-go.exe # Windows

#Exemplo de uso
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

Tecnologias Utilizadas

Go (Golang)

Pacote fmt para entrada e saída

Pacote math para operações matemáticas avançadas

Modularização com pacotes próprios