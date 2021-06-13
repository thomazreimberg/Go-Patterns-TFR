package tipo

import (
	"fmt"
	//"golang.org/x/tools/go/analysis/passes/nilfunc"
)

type Pessoa struct {
	nome string
}

type Familia struct {
	pai Pessoa
	mae Pessoa
}

func E1_DeclaracaoC() {
	//
	var pessoa Pessoa
	pessoa.nome = "Bruno"
	fmt.Println(pessoa.nome)

	var x1 string = pessoa.nome
	fmt.Println(x1)

	//
	pessoa2 := Pessoa{nome: "Bruno"}
	fmt.Println(pessoa2.nome)

	//
	var familia Familia

	familia.pai.nome = "Junior"
	familia.mae.nome = "Maria"

	var x2 string = familia.pai.nome
	var x3 string = familia.mae.nome
	fmt.Println(x2)
	fmt.Println(x3)

	//
	familia1 := Familia{pai: Pessoa{nome: "Junior"}, mae: Pessoa{nome: "Maria"}}
	fmt.Println(familia1.pai.nome)

	//
	var pessoa3 Pessoa = pessoa
	pessoa3.nome = "Ingrid"
	fmt.Println(pessoa3.nome)

}

func E2_ValoresNull() {
	var pessoa Pessoa // Inicializa com valores padrão

	fmt.Println(pessoa.nome)
	pessoa.nome = "Aline"

	fmt.Println(pessoa.nome)
}

func E3_TiposAnonimos() {
	pessoa := struct {
		nome string
	}{
		nome: "João",
	}

	fmt.Println(pessoa.nome)

	// familia := struct{
	// 	pai struct
	// 	mae struct
	// }{
	// 	pai: struct{nome string}{nome: "a"},
	// 	mae: struct{nome string}{nome: "a"},
	// }
}

func E4_AlteracaoC() {
	p1 := Pessoa{nome: "Bruno"}
	fmt.Println(p1.nome)
	p1.nome = "Bruno O."
	fmt.Println(p1.nome)

	x1 := "Oliveira"
	p1.nome = x1
	fmt.Println(p1.nome)
	p1.nome = "Bruno" + x1
	fmt.Println(p1.nome)

	var x2 = p1.nome
	fmt.Println(x2)
}

func E5_Comparacao() {
	p1 := Pessoa{nome: "Bruno"}
	p2 := Pessoa{nome: "Bruno"}

	var x1 = p1 == p2
	fmt.Println(x1)
	//
	p3 := Pessoa{nome: "Bruno"}
	p4 := p3
	p3.nome = "Ingrid"

	var x3 = p3 == p4
	fmt.Println(x3)
}
