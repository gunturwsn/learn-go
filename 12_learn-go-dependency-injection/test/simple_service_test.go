package test

import (
	"fmt"
	"learn-go-dependency-injection/simple"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService := simple.InitializedService()
	fmt.Println(simpleService.SimpleRepository)
}
