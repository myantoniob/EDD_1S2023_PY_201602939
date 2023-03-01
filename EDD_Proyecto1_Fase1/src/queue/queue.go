package queue

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type Item struct {
	Name     string
	LastName string
	Carnet   int
	Password string
}

type Node struct {
	name     string
	lastName string
	carnet   int
	password string
	next     *Node
}

type Queue struct {
	front  *Node
	rear   *Node
	length int
}

func NewNode(name string, lastName string, carnet int, password string, next *Node) *Node {
	return &Node{name: name, lastName: lastName, carnet: carnet, password: password, next: next}
}

func (queue *Queue) Enqueue(name string, lastName string, carnet int, password string) {
	//fmt.Println(name, lastName, carnet, password)
	newItem := NewNode(name, lastName, carnet, password, queue.front)
	if queue.isEmpty() {
		queue.front = newItem
		queue.rear = newItem
	} else {
		queue.rear.next = newItem
		queue.rear = newItem
	}
	queue.length = queue.length + 1
}
func (queue *Queue) Len() int {
	return queue.length
}
func (queue *Queue) isEmpty() bool {
	return queue.Len() == 0
}
func (queue *Queue) Dequeue() Item {
	if queue.isEmpty() {
		return Item{}
	}
	valueToDequeue := Item{Name: queue.front.name, LastName: queue.front.lastName, Carnet: queue.front.carnet, Password: queue.front.password} //queue.front
	queue.front = queue.front.next
	queue.length--
	return valueToDequeue
}
func (queue *Queue) Front() Item {
	if queue.isEmpty() {
		return Item{}
	}
	item := Item{Name: queue.front.name, LastName: queue.front.lastName, Carnet: queue.front.carnet, Password: queue.front.password}
	return item
}

func (r *Queue) Graficar() {
	fmt.Println("* Impresion")
	nombre_archivo_dot := "./dot/cola.dot"
	nombre_imagen := "./images/cola.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	auxiliar := r.front
	contador := 0
	for i := 0; i < r.Len(); i++ {
		texto = texto + "nodo" + strconv.Itoa(i) + "[label=\"{|" + "Carnet: " + strconv.Itoa(auxiliar.carnet) + " Name: " + auxiliar.name + " " + auxiliar.lastName + "|}\"];\n"
		auxiliar = auxiliar.next
	}
	texto += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < r.Len(); i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "-> nodo" + strconv.Itoa(c) + ";\n"
		//texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
		contador = c
	}

	texto += "nodo" + strconv.Itoa(contador) + "->nodonull2;\n"
	texto += "}"

	crearArchivoDot(nombre_archivo_dot)
	escribirArchivoDot(texto, nombre_archivo_dot)
	ejecutar(nombre_imagen, nombre_archivo_dot)

}

// *************************

func crearArchivoDot(nombre_archivo string) {
	//Verifica que el archivo existe
	var _, err = os.Stat(nombre_archivo)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
		var file, err = os.Create(nombre_archivo)
		if err != nil {
			return
		}
		defer file.Close()
	}
	fmt.Println("Archivo creado exitosamente", nombre_archivo)
}

func escribirArchivoDot(contenido string, nombre_archivo string) {
	// Abre archivo usando permisos READ & WRITE
	var file, err = os.OpenFile(nombre_archivo, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	// Escribe algo de texto linea por linea
	_, err = file.WriteString(contenido)
	if err != nil {
		return
	}
	// Salva los cambios
	err = file.Sync()
	if err != nil {
		return
	}
	fmt.Println("Archivo actualizado existosamente.")
}

func ejecutar(nombre_imagen string, archivo_dot string) {
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tjpg", archivo_dot).Output()
	mode := 0777
	_ = ioutil.WriteFile(nombre_imagen, cmd, os.FileMode(mode))
}

// *************************
