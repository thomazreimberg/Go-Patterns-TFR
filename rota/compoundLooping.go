package rota

import "fmt"

func E1_Somar_Matriz(matriz [][]int) int {
	 var somar int
	 for iAux, i := range matriz {
		    for jAux := range i {
            somar += matriz[iAux][jAux]
        }
    }
	fmt.Println(somar)
	 return somar
}

func E2_Ordenar(numeros []int)[]int{
	for i := 0; i < len(numeros); i++ {
		for j := 0; j < len(numeros); j++ {
			if numeros[j] < numeros[i] {
				aux := numeros[i]
				numeros[i] = numeros[j]
				numeros[j] = aux
			}
		}
	}
	fmt.Println(numeros)
	return numeros
}
