package queue

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
