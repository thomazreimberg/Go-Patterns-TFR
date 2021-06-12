package rota

import "fmt"

func Run() {
	fmt.Println("# Rota #")

	fmt.Println("Simple")
	E1_Media()

	fmt.Println("Conditional")
	E1_IF_Aprovado()
	E2_IF_Aprovado()
	E3_Switch_Avaliacao()
	E4_Switch_Avaliacao()

	fmt.Println("CompoundConditional")
	E1_ParOuImpar()
	E2_RecomendacaoMusica()
	E3_OqueFazer()

	fmt.Println("Looping")
	E1_FOR_SomarNumeros()
	E2_FOREACH_FiltrarPares()
	E3_CONTINUE_GerarSequenciaPar()
	E4_BREAK_TodosPares()
	E5_WHILE_ProximaRaizInteira()
	E6_DOWHILE_Somar()

	
}
