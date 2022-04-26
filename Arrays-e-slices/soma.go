package main

func Soma(numeros []int) int {
    soma := 0

	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

// Para receber o valor de um array em uma posição específica, basta usar a sintaxe array[índice].
// O range permite que você percorra um array. Sempre que é chamado,
// retorna dois valores: o índice e o valor.
// o valor do índice foi ignorado usando o _ blank identifier.