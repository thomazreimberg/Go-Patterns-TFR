//Thomaz
package funcao

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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

type Response struct {
	Cep string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf string `json:"uf"`
	Ibge string `json:"ibge"`
	Gia string `json:"gia"`
	Ddd string `json:"ddd"`
	Siafi string `json:"siafi"`
}

func E4_ComoParametro_Correio(cep string, callback func(x Response)) {
	client := &http.Client{} //Criei um client para ter controle sobre cabaçalhos HTTP, política de redirecionamento e outras configs
	//se não não querer fazer esse controle http.Get (" http://example.com/ ") já resolveria o problema
	req, err := http.NewRequest("GET", "https://viacep.com.br/ws/"+ strings.TrimSpace(cep)+"/json/", nil) //Monta a requisição

	if (err != nil) {
		fmt.Println(err.Error())
	}
	//Adicionando cabeçalhos
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	//Enviando requisição
	resp, err := client.Do(req)

	if (err != nil) {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if (err != nil) {
		fmt.Println(err.Error())
	}
	
	var response Response
	json.Unmarshal(bodyBytes, &response) //& antes de uma variável significa algo como "use o endereço dessa variável para a ação x". Fazendo isso, afetará diretamente a variável
	
	defer callback(response)
}
