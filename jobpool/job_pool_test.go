package jobpool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JobPool(t *testing.T) {
	t.Run("100 Jobs runs 100 Jobs", func(t *testing.T) {
		count := 0
		action := func() error {
			count++
			return nil
		}
		errHandler := func(err error) {}

		jp := NewJobPool(action, errHandler, 10, 100)

		jp.Start()

		assert.Equal(t, 100, count)
	})
}
