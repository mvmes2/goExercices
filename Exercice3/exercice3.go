package Exercice3

import "fmt"

/*
Um professor de programação está corrigindo as avaliações de seus alunos na disciplina de
Programação I para fornecer-lhes o feedback correspondente. Um dos pontos do exame é
declarar diferentes variáveis.
Ajude o professor com as seguintes questões:
1. Verifique quais dessas variáveis declaradas pelo aluno estão corretas.
2. Corrigir as incorrectas.

var 1nome string
var sobrenome string
var int idade
1sobrenome := 6
var licenca_para_dirigir = true
var estatura da pessoa int
quantidadeDeFilhos := 2

*/

func FixText() {
	fmt.Println("var 1nome string / incorreto => var nome string")
	fmt.Println("var sobrenome string / correto")
	fmt.Println("var int idade / incorreto => var idade int")
	fmt.Println("1sobrenome := 6 / incorreto => sobrenome := 6")
	fmt.Println("var licenca_para_dirigir = true /correto")
	fmt.Println("var estatura da pessoa int / incorreto => var estaturaDaPessoa int")
	fmt.Println("quantidadeDeFilhos := 2 / correto")
}
