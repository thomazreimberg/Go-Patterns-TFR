package rota

import (
	"fmt"
	"time"
)

func Run() {
	fmt.Println("# Rota #")

	fmt.Println("Simple")
	E1_Media(7, 7, 10)

	fmt.Println("Conditional")
	fmt.Println(E1_IF_Aprovado(3.3))
	fmt.Println(E2_IF_Aprovado(6, 29))
	fmt.Println(E3_Switch_Avaliacao("A"))
	fmt.Println(E4_Switch_Avaliacao("E"))

	fmt.Println("CompoundConditional")
	E1_ParOuImpar(10, 3, "Par", "Ímpar")
	E2_RecomendacaoMusica("Brasil", "Rock")
	E3_OqueFazer(time.Date(2021, 6, 13, 16, 0, 0, 0, time.UTC))

	fmt.Println("Looping")
	E1_FOR_SomarNumeros()
	E2_FOREACH_FiltrarPares()
	E3_CONTINUE_GerarSequenciaPar()
	E4_BREAK_TodosPares()
	E5_WHILE_ProximaRaizInteira()
	E6_DOWHILE_Somar()

	
}
