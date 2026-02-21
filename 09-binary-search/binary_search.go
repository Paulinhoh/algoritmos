package binarysearch

func binarySearch(elementos []int, alvo int) int {
	inicio := 0
	fim := len(elementos) - 1

	for inicio <= fim {
		meio := (inicio + fim) / 2
		if elementos[meio] == alvo {
			return meio
		}
		if elementos[meio] < alvo {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}
	return -1 // Elemento não encontrado
}
