package main

import (
	_ "bufio"
	"fmt"
	_ "models"
	_ "os"
	"unsafe"
)

func main() {
	//disp :=Dis
	//scanner:=*bufio.Scanner
	//tokenize :=func(string) []string
	//outSeparator:=rune
	//scanner := bufio.NewScanner(os.Stdin)
	//bufio.SplitFunc()
	//bufio.ScanLines()
	errT := func(err error) string {
		if err != nil {
			return fmt.Sprintf("Error occurred: %v\n", err)
		}
		return "Phone was successfully changed!\n"
	}
	fmt.Printf("Type - %T, value - %v, size - %v \n", errT, errT, unsafe.Sizeof(errT))
	var errT1 any = errT
	fmt.Printf("Type - %T, value - %v, size - %v \n", errT1, errT1, unsafe.Sizeof(errT1))

}
