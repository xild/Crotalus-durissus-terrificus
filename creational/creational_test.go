package creational_test

import (
	"testing"

	"github.com/xild/Crotalus-durissus-terrificus/creational"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	s1 := creational.Singleton()
	s1["this"] = "that"

	s2 := creational.Singleton()
	assert.Equal(t, "that", s2["this"])
}
