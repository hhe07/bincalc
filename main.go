package main


import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var res = 0 // init result
	scanner := bufio.NewScanner(os.Stdin) // init input scanner
	for true{
		// actual loop
		// take input
		fmt.Print("> ")
		scanner.Scan()
		x := scanner.Text()

		// shortcuts
		if x == "exit"{
			// case: exit: stop loop
			break
		} else if x == ""{
			// case: none: skip
			continue
		} else if x == "x"{
			// case: "x": print result
			fmt.Println(res)
			continue
		}

		// tokenize input
		tkns, err := Tokenizer(x)
		if err != nil{
			fmt.Println(err)
			break
		}
		// get solution by shunting yard
		res, err = ShuntYard(tkns, res)
		if err != nil{
			fmt.Println(err)
			break
		}
		// print result
		fmt.Println(res)


	}

}