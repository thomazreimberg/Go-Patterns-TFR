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
	E1_AreaRetanguloCP()
	E2_RetangulosIguais()
	E3_RetangulosIguais()

	fmt.Println("Packages")
	E1_CalcularCompra()

	fmt.Println("HighOrderFunction")
	E1_FuncoesAnonimas()
	E2_ComoParametro_Filter()
	E3_ComoParametro_Map()
	E4_ComoParametro_Correio()
}
