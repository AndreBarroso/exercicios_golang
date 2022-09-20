package strutsHerancaFuncoes

import "sort"

type Pessoa struct {
	nome   string
	altura float32
	idade  int
}

type Faculdade struct {
	nomeFaculdade string
	estado        string
}

type Estudante struct {
	Pessoa
	Faculdade
	nota float64
}

func CriaEstudante() []Estudante {
	arrayPessoas := [15]Pessoa{
		{nome: "Andre", altura: 1.70, idade: 33},
		{nome: "Mayara", altura: 1.55, idade: 35},
		{nome: "Gabriel", altura: 1.68, idade: 27},
		{nome: "Jairo", altura: 1.69, idade: 65},
		{nome: "Vilma", altura: 1.56, idade: 63},
		{nome: "Jaci", altura: 1.54, idade: 67},
		{nome: "Felipe", altura: 1.86, idade: 37},
		{nome: "Milton", altura: 1.85, idade: 39},
		{nome: "Jota", altura: 1.77, idade: 70},
		{nome: "Saozinha", altura: 1.53, idade: 69},
		{nome: "Thomas", altura: 1.80, idade: 38},
		{nome: "Everton", altura: 1.81, idade: 45},
		{nome: "Rogério", altura: 1.75, idade: 40},
		{nome: "Viviane", altura: 1.79, idade: 42},
		{nome: "Enio", altura: 1.74, idade: 33},
	}

	faculdades := [15]Faculdade{
		{nomeFaculdade: "UFMG", estado: "MG"},
		{nomeFaculdade: "UFRJ", estado: "RJ"},
		{nomeFaculdade: "USP", estado: "SP"},
		{nomeFaculdade: "UFVJM", estado: "MG"},
		{nomeFaculdade: "UFRS", estado: "RS"},
		{nomeFaculdade: "UNICAMP", estado: "SP"},
		{nomeFaculdade: "UEMG", estado: "MG"},
		{nomeFaculdade: "UFPA", estado: "PA"},
		{nomeFaculdade: "UNIRIO", estado: "RJ"},
		{nomeFaculdade: "UFV", estado: "MG"},
		{nomeFaculdade: "UFBA", estado: "BA"},
		{nomeFaculdade: "UFSC", estado: "SC"},
		{nomeFaculdade: "UFMT", estado: "MT"},
		{nomeFaculdade: "UFF", estado: "RJ"},
		{nomeFaculdade: "UFPR", estado: "PR"},
	}

	estudantes := []Estudante{}
	for i := 0; i < 14; i++ {

		nota := float64(i)/2 + 1
		pessoa := Pessoa{nome: arrayPessoas[i].nome, altura: arrayPessoas[i].altura, idade: arrayPessoas[i].idade}
		faculdade := Faculdade{nomeFaculdade: faculdades[i].nomeFaculdade, estado: faculdades[i].estado}
		estudante := Estudante{pessoa, faculdade, nota}

		estudantes = append(estudantes, estudante)
	}
	return estudantes
}

func Ordena (array []Estudante) {
	sort.SliceStable(array, func(i, j int) bool {
		return array[i].nome < array[j].nome
	})
}

func SituacaoAlunos (array []Estudante) ([]Estudante, []Estudante, []Estudante){
	aprovados := []Estudante{}
	recuperacao := []Estudante{}
	reprovados := []Estudante{}

	for i:= 0; i < len(array); i++ {
		if array[i].nota >= 7 {
			aprovados = append(aprovados, array[i])
		} else if array[i].nota < 7 && array[i].nota >=6 {  
			recuperacao = append(recuperacao, array[i])
		} else if array[i].nota < 6 {  
			reprovados = append(reprovados, array[i])
		} 
	}

	return aprovados, recuperacao, reprovados
}


func EliminaReprovados (array []Estudante) []Estudante{
	auxArray := []Estudante{}

	for i:= 0; i < len(array); i++ {
		 if array[i].nota > 6 {  
			auxArray = append(auxArray, array[i])
		} 
	}
	
	return auxArray
}


