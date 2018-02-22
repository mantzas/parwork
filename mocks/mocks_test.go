package mocks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMockWork_ID(t *testing.T) {
	req := require.New(t)
	m := MockWork{}
	req.Equal("1", m.ID())
}

func TestMockWork_Do(t *testing.T) {
	req := require.New(t)
	m := MockWork{}
	req.NotPanics(m.Do)
}

func TestMockWork_GetError(t *testing.T) {
	req := require.New(t)
	m := MockWork{}
	req.NoError(m.GetError())
}

func TestGenerator(t *testing.T) {
	req := require.New(t)
	req.NotNil(Generator())
}

func TestReporter(t *testing.T) {
	req := require.New(t)
	req.NotPanics(func() { Reporter(MockWork{}) })
}
