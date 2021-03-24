package mains

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Crear un scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Llamar a Scanner
	fmt.Print("Introdusca algo :")
	scanner.Scan()

	// Capturar value del user
	operacion := scanner.Text()

	// Imprimar text
	fmt.Print("El data es ", operacion)

	// Metodo split de los strings

	operador := "/"
	valores := strings.Split(operacion, operador)

	fmt.Println("\n\nEstos son los valores ingresados: ", valores)
	fmt.Println("\nPrimer y segundo valor sumados como texto: ", valores[0]+valores[1])
	fmt.Println("\n\n\n_")

	// Cast valores from text to number
	// strconv.Atoi convierte string a entero

	/* Atoi me devuelve dos valores
		_ = blank identifier

	Sirve como placeholder anónimo, oser asigneselo a nada, el problema es que si
	no lo pongo Go me bota un error, y se pongo una variable y no la utilizo back me
	vota error
	*/

	// ESTO ME VOTA ERROR tan raro
	// operador1, err := strconv.Atoi(valores[0]) // err en un nil ???? WTF
	// operador2, err2 := strconv.Atoi(valores[1])
	// // No era nill siempre, el segundo valor es el ERROR Ahhhhhhh xd:D

	// // fmt.Println("\nCuale es el otro valor entonces :  ", err)
	// if err != nil && err2 != nil {
	// 	fmt.Println("\nSuma de los dos operadores matematicamente: ", operador1+operador2)
	// } else {
	// 	print("Intrusca un valores validos ", err, err2)
	// 	print("Values  ", operador1, operador2)
	// }

	operator1, err1 := strconv.Atoi(valores[0])
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(operator1)
	}
	operator2, err2 := strconv.Atoi(valores[1])

	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(operator2)
	}

	switch operador {
	case "+":
		fmt.Println(operator1 + operator2)
	case "-":
		fmt.Println(operator1 - operator2)
	case "*":
		fmt.Println(operator1 * operator2)
	case "/":
		fmt.Println(operator1 / operator2)
	default:
		fmt.Println(operator1 + operator2)
	}

}
