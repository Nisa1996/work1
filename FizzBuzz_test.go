package FizzBuzz

import (

	"testing"
	"github.com/stretchr/testify/assert"

)

func TestFizzBuzz(t *testing.T) {

    assert.Equal(t, FizzBuzz(1) ,"1")
		assert.Equal(t, FizzBuzz(2) ,"2")

    assert.Equal(t, FizzBuzz(3),"Fizz")
		assert.Equal(t, FizzBuzz(6),"Fizz")

    assert.Equal(t, FizzBuzz(5),"Buzz")
		assert.Equal(t, FizzBuzz(10),"Buzz")

    assert.Equal(t, FizzBuzz(15),"FizzBuzz")

}
