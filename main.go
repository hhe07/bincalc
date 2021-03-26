package main


import (
	"fmt"
	"bufio"
	"os"
)
// todo: if time: take args and allow for verbosity?

func main() {
	var res = 0
	scanner := bufio.NewScanner(os.Stdin)
	for true{
		// take input
		fmt.Print("> ")
		scanner.Scan()
		x := scanner.Text()
		if x == "exit"{
			break
		} else if x == ""{
			continue
		} else if x == "x"{
			fmt.Println(res)
			continue
		}
		tkns, err := Tokenizer(x)
		if err != nil{
			fmt.Println(err)
			break
		}
		res, err = ShuntYard(tkns, res)
		if err != nil{
			fmt.Println(err)
			break
		}
		fmt.Println(res)


	}

}