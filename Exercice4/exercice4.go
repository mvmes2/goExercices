package Exercice4

import "fmt"

/*
Exercício 4 - Tipos de dados

Um estudante de programação tentou fazer declarações de variáveis com seus respectivos
tipos em Go mas teve várias dúvidas ao fazê-lo. A partir disso, ele nos deu seu código e
pediu a ajuda de um desenvolvedor experiente que pode:
● Revisar o código e realizar as devidas correções.

var sobrenome string = "Silva"
var idade int = "25"
boolean := "false";
var salario string = 4585.90
var nome string = "Fellipe"

*/

func CodeReview() {
	fmt.Println("var sobrenome string = \"Silva\" / correto")
	fmt.Println("var idade int = \"25\" / incorreto => var idade int = 25")
	fmt.Println("boolean := \"false\"; / incorreto => boolean := false")
	fmt.Println("var salario string = 4585.90 / incorreto => var salario float = 4585.90")
	fmt.Println("var nome string = \"Fellipe\" /correto")
}
