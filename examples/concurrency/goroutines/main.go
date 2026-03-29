package main

import (
	"time"
	"fmt"
	"sync"
)

func  ProcessSim(n int, wg *sync.WaitGroup) int {
	defer wg.Done()
	sq := n * n
	fmt.Printf("Processed number: %d, square: %d\n", n, sq)
	return sq
}


func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	
	go func() {
		go ProcessSim(25, &wg)
		wg.Add(1)
		time.Sleep(1 * time.Second)
		go ProcessSim(35, &wg)
		go ProcessSim(55, &wg)
		time.Sleep(2 * time.Second)
	}()
	wg.Wait()
	fmt.Println("All goroutines complete.")


	
}
