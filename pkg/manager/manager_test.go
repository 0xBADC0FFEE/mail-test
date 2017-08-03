package manager

import "testing"
import "time"

func TestManager(t *testing.T) {
	maxWorkers := 5
	jobs := 50

	results := make(chan struct{})

	manager := New(maxWorkers)

	max := 0
	current := 0
	completed := 0

	manager.Wake()

	go func() {
		for i := 1; i < jobs+1; i++ {

			func(i int) {

				manager.Add(func() {
					current++
					if max < current {
						max = current
					}
					time.Sleep(time.Millisecond * 100)
					current--
					completed++
					t.Logf("completed #%d, current %d, max %d", i, current, max)
					results <- struct{}{}
				})

			}(i)

		}
	}()

	resultsCount := 0

	for range results {
		resultsCount++

		if resultsCount >= jobs {

			if max != maxWorkers {
				t.Errorf("Incorrect concurrency. Expected %d got %d", maxWorkers, max)
			}

			if completed < jobs {
				t.Errorf("Not all jobs completed. Expected %d got %d", jobs, completed)
			}

			return
		}

	}

}
