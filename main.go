//go mod init booking-app
package main
//need to declare the execution point - need to give the go compiler a starting point
import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	
	fmt.Printf("conferenceTickets is a %T, remainingTickets is %T, confenrenceName is %T\n",conferenceTickets, remainingTickets, conferenceName)//string interpolation
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are avaiable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//user input scan from the fmt library
	//using pointers - values are stored in memories on your computer - when you reference that value using the variable name
	// in needs to know that memory address - pointer is a "special variable"
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)
	
	fmt.Println("Enter your email address")
	fmt.Scan(&email)
	
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v", remainingTickets, conferenceName)
}