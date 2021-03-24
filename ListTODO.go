package main

import "fmt"

// In slice es mejor que un array en Go
type taskList struct {
	tasks []*task
}

func (t *taskList) appendTask(tl *task) {
	// append toma dos valores, la lista, y el elemento a agregar
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) removeTask(index int) {

	/*
		[2,5,6,8,9,11,74]
		Si quiero quitar el ochco then
			armo un nuevo slice
			1. [2,5,6]
			2. [9,11,74]
			Despues los uno
			con append
	*/
	print("QUE POS ES ESA DE index ??? ", t.tasks[index+1:])

	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

type task struct {
	name        string
	description string
	completed   bool
}

func (t *task) markAsCompleted() {
	t.completed = true
}

func (t *task) updateName(name string) {
	t.name = name
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func main() {
	t1 := &task{
		name:        "Terminar Curso de Go",
		description: "Terminar el Curso de Go en Platzi en las proximas dos semanas",
	}
	t2 := &task{
		name:        "Terminar Curso de JavaScript",
		description: "Terminar mi curso de Async/Await en JavaScript",
	}

	// En las llaves pones los elementos del slice
	list := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	//  "%+v Imprima llave , valor
	fmt.Printf("%+v\n", *list.tasks[0])
	fmt.Printf("%+v\n", *list.tasks[1])
	list.tasks[1].markAsCompleted()
	// fmt.Printf("%+v\n", *list.tasks[1])
	t3 := &task{
		name:        "Construir mi propio servidor web",
		description: "Construir mi propio web server utilizando Go",
	}
	list.appendTask(t3)
	// fmt.Printf("%+v\n", *list.tasks[2])
	list.removeTask(1)
	// fmt.Printf("%+v\n", *list.tasks[1])
	fmt.Printf("\n Longitud:  ", len(list.tasks))
	var buffer [256]byte
	var slice []byte = buffer[100:150]
	fmt.Printf("\n Longitud:  ", slice)
}
