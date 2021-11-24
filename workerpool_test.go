package workerpool_test

import (
	"fmt"
	"log"

	"github.com/magodo/workerpool"
)

func ExampleWorkerPool() {
	wp := workerpool.NewWorkPool(2)
	sum := 0
	wp.Run(func(i interface{}) error {
		sum += i.(int)
		return nil
	})
	for i := 0; i < 10; i++ {
		i := i
		wp.AddTask(func() (interface{}, error) { return i, nil })
	}
	if err := wp.Done(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
	// Output: 45
}
