package Exercice7

import "fmt"

/*
Exercício 3 - A que mês corresponde?

Faça uma aplicação que contenha uma variável com o número do mês.
1. Dependendo do número, imprima o mês correspondente em texto.
2. Ocorre a você que isso pode ser resolvido de maneiras diferentes? Qual você
escolheria e porquê?
Ex: 7 de julho.

R: Escolheria fazer um map, pois map é mais performático que arrays, e de um modo simples podemos verificar
a existencia ou nao de uma chave sem a necessidade de um looping para percorre-lo usando o comma idiom
*/

func NumeroDoMes(mes uint) {
	var mesMap = map[uint]string{
		1:  "Janeiro",
		2:  "Fevereiro",
		3:  "Março",
		4:  "Abril",
		5:  "Maio",
		6:  "Junho",
		7:  "Julho",
		8:  "Agosto",
		9:  "Setembro",
		10: "Outubro",
		11: "Novembro",
		12: "Dezembro",
	}
	value, ok := mesMap[mes]
	if ok {
		fmt.Println(mes, "de:", value)
	} else {
		fmt.Println("Não existe nenhum mês com este número!")
	}
}
