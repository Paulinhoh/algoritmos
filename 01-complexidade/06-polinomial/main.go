package main

import "fmt"

func main() {

}

func printValues(n int) {
	for i := 1; i <= n; i++ { // for{} -> n
		for j := 1; j <= n; j++ { // for{for{}} -> n²
			print(fmt.Sprintf("i = %d, j = %d", i, j))
		}
		print("finalizando")
	}

	print("fora do loop")
	// T = 9n² + 11n + 5 -> O(n²)
}
