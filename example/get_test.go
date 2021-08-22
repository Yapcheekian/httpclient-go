package example

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	endpoints, err := GetEndpoints()

	fmt.Println(err)
	fmt.Println(endpoints)
}
