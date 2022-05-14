package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const ultimaPalavra = "Vai"
const inicioContagem = 3

// Contagem imprime uma contagem de 3 para a saída
func Contagem(saida io.Writer) {
    for i := inicioContagem; i > 0; i-- {
        fmt.Fprintln(saida, i)
    }

	time.Sleep(1 * time.Second)
    fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	Contagem(os.Stdout)
}

// É importante ser capaz de dividir os requerimentos da menor forma que conseguir para você ter um software funcionando.
// Contagem precisa escrever dados em algum lugar e o io.Writer é a forma de capturarmos essa saída como uma interface em Go.
// Na main, o os.Stdout é enviado como parâmetro para nossos usuários verem a contagem regressiva impressa no terminal.
// No teste, o bytes.Buffer é enviado como parâmetro para que nossos testes possam capturar que dado está sendo gerado.