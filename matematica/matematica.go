package matematica

import "math"

func Somar(a, b int) int {
	return a + b
}

func Subtrair(a, b int) int {
	return a - b
}

func Multiplicar(a, b int) int {
	return a * b
}

func Dividir(a, b int) float64 {
	return float64(a) / float64(b)
}

func Potencia(a, b int) float64 {
	return math.Pow(float64(a), float64(b))
}

func RaizQuadrada(a float64) float64 {
	return math.Sqrt(a)
}

func EquacaoSegundoGrau(a, b, c float64) (float64, float64, float64) {
	delta := (b * b) - 4*a*c

	if delta < 0 {
		return math.NaN(), math.NaN(), delta
	}

	x1 := (-b + math.Sqrt(delta)) / (2 * a)
	x2 := (-b - math.Sqrt(delta)) / (2 * a)
	return x1, x2, delta
}
