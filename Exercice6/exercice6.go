package Exercice6

import (
	"fmt"
)

/*
Exercício 2 - Empréstimo

Um banco deseja conceder empréstimos a seus clientes, mas nem todos podem acessá-los.
Para isso, possui certas regras para saber a qual cliente pode ser concedido. Apenas
concede empréstimos a clientes com mais de 22 anos, empregados e com mais de um ano
de atividade. Dentro dos empréstimos que concede, não cobra juros para quem tem um
salário superior a US$ 100.000.
É necessário fazer uma aplicação que possua essas variáveis e que imprima uma mensagem
de acordo com cada caso.
Dica: seu código deve ser capaz de imprimir pelo menos 3 mensagens diferentes.
*/

func ClientesEmprestimo() {
	clientes := []map[string]interface{}{
		{
			"Nome":       "João",
			"Idade":      21,
			"Atividade:": 1,
			"Salário":    100000,
		},
		{
			"Nome":      "José",
			"Idade":     25,
			"Atividade": 1,
			"Salário":   100000,
		},
		{
			"Nome":      "Carlos",
			"Idade":     27,
			"Atividade": 2,
			"Salário":   99000,
		},
		{
			"Nome":      "Antônio",
			"Idade":     35,
			"Atividade": 5,
			"Salário":   150000,
		},
	}

	for _, item := range clientes {
		nome := item["Nome"]
		fmt.Println("Cliente:", nome)
		idade := item["Idade"].(int)
		if idade <= 22 {
			fmt.Println("Não possui Empréstimo disponível (idade insuficiente)")
			continue
		}
		atividade := item["Atividade"].(int)
		if atividade <= 1 {
			fmt.Println("Não possui empréstimo disponível (atividade insuficiente)")
			continue
		}
		salario := item["Salário"].(int)
		if salario > 100000 {
			fmt.Println("Possui empréstimo disponível sem juros")
		} else {
			fmt.Println("Possui empréstimo disponível com juros")
		}
	}

}
