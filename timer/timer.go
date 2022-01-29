package timer

import (
	"fmt"
	"strconv"
	"time"
)

func Sleep(n int) {
	timer := time.NewTimer(time.Duration(n) * time.Second)

	fmt.Println(strconv.Itoa(n) + " second wait")

	<-timer.C

	fmt.Println("finish")
}
