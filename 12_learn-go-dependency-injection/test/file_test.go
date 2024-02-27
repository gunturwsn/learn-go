package test

import (
	"github.com/stretchr/testify/assert"
	"learn-go-dependency-injection/simple"
	"testing"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Database")
	assert.NotNil(t, connection)

	cleanup()
}
