package operacoesAritimeticas

func CalculaLucro(valorVenda, custo float64) (float64) {
	custoFinal := custo + custo*.2
	return valorVenda - custoFinal
} 