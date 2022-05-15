package context3

import (
	"context"
	"fmt"
	"net/http"
)

// Store busca dados
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server retorna um handler para chamar a Store
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: registre o erro como você quiser
		}

		fmt.Fprint(w, data)
	}
}

// Pacote context usado para ajudar a gerenciar processos de longa duração.