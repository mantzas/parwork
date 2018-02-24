package mocks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMockWork_ID(t *testing.T) {
	require := require.New(t)
	m := MockWork{}
	require.Equal("1", m.ID())
}

func TestMockWork_Do(t *testing.T) {
	require := require.New(t)
	m := MockWork{}
	require.NotPanics(m.Do)
}

func TestMockWork_GetError(t *testing.T) {
	require := require.New(t)
	m := MockWork{}
	require.NoError(m.GetError())
}

func TestGenerator(t *testing.T) {
	require := require.New(t)
	require.NotNil(Generator())
}

func TestReporter(t *testing.T) {
	require := require.New(t)
	require.NotPanics(func() { Reporter(MockWork{}) })
}
