package learn_go_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0

	for i := 1; i <= 1000; i++ {
		var group sync.WaitGroup
		group.Add(1)

		go func() {
			defer group.Done()
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
			}
		}()

		group.Wait()
	}
	fmt.Println("Counter =", x)
}
