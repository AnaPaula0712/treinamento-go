package corredor

import (
	"fmt"
	"net/http"
	"time"
)

var limiteDeDezSegundos = 10 * time.Second

// Corredor compara os tempos de resposta de a e b, retornando o mais rápido com tempo limite de 10s
func Corredor(a, b string) (vencedor string, error error) {
	return Configuravel(a, b, limiteDeDezSegundos)
}

// Configuravel compara os tempos de resposta de a e b, retornando o mais rápido
func Configuravel(a, b string, tempoLimite time.Duration) (vencedor string, erro error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(tempoLimite):
		return "", fmt.Errorf("tempo limite de espera excedido para %s e %s", a, b)
	}
}

func ping(URL string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(URL)
		close(ch)
	}()
	return ch
}

// select nos ajuda a sincronizar os processos de forma mais fácil e clara.
// O que o select permite fazer é aguardar múltiplos canais.
// O primeiro a enviar um valor "vence" e o código abaixo do case é executado.
// Foi usado ping em nosso select para configurar um canal para cada uma de nossas URLs.
// Qualquer um que enviar para esse canal primeiro vai ter seu código executado no select, que resultará nessa URL sendo retornada (que será a vencedora).
// time.After é uma função muito útil quando usamos select, ele retorna um chan (como ping) e enviará um sinal após a quantidade de tempo definida.
// Select - resumo:
// 1) Ajuda você a escutar vários canais.
// 2) Às vezes você pode precisar incluir time.After em um de seus cases para prevenir que seu sistema fique bloqueado para sempre.
// httptest - resumo:
// 1) Uma forma conveniente de criar servidores de teste para que se tenha testes confiáveis e controláveis.
// 2)Usa as mesmas interfaces que servidores net/http reais, o que torna seu sistema consistente e gera menos coisas para você aprender.