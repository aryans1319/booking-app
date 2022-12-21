package main
//FMT stands for Format Package
import "fmt"

func main (){
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50;	

	
	fmt.Println("Welcome to", conferenceName,"our booking application")
	fmt.Println("We have total of", conferenceTickets, "and this", remainingTickets, "are left")
	fmt.Println("Get your tickets here")

	// Another way to print
	fmt.Printf("Welcome here to %v my booking application\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	// Data types in Go
	/*
		[Strings, Integers, Boolean, Maps, Arrays]
		Go can infer a type when we assign a value to 
		any variable 
		Eg: const totalTickets = 50
			var name = "Aryan"
		But, if we don't assign a value to a variable we need
		to define the data type	
		Eg: var userName string
			var totalTickets int
	*/
	var userName string
	var totalTickets int

	userName = "Aryan"
	totalTickets = 50

	fmt.Printf("User %v booked %v tickets for the conference\n", userName, totalTickets);

	/*
		Advantage of GO: 
			discover mistakes at compile time and not at run time
	*/

	// Now how to check the data-type of a variable in Go

	fmt.Printf("Conference Tickets is %T, Username is %T\n", conferenceTickets, userName)

	// Alternative way to write this
	//	var myName string = "Aryan"
	myName := "Aryan"; 
	// this way doesn't work for constants and when we explicitly want to 
	// declare a var for eg: var remainingTickets uint = 50
	fmt.Printf("%v", myName)
}