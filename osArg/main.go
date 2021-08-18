package main;

import(
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)
	
	fmt.Println("Path : ", os.Args[0])
	fmt.Println("1st arg : ", os.Args[1])
	fmt.Println("2nd arg : ", os.Args[2])
	fmt.Println("3rd arg : ", os.Args[3])
	
	fmt.Println("Total args : ", len(os.Args))
}