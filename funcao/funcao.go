package funcao

import (
	"fmt"
	"time"
)

func Run() {
	fmt.Println("# Funcão #")

	fmt.Println("Primitive")
	E1_OperadoresMatematicos()
	E2_OperadoresRelacionais()
	E3_OperadoresLogicos()
	E4_OutrosOperadores()

	fmt.Println("HighLevel")
	E1_FuncoesParaNumeros()
	E2_FuncoesParaTextos()
	E3_FuncoesParaData()
	E4_FuncoesParaTiposRecursivos()
	E5_FuncoesParaConversao()

	fmt.Println("Procedure")
	fmt.Println(E1_AreaRetangulo(10,5))
	fmt.Println(E2_Media(4, 6, 8))
	fmt.Println(E3_Hipontenusa(3, 4))
	fmt.Println(E4_PrimeiroNome("Thomaz Henrique"))
	fmt.Println(E5_PrimeiroDiaProximoMes(time.Now()))
	fmt.Println(E6_AreaRetangulo(Retangulo { Altura: 10, Base:  5 }))
	fmt.Println(E7_Bhaskara(4, 2, -6))

	fmt.Println("CompoundProcedure")
	fmt.Println(E1_AreaRetanguloCP(2, 10))
	fmt.Println(E2_RetangulosIguais(Retangulo{ Altura: 2, Base: 10 }, Retangulo{ Altura: 4, Base: 5 }))
	fmt.Println(E3_RetangulosIguais([]Retangulo {
		{Altura: 2, Base: 10},
		{Altura: 4, Base: 5},
		{Altura: 1, Base: 20},
	}))

	fmt.Println("Packages")//Parou em package
	E1_CalcularCompra(Pedido{Valor: 1000, AnosGarantia: 2, DistanciaEntregaKm: 10, Parcelas: 10}, 20.0)

	fmt.Println("Recursive")
	fmt.Println(E1_Fatorial_Linear(5))
	fmt.Println(E2_Fatorial_Iterativo(5))

	fmt.Println("HighOrderFunction")
	E1_FuncoesAnonimas()
	filtro := func(x int) bool {
		return x > 4
	}
	fmt.Println(E2_ComoParametro_Filter([]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, filtro))
	mapper := func(x int) int {
		return x * 4
	}
	fmt.Println(E3_ComoParametro_Map([]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, mapper))
	fmt.Println("E4_ComoParametro_Correio")
	E4_ComoParametro_Correio("04895020", func(x Response) { fmt.Println(x) })
}