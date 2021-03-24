package mains //package name
// package	otroPaquete

import "fmt" // librerías necesarias para importar el paquete

// func types
func privada() {
	fmt.Println("Ejecutar lógica que no necesita ser exportada (pertenece solo a este paquete)")
}

func Publica() {
	fmt.Println("Lógica que quiero exportar a otros paquetes")
}

// defer
func printHelloWMarte() {
	// es la utliza instruccion que hace apenas termina, servirair como un return ???
	defer fmt.Println("Marte We are lives!")
	fmt.Println("Hello")

	fmt.Println("defer ejecuta un fragmento de código hasta que la función haya terminado")
}
func main() {
	// Variables
	fmt.Println("Strings")
	var mensaje string = "Hola MARTE"

	// Con polla, Go interpreta el dato como --> php o js
	easyMessage := "Hola MARTE usando := polla"

	fmt.Println(mensaje)
	fmt.Println(easyMessage)

	// Puedes comentar usando "//"

	// float numbers
	fmt.Println("\n Floats")
	a := 10.
	var b float64 = 3
	fmt.Println(a / b)

	fmt.Println("\n integers")
	//integer numbers
	var c int = 10
	d := 3
	fmt.Println(c / d)

	// boolean
	fmt.Println("\n Boleans")
	var x bool = true
	y := false
	fmt.Println(x || y)
	fmt.Println(x && y)
	fmt.Println(!x)

	fmt.Println("\n Methods privada")
	//Lógica privada
	privada()

	fmt.Println("\n Methods Publica")
	//Lógica Publica
	// PARA QUE UNA FUNCION SEA PUBLICA, HAY QUE INICIARLA CON LA
	// PRIMERA LETRA EN MAYUSCULA
	Publica()

	fmt.Println("\n Methods printHelloWMarte defer")
	//Función "defer"
	printHelloWMarte()
}
