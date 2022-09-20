package main

import (
	"fmt"
	"modulo/strutsHerancaFuncoes"
)

func main() {
	estudantes := strutsHerancaFuncoes.CriaEstudante()

	strutsHerancaFuncoes.Ordena(estudantes)

	// aprovados, _, _ := strutsHerancaFuncoes.SituacaoAlunos(estudantes)

	// for _, aprovado := range aprovados {
	// 	fmt.Println(aprovado)
	// }

	as := strutsHerancaFuncoes.EliminaReprovados(estudantes)

	for _, a := range as {
		fmt.Println(a)
	}
}
