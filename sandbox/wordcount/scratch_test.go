package wordcount

import (
	"reflect"
	"testing"
)

func TestBuildWordCountMap(t *testing.T){

	type test struct {
		input string 
		expected map[string] int
	}

	tests := []test {
		{input: "hello world", expected: map[string] int {"hello":1, "world": 1}},
		{input: "hello hello world", expected: map[string] int {"hello":2, "world": 1}},
		{input: "hello there there world world world", expected: map[string] int {"hello":1, "world": 3, "there": 2}},
		{input: "hello world", expected: map[string] int {"hello":1, "world": 1}},
	}


	for _, testCase := range tests {
		got := BuildWordCountMap(testCase.input)
		if !reflect.DeepEqual(got, testCase.expected) {
			t.Errorf("got %v, wanted %v", got, testCase.expected)
		}
		}
	}
	
