package test

import (
	"fmt"
	"learn-go-dependency-injection/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializedService()
	fmt.Println(err)
	fmt.Println(simpleService.SimpleRepository)
}
