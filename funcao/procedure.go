//Thomaz
package funcao

import (
	"math"
	"strings"
	"time"
)

type Retangulo struct {
	Altura int
	Base int
}

type Bhaskara struct {
	X1 float64
	X2 float64
}

func E1_AreaRetangulo(a int, b int) int {
	return a * b
}

func E2_Media(a float64, b float64, c float64) float64 {
	return (a + b + c) / 3
}

func E3_Hipontenusa(co float64, ca float64) float64 {
	return math.Sqrt(math.Pow(ca, 2) + math.Pow(co, 2))
}

func E4_PrimeiroNome(nome string) string {
	p := nome[0 : strings.Index(nome, " ")]
	return p
}

func E5_PrimeiroDiaProximoMes(dia time.Time) time.Time {
	return dia.AddDate(0, 1, 0).AddDate(0, 0, -dia.Day() + 1)
}

func E6_AreaRetangulo(ret Retangulo) int {
	return ret.Altura * ret.Base
}

func E7_Bhaskara(a int, b int, c int) Bhaskara {
	//delta := math.Pow(b, 2) - 4 * a * c
	delta := (math.Pow(float64(b), 2) - float64(4 * a * c))
	x := Bhaskara {
		X1: (-float64(b) + delta) / float64((2 * a)),
		X2: (-float64(b) - delta) / float64((2 * a)),
	}

	return x
}