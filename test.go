package main

import (
	"fmt"
	"goft_redis/gedis"
)

func main() {
	fmt.Println(gedis.NewStringOperation().Get("name").Unwrap())

}
