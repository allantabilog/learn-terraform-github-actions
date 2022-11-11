package embedded

import (
	"fmt"
	"testing"
)

func TestInterface(t *testing.T) {

	var a = NewStudent()
	fmt.Println(a.GetName())
	fmt.Println(a.GetAge())
	fmt.Println(a.GetScore())
	fmt.Println(a.GetSchoolName())

}
