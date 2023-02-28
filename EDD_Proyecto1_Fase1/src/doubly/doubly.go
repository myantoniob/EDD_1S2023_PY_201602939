package doubly

import (
	"fmt"
)

type node struct {
	data string

	name     string
	lastName string
	carnet   int
	password string

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
		name: name, lastName: lastName,
		carnet: carnet, password: password,
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
		name: name, lastName: lastName,
		carnet: carnet, password: password,
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
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.carnet, temp.prev, temp.next)
		temp = temp.next
	}
	fmt.Println()
	return nil
}

func (d *DoublyList) TraverseReverse() error {
	if d.head == nil {
		return fmt.Errorf("TraverseError: List is empty")
	}
	temp := d.tail
	for temp != nil {
		fmt.Printf("value = %v, prev = %v, next = %v\n", temp.carnet, temp.prev, temp.next)
		temp = temp.prev
	}
	fmt.Println()
	return nil
}

func (d *DoublyList) Size() int {
	return d.len
}
