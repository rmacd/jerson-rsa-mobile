package rsa

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFastRSA_Generate1024(t *testing.T) {

	instance := NewFastRSA()
	output, err := instance.Generate(1024)
	assert.NoError(t, err)

	t.Log("output:", output)
}
func TestFastRSA_Generate2048(t *testing.T) {

	instance := NewFastRSA()
	output, err := instance.Generate(2048)
	assert.NoError(t, err)

	t.Log("output:", output)
}
func TestFastRSA_Generate4096(t *testing.T) {

	instance := NewFastRSA()
	output, err := instance.Generate(4096)
	assert.NoError(t, err)

	t.Log("output:", output)
}
