package Exercice8

import "fmt"

/*
Exercício 4 - Quantos anos tem...
Um funcionário de uma empresa deseja saber o nome e a idade de um de seus funcionários.
De acordo com o mapa abaixo, ajude a imprimir a idade de Benjamin.

var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

Por outro lado, você também precisa:
- Saiba quantos de seus funcionários têm mais de 21 anos.
- Adicione um novo funcionário à lista, chamado Federico, que tem 25 anos.
- Excluir Pedro do mapa.
*/

func QuantosAnosTem() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println("A idade de Benjamin é:", employees["Benjamin"])

	var countMaiorDe21 uint = 0

	for _, item := range employees {
		if item > 21 {
			countMaiorDe21++
		}
	}

	var funcionarios []string

	fmt.Println(countMaiorDe21, "funcionários tem mais de 21 anos no quadro atual de funcionarios")
	fmt.Println("Pedro foi demitido do quadro de funcionarios, e Frederico foi Admitido")
	employees["Frederico"] = 25
	delete(employees, "Pedro")
	for index := range employees {
		funcionarios = append(funcionarios, index)
	}
	fmt.Println("Quadro de funcionarios sem o Pedro e com o Frederico:")
	for _, item := range funcionarios {
		fmt.Println(item)
	}

}
