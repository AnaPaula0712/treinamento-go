package main

import "fmt"

const espanhol = "espanhol"
const francês = "francês"
const alemão = "alemão"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaAlemao = "Hallo, "

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixoSaudacao(idioma) + nome
}

func prefixoSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case francês:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case alemão:
		prefixo = prefixoOlaAlemao
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

func main() {
    fmt.Println(Ola("Mundo", ""))
}