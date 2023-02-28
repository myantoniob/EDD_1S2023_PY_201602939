package main

import (
	"bufio"
	"doubly"
	"fmt"
	"os"
	"queue"
	"strconv"
	"strings"
)

type Customer struct {
	username string
	password string
}

var cola queue.Queue
var listaDoble doubly.DoublyList

func main() {

	/*cola := queue.Queue{}
	cola.Enqueue("mynor", "ramirez", 2016, "pass")
	fmt.Println(cola.Dequeue())
	listaDoble.InsertFront("juan", "perez", 2016, "sdssd")
	listaDoble.TraverseForward()*/

	// Choos log in / or Report or exit
	selected := menu()
	if selected == 1 {
		client := login()
		if client.username == "admin" && client.password == "admin" {
			adminDashboard()
		}

	} else if selected == 2 {
		fmt.Println("Reportes")
	} else {

	}
}

func menu() int {
	reader := bufio.NewReader(os.Stdin)
	exit := true
	for exit {

		fmt.Println("****** EDD GoDrive ******")
		fmt.Println("1. Log in")
		fmt.Println("2. Reports")
		fmt.Println("3. Exit")
		fmt.Println("*************************")
		fmt.Print("Enter a # : ")
		//fmt.Scanln(&option)
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)
		switch option {
		case "1":
			return 1
		case "2":
			return 2
		case "3":
			return 3
		default:
			for i := 0; i < 6; i++ {
				fmt.Println("")
			}
			fmt.Println("Choose a valid option :( ")
		}

	}
	return -1
}

func login() Customer {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("*************")
	fmt.Print("User name: ")
	userame, _ := reader.ReadString('\n')
	userame = strings.TrimSpace(userame)
	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)
	fmt.Println("*************")
	for i := 0; i < 6; i++ {
		fmt.Println("")
	}

	return Customer{username: userame, password: password}

}

func adminDashboard() {
	exit := true
	for exit {
		fmt.Println("**** Administrator Dashboard ****")
		fmt.Println("1. Pending students")
		fmt.Println("2. Accepted students")
		fmt.Println("3. New students")
		fmt.Println("4. Students Bulk load")
		fmt.Println("5. Exit")
		fmt.Println("*********************************")
		fmt.Print("Enter a #: ")
		reader := bufio.NewReader(os.Stdin)
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			pendingStudents()
			break
		case "2":
			acceptedStudents()
			break
		case "3":
			newStudents()

		case "4":
			studentsBulkLoad()
			break
		case "5":
			fmt.Println(" ")
			exit = false
			break
		default:
			fmt.Println("Enter a valid option :( ")
		}
	}

}

func pendingStudents() {
	fmt.Println("**** Pending Students ****")

	fmt.Println("**** Pending #: ", cola.Len(), "****")
	actual := cola.Dequeue()
	fmt.Println("Current student: ", actual.Name, " ", actual.LastName)
	fmt.Println("1. Accept student")
	fmt.Println("2. Reject student")
	fmt.Println("3. Return to menu")
	fmt.Print("Enter a #: ")
	reader := bufio.NewReader(os.Stdin)
	option, _ := reader.ReadString('\n')
	option = strings.TrimSpace(option)

	switch option {
	case "1":
		listaDoble.InsertFront(actual.Name, actual.LastName, actual.Carnet, actual.Password)
		fmt.Println("Student accepted :) ")
		listaDoble.TraverseForward()
	case "2":
		fmt.Println("Student rejected :( ")
	case "3":
		fmt.Println(" ")
	default:
		fmt.Println("Choose a valid option :(")
	}
}

func acceptedStudents() {}

func newStudents() bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("**** Create a new student ****")

	fmt.Print("Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Last Name: ")
	lastName, _ := reader.ReadString('\n')
	lastName = strings.TrimSpace(lastName)

	fmt.Print("Carnet : ")
	carnet, _ := reader.ReadString('\n')
	carnet = strings.TrimSpace(carnet)
	carnet1, _ := strconv.Atoi(carnet)

	fmt.Print("Password : ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	//To add a node to the queue
	cola.Enqueue(name, lastName, carnet1, password)
	return true
}
func studentsBulkLoad() {}
