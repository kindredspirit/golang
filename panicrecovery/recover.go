package main

import "fmt"

func main() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Cops and Robbers")
	fmt.Println("")
	cop()
	fmt.Println("Cop returned to station unharmed.")	
	fmt.Println("")
	fmt.Println("")
}


/*************************************************************************************
Function cop() - demonstrates defered execution and recovery from panic.
		 once the robbers panic, they run and the cop chases after them.
		 cop catches them and returns them to the police station for booking.
*************************************************************************************/
func cop() {
	defer func() {
		// at first tried using "null" but "nil" is the only one that works
		if r := recover(); r != nil {
			fmt.Println("Cop chases down robber...", r)
		}
	}()

	fmt.Println("A group of robbers are coming...")
	robber(0)
	fmt.Println("Robbers are cornered and have been caught by cop!")
}


/***************************************************************************
Function robber() - demonstrates panic after iterating through each robber.
		    each robber assigns the next robber to be the lookout for a cop.
		    after the 4th robber sees a cop, they all panic.
Params 		  - i int: a robber represented as an int
***************************************************************************/
func robber(i int) {
	if i > 3 {
		fmt.Println("Robber ",i," says: 'Shit guys, it's A KERP! Let's get outta here!'")
		panic(fmt.Sprintf("%v",i))
	}

	defer fmt.Println("Robber ",i," flees from scene.")
	fmt.Println("Alright you check for cops, robber ",i)
	robber(i+1)
}
