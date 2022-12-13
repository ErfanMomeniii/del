package del

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEventWithoutListener(t *testing.T) {
	is := assert.New(t)
	e := NewEvent("first")

	is.Equal(e, Event{name: "first"})
}
