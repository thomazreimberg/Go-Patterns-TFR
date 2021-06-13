package tipo

import "fmt"

func Run() {
	fmt.Println("# Tipo #")

	fmt.Println("Primitive")
	E1_Declaracao()
	E2_Alteracao()
	E3_Conversao()
	E4_Coercion()
	E5_Inferencia()

	fmt.Println("Compound")
	E1_DeclaracaoC()
	E2_ValoresNull()
	E3_TiposAnonimos()
	E4_AlteracaoC()
	E5_Comparacao()

	fmt.Println("Recursive")
	E1_Vetor()
	E2_Slice()
	E3_Map()

}
