package add

import "testing"

func TestAddition(t *testing.T) {
	if Addition(4, 5) != 9 {
		t.Error("Expected 4 + 5 to equal 9")
	}
	if Addition(0, -1) != -1 {
		t.Error("Expected 0 + (-1) to equal -1")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		a        int
		b        int
		expected int
	}{
		{2, 2, 4},
		{-1, 1, 0},
		{0, 4, 4},
		{-5, -3, -8},
		{99999, 2, 100001},
	}

	for _, test := range tests {
		if output := Addition(test.a, test.b); output != test.expected {
			t.Error("Test Failed: {} a ,+, {}b, {} expected, recieved: {}", test.a, test.b, test.expected, output)
		}
	}
}
