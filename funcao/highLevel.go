// Thomaz
package funcao

import (
	"fmt"
	"math"
	"strings"
	"time"
	"strconv"
	"github.com/ahmetb/go-linq/v3"
)

func E1_FuncoesParaNumeros() {
	var x float64 = 10.1234
	var y float64 = 9

	var x1 float64 = math.Floor(x)
	fmt.Println(x1)
	var x2 float64 = math.Ceil(x)
	fmt.Println(x2)
	var x3 float64 = math.Trunc(x)
	fmt.Println(x3)
	var x4 float64 = math.Round(x)
	fmt.Println(x4)
	var x5 float64 = math.Abs(x)
	fmt.Println(x5)
	var x6 float64 = math.Sqrt(x)
	fmt.Println(x6)
	var x7 float64 = math.Pow(y, 2)
	fmt.Println(x7)
}

func E2_FuncoesParaTextos() {
	var msg string = "Olá! Tudo bem?"

	var x1 string = strings.ToUpper(msg)
	fmt.Println(x1)
	var x2 string = strings.ToLower(msg)
	fmt.Println(x2)
	var x3 string = strings.TrimSpace(msg)
	fmt.Println(x3)
	var x4 string = strings.ReplaceAll(msg, "Olá", "Oiee")
	fmt.Println(x4)
	var x5 string = msg[5:]
	fmt.Println(x5)
	var x6 int = strings.Index(msg, " ")
	fmt.Println(x6)
	var x7 int = strings.LastIndex(msg, " ")
	fmt.Println(x7)
	var x8 bool = strings.Contains(msg, "bem?")
	fmt.Println(x8)
	var x9 int = len(msg)
	fmt.Println(x9)

	var x10 []string = []string{"Olá,", "tudo", "bem?"}
	fmt.Println(x10)
	var x11 string = strings.Join(x10, "-")
	fmt.Println(x11)

}

func E3_FuncoesParaData() {
	var x1 time.Time = time.Now()
	fmt.Println(x1)

	x2, _ := time.Parse(time.RFC822, "01 Jan 01 10:00 UTC")
	var x3 time.Time = x2.AddDate(0, 0, 1) //Day
	fmt.Println(x3)
	var x4 time.Time = x2.AddDate(0, 1, 0) //Month
	fmt.Println(x4)
	var x5 time.Time = x2.AddDate(1, 0, 0) //Year
	fmt.Println(x5)

	var x6 int = x2.Day()
	fmt.Println(x6)
	var x7 int = int(x2.Month())
	fmt.Println(x7)
	var x8 int = x2.Year()
	fmt.Println(x8)

	var x9 time.Weekday = x2.Weekday()
	fmt.Println(x9)
}

type Carro struct {
	ano int
	dono, modelo string
}

func E4_FuncoesParaTiposRecursivos() {
	fmt.Println("A linguagem Go não dá suporte para funções LINQ. Tenha acesso a essas funções por meio do packages externas: https://github.com/ahmetb/go-linq")
	//Importa: go get github.com/ahmetb/go-linq/v3
	//Pega a versão mais recente dos arquivos baixados: go get gopkg.in/ahmetb/go-linq.v3
	var donos []string
	carros := []Carro{
		{ano: 2000, dono: "User 1", modelo: "Fiat" },
		{ano: 2002, dono: "User 2", modelo: "Fiat 2" },
		{ano: 2000, dono: "User 3", modelo: "Fiat" },
	}
	
	linq.From(carros).Where(func(c interface{}) bool{
		return c.(Carro).ano >= 2001
	}).Select(func(c interface{}) interface{} {
		return c.(Carro).dono
	}).ToSlice(&donos)
	
	fmt.Println(donos)
}

func E5_FuncoesParaConversao() {
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
