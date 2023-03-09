package application

import "fmt"

func Monitor(sites []string, tries int, delay int) {
	fmt.Println("EXEC MONITOR")
	fmt.Println("sites", sites)
	fmt.Println("tries:", tries)
	fmt.Println("delay:", delay)
}
