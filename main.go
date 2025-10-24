package main

import (
	"fmt"
	"modulo/matematica"
)

func main() {
	for {
		fmt.Println("===================================")
		fmt.Println("         🧮 CALCULADORA GO         ")
		fmt.Println("===================================")
		fmt.Println("1 - Somar")
		fmt.Println("2 - Subtrair")
		fmt.Println("3 - Multiplicar")
		fmt.Println("4 - Dividir")
		fmt.Println("5 - Potência")
		fmt.Println("6 - Raiz Quadrada")
		fmt.Println("7 - Equação do 2º Grau")
		fmt.Println("0 - Sair")
		fmt.Println("===================================")
		fmt.Print("Escolha uma opção: ")

		var opcao int
		fmt.Scanln(&opcao)

		if opcao == 0 {
			fmt.Println("Saindo... 👋")
			break
		}

		switch opcao {
		case 1:
			var a, b int
			fmt.Print("Digite o primeiro número: ")
			fmt.Scanln(&a)
			fmt.Print("Digite o segundo número: ")
			fmt.Scanln(&b)
			fmt.Printf("Resultado: %d\n", matematica.Somar(a, b))

		case 2:
			var a, b int
			fmt.Print("Digite o primeiro número: ")
			fmt.Scanln(&a)
			fmt.Print("Digite o segundo número: ")
			fmt.Scanln(&b)
			fmt.Printf("Resultado: %d\n", matematica.Subtrair(a, b))

		case 3:
			var a, b int
			fmt.Print("Digite o primeiro número: ")
			fmt.Scanln(&a)
			fmt.Print("Digite o segundo número: ")
			fmt.Scanln(&b)
			fmt.Printf("Resultado: %d\n", matematica.Multiplicar(a, b))

		case 4:
			var a, b int
			fmt.Print("Digite o dividendo: ")
			fmt.Scanln(&a)
			fmt.Print("Digite o divisor: ")
			fmt.Scanln(&b)
			fmt.Printf("Resultado: %.2f\n", matematica.Dividir(a, b))

		case 5:
			var a, b int
			fmt.Print("Digite a base: ")
			fmt.Scanln(&a)
			fmt.Print("Digite o expoente: ")
			fmt.Scanln(&b)
			fmt.Printf("Resultado: %.2f\n", matematica.Potencia(a, b))

		case 6:
			var a float64
			fmt.Print("Digite o número: ")
			fmt.Scanln(&a)
			fmt.Printf("Resultado: %.2f\n", matematica.RaizQuadrada(a))

		case 7:
			var a, b, c float64
			fmt.Print("Digite o valor de a: ")
			fmt.Scanln(&a)
			fmt.Print("Digite o valor de b: ")
			fmt.Scanln(&b)
			fmt.Print("Digite o valor de c: ")
			fmt.Scanln(&c)
			x1, x2, delta := matematica.EquacaoSegundoGrau(a, b, c)
			if delta < 0 {
				fmt.Println("A equação não possui raízes reais.")
			} else {
				fmt.Printf("x1 = %.2f, x2 = %.2f, delta = %.2f\n", x1, x2, delta)
			}

		default:
			fmt.Println("Opção inválida, tente novamente.")
		}

		fmt.Println()
	}
}
