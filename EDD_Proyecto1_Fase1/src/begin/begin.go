package main

import (
	"bufio"
	"doubly"
	"encoding/csv"
	"fmt"
	"os"
	"queue"
	"stack"
	"strconv"
	"strings"
	"time"
)

type Customer struct {
	username string
	password string
}

var cola queue.Queue
var listaDoble doubly.DoublyList
var pila stack.Stack

func main() {

	/*cola := queue.Queue{}
	listaDoble.TraverseForward()*/

	// Choos log in / or Report or exit
	exit := true
	for exit {
		selected := menu()
		if selected == 1 {
			client := login()
			if client.username == "admin" && client.password == "admin" {
				adminDashboard()
			}

		} else if selected == 2 {
			fmt.Println("Reportes")
			readAdminStack()
		} else {
			fmt.Println("Have a nice day :) ")
			exit = false
		}
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
			for i := 0; i < 6; i++ {
				fmt.Println("")
			}
			return 1
		case "2":
			for i := 0; i < 6; i++ {
				fmt.Println("")
			}
			return 2
		case "3":
			for i := 0; i < 6; i++ {
				fmt.Println("")
			}
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
			for i := 0; i < 6; i++ {
				fmt.Println(" ")
			}
			pendingStudents()

		case "2":
			for i := 0; i < 6; i++ {
				fmt.Println(" ")
			}
			acceptedStudents()

		case "3":
			for i := 0; i < 6; i++ {
				fmt.Println(" ")
			}
			newStudents()

		case "4":
			for i := 0; i < 6; i++ {
				fmt.Println(" ")
			}
			studentsBulkLoad()

		case "5":
			for i := 0; i < 6; i++ {
				fmt.Println(" ")
			}
			exit = false

		default:
			fmt.Println("Enter a valid option :( ")
		}
	}

}

func pendingStudents() {
	fmt.Println("**** Pending Students ****")
	exit := true
	for exit {
		fmt.Println("**** Pending #: ", cola.Len(), "****")
		actual := cola.Front()
		if actual.Carnet != 0 {
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

				aux := cola.Dequeue()
				ct := time.Now()
				ms := fmt.Sprintf("Student accepted: %s %s %v/%v/%v  %v:%v", aux.Name, aux.LastName, ct.Day(), ct.Month(), ct.Year(), ct.Hour(), ct.Minute())
				//ms := fmt.Sprintf("Student accepted: ", aux.Name, " ", ct.Day(), "/", ct.Month(), "/", ct.Year(), " ", ct.Hour(), ":", ct.Minute())
				pila.Push(ms)

				for i := 0; i < 6; i++ {
					fmt.Println(" ")
				}
				fmt.Println("Student accepted :) ")
				//listaDoble.TraverseForward()
			case "2":
				aux := cola.Dequeue()
				ct := time.Now()
				ms := fmt.Sprintf("Student rejected: %s %s %v/%v/%v  %v:%v", aux.Name, aux.LastName, ct.Day(), ct.Month(), ct.Year(), ct.Hour(), ct.Minute())
				//ms := fmt.Sprintf("Student rejected: ", aux.Name, " ", ct.Day(), "/", ct.Month(), "/", ct.Year(), " ", ct.Hour(), ":", ct.Minute())
				pila.Push(ms)

				for i := 0; i < 6; i++ {
					fmt.Println(" ")
				}
				fmt.Println("Student rejected :( ")

			case "3":
				fmt.Println(" ")
				exit = false
			default:
				fmt.Println("Choose a valid option :(")
			}
		}

	}

}

func acceptedStudents() {

	fmt.Println("***** Students List *****")
	fmt.Println("Size: ", listaDoble.Size())
	listaDoble.TraverseForward()
}

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
	ct := time.Now()
	ms := fmt.Sprintf("New student created: %s %s %v/%v/%v  %v:%v", name, lastName, ct.Day(), ct.Month(), ct.Year(), ct.Hour(), ct.Minute())
	pila.Push(ms)

	for i := 0; i < 6; i++ {
		fmt.Println(" ")
	}
	return true
}

type Preload struct {
	name     string
	lastname string
	carnet   int
	password string
}

func studentsBulkLoad() {
	archivo := readCsvFile("prueba.csv")
	for i := 1; i < len(archivo); i++ {
		var preload Preload
		for j := 0; j < len(archivo[i]); j++ {
			if j == 0 {
				preload.carnet, _ = strconv.Atoi(archivo[i][j])
			}
			var res []string
			if j == 1 {
				res = strings.Split(archivo[i][j], " ")
				preload.name = res[0]
				preload.lastname = res[1]
			}

			if j == 2 {
				preload.password = archivo[i][j]
			}
		}
		// Go to the pending queue
		cola.Enqueue(preload.name, preload.lastname, preload.carnet, preload.password)
	}
	ct := time.Now()

	ms := fmt.Sprintf("Students BulkLoad was done %v/%v/%v  %v:%v", ct.Day(), ct.Month(), ct.Year(), ct.Hour(), ct.Minute())
	//ms := fmt.Sprintf("Students were added ", ct.Day(), "/", ct.Month(), "/", ct.Year(), " ", ct.Hour(), ":", ct.Minute())
	pila.Push(ms)

	fmt.Println("<- New students were added to the queue ->")
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func readAdminStack() {
	pila.ReadStack()
}
