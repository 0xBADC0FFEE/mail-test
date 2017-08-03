package manager

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestManager(t *testing.T) {
	maxWorkers := 5
	jobs := 500

	results := make(chan struct{})

	manager := New(maxWorkers)

	var max int32 = 0
	var current int32 = 0

	manager.Wake()

	go func() {
		for i := 1; i < jobs+1; i++ {

			func(i int) {

				manager.Add(func() {
					atomic.AddInt32(&current, 1)

					if atomic.LoadInt32(&max) < atomic.LoadInt32(&current) {
						atomic.StoreInt32(&max, atomic.LoadInt32(&current))
					}

					time.Sleep(time.Millisecond * 100)
					t.Logf("completed #%d, current %d, max %d", i, current, max)
					atomic.AddInt32(&current, -1)
					results <- struct{}{}
				})

			}(i)

		}
	}()

	resultsCount := 0

	for range results {
		resultsCount++

		if resultsCount >= jobs {

			if max != int32(maxWorkers) {
				t.Errorf("Incorrect concurrency. Expected %d got %d", maxWorkers, max)
			}

			return
		}

	}

}
