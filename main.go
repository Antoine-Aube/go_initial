//go mod init booking-app
package main
//need to declare the execution point - need to give the go compiler a starting point
import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50
	
	fmt.Printf("conferenceTickets is a %T, remainingTickets is %T, confenrenceName is %T\n",conferenceTickets, remainingTickets, conferenceName)//string interpolation
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are avaiable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask the suer for their name
	userName = "Tom"
	userTickets = 2

	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)
}