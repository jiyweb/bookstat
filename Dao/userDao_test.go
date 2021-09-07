package dao

import (
	"fmt"
	"testing"
)

func TestCheckUserNameAndPassword(t *testing.T) {
	CheckUserNameAndPassword("admin99", "123456")

	fmt.Println("user=", user)
}
