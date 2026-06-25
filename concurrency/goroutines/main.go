package main

import (
	"time"
	"fmt"
	"sync"
)

func ProcessSim(n int, wg *sync.WaitGroup) int {
	defer wg.Done()
	sq := n * n
	fmt.Printf("Processed number: %d, square: %d\n", n, sq)
	return sq
}

func WaitFor(wg *sync.WaitGroup) {
	wg.Wait()
	fmt.Println("All processes completed.")
}


func spellName(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, char := range name {
		wg.Add(1)
		fmt.Printf("%c\n", char)
		time.Sleep(500 * time.Millisecond)
	}
}


func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	spellName("Keith", &wg)
	
	go func() {
		go ProcessSim(25, &wg)
		wg.Add(1)
		time.Sleep(1 * time.Second)
		go ProcessSim(35, &wg)
		go ProcessSim(55, &wg)
		time.Sleep(2 * time.Second)
	}()


	
}
