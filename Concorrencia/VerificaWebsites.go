package concorrencia

// VerificadorWebsite verifica uma URL, retornando uma booleana
type VerificadorWebsite func(string) bool
type resultado struct {
    string
    bool
}

// VerificaWebsites recebe um VerificadorWebsite e um slice de URLs e retorna um map
// de URLs com o resultado da verificação de cada URL com a função VerificadorWebsite
// cada iteração do loop vai iniciar uma nova goroutine, concorrente com o processo atual (a função VerificadorWebsite),
// e cada uma vai adicionar seu resultado ao map de resultados.
func VerificaWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
	resultados := make(map[string]bool)
	canalResultado := make(chan resultado)

	for _, url := range urls {
		go func (u string)  {
			canalResultado <- resultado{u, vw(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
        resultado := <-canalResultado
        resultados[resultado.string] = resultado.bool
    }

	return resultados
}

// concorrência que, para fins dessa situação, significa "fazer mais do que uma coisa ao mesmo tempo".
// Ao invés de esperar por um website responder antes de enviar uma requisição para o próximo website,
// vamos dizer para nosso computador fazer a próxima requisição enquanto espera pela primeira.
// Quando temos que esperar algo acabar para terminar seu trabalho ela é uma operação de bloqueante.
// Uma operação que não bloqueia no Go vai rodar em um processo separado, chamado de goroutine.
// Para dizer ao Go começar uma nova goroutine, transformamos a chamada de função em uma declaração go colocando a palavra-chave go na frente da função: go fazAlgumaCoisa().
// a única forma de começar uma goroutine é colocar go na frente da chamada de função, costumamos usar funções anônimas quando queremos iniciar uma goroutine.
// Uma função anônima literal é bem parecida com uma declaração de função normal, mas (obviamente) sem um nome.
// Funções anônimas têm várias funcionalidades úteis, como por exemplo:
// 1) Podem ser executadas assim que fazemos sua declaração - que é o () no final da função anônima.
// 2) Elas mantém acesso ao escopo léxico em que são definidas.
// Canais são uma estrutura de dados em Go que pode receber e enviar valores.
// Junto do map resultados, há agora o canalResultados, criado usando make.
// O chan resultado é o tipo do canal - um canal de resultado.
// O tipo novo, resultado, foi criado para associar o retorno de VerificadorWebsite com a URL sendo verificada, que contém uma string e um bool.
// Dentro, estamos usando uma expressão de recebimento, que atribui um valor recebido por um canal a uma variável.
// Essa expressão também usa o operador <-, mas com os dois operandos em posições invertidas:
// o canal agora fica à direita e a variável que está recebendo o valor dele fica à esquerda.