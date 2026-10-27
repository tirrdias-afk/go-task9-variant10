package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Now компьютердің ағымдағы жергілікті уақытын алады.
	now := time.Now()
	fmt.Println("Ағымдағы уақыт:", now.Format("2006-01-02 15:04:05 -07:00"))
}
