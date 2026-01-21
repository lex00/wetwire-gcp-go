// Test fixture: Non-resource variables that should be ignored
package testdata

import "fmt"

var Name = "not-a-resource"

var Count = 42

var Config = map[string]interface{}{
	"key": "value",
}

func DoSomething() {
	fmt.Println("This is a function, not a resource")
}
