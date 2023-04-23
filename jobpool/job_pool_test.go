package jobpool

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JobPool(t *testing.T) {
	t.Run("100 Jobs runs 100 Jobs", func(t *testing.T) {
		count := 0
		var mu sync.Mutex
		action := func() error {
			mu.Lock()
			count++
			mu.Unlock()
			return nil
		}
		errHandler := func(err error) {}

		jp := NewJobPool(action, errHandler, 10, 100)

		jp.Start()

		assert.Equal(t, 100, count)
	})
}
