package Exercice2

import "fmt"

/*
Uma empresa de meteorologia quer ter um sistema onde possa ter a temperatura, umidade e
pressão atmosférica de diferentes lugares.
1. Declare 3 variáveis especificando o tipo de dado delas, como valor elas devem
possuir a temperatura, umidade e pressão atmosférica de onde você se encontra.
2. Imprima os valores no console.
3. Quais tipos de dados serão atribuídos a essas variáveis?
r: number, string, string.
*/

var (
	temperatura int8   = 14
	umidade     string = "81%"
	pa          string = "1016hPa"
)

func Metereologia() {
	fmt.Println("A temperatura em São Paulo é de:", temperatura, "Umidade:", umidade, "Pressão Atmosférica:", pa)
}
