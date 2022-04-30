package main

import "testing"

//  retorna o perímetro de um retângulo
func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
	}
}

// retorna a área de uma forma
func TestArea(t *testing.T) {
	testesArea := []struct {
		nome     string
		forma    Forma
		temArea float64
	}{
		{nome: "Retângulo", forma: Retangulo{Largura: 10.0, Altura: 10.0}, temArea: 100.0},
		{nome: "Círculo", forma: Circulo{Raio: 10}, temArea: 314.1592653589793},
		{nome: "Triangulo", forma: Triangulo{Base: 12, Altura:6}, temArea: 36.0},
	}

	for _, tt := range testesArea {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := tt.forma.Area()
			if resultado != tt.temArea {
				t.Errorf("%#v resultado %.2f esperado %.2f", tt.forma, resultado, tt.temArea)
			}
		})
	}
}

// O f é para o float64 e o .2 significa imprimir duas casas decimais.
// Um método é uma função com um receptor. Uma declaração de método vincula
// um identificador e o nome do método a um método e associa o método com
// o tipo base do receptor.

// Interfaces são um conceito muito poderoso em linguagens de programação estaticamente tipadas, como Go,
// porque permitem que você crie funções que podem ser usadas com diferentes tipos e permite a criação de
// código altamente desacoplado, mantendo ainda a segurança de tipos.

// Em Go a resolução de interface é implícita. Se o tipo que você passar combinar com o que a interface
// está esperando, o código será compilado.

// Table driven tests são úteis quando você quer construir uma lista de casos de testes
// que podem ser testados da mesma forma.