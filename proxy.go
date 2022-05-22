package main

import (
	"exercices/Exercice1"
	"exercices/Exercice2"
	"exercices/Exercice3"
	"exercices/Exercice4"
	"exercices/Exercice5"
	"exercices/Exercice6"
	"exercices/Exercice7"
	"exercices/Exercice8"
	"fmt"
)

func ConsoleView() {
	fmt.Println("///////////Exercício 1://////////////")
	Exercice1.PrintName()
	fmt.Println("///////////Exercício 2://////////////")
	Exercice2.Metereologia()
	fmt.Println("///////////Exercício 3://////////////")
	Exercice3.FixText()
	fmt.Println("///////////Exercício 4://////////////")
	Exercice4.CodeReview()
	fmt.Println("///////////Exercício 5://////////////")
	Exercice5.SoletraPalavra("tomiouka")
	fmt.Println("///////////Exercício 6://////////////")
	Exercice6.ClientesEmprestimo()
	fmt.Println("///////////Exercício 7://////////////")
	Exercice7.NumeroDoMes(10)
	fmt.Println("///////////Exercício 8://////////////")
	Exercice8.QuantosAnosTem()

}
