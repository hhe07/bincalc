package main


import (
	"fmt"
)
// todo: if time: take args and allow for verbosity?

func main() {
	var res = 0
	for true{
		// take input
		var x string
		fmt.Print("> ")
		fmt.Scanln(&x)
		if x == ""{
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
		res, err = ShuntYard(tkns, res) // should previous res be a ptr?
		if err != nil{
			fmt.Println(err)
			break
		}
		fmt.Println(res)

	}
	
	//STest()
}