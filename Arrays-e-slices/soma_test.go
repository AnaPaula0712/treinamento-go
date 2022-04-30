package main

import "testing"
import "reflect"

func TestSomaTodoOResto(t *testing.T) {

    verificaSomas := func(t *testing.T, resultado, esperado []int) {
        t.Helper()
        if !reflect.DeepEqual(resultado, esperado) {
            t.Errorf("resultado %v, esperado %v", resultado, esperado)
        }
    }

    t.Run("faz a soma do resto", func(t *testing.T) {
        resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
        esperado := []int{2, 9}
        verificaSomas(t, resultado, esperado)
    })

    t.Run("soma slices vazios de forma segura", func(t *testing.T) {
        resultado := SomaTodoOResto([]int{}, []int{3, 4, 5})
        esperado := []int{0, 9}
        verificaSomas(t, resultado, esperado)
    })

    // func TestSomaTudo(t *testing.T) {

    //     recebido := SomaTudo([]int{1,2}, []int{0,9})
    //     esperado := []int{3, 9}

    //     if !reflect.DeepEqual(recebido, esperado) {
    //         t.Errorf("resultado %v esperado %v", recebido, esperado)
    //     }
    // }

}

// O Go não te deixa usar operadores de igualdade com slices.
// É possível escrever uma função que percorre cada slice recebido e esperado e
// verificar seus valores, mas por praticidade podemos usar o reflect.DeepEqual
// que é útil para verificar se duas variáveis são iguais.
// Para usar o reflect.DeepEqual necessário importa o "reflect".
// É importante saber que o reflect.DeepEqual não tem "segurança de tipos",
// O código compila mesmo se você tentar comparar um string com um slice, por exemplo.