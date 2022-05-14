package main

import (
	"bytes"
	"testing"
)

func TestContagem(t *testing.T) {
	buffer := &bytes.Buffer{}

	Contagem(buffer)

	resultado := buffer.String()
	esperado := `3
	2
	1
	Vai!`

	if resultado != esperado {0
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

// A sintaxe de aspas simples é outra forma de criar uma string,
// mas te permite colocar coisas como linhas novas, o que é perfeito para nosso teste.