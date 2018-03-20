package math

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestGcd(t *testing.T) {
	assert.Equal(t, 3, Gcd(12, 15))
	assert.Equal(t, 23, Gcd(23, 23))
	assert.Equal(t, 1, Gcd(1, 550))
	assert.Equal(t, 16, Gcd(96, 16))
	assert.Equal(t, 55, Gcd(0, 55))
}


func TestLcm(t *testing.T) {
	assert.Equal(t, 12, Lcm([]int{12, 6}))
	assert.Equal(t, 48, Lcm([]int{12, 16}))
	assert.Equal(t, 130, Lcm([]int{65, 10, 5}))
	assert.Equal(t, 65, Lcm([]int{65}))
	assert.Equal(t, 31416, Lcm([]int{77, 56, 34, 21, 4}))
	assert.Equal(t, 4, Lcm([]int{4, 2, 1}))
}
