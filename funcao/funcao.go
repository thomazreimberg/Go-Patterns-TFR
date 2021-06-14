// Felipe
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
	fmt.Println(E1_AreaRetangulo(10,10))
	fmt.Println(E2_Media(10, 10, 10))
	fmt.Println(E3_Hipontenusa(4, 5))
	fmt.Println(E4_PrimeiroNome("Thomaz Henrique"))
	fmt.Println(E5_PrimeiroDiaProximoMes(time.Now()))
	fmt.Println(E6_AreaRetangulo(Retangulo { Altura: 10, Base:  10 }))
	fmt.Println(E7_Bhaskara(10, 10, 10))

	fmt.Println("CompoundProcedure")
	fmt.Println(E1_AreaRetanguloCP(2, 10))
	fmt.Println(E2_RetangulosIguais(Retangulo{ Altura: 2, Base: 10 }, Retangulo{ Altura: 4, Base: 5 }))
	fmt.Println(E3_RetangulosIguais([]Retangulo {
		{Altura: 10, Base: 20},
		{Altura: 20, Base: 10},
		{Altura: 200, Base: 1},
	}))

	fmt.Println("Packages")
	E1_CalcularCompra(Pedido{Valor: 1000, AnosGarantia: 2, DistanciaEntregaKm: 10, Parcelas: 10}, 20.0)

	fmt.Println("HighOrderFunction")
	E1_FuncoesAnonimas()
	E2_ComoParametro_Filter()
	E3_ComoParametro_Map()
	E4_ComoParametro_Correio()
}
