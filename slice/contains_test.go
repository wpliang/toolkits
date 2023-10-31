package slice

import "testing"

type TestCase[T any] struct {
	paramSlice    []T
	value         T
	expectedValue bool
}

func TestContainsString(t *testing.T) {
	testStringCases := []TestCase[string]{
		{
			paramSlice:    []string{"a", "b"},
			value:         "a",
			expectedValue: true,
		},
		{
			paramSlice:    []string{"a", "b"},
			value:         "c",
			expectedValue: false,
		},
	}
	for _, testStringCase := range testStringCases {
		t.Run("test string slice", func(t *testing.T) {
			result := ContainsString(testStringCase.paramSlice, testStringCase.value)
			if result != testStringCase.expectedValue {
				t.Errorf("expected != result:\n%v\n%v", testStringCase.expectedValue, result)
			}
		})
	}
}

func TestContains(t *testing.T) {
	testStringCases := []TestCase[string]{
		{
			paramSlice:    []string{"a", "b"},
			value:         "a",
			expectedValue: true,
		},
		{
			paramSlice:    []string{"a", "b"},
			value:         "c",
			expectedValue: false,
		},
	}
	for _, testStringCase := range testStringCases {
		t.Run("test string slice", func(t *testing.T) {
			result := Contains(testStringCase.paramSlice, testStringCase.value)
			if result != testStringCase.expectedValue {
				t.Errorf("expected != result:\n%v\n%v", testStringCase.expectedValue, result)
			}
		})
	}

	testIntCases := []TestCase[int]{
		{
			paramSlice:    []int{1, 2},
			value:         1,
			expectedValue: true,
		},
		{
			paramSlice:    []int{1, 2},
			value:         3,
			expectedValue: false,
		},
	}
	for _, testCase := range testIntCases {
		t.Run("test int slice", func(t *testing.T) {
			result := Contains(testCase.paramSlice, testCase.value)
			if result != testCase.expectedValue {
				t.Errorf("expected != result:\n%v\n%v", testCase.expectedValue, result)
			}
		})
	}

	testNilCases := []TestCase[int]{
		{
			paramSlice:    nil,
			value:         1,
			expectedValue: false,
		},
	}
	for _, testCase := range testNilCases {
		t.Run("test nil slice", func(t *testing.T) {
			result := Contains(testCase.paramSlice, testCase.value)
			if result != testCase.expectedValue {
				t.Errorf("expected != result:\n%v\n%v", testCase.expectedValue, result)
			}
		})
	}
}

type TestSliceCase[T any] struct {
	source        []T
	target        []T
	expectedValue bool
}

func TestHasCommon(t *testing.T) {
	testStringCases := []TestSliceCase[string]{
		{
			source:        []string{"a", "b"},
			target:        []string{"a", "c"},
			expectedValue: true,
		},
		{
			source:        []string{"a", "b"},
			target:        []string{"c", "d"},
			expectedValue: false,
		},
	}
	for _, testCase := range testStringCases {
		t.Run("test string slice", func(t *testing.T) {
			result := HasCommon(testCase.source, testCase.target)
			if result != testCase.expectedValue {
				t.Errorf("expected != result:\n%v\n%v", testCase.expectedValue, result)
			}
		})
	}

	testNilCases := []TestSliceCase[string]{
		{
			source:        nil,
			target:        []string{"a", "c"},
			expectedValue: false,
		},
		{
			source:        []string{"a", "b"},
			target:        nil,
			expectedValue: false,
		},
		{
			source:        nil,
			target:        nil,
			expectedValue: false,
		},
	}
	for _, testCase := range testNilCases {
		t.Run("test nil slice", func(t *testing.T) {
			result := HasCommon(testCase.source, testCase.target)
			if result != testCase.expectedValue {
				t.Errorf("expected != result:\n%v\n%v", testCase.expectedValue, result)
			}
		})
	}
}

func TestEquals(t *testing.T) {
	testStringCases := []TestSliceCase[string]{
		{
			source:        []string{"a", "b"},
			target:        []string{"a", "c"},
			expectedValue: false,
		},
		{
			source:        []string{"a", "b"},
			target:        []string{"a", "b"},
			expectedValue: true,
		},
	}
	for _, testCase := range testStringCases {
		t.Run("test string slice", func(t *testing.T) {
			result := Equals(testCase.source, testCase.target)
			if result != testCase.expectedValue {
				t.Errorf("expected != result:\n%v\n%v", testCase.expectedValue, result)
			}
		})
	}

	testNilCases := []TestSliceCase[string]{
		{
			source:        nil,
			target:        []string{"a", "c"},
			expectedValue: false,
		},
		{
			source:        []string{"a", "b"},
			target:        nil,
			expectedValue: false,
		},
		{
			source:        nil,
			target:        nil,
			expectedValue: true,
		},
	}
	for _, testCase := range testNilCases {
		t.Run("test nil slice", func(t *testing.T) {
			result := Equals(testCase.source, testCase.target)
			if result != testCase.expectedValue {
				t.Errorf("expected != result:\n%v\n%v", testCase.expectedValue, result)
			}
		})
	}
}

type TestDiffCase[T any] struct {
	source        []T
	target        []T
	expectedValue []T
}

func TestDiffSlice(t *testing.T) {
	testStringCases := []TestDiffCase[string]{
		{
			source:        []string{"a", "b"},
			target:        []string{"a", "c"},
			expectedValue: []string{"b"},
		},
		{
			source:        []string{"a", "b"},
			target:        []string{"c", "d"},
			expectedValue: []string{"a", "b"},
		},
	}
	for _, testCase := range testStringCases {
		t.Run("test string slice", func(t *testing.T) {
			result := DiffSlice(testCase.source, testCase.target)
			if !Equals(result, testCase.expectedValue) {
				t.Errorf("expected != result:\n%v\n%v", testCase.expectedValue, result)
			}
		})
	}

	testNilCases := []TestDiffCase[string]{
		{
			source:        nil,
			target:        []string{"a", "c"},
			expectedValue: nil,
		},
		{
			source:        []string{"a", "b"},
			target:        nil,
			expectedValue: []string{"a", "b"},
		},
		{
			source:        nil,
			target:        nil,
			expectedValue: nil,
		},
	}
	for _, testCase := range testNilCases {
		t.Run("test nil slice", func(t *testing.T) {
			result := DiffSlice(testCase.source, testCase.target)
			if !Equals(result, testCase.expectedValue) {
				t.Errorf("expected != result:\n%v\n%v", testCase.expectedValue, result)
			}
		})
	}
}
