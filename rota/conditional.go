//Thomaz
package rota

func E1_IF_Aprovado(media float64) string {
	situacao := "Reprovado"

	if media >= 5 {
		situacao = "Aprovado"
	} else if media >= 3.5 {
		situacao = "Reprovado"
	} else {
		situacao = "Reprovado"
	}

	return situacao
}

func E2_IF_Aprovado(media float64, faltas int) string {
	situacao := "Reprovado"

	if media < 5 && faltas >= 25 {
		situacao = "Reprovado por nota e falta"
	} else if faltas > 25 {
		situacao = "Reprovado por falta"
	} else if media < 5 {
		situacao = "Reprovado por nota"
	} else {
		situacao = "Aprovado"
	}

	return situacao
}

func E3_Switch_Avaliacao(nota string) string {
	situacao := ""

	switch nota {
	case "A":
		situacao = "Ótimo"
	case "B":
		situacao = "Bom"
	default:
		situacao = "Ruim"
	}

	return situacao
}

func E4_Switch_Avaliacao(nota string) string {
	situacao := ""

	switch nota {
	case "A":
	case "B":
		situacao = "Ótimo"
	case "C":
		situacao = "Bom"
	case "D":
	case "E":
		situacao = "Ruim"
	default:
		situacao = "Nota inválida"
	}

	return situacao
}
