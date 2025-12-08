// Package generators demonstrates simple generator functions.
package generators


var resume chan int

func Integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()
	return yield
}

func GenerateInteger() int {
	return <-resume
}


