package operacoesAritimeticas

func VerificaPar(array [3]int) bool {
	for i := 0; i < 3; i++ {
		if array[i]%2 == 0 {
			return true
		}

	}
	return false
}
