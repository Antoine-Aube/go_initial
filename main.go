//go mod init booking-app
package main
//need to declare the execution point - need to give the go compiler a starting point
import "fmt"

func main() {
	var conferenceName = "Go Conference"
													//string interpolation
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("Get your tickets here to attend")

	
}