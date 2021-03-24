package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parse(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// Struct como en C
type calc struct{}

/* Las Receiver Function le da la propiedad al struct de tener
el método dentro de el, es decir cuando instanciemos a
la struct vamos a poder llamar a los métodos dentro de
la Receiver Function*/
func (calc) operate(entrada string, operador string) int {

	entradalimpia := strings.Split(entrada, operador)
	operador1 := parse(entradalimpia[0])
	operador2 := parse(entradalimpia[1])

	print("Y el intA que ?? ")
	total := 0

	switch operador {
	case "+":
		total = operador1 + operador2
	case "-":
		total = operador1 - operador2
	case "*":
		total = operador1 * operador2
	case "/":
		total = operador1 / operador2
	default:
		total = 11111111
	}

	return total
}

func main() {
	print("Digite Entrada \t")
	entrada := leerEntrada()
	print("\n Digite Operador: ")
	operador := leerEntrada()

	// Instancia de calc ???
	c := calc{}
	println(c.operate(entrada, operador))
}
