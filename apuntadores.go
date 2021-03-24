package main

func main() {
	value := 55
	println("El valor de X es ", value)

	changue_vlaue(value)

	println("Direccion en Memoria:  ", &value)
	println("El valor de X after es:  ", value)
	reference := &value
	println("Valor de referencia es:  ", *reference)

}
func changue_vlaue(a int) {
	println("Direccion en Memoria de a :  ", &a)
	a = 36
}
