package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
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
		fmt.Println("4. Testing struck")
		fmt.Println("5. Exit")
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
		case 4:
			testStruct()
			break
		}

		if cek == 5 {
			break
		}
	}
}

func testStruct() {
	// var emp employe
	// emp.firstname = "Aldi"
	// emp.lastname = "Nugroho"
	// emp.salary = 20000000
	// emp.fulltime = true

	// fmt.Println("Firstname : " + emp.firstname)
	// fmt.Println("Lastname : " + emp.lastname)
	// fmt.Println("Salary : " + strconv.Itoa(emp.salary))
	// fmt.Println("Job Type : " + strconv.FormatBool(emp.fulltime))
	// fmt.Println("please type enter to JSON...")
	// fmt.Scanln()
	type employe struct {
		Firstname string
		Lastname  string
		Created   time.Time
	}

	list := employe{
		Firstname: "Aldi",
		Lastname:  "Nugroho",
		Created:   time.Now(),
	}

	var jsonData []byte
	jsonData, err := json.MarshalIndent(list, "", "    ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
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
