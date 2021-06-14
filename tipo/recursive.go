// Felipe
package tipo

import "fmt"

func E1_Vetor() {
	//
	var vetor [3]int
	vetor[0] = 10
	vetor[1] = 20
	vetor[2] = 30

	//
	var a int = vetor[0]
	var b int = vetor[1]
	var c int = vetor[2]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//
	x := len(vetor)
	fmt.Println(x)
}

func E2_Slice() {
	var slice []int
	slice = append(slice, 10, 20, 30)

	//
	var a int = slice[0]
	var b int = slice[1]
	var c int = slice[2]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//
	x := len(slice)
	fmt.Println(x)

}

func E3_Map() {
	var dic map[int]string = make(map[int]string)
	dic[1] = "C#"
	dic[2] = "Javascript"
	dic[3] = "Python"

	//
	var a string = dic[1]
	var b string = dic[2]
	var c string = dic[3]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//
	var x int = len(dic)
	fmt.Println(x)

	//
	dic[3] = "Java"
	fmt.Println(dic[3])
	//
	delete(dic, 3)
	fmt.Println(dic[3])
}
