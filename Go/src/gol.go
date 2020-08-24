package main

import (
	"fmt"
	"math/rand"
	"time"
	"log"
	"os"
)

func life(rows int, cols int, some float64, generations int) {
	var now = make([]int, rows*cols)
	for c:=0; c < rows*cols; c++ {
		if rand.Float64() < some {
			now[c] = 1
		} else {
			now[c] = 0
		}
	}
	live(now, rows, generations)
}

func live(a []int, r int, gen int) int {
	var b = make([]int, len(a))
	var neighbors int
	if gen < 1 {
		return 0
	}
	sleep()
	fmt.Println("\n \n Generation ", gen)
	for c:=0; c < len(a); c++ {
		if a[c] == 1 {
			fmt.Print("o")
		} else {
			fmt.Print(" ")
		}
		if c % r == 0 {
			fmt.Println()
		}
	}
	for c:=52; c <= len(a) - 51; c++ {
		neighbors = a[c-1] + a[c+1] + a[c-r-1] + a[c-r] + a[c-r+1] + a[c+r-1] + a[c+r] + a[c+51]
		b[c] = a[c]
		if a[c] == 0 {
			if neighbors == 3 {
				b[c] = 1
			} else {
				b[c] = 0
			}
		} else {
			if neighbors == 2 || neighbors == 3 { 
				b[c] = 1
			} else {
				b[c] = 0
			}
		}
	}
	return live(b, r, gen-1)
}

func sleep() {
	time.Sleep(1)
}

func main() {
	file, err := os.OpenFile("Attempts", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	log.SetOutput(file)

	log.Println("The Game of Life begins!")

	sleep()
	life(50, 20, 0.4, 3)

	log.Println("Successfully compiled!")
}