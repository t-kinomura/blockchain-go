package main

func main() {
	bc := NewBlockchain()

	bc.AddTransaction("transaction1")
	bc.AddTransaction("transaction2")
	bc.CreateBlock()

	bc.AddTransaction("transaction3")
	bc.AddTransaction("transaction4")
	bc.CreateBlock()

	bc.Print()
}


