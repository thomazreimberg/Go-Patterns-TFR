package funcao

import (
	"fmt"
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
	E1_AreaRetangulo()
	E2_Media()
	E3_Hipontenusa()
	E4_PrimeiroNome()
	E5_PrimeiroDiaProximoMes()
	E6_AreaRetangulo()
	E7_Bhaskara()

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
