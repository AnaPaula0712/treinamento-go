package main

// import (
// 	"errors"
// )

type Dicionario map[string]string

const (
    ErrNaoEncontrado = ErrDicionario("não foi possível encontrar a palavra que você procura")
    ErrPalavraExistente = ErrDicionario("não é possível adicionar a palavra pois ela já existe")
	ErrPalavraInexistente = ErrDicionario("não foi possível atualizar a palavra pois ela não existe")
)

type ErrDicionario string

func (e ErrDicionario) Error() string {
    return string(e)
}

func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]
	if !existe {
		return "", ErrNaoEncontrado
	}

	return definicao, nil
}

func (d Dicionario) Adiciona(palavra, definicao string) error {
	_, err := d.Busca(palavra)
    switch err {
    case ErrNaoEncontrado:
        d[palavra] = definicao
    case nil:
        return ErrPalavraExistente
    default:
        return err

    }

	return nil
}

func (d Dicionario) Atualiza(palavra, definicao string) error {
	_, err := d.Busca(palavra)
    switch err {
    case ErrNaoEncontrado:
        return ErrPalavraInexistente
    case nil:
        d[palavra] = definicao
    default:
        return err

    }

	return nil
}

func (d Dicionario) Deleta(palavra string) {
	delete(d, palavra)
}

// Maps te permitem armazenar itens de forma parecida com a de um dicionário.
// Você pode pensar na chave como a palavra e o valor como a definição.
// Obter um valor de um map é igual a obter um valor de um array: map[chave]
// Adicionar coisas a um map também é bem semelhante a um array.
// Você só precisar especificar uma chave e definir qual é seu valor.
// É possível modificar o map sem passá-lo como ponteiro. Isso é porque o map é um tipo referência.
// Isso significa que ele contém uma referência à estrutura de dado que estamos utilizando, assim como um ponteiro.
// Logo, quando criamos passamos o map como parâmetro, estamos alterando o map original e não sua cópia.
//  você nunca deve inicializar um map vazio, como: var m map[string]string
// Mas você pode inicializar um map vazio usando a palavra-chave make para criar um map para você: dicionario = make(map[string]string)
// Go tem uma função nativa chamada delete que funciona em maps.
// Ela leva dois argumentos: o primeiro é o map e o segundo é a chave a ser removida.
// A função delete não retorna nada, e baseamos nosso método Deleta nesse conceito.
// Já que deletar um valor não tem nenhum efeito, diferentemente dos nossos métodos Atualiza e Adiciona, não precisamos implementar erros.