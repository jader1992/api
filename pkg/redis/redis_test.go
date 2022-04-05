package redis

import (
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	fmt.Println(Redis)
	fmt.Println(Redis.Get("name"))
}
