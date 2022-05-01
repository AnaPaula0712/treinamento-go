package main

import "testing"

func TestCarteira(t *testing.T) {
	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))

		confirmaSaldo(t, carteira, Bitcoin(0))
	})


	t.Run("Retirar com saldo suficiente", func(t *testing.T) {
        carteira := Carteira{Bitcoin(20)}
        erro := carteira.Retirar(Bitcoin(10))

		confirmaSaldo(t, carteira, Bitcoin(10))
		confirmaErroInexistente(t, erro)
    })

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))

		confirmaSaldo(t, carteira, saldoInicial)
		confirmaErro(t, erro, ErroSaldoInsuficiente)
	})
}

func confirmaSaldo(t *testing.T, carteira Carteira, esperado Bitcoin) {
    t.Helper()
    resultado := carteira.Saldo()

    if resultado != esperado {
        t.Errorf("resultado %s, esperado %s", resultado, esperado)
    }
}

func confirmaErroInexistente(t *testing.T, resultado error) {
	t.Helper()
	if resultado != nil {
		t.Fatal("erro inesperado recebido")
	}
}

func confirmaErro(t *testing.T, resultado error, esperado error) {
    t.Helper()
    if resultado == nil {
        t.Fatal("esperava um erro, mas nenhum ocorreu")
    }

    if resultado != esperado {
        t.Errorf("erro resultado %s, erro esperado %s", resultado, esperado)
    }
}

// PONTEIRO é uma varárel que armazena um endereço de memória.
// Descobrir o endereço desse bit de memória usando &nomeDaVariável.
// AI usar o asterisco (*) você está fazendo o de reference...
// vc tem o endereço e vc quer saber o que tem dentro daquele endereço.
// Em Go, se você quiser indicar um erro,
// sua função deve retornar um err para que quem a chamou possar verificá-lo e tratá-lo.
// Verificamos se um erro foi retornado falhando o teste se o valor for nil.
// Usamos o t.Fatal que interromperá o teste se for chamado.
// Isso é feito porque não queremos fazer mais asserções no erro retornado, se não houver um.
// Sem isso, o teste continuaria e causaria erros por causa do ponteiro nil.