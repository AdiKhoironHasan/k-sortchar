package main

import "testing"

type sortCharTest struct {
	argument, vowels, consonant string
}

var sortCharTests = []sortCharTest{
	sortCharTest{"Sample Case", "aaee", "ssmplc"},
	sortCharTest{"Next Case", "eea", "nxtcs"},
}

func TestSortChar(t *testing.T) {
	for i := 0; i < 100000; i++ {
		for _, test := range sortCharTests {
			if outputVow, outputCon := sortChar(test.argument); outputVow+outputCon != test.vowels+test.consonant {
				t.Errorf("Output %q not equal to expected %q", outputVow+outputCon, test.vowels+test.consonant)
			}
		}
	}
}

func BenchmarSortChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortChar("Sample Case")
	}
}
