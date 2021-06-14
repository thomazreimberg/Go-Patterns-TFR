//Thomaz
package funcao

import "fmt"

func E1_FuncoesAnonimas() {
	soma := func(a int, b int) int {
		return a + b
	}

	fmt.Println(soma(10, 20))

	func() {
		fmt.Println("Função anônima #1")
	}()

	func(msg string) {
		fmt.Println(msg)
	}("Função anônima #2")
}

func E2_ComoParametro_Filter(lista []int, filtro func(x int) bool) []int {
	var newList []int

	for item := range lista {
		if (filtro(item)) {
			newList = append(newList, item)
		}
	}

	return newList
}

func E3_ComoParametro_Map(lista []int, mapper func(x int) int) []int {
	var newList []int

	for item := range lista {
		newList = append(newList, mapper(item))
	}
	
	return newList
}

func E4_ComoParametro_Correio() {

}
