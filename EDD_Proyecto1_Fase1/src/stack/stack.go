package stack

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type item struct {
	value string // hold any data type
	next  *item
}
type Stack struct {
	top  *item
	size int
}

func (stack *Stack) Push(v string) {
	stack.top = &item{
		value: v,
		next:  stack.top,
	}
	stack.size++
}
func (stack *Stack) Len() int {
	return stack.size
}
func (stack *Stack) isEmpty() bool {
	return stack.Len() == 0
}
func (stack *Stack) Pop() string {
	if stack.isEmpty() {
		return ""
	}

	valueToPop := stack.top.value
	stack.top = stack.top.next
	stack.size--
	return valueToPop
}
func (stack *Stack) Peek() string {
	if stack.isEmpty() {
		return ""
	}
	return stack.top.value
}

func (stack *Stack) ReadStack() {
	temp := stack.top
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
	fmt.Println(" ")
}

func (r *Stack) Graficar() {
	fmt.Println("* Impresion")
	nombre_archivo_dot := "./dot/pila.dot"
	nombre_imagen := "./images/pila.jpg"
	texto := "digraph lista{\n"
	texto += "rankdir=LR;\n"
	texto += "node[shape = record];\n"
	texto += "nodonull1[label=\"null\"];\n"
	texto += "nodonull2[label=\"null\"];\n"
	auxiliar := r.top
	contador := 0
	for i := 0; i < r.Len(); i++ {
		texto = texto + "nodo" + strconv.Itoa(i) + "[label=\"{|" + "Info: " + auxiliar.value + "|}\"];\n"
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
