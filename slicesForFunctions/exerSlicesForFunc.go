package slicesforfunctions

import (
	"fmt"
	"strconv"
)

var slice1 = []int{5, 9, 3, 19, 70, 8, 100, 2, 35, 27}

func ImprimeValoresSlice() {

	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}
}

func SomaValoresSlice() int {
	var soma = 0
	for i := 0; i < len(slice1); i++ {
		soma += slice1[i]
	}
	return soma
}

func MediaAritmetica() int {
	var soma = 0
	for i := 0; i < len(slice1); i++ {
		soma += slice1[i]
	}
	return soma / len(slice1)
}

func MaiorValor() int{
	var maior = 0
	for i := 0; i < len(slice1); i++ {
		if(slice1[i] > maior) {
			maior = slice1[i]
		}
	}
	return maior
}

func NumeroImpares() string{
	var impares = 0
	for i := 0; i < len(slice1); i++ {
		if(slice1[i]%2 !=0 ) {
			impares += 1
		}
	}
	
	if(impares < 1) {
		return "nenhum valor ímpar encontrado"
	}
	return strconv.Itoa(impares)
}


func MenorValor() int{
	var menor = slice1[0]
	for i := 0; i < len(slice1); i++ {
		if(slice1[i] < menor) {
			menor = slice1[i]
		}
	}
	return menor
}

func CriaSlice() []int{
	var slice2 = []int{}
	for i := 1; i <= 25; i++ {
		slice2 = append(slice2, i)
	}
	return slice2
}

func DivideIndices(slice []int) {
	for _, numero := range slice {
		var divisao float64 = float64(numero)/2.0
		println(divisao)
	}
}