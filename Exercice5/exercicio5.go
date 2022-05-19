package Exercice5

import (
	"fmt"
	"strings"
)

/*
Exercício 1 - Letras de uma palavra

A Real Academia Brasileira quer saber quantas letras tem uma palavra e então ter cada uma
das letras separadamente para soletrá-la. Para isso terão que:
1. Crie uma aplicação que tenha uma variável com a palavra e imprima o número de
letras que ela contém.
2. Em seguida, imprimi cada uma das letras.
*/

func SoletraPalavra(palavra string) {
	teste := strings.Split(palavra, "")
	fmt.Println(palavra, "tem", len(teste), "letras")
	for _, i := range teste {
		fmt.Println(i)
	}
}
