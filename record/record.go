package record

import (
	"fmt"
	"time"
)

// Record is record of ko-bai data
type Record struct {
	Name string
	Date time.Time
}

// PutSlack is post slack
func (h *Record) PutSlack() {
	fmt.Printf("%v!", h)
}
