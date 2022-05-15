package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper te permite definir pausas
type Sleeper interface {
	Pausa()
}

// SleeperConfiguravel é uma implementação de Sleepr com uma pausa definida
type SleeperConfiguravel struct {
	duracao time.Duration
	pausa   func(time.Duration)
}

// Pausa vai pausar a execução pela Duração definida
func (s *SleeperConfiguravel) Pausa() {
	s.pausa(s.duracao)
}

const ultimaPalavra = "Vai!"
const inicioContagem = 3

// Contagem imprime uma contagem de 3 para a saída
func Contagem(saida io.Writer, sleeper Sleeper) {
    for i := inicioContagem; i > 0; i-- {
		sleeper.Pausa()
        fmt.Fprintln(saida, i)
    }

	sleeper.Pausa()
    fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperConfiguravel{1 * time.Second, time.Sleep}
	Contagem(os.Stdout, sleeper)
}

// É importante ser capaz de dividir os requerimentos da menor forma que conseguir para você ter um software funcionando.
// Contagem precisa escrever dados em algum lugar e o io.Writer é a forma de capturarmos essa saída como uma interface em Go.
// Na main, o os.Stdout é enviado como parâmetro para nossos usuários verem a contagem regressiva impressa no terminal.
// No teste, o bytes.Buffer é enviado como parâmetro para que nossos testes possam capturar que dado está sendo gerado.