package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	// pakai strconv ubah tipe data
)

func main() {
	mainTosee()
	/* scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("input is : %s\n", line)

		fmt.Println("Hello World !")
	}

	// error counter
	err := scanner.Err()
	if err != nil {
		fmt.Println("Error encountered:", err)
	} */

	/* SIMPLE HERE */
	// scanner := bufio.NewScanner(os.Stdin)

	// fmt.Print("Your Name : ")
	// scanner.Scan()
	// instxt := scanner.Text()

	// fmt.Print("Born Year : ")
	// scanner.Scan()
	// ins, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	// fmt.Println("Your Name is : ", myname(instxt))
	// fmt.Printf("Your age will be %d in 2020\n", 2020-ins)
}

func mainTosee() {
	cek := -1
	for {
		breaker()
		fmt.Println("WELCOME TO GOLANG :)")
		fmt.Println("1. Input your data")
		fmt.Println("2. Input your Blood type")
		fmt.Println("3. Check are you a vampire ?")
		fmt.Println("4. Exit")
		fmt.Print("Choose : ")

		fmt.Scan(&cek)
		switch cek {
		case 1:
			data()
			break
		case 2:
			blood()
			break
		case 3:
			vampire()
			break
		}
		if cek == 4 {
			break
		}
	}
}

func vampire() {
	cekVamp := rand.Intn(10)
	if cekVamp%2 == 0 {
		fmt.Println("You Are not a Vampire")
	} else {
		fmt.Println("You Are a Vampire 🦇")
	}
	fmt.Println("please type enter...")
	fmt.Scanln()
	breaker()
}

func blood() {
	scan := bufio.NewScanner(os.Stdin)

	// input blood
	fmt.Print("Your blood type : ")
	scan.Scan()
	blood := scan.Text()

	fmt.Println("My blood type is ", blood)

	fmt.Println("please type enter...")
	fmt.Scanln()
	breaker()
}

func data() {
	scan := bufio.NewScanner(os.Stdin)

	// input name
	fmt.Print("Your Name : ")
	scan.Scan()
	name := scan.Text()

	// input address
	fmt.Print("Your Age : ")
	scan.Scan()
	Age := scan.Text()

	fmt.Println("My Name is ", myname(name))
	fmt.Println("My Age is ", Age)

	fmt.Println("please type enter...")
	fmt.Scanln()
	breaker()
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
