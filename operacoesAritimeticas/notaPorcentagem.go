package operacoesAritimeticas

import "fmt"

type regraNota struct {
    valor int
    conteito  string
}

func NotaConceito(nota int) {
	regras:= [4]regraNota{{valor: 90, conteito: "A"}, {valor: 80, conteito: "B"}, {valor: 70, conteito: "C"}, {valor: 60, conteito: "D"},}
	index := 0
	for contador := 90; contador >= 60; contador-=10{
		if(nota >= contador) {
			fmt.Println(regras[index].conteito)
			return
		}
		index ++
	}
	fmt.Printf("Nota invalida")
}