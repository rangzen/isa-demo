package logs

import "fmt"

var Stdout = func(s string) {
	fmt.Println(s)
}
