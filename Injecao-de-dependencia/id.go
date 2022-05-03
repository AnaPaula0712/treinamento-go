package main

import (
	"fmt"
	"io"
	"net/http"
)

func Cumprimenta(escritor io.Writer, nome string) {
    fmt.Fprintf(escritor, "Olá, %s", nome)
}

func HandlerMeuCumprimento(w http.ResponseWriter, r *http.Request) {
    Cumprimenta(w, "mundo")
}

func main() {
    err := http.ListenAndServe(":5000", http.HandlerFunc(HandlerMeuCumprimento))

    if err != nil {
        fmt.Println(err)
    }
}

// o fmt.Fprintf te permite passar um io.Writer, que sabemos que o os.Stdout e bytes.Buffer implementam.
// Quando se cria um handler HTTP, você recebe um http.ResponseWriter e o http.Request que é usado para fazer a requisição.
// Quando implementa seu servidor, você escreve sua resposta usando o escritor.
// O http.ResponseWriter também implementa o io.Writer e é por isso que podemos reutilizar nossa função Cumprimenta dentro do nosso handler.