// +build unit

package utils

import (
	"github.com/ccdle12/Blocksage/go-crawler/test-utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestConvNodeResToBlock is a function that will convert a NodeRespone to a Block.
func TestConvNodeResToBlock(t *testing.T) {
	assert := assert.New(t)

	nodeRes := testutils.NodeResponse

	block, err := ConvNodeResToBlock(nodeRes)

	assert.NoError(err, "There should not be an error when converting Node Response to a Block")
	assert.NotNil(block, "block should not be nil")
}

// TestEmptyString will test the util function if a string is empty/zero valued
func TestEmptyString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  []string
		output bool
	}{
		{[]string{"hello"}, false},
		{[]string{""}, true},
		{[]string{"3421"}, false},
		{[]string{""}, true},
		{[]string{""}, true},
		{[]string{"21321", "asdsa", "Sdf", "Fff"}, false},
		{[]string{"21321", "asdsa", "", "Fff"}, true},
		{[]string{"", "asdsa", "", "Fff"}, true},
		{[]string{"", "", "", "Fff"}, true},
		{[]string{"", "", "", ""}, true},
		{[]string{"", ""}, true},
	}

	for _, eachTest := range tests {
		result := EmptyString(eachTest.input...)
		assert.Equal(eachTest.output, result, "The tests for empty string should match the output and results")
	}
}

// TestEmptyStringNotUsingSlices will test EmptyString(), sending comma separated arguments
func TestEmptyStringNotUsingSlices(t *testing.T) {
	assert := assert.New(t)

	result1 := EmptyString("", "sdf", "dfdf", "", "123214")
	result2 := EmptyString("asdf", "sdf", "dfdf", "", "123214")
	result3 := EmptyString("asdf", "sdf", "dfdf", "fdsaf", "123214")
	result4 := EmptyString("")
	result5 := EmptyString("adsf")

	assert.True(result1, "result1 should be true since we have sent empty strings to the function")
	assert.True(result2, "result2 should be true since we have sent empty strings to the function")
	assert.False(result3, "result3 should be false since we have sent empty strings to the function")
	assert.True(result4, "result4 should be true since we have sent empty strings to the function")
	assert.False(result5, "result5 should be false since we have sent empty strings to the function")
}
