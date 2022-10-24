package time

import (
	"fmt"
	"testing"
)

func TestGetCurrentTime(t *testing.T) {
	t.Run("Test get current time", func(t *testing.T) {
		fmt.Println(GetCurrentTime())
	})
}