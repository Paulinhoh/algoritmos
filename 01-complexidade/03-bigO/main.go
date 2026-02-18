/*
Notações BigO (O):
	É uma forma simples de definir o limite SUPERIOR (upper bound) de um algoritmo de execução. Isso exemplifica qual o tempo maximo que uma aplicação levaria para execultar.

	POR EXEMPLO: se você disser que um algoritmo executa em no maximo 100s em seu pior tempo, isso quer dizer que temos o upper bound, onde não teriamos cenarios onde demoraria mais do que 100s.

Regras da notação BigO:
	Imagine que temos um computador com um unico pocessador e ele executa uma tarefa de cada vez.

	* cada ação de "assign" (var int x = 5) custa 1 unidade de tempo
	* cada ação de "return" (return x) custa 1 unidade de tempo
	* cada operação aritmética (x + y) custa 1 unidade de tempo
	* cada operação lógica (x && y) custa 1 unidade de tempo
	* outras pequenas operações também levam 1 unidade de tempo
	* para outras ordem menores, sempre desconsiderar:
		* T = n² + 3n + 1 ==> O(n²)
	* desconsiderar constantes:
		* T = 3n² + 6n + 1 ==> O(n²)
*/

package main

func main() {

}
