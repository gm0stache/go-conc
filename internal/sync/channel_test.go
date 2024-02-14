package sync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	// arrange
	ch := NewChannel[int](5)

	ch.Send(1)
	ch.Send(2)

	// act
	recv1 := ch.Receive()
	recv2 := ch.Receive()

	// assert
	assert.Equal(t, recv1, 1)
	assert.Equal(t, recv2, 2)
}
