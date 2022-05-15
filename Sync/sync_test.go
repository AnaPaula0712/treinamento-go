package v1

import (
	"sync"
	"testing"
)

func TestContador(t *testing.T) {
	t.Run("incrementar o contador 3 vezes resulta no valor 3", func(t *testing.T) {
		contador := NovoContador()
		contador.Incrementa()
		contador.Incrementa()
		contador.Incrementa()

		verificaContador(t, contador, 3)
	})

	t.Run("roda concorrentemente em segurança", func(t *testing.T) {
		contagemEsperada := 1000
		contador := NovoContador()

		var wg sync.WaitGroup
		wg.Add(contagemEsperada)

		for i := 0; i < contagemEsperada; i++ {
			go func(w *sync.WaitGroup) {
				contador.Incrementa()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		verificaContador(t, contador, contagemEsperada)
	})
}

func verificaContador(t *testing.T, resultado *Contador, esperado int) {
	t.Helper()
	if resultado.Valor() != esperado {
		t.Errorf("resultado %d, esperado %d", resultado.Valor(), esperado)
	}
}

// Estamos usando sync.WaitGroup,
// que é uma maneira simples de sincronizar processos concorrentes.
// Um WaitGroup aguarda por uma coleção de goroutines terminar seu processamento.
// A goroutine principal faz a chamada para o Add definir o número de goroutines que serão esperadas.
// Então, cada uma das goroutines é executada e chama Done quando termina sua execução.
// Ao mesmo tempo, Wait pode ser usado para bloquear a execução até que todas as goroutines tenham terminado.
// Ao esperar por wg.Wait() terminar sua execução antes de fazer nossas asserções,
// podemos ter certeza que todas as nossas goroutines tentaram chamar o Incrementa no Contador.
// Pacote Sync - Resumo:
// 1) Mutex nos permite adicionar travas aos nossos dados.
// 2) WaitGroup é uma maneira de esperar as goroutines terminarem suas tarefas.
// Um erro comum de um iniciante em Go é usar demais os channels e goroutines apenas porque é possível.
// Não tenha medo de usar um sync.Mutex se for uma solução melhor para o seu problema. Resumindo:
// 1) Use channels quando for passar a propriedade de um dado
// 2) Use mutexes para gerenciar estados