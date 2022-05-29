package main

import (
	"fmt"
	"goft_redis/gedis"
	"time"
)

func main() {
	fmt.Println(gedis.NewStringOperation().Set("name", "hsezr", gedis.WithExpire(time.Second * 10)))
}
