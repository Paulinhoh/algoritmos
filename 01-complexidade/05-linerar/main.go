package main

func main() {

}

func calculate(n int) int {
	sum := 0                  // 1
	for i := 1; i <= n; i++ { // 1+3(n+1)+3n
		sum += i // n(1+1+1+1) = 4n
	}
	return sum // 1+1
	// T = 10n + 7 -> O(n)
}
