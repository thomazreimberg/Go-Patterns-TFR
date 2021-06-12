package tipo

import (
	"fmt"
	"strconv"
)

func E1_Declaracao() {
	//
	var x0 bool = true
	fmt.Println(x0)

	var x1 int = 10
	fmt.Println(x1)

	var x2 int8 = 10
	fmt.Println(x2)
	var x3 int16 = 10
	fmt.Println(x3)
	var x4 int32 = 10
	fmt.Println(x4)
	var x5 int64 = 10
	fmt.Println(x5)

	var x6 float32 = 20.19
	fmt.Println(x6)
	var x7 float64 = 20.19
	fmt.Println(x7)

	var x8 string = "Linguagem Go"
	fmt.Println(x8)

	var x9, x10 string = "Teste2", "Teste1"
	fmt.Println(x9)
	fmt.Println(x10)
	var x11, x12 int = 10, 11
	fmt.Println(x11, x12)

	var x13 int = x11
	fmt.Println(x13)

	var x14 = x11 + x12
	fmt.Println(x14)

	const x15 int = 10
	fmt.Println(x15)

	var x16 int = int(^uint(0) >> 1) //MaxValue
	fmt.Println(x16)
	var x17 int = -x16 - 1
	fmt.Println(x17)

	x18 := "Teste string"
	fmt.Println(x18)
	x19 := 10
	fmt.Println(x19)
	x20 := 10.20
	fmt.Println(x20)
}

func E2_Alteracao() {
	var x1 int = 10
	x1 = 15
	fmt.Println(x1)

	var x2 string = "oie"
	x2 = "olá"
	fmt.Println(x2)

	var x3 int = 10
	x3 = x1 + 5
	fmt.Println(x3)
	//x3 = x1 = 20 => error
}

func E3_Conversao() {
	var x string = "10.5"
	var y int = 10

	x0, _ := strconv.ParseBool(x)
	fmt.Println(x0)

	x1, _ := strconv.Atoi(x)
	fmt.Println(x1)
	x2, _ := strconv.ParseInt(x, 8, 8)
	fmt.Println(x2)
	x3, _ := strconv.ParseInt(x, 16, 16)
	fmt.Println(x3)
	x4, _ := strconv.ParseInt(x, 32, 32)
	fmt.Println(x4)
	x5, _ := strconv.ParseInt(x, 64, 64)
	fmt.Println(x5)

	x6, _ := strconv.ParseFloat(x, 32)
	fmt.Println(x6)
	x7, _ := strconv.ParseFloat(x, 64)
	fmt.Println(x7)

	x8 := string(x)
	fmt.Println(x8)

	x9 := strconv.Itoa(y)
	fmt.Println(x9)
}

func E4_Coercion() {
	//var x = "10" + 10
	fmt.Println("Não existe Coercion na linguagem Go.")
}

func E5_Inferencia() {

}
