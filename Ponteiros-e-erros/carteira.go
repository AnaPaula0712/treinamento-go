package main

import "fmt"
import "errors"

// Bitcoin representa o número de Bitcoins
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Carteira armazena o número de bitcoins que uma pessoa tem
type Carteira struct {
	saldo Bitcoin
}

// Depositar vai adicionar Bitcoins à carteira
func (c Carteira) Depositar(quantidade Bitcoin){
	c.saldo += quantidade
}

// ErroSaldoInsuficiente significa que uma carteira não tem Bitcoins suficientes para fazer uma retirada
var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

// Retirar vai tirar Bitcoins da carteira
func (c *Carteira) Retirar(quantidade Bitcoin) error {

	if quantidade > c.saldo {
		return ErroSaldoInsuficiente
	}

	c.saldo -= quantidade
	return nil
}

// Saldo retorna o número de Bitcoins que uma carteira tem
func (c *Carteira) Saldo() Bitcoin{
	return c.saldo
}

// Em Go, se uma variável, tipo, função e etc, começam com uma letra minúsculo,
// então esta será privada para outros pacotes que não seja o que a definiu.
// O \n é um caractere de escape queeadiciona uma nova linha após imprimir o endereço de memória.
// Conseguimos acessar o ponteiro para algo com o símbolo de endereço &.
//  Ponteiros nos permitem apontar para alguns valores e então mudá-los.
// Go permite criarmos novos tipos a partir de tipos existentes.
// A sintaxe é type MeuNome TipoOriginal
//  A palavra-chave var no escopo do arquivo nos permite definir valores globais para o pacote.