package time

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/beevik/ntp"
)

// GetCurrentTime ...
func GetCurrentTime() time.Time {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		_, err := io.WriteString(os.Stderr, err.Error())
		if err != nil {
			fmt.Printf("Error during stderr write: %s", err)
		}
		fmt.Println("Error getting current time:", err)
	}

	return time
}