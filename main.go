package main

import (
	"fmt"
	"github.com/Go-Patterns-TFR/tipo"
	"github.com/Go-Patterns-TFR/funcao"
	"github.com/Go-Patterns-TFR/rota"
)

func main() {
	fmt.Println("## Go ##")
	
	tipo.Run()
	funcao.Run()
	rota.Run()
}
