package doubly

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"stack"
	"strconv"
	"time"
)

type node struct {
	//data string

	Name     string
	LastName string
	Carnet   int
	Password string

	StStack stack.Stack

	prev *node
	next *node
}

type DoublyList struct {
	len  int
	tail *node
	head *node
}

func initDoublyList() *DoublyList {
	return &DoublyList{}
}

func (d *DoublyList) InsertFront(name string, lastName string, carnet int, password string) {
	newNode := &node{
		Name: name, LastName: lastName,
		Carnet: carnet, Password: password,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.len++
	return
}

func (d *DoublyList) InsertEnd(name string, lastName string, carnet int, password string) {
	newNode := &node{
		Name: name, LastName: lastName,
		Carnet: carnet, Password: password,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		d.tail = newNode
	}
	d.len++
	return
}
func (d *DoublyList) TraverseForward() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := d.head
	for temp != nil {
		//fmt.Printf("value = %v, prev = %v, next = %v\n", temp.carnet, temp.prev, temp.next)
		fmt.Println("Name:", temp.Name, " ", temp.LastName, ", Carnet: ", temp.Carnet)
		temp = temp.next
	}
	fmt.Println()
	return nil
}

func (d *DoublyList) Search(carne int, pass string) node {

	if d.head == nil {
		return node{}
	}

	temp := d.head
	for temp != nil {
		if temp.Carnet == carne && temp.Password == pass {
			ct := time.Now()
			ms := fmt.Sprintf("# Se inicio sesion %s %s \n %v/%v/%v  %v:%v", temp.Name, temp.LastName, ct.Day(), ct.Month(), ct.Year(), ct.Hour(), ct.Minute())

			temp.StStack.Push(ms)
			return *temp
		}
		temp = temp.next
	}

	return node{}
}

func (d *DoublyList) TraverseReverse() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := d.tail
	for temp != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.Carnet, temp.prev, temp.next)
		temp = temp.prev
	}
	fmt.Println()
	return nil
}

func (d *DoublyList) Size() int {
	return d.len
}

// ---------------------------------

func (r *DoublyList) Graficar() {
	fmt.Println("Impresion")
	nombre_archivo_dot := "./dot/listadoble.dot"
	nombre_imagen := "./images/listadoble.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	auxiliar := r.head
	contador := 0
	for i := 0; i < r.Size(); i++ {
		texto = texto + "nodo" + strconv.Itoa(i) + "[label=\"{|" + "Carne: " + strconv.Itoa(auxiliar.Carnet) + " Name: " + auxiliar.Name + " " + auxiliar.LastName + "|}\"];\n"
		auxiliar = auxiliar.next
	}
	texto += "nodonull1->nodo0 [dir=back];\n"
	for i := 0; i < r.Size(); i++ {
		c := i + 1
		texto += "nodo" + strconv.Itoa(i) + "->nodo" + strconv.Itoa(c) + ";\n"
		texto += "nodo" + strconv.Itoa(c) + "->nodo" + strconv.Itoa(i) + ";\n"
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
