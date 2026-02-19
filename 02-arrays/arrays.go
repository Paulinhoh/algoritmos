package main

import "fmt"

func main() {
	array := []int{5, 2, 0, 3, 4, 6}
	printarArray(array)

	oddNumbers := removeEvenNumbersFromArray(array)
	fmt.Println(oddNumbers)

	reversedArray := reverseArray(array)
	fmt.Println(reversedArray)

	moveZeros := moveZeros(array)
	fmt.Println(moveZeros)

	word := "arara"
	palindrome := isPalindrome(word)
	fmt.Printf("%s é palindromo: %v", word, palindrome)
}

func printarArray(array []int) {
	for _, item := range array {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

func removeEvenNumbersFromArray(array []int) (oddNumbers []int) {
	for _, item := range array {
		if item%2 != 0 {
			oddNumbers = append(oddNumbers, item)
		}
	}

	return
}

func reverseArray(array []int) []int {
	start := 0
	end := len(array) - 1

	for start < end {
		array[start], array[end] = array[end], array[start]
		start++
		end--
	}

	return array
}

func moveZeros(array []int) []int {
	j := 0
	arrayLen := len(array)
	for i := 0; i < arrayLen; i++ {
		if array[i] != 0 && array[j] == 0 {
			array[i], array[j] = array[j], array[i]
		}
		if array[j] != 0 {
			j++
		}
	}

	return array
}

func isPalindrome(word string) bool {
	start := 0
	end := len(word) - 1

	for start < end {
		letterStart := fmt.Sprintf("%c", word[start])
		letterEnd := fmt.Sprintf("%c", word[end])

		if letterStart != letterEnd {
			return false
		}

		start++
		end--
	}

	return true
}
