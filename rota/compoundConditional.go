// Felipe
package rota

import (
	"fmt"
	"time"
)

func E1_ParOuImpar(jogador1, jogador2 int, escolhaJ1, escolhaJ2 string) string {

	soma := jogador1 + jogador2
	var resultado string

	if soma%2 == 0 {
		resultado = "Par"
	} else {
		resultado = "Ímpar"
	}

	var situacao string

	if resultado == "Par" {

		if escolhaJ1 == "Par" {
			situacao = "Deu Par! Jogador 1 venceu!"
		} else {
			situacao = "Deu Par! Jogador 2 venceu!"
		}
	} else {
		if escolhaJ1 == "Ímpar" {
			situacao = "Deu Ímpar! Jogador 1 venceu!"
		} else {
			situacao = "Deu Ímpar! Jogador 2 venceu!"
		}
	}
	fmt.Println(situacao)
	return situacao
}

func E2_RecomendacaoMusica(pais, estilo string) string {
	var musica string
	switch pais {
	case "Brasil":
		switch estilo {
		case "Rock":
			musica = "Quase sem querer"
			break

		case "Pop":
			musica = "Cor de Marte"
			break
		}
		break

	case "EUA":
		musica = "Any EUA music"
		break
	}
	fmt.Println(musica)
	return musica
}

func E3_OqueFazer(dia time.Time) string{
	var oqueFazer string

	switch dia.Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":

			if (dia.Hour() >= 8 && dia.Hour() <= 17) {
				oqueFazer = "Trabalhar"
				fmt.Println(oqueFazer)
			} else {
				oqueFazer = "Dormir"
				fmt.Println(oqueFazer)
			}
		break
		
		case "Saturday", "Sunday":
			
			if dia.Hour() >= 8 && dia.Hour() <= 17 {
				oqueFazer = "It's party time!"
				fmt.Println(oqueFazer)
			} else {
				oqueFazer = "Dormir" 
				fmt.Println(oqueFazer)
			}
		break
	}
	return oqueFazer
}
