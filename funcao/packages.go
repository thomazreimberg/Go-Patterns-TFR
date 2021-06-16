// Felipe
package funcao

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

type Pedido struct {
	Valor              float64
	AnosGarantia       int
	DistanciaEntregaKm int
	Parcelas           int
}

type Nota struct {
	ValorFinal float64
	NotaFiscal string
}

func E1_CalcularCompra(pedido Pedido, cupom float64) Nota {
	nota := Nota{
		ValorFinal: E5_ValorJuros(pedido.Valor-
			E2_ValorDesconto(pedido.Valor, cupom)+
			E3_ValorGarantia(pedido.Valor, float64(pedido.AnosGarantia))+
			E4_ValorFrete(pedido.DistanciaEntregaKm), pedido.AnosGarantia),
		NotaFiscal: time.Now().String(),
	}
	fmt.Println(nota)
	return nota
}

func E2_ValorDesconto(valorCompra, cupom float64) float64 {
	return valorCompra * cupom / 100.0
}

func E3_ValorGarantia(valorCompra, anos float64) float64 {
	return valorCompra * 0.05 * anos
}

func E4_ValorFrete(distancia int) float64 {
	return float64(distancia / 10 * 3)
}

func E5_ValorJuros(valor float64, parcelas int) float64 {
	return valor * math.Pow(1.05, float64(parcelas))
}

type Aluno struct {
	nome     string
	curso    string
	semestre int
}

func (i Aluno) Mensagem() {
	fmt.Println(i.nome + " está cursando " + i.curso + " no " + strconv.Itoa(i.semestre) + "º semestre")
}

func (i *Aluno) AtualizarSemestre(semestre int) {
	i.semestre = semestre
}
