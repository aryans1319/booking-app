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
}