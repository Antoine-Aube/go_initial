//go mod init booking-app
package main
//need to declare the execution point - need to give the go compiler a starting point
import ("fmt"
				"strings"
			)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)
	
	fmt.Printf("conferenceTickets is a %T, remainingTickets is %T, confenrenceName is %T\n",conferenceTickets, remainingTickets, conferenceName)//string interpolation
	fmt.Printf("We have a total of %v tickets and %v are avaiable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//define array length, element types, and then can prepopulate with elements of just empty
	// var bookings = [50]string{}

	//arr syntax - need to assign number of elements in order to make it an array. 
	// var bookings []string

	//Can also use conditionals after the for loop declarations like a while loop in ruby
	// for remainingTickets > 0 && len(bookings) < 50 {


	for {

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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets

		
		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firstName + " " + lastName)
		
			fmt.Printf("The whole slice: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Array type: %T\n", bookings)
			fmt.Printf("Array Length: %v\n", len(bookings))
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			//the blank identifier is a space holder for the index, since we are not using hte index within the loop
			for _, booking := range bookings {
				//fields key word from the string package that we import above

				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			//using a conditional to see if tickets have run out
			//can save a bool into a variable, haven't done this in other languages before

			var noTicketsRemaining bool = remainingTickets == 0
			if noTicketsRemaining {

				fmt.Println("Our conference is booked out please come back next year")
				break 
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("Invalid email")
			}

			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid")
			}
			//key word continue says to keep going with the rest of hte program rather than exiting
		}
	}
	
}

func greetUsers(name string, tickets int, remaining uint) {
	fmt.Printf("Welcome to %v booking application  \n\n", name)
	fmt.Printf("We have a total of %v tickets and %v are avaiable.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	
}