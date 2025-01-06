package main

import (
	"fmt"
	"time"
)
func AutoLogout(refresh chan rune, logOffTimeMS time.Duration) {

	tmr := time.NewTimer(5 * time.Second)
	tick := time.Tick(1 * time.Second)
	var i int = 0

	for {
		select {
		case <-tick:
			i++
			fmt.Println(time.Now() , " Tick...", logOffTimeMS)
			if i > 10 {
				fmt.Println("True...", time.Now())

			}else {
				tmr.Reset(5 *time.Second)
			}
			break
		case <- tmr.C:
			fmt.Println("Restart")
			break
		}
	}
}

func main() {

	fmt.Println("Hello World.")
	fmt.Println("Go Building could make seriously trouble.")

	refresh := make (chan rune)


 AutoLogout(refresh, 5)
}
