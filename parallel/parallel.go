package parallel

import (
	"fmt"
	"math/rand"
	"play/timer"
	"time"
)

func AsyncFunc() (int, error) {
	rand.Seed(time.Now().UnixNano())

	result := rand.Intn(10) + 1

	timer.Sleep(result)

	if result >= 5 {
		return result, fmt.Errorf("over 5")
	}
	return result, nil
}
