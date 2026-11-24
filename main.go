package main

func main() { //just test
	bc := NewBlockChain()
	defer bc.db.Close()
	cli := CLI{bc}
	cli.Run()
}
