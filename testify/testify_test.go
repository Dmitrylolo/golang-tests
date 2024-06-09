package testify_test

import (
	"testing"

	"github.com/Dmitrylolo/golang-tests/testify"
	"github.com/stretchr/testify/assert"
)

func TestFilterUnique(t *testing.T) {
	developers := []testify.Developer{
		{"John"},
		{"Jane"},
		{"John"},
		{"Jane"},
		{"Jack"},
		{"Jill"},
		{"Jack"},
	}
	expected := []string{"John", "Jane", "Jack", "Jill"}

	result := testify.FilterUnique(developers)

	assert.Equal(t, expected, result)

	// if !reflect.DeepEqual(result, expected) {
	// 	t.Fail()
	// }
}

func TestNegativeFilterUnique(t *testing.T) {
	developers := []testify.Developer{
		{"John"},
		{"Jane"},
		{"John"},
		{"Jane"},
		{"Jack"},
		{"Jill"},
		{"Jack"},
	}
	expected := []string{"John", "Jane"}

	result := testify.FilterUnique(developers)
	assert.NotEqual(t, expected, result)
}
