package main

// GO é fortemente tipado

const a = "Hello, World"

// Essa declarações são de escopo global
var (
	b bool    = true      // Valor padrão inferido false
	c int     = 10        // Valor padrão inferido 0
	d string  = "Gustavo" // Valor padrão inferido vazio
	e float64 = 1.2       // // Valor padrão inferido +0.000000e+000
)

func main() {
	//	var a string = "X" // Escopo local
	// Podemos ultilizar shordHand para declarar variáveis em GO. Só podemos usar isso na primeira declaração da variável
	a := "X"
	println(a) // Em variáveis local temos que ultizalas, pois o não uso delas causa erro de compilação
}

func x() {

}
