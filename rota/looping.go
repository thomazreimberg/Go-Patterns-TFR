//Thomaz
package rota

import "math"

func E1_FOR_SomarNumeros(fim int) int {
	soma := 0
	for i := 0; i <= fim; i++ {
		soma += i
	}

	return soma
}

func E2_FOREACH_FiltrarPares(numeros []int) []int {
	var pares []int
	//0: "Carro 1"
	//1: "Carro 2"
	for _, item := range numeros {
		if (item % 2 == 0) {
			pares = append(pares, item)
		}
	}

	return pares
}

func E3_CONTINUE_GerarSequenciaPar(fim int) []int {
	var seq []int

	for i := 0; i <= fim; i++ {
		if i % 2 == 0 {
			continue
		}
		seq = append(seq, i)
	}

	return seq
}

func E4_BREAK_TodosPares(numeros []int) bool {
	pares := true

	for _, item := range numeros {
		if (item % 2 != 0) {
			pares = false
			break
		}
	}

	return pares
}

func E5_WHILE_ProximaRaizInteira(numero int) int {
	resto := 1

	for resto != 0 {
		numero ++
		raiz := int(math.Floor(math.Sqrt(float64(numero))))
		resto = raiz % 1
	}

	return numero
}

func E6_DOWHILE_Somar(inicio int, fim int) []int {
	var seq []int

	for {
		seq = append(seq, inicio + 1)
		
        if inicio >= fim {
			break
        }
        inicio ++
    }
	
	return seq
}
