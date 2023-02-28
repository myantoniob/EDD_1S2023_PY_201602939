module begin

go 1.20

replace queue => ../queue

require (
	doubly v0.0.0-00010101000000-000000000000
	queue v0.0.0-00010101000000-000000000000
	stack v0.0.0-00010101000000-000000000000
)

replace doubly => ../doubly

replace stack => ../stack
