package main

import (

	"modulo/operacoesAritimeticas"
)

func main() {
	a:=1
	b:=2
	a, b = operacoesAritimeticas.Inverte(a, b)

	println(a)
	println(b)

	
}