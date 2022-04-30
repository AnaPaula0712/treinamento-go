package main

import "math"

// Forma é implementado por qualquer coisa que possa dizer qual é sua área
type Forma interface {
	Area() float64
}

type Retangulo struct {
	Largura float64
	Altura float64
}

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}
//  nomeDoReceptor = r TipoDoReceptor = Retangulo NomeDoMetodo = Area  argumentos = ()
func (r Retangulo) Area() float64 {
	return (r.Largura * r.Altura)
}
// func Area(retangulo Retangulo) float64 {
// 	return (retangulo.Largura * retangulo.Altura)
// }

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * (c.Raio * c.Raio)
}

type Triangulo struct {
	Base float64
	Altura float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura)/2
}
// Uma struct é apenas uma coleção nomeada de campos onde você pode armazenar dados.
// Você pode acessar os campos de uma struct com a sintaxe minhaStruct.campo.
// A sintaxe para declaração de métodos é quase a mesma que usamos para funções e isso acontece porque eles são muito parecidos.
// A única diferença é a sintaxe para o método receptor: func (nomeDoReceptor TipoDoReceptor) NomeDoMetodo(argumentos).
// É uma convenção em Go que a variável receptora seja a primeira letra do tipo em minúsculo.
// A constante Pi está no pacote math (lembre-se de importá-lo).