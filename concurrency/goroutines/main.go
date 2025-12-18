package main

import (
	"fmt"
	"sync"
)

func StringWorker(s string, wg *sync.WaitGroup) string {
	defer wg.Done()
	// Simulate some processing
	result := fmt.Sprintf("Processed string: %s", s)
	return result
}


	var mutex = &sync.Mutex{}

func main() {

	go StringWorker("Hello, World!", &sync.WaitGroup{})
	var wg sync.WaitGroup
	wg.Add(1)
	
	mutex.Lock()
	result := StringWorker("Hello, World!", &wg)
	wg.Wait()
	fmt.Println(result)
	
}
