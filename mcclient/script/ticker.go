package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			t := time.Now()
			fmt.Printf("counter: %d, time: %s\n", counter, t.Format("03:04:05"))
			counter += 1
			if counter%5 == 0 {
				time.Sleep(15 * time.Second)
			} else {
				time.Sleep(time.Second)
			}
		}
	}
}
