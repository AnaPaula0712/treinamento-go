package v1

import "sync"

// Contador incrementa um número
type Contador struct {
	mu sync.Mutex
	valor int
}

// NovoContador retorna um novo Contador
func NovoContador() *Contador {
	return &Contador{}
}

// Incrementa o contador
func (c *Contador) Incrementa() {
	c.mu.Lock()
    defer c.mu.Unlock()
	c.valor++
}

// Valor retorna a contagem atual
func (c *Contador) Valor() int {
	return c.valor
}


// Um Mutex é uma trava de exclusão mútua.
// O valor zero de um Mutex é um Mutex destravado.
// Isso significa que qualquer goroutine chamando Incrementa vai receber a trava em Contador se for a primeira chamando essa função.
// Todas as outras goroutines vão ter que esperar por essa primeira execução até que ele esteja Unlock, ou destravado,
// antes de ganhar o acesso à instância de Contador alterada pela primeira chamada de função.
// Um Mutex não deve ser copiado depois do primeiro uso.