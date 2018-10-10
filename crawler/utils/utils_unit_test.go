// +build unit

package utils

import (
	"github.com/ccdle12/Blocksage/crawler/test-utils"
	"github.com/stretchr/testify/suite"
	"testing"
)

// ===========================================================
// Testing Suite
// ===========================================================
type UtilsUnitSuite struct {
	suite.Suite
}

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestSuiteUtilsUnit(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(UtilsUnitSuite))
}

// ===========================================================
// Unit Tests
// ===========================================================
// TestConvNodeResToBlock is a function that will convert a NodeRespone to a Block.
func (suite *UtilsUnitSuite) TestConvNodeResToBlock() {
	nodeRes := testutils.NodeResponse
	block, err := ConvNodeResToBlock(nodeRes)

	suite.NoError(err, "There should not be an error when converting Node Response to a Block")
	suite.NotNil(block, "block should not be nil")
}

// TestEmptyString will test the util function if a string is empty/zero valued
func (suite *UtilsUnitSuite) TestEmptyString() {
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
		suite.Equal(eachTest.output, result, "The tests for empty string should match the output and results")
	}
}

// TestEmptyStringNotUsingSlices will test EmptyString(), sending comma separated arguments
func (suite *UtilsUnitSuite) TestEmptyStringNotUsingSlices() {

	tests := []struct {
		actual   bool
		expected bool
	}{
		{EmptyString("", "sdf", "dfdf", "", "123214"), true},
		{EmptyString("asdf", "sdf", "dfdf", "", "123214"), true},
		{EmptyString("asdf", "sdf", "dfdf", "fdsaf", "123214"), false},
		{EmptyString(""), true},
		{EmptyString("adsf"), false},
	}

	for _, eachTest := range tests {
		suite.Equal(eachTest.actual, eachTest.expected, "actual should match expected")
	}
}

// TestNodeClientAddressFormat will test that when passed different formats of addresses it will conform to the correct format.
func (suite *UtilsUnitSuite) TestAddressFormat() {
	tests := []struct {
		input    string
		expected string
	}{
		{"12345:8432", "http://12345:8432"},
		{"http://12343:8080", "http://12343:8080"},
	}

	for _, t := range tests {
		result := FormatAddress(t.input)
		suite.Equal(t.expected, result)
	}
}
