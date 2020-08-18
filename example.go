package ctime

import (
	"fmt"
	"os"
	"time"
)

func main() {

	const year, month, day, hour, min, sec, nsec = 2004, 10, 5, 16, 30, 44, 44

	ctime, err := NewTime(time.Date(year, month, day, hour, min, sec, nsec, time.FixedZone("", 0)))

	// The ctime and err return values from NewTime are
	if err == nil {
		fmt.Fprintf(os.Stderr, ctime.String(), err)
	}

	//Output:
	//
	//"2004-10-05T16:30 -> Tuesday"
}
