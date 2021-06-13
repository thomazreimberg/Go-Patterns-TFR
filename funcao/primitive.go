package funcao

import (
	"fmt"
	"reflect"
)

func E1_OperadoresMatematicos() {

	var x1 int = 10 + 10
	var x2 int = 10 - 10
	var x3 int = 10 * 10
	var x4 int = 10 / 10
	var x5 int = 10 % 10
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(x4)
	fmt.Println(x5)

	//var x6 int = x1++
	//var x7 int = ++x1
	x3 += 10
	fmt.Println(x3)
	x3 -= 05
	fmt.Println(x3)
	x3 *= 02
	fmt.Println(x3)
	x3 /= 02
	fmt.Println(x3)

	var x8 int = x1 + x2
	var x9 int = x1 + x2*x3
	var x10 int = (x1 + x2) * x3
	fmt.Println(x8)
	fmt.Println(x9)
	fmt.Println(x10)

}

func E2_OperadoresRelacionais() {

	var x1 bool = 10 > 5
	var x2 bool = 10 < 5
	var x3 bool = 10 >= 10
	var x4 bool = 10 <= 10
	var x5 bool = 10 == 10
	var x6 bool = 10 != 10
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(x4)
	fmt.Println(x5)
	fmt.Println(x6)

	var x7 int = 10
	var x8 int = 5

	var x9 bool = x7 > 10
	var x10 bool = x7 > x8
	fmt.Println(x9)
	fmt.Println(x10)

}

func E3_OperadoresLogicos() {

	var x1 bool = true && true
	var x2 bool = true && false
	var x3 bool = true || true
	var x4 bool = true || false
	var x5 bool = true || true
	var x6 bool = !false
	var x7 bool = !true
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(x4)
	fmt.Println(x5)
	fmt.Println(x6)
	fmt.Println(x7)

	var x8 bool = 10 > 5 && 20 < 10
	var x9 bool = 10 > 5 || 20 < 10
	fmt.Println(x8)
	fmt.Println(x9)

	var x10 int = 10
	var x11 int = 5
	fmt.Println(x10)
	fmt.Println(x11)

	var x12 bool = x10 > 5 && x11 < x10
	var x13 bool = x10 > 5 || x11 < x10
	fmt.Println(x12)
	fmt.Println(x13)

}

func E4_OutrosOperadores() {

	var x1 string = "Olá! "
	var x2 string = "Tudo bem?"
	var x3 string = x1 + x2
	fmt.Println(x3)

	x5 := []int{1, 2, 3, 4, 5}
	var x6 int = x5[2]
	fmt.Println(x6)

	// var x7 string = x6 > 1 ? "A" : "B"
	x10 := reflect.TypeOf(10 + 20)
	fmt.Println(x10)

	var x16 int = 10       //00001010		(10)
	var x17 int = x16 << 2 // 00010100     (20)
	fmt.Println(x16)
	fmt.Println(x17)

	var x18 int = 10       // 00001010     (10)
	var x19 int = x18 >> 2 // 00000010     (2)
	fmt.Println(x18)
	fmt.Println(x19)

	var x20 int = 10        // 00001010     (10)
	var x21 int = 6         // 00000110     (6)
	var x22 int = x20 | x21 // 00000010     (2)
	var x23 int = x20 & x21 // 00001110     (14)
	fmt.Println(x20)
	fmt.Println(x21)
	fmt.Println(x22)
	fmt.Println(x23)

}
