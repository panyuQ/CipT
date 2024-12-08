package NoKey

import (
	"fmt"
	"testing"
)

func TestGetMethods(t *testing.T) {
	fmt.Println(GetMethods(true))
	fmt.Println(GetMethods(false))
}
