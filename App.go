package main

import (
	"bufio"
	"fmt"
	"os"

	// pakai strconv ubah tipe data
	"strconv"
)

func main() {
	breaker()
	fmt.Println("WELCOME TO GOLANG :)")

	// scanner := bufio.NewScanner(os.Stdin)
	// if scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Printf("input is : %s\n", line)

	// 	fmt.Println("Hello World !")
	// }

	// // error counter
	// err := scanner.Err()
	// if err != nil {
	// 	fmt.Println("Error encountered:", err)
	// }

	/* SIMPLE HERE */
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Your Name : ")
	scanner.Scan()
	instxt := scanner.Text()

	fmt.Print("Born Year : ")
	scanner.Scan()
	ins, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Println("Your Name is : ", myname(instxt))
	fmt.Printf("Your age will be %d in 2020\n", 2020-ins)
}

func myname(txt string) string {
	plus := txt + " Rais"
	return plus
}

func breaker() {
	for i := 0; i < 20; i++ {
		fmt.Printf("\n")
	}
}
