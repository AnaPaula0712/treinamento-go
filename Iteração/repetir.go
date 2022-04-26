package iteracao

const quantidadeRepeticoes = 5

// Repetir retorna o caractere repetido 5 vezes
func Repetir(caractere string) string {
	var repeticoes string
	// estágios do for: inicialização, condição e pós
	for i:= 0; i < quantidadeRepeticoes; i++ {
		repeticoes += caractere
	}
	return repeticoes
}

// inicialização: O i vai começar com valor 0
// condição: O código dentro das {} vai rodar enquanto o valor de i for menor que 5
// pós: Em cada looping do for o i será incrementado em 1 (uma unidade)
// As 3 instruções ou estágios do for são separadas por ponto e vírgula (;)
// O operador adicionar & atribuir += adiciona o valor que está à direita no valor que esta à esquerda e atribui o resultado ao valor da esquerda.