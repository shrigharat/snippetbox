package main

import (
	"testing"
	"time"

	"snippetbox.shrishail.dev/internal/assert"
)

func TestHumanDate(t *testing.T) {
	testCases := []struct {
		name           string
		input          time.Time
		expectedOutput string
	}{
		{
			name:           "UTC",
			input:          time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC),
			expectedOutput: "17 Mar 2024 at 10:15",
		},
		{
			name:           "Empty",
			input:          time.Time{},
			expectedOutput: "",
		},
		{
			name:           "CET",
			input:          time.Date(2024, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			expectedOutput: "17 Mar 2024 at 09:15",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			output := humanDate(testCase.input)

			assert.Equal(t, output, testCase.expectedOutput)
		})
	}
}
