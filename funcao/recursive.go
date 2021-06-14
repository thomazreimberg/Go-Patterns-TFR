// Thomaz
package funcao

func E1_Fatorial_Linear(x int) int {
	if (x <= 1) {
		return 1
	} else {
		return x * E1_Fatorial_Linear(x - 1)
	}
}

func E2_Fatorial_Iterativo(x int) int {
	return e2_Fatorial_IterativoHelper(x, 1)
}

func e2_Fatorial_IterativoHelper(x int, fat int) int {
	if (x == 1) {
		return fat
	} else {
		return e2_Fatorial_IterativoHelper(x - 1, fat * x)
	}
}