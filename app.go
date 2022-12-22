package main
//FMT stands for Format Package
import "fmt"

func main (){
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50;	

	
	fmt.Println("Welcome to", conferenceName,"our booking application")
	fmt.Println("We have total of", conferenceTickets, "and this", remainingTickets, "are left")
	fmt.Println("Get your tickets here")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// Arrays
	// var bookings [50]string

	//Slices
	var bookings []string 

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName);

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName);

	fmt.Println("Enter your email: ")
	fmt.Scan(&email);

	fmt.Println("Enter user tickets: ")
	fmt.Scan(&userTickets);

	remainingTickets = remainingTickets - userTickets;
	
	// bookings[0] = firstName + " " + lastName 
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value is: %v\n", bookings[0]);
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(bookings));

	fmt.Printf("thank you %v %v for booking %v tickets!\n", firstName, lastName, userTickets);

	fmt.Printf("Remaining Tickets: %v\n", remainingTickets);

	// Slices - an abstraction of array, dynamic in size
	
}