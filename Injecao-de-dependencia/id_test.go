package main

import (
	"bytes"
	"testing"
)

func TestCumprimenta(t *testing.T) {
    buffer := bytes.Buffer{}
    Cumprimenta(&buffer, "Chris")

    resultado := buffer.String()
    esperado := "Olá, Chris"

    if resultado != esperado {
        t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
    }
}

// O tipo buffer do pacote bytes implementa a interface Writer.
// Refatoramos o código para que pudéssemos controlar para onde os dados eram escritos injetando uma dependência que nos permitiu:
// 1) Testar nosso código: se você não consegue testar uma função de forma simples, geralmente é porque dependências estão acopladas em uma função ou estado global.
// 2) Separar nossas preocupações, desacoplando onde os dados vão de como gerá-los. Se você já achou que um método/função tem responsabilidades demais a injeção de dependência será útil.
// 3) Permitir que nosso código seja reutilizado em contextos diferentes.