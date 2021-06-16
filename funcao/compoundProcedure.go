//Thomaz
package funcao

func E1_AreaRetanguloCP(a int, b int) int {
	return a * b
}

func E2_RetangulosIguais(r1 Retangulo, r2 Retangulo) bool {
	return E1_AreaRetanguloCP(r1.Altura, r1.Base) == E1_AreaRetanguloCP(r2.Altura, r2.Base)
}

func E3_RetangulosIguais(retangulos []Retangulo) bool {
	actualRet := retangulos[0]
	for _, ret := range retangulos {
		if (E1_AreaRetanguloCP(actualRet.Altura, actualRet.Base) != E1_AreaRetanguloCP(ret.Altura, ret.Base)) {
			return false
		}
	}

	return true
}
