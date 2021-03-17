package area

import "math"

// Pi proporção numérica , perímetro de uma circunferência
const Pi = 3.1416

// Circ Calcula área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect Calcula area de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visível
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
