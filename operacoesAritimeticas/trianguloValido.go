package operacoesAritimeticas

func TrianguloValido(a, b, c int) (string) {
	if a < 0 || b < 0 || c < 0 {
		return "Um dos angulos é negativo por tanto inálido"
	} 
	if a + b + c  != 180 {
		return "Triângulo inválido"
	}
	return "triangulo válido"
}