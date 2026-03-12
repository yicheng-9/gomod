package greet
import "fmt"
func Greet(name string) string {
	fmt.Println("Greeting", name)
	return "Hello, " + name
}