package main

import "fmt"

func main() {
	var (
		nickname string
		major    string
		class    string
		nim      string
	)
	fmt.Print("Insert Your Nickname: ")
	fmt.Scanln(&nickname)

	fmt.Print("What's Your Major: ")
	fmt.Scanln(&major)

	fmt.Print("Insert Your Class: ")
	fmt.Scanln(&class)

	fmt.Print("Insert Your NIM: ")
	fmt.Scanln(&nim)

	fmt.Println("")
	fmt.Println("Haii!, you guys can call me " + nickname + ", my major is " + major + ", from class " + class + ", and my NIM is" + nim)
}
