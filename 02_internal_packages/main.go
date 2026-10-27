package main

import (
	"fmt"
	"internal10/internal/greetings"
	"internal10/internal/mathutils"
)

func main() {
	fmt.Println(greetings.Hello("Диас"))
	fmt.Printf("10 + 5 = %d\n", mathutils.Add(10, 5))
	fmt.Printf("10 * 5 = %d\n", mathutils.Multiply(10, 5))
	fmt.Println("Екі internal пакет бір модульден қолданылды.")
}
