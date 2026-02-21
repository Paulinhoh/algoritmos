package sort

func quickSort(numeros []int) []int {
	if len(numeros) <= 1 {
		return numeros
	}

	n := make([]int, len(numeros))
	copy(n, numeros)

	indicePivo := len(n) / 2
	pivo := n[indicePivo]
	n = append(n[:indicePivo], n[indicePivo+1:]...)

	menores, maiores := particionar(n, pivo)

	return append(
		append(quickSort(menores), pivo),
		quickSort(maiores)...)
}

func particionar(n []int, pivo int) (menores []int, maiores []int) {
	for _, v := range n {
		if v <= pivo {
			menores = append(menores, v)
		} else {
			maiores = append(maiores, v)
		}
	}
	return menores, maiores
}
