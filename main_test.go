package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getSheetName(t *testing.T) {
	testCases := []struct {
		fileName  string
		sheetName string
	}{
		{
			fileName:  "some/path/to/header.csv",
			sheetName: "header",
		},
		{
			fileName:  "some/path with spaces/to/header.csv",
			sheetName: "header",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.fileName, func(t *testing.T) {
			name := getSheetName(tc.fileName)
			assert.Equal(t, tc.sheetName, name)
		})
	}
}
