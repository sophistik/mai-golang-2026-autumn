package timedemo

import (
	"fmt"
	"time"
)

func demo() error {
	// ЭТАЛОН: 1 2 3 4 5 6 7 = месяц день час мин сек год зона
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	t, err := time.Parse(time.RFC3339, "2026-10-07T12:00:00Z")
	if err != nil {
		return err
	}
	fmt.Println(t.Format(time.RFC3339)) // json пишет time.Time так же
	return nil
}
