package greet
import "fmt"
//打招呼
func Greet(name string) string {
	fmt.Println("Greeting", name)
	return "Hello, " + name
}
