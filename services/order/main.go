//go:generate echo -e "\033[1;34mGenerating sqlc...\033[0m"
//go:generate sqlc generate

package main

import "fmt"

func main() {
	fmt.Println("Order Service")
}
