package main

func Soma(numeros []int) int {
    soma := 0

	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

// func SomaTudo(numerosParaSomar ...[]int) (somas []int) {
//     quantidadeDeNumeros := len(numerosParaSomar)
// 	somas = make([]int, quantidadeDeNumeros)

// 	for i, numeros := range numerosParaSomar {
// 		somas[i] = Soma(numeros)
// 	}

// 	return
// }

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
    var somas []int
	for _, numeros := range numerosParaSomar {
		if len(numeros) == 0 {
            somas = append(somas, 0)
        } else {
            final := numeros[1:]
            somas = append(somas, Soma(final))
        }
	}

	return somas
}

// Para receber o valor de um array em uma posição específica, basta usar a sintaxe array[índice].
// O range permite que você percorra um array. Sempre que é chamado,
// retorna dois valores: o índice e o valor.
// o valor do índice foi ignorado usando o _ blank identifier.
// O make te permite criar um slice com
// uma capacidade inicial de len de numerosParaSomar que precisamos percorrer.
// Slices podem ser "fatiados"! A sintaxe usada é slice[inicio:final].
// Se for omitido o valor de um dos lados dos : ele captura tudo do lado omitido.
// Aqui, quando usamos numeros[1:], estamos dizendo "pegue da posição 1 até o final".