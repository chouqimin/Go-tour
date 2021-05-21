package word

import "testing"

type testData struct {
	before string
	after  string
}

// 單字全部轉為大寫
func TestToUpper(t *testing.T) {
	tests := []testData{
		testData{before: "abc", after: "ABC"},
		testData{before: "DEF", after: "DEF"},
		testData{before: "abcDEF", after: "ABCDEF"},
		testData{before: "你好嗎?", after: "你好嗎?"},
	}
	for _, test := range tests {
		got := ToUpper(test.before)
		if got != test.after {
			t.Errorf("給的值為\"%s\"，想要的值為\"%s\", 但得到的值為\"%s\"", test.before, test.after, got)
		}
	}
}

// 單字全部轉為小寫
func TestToLower(t *testing.T) {
	tests := []testData{
		testData{before: "abc", after: "abc"},
		testData{before: "DEF", after: "def"},
		testData{before: "abcDEF", after: "abcdef"},
		testData{before: "你好嗎?", after: "你好嗎?"},
	}
	for _, test := range tests {
		got := ToLower(test.before)
		if got != test.after {
			t.Errorf("給的值為\"%s\"，想要的值為\"%s\", 但得到的值為\"%s\"", test.before, test.after, got)
		}
	}
}

// 底線單字轉大寫駝峰單字
func TestUnderscoreToUpperCamelCase(t *testing.T) {
	tests := []testData{
		testData{before: "abc", after: "Abc"},
		testData{before: "DEF", after: "Def"},
		testData{before: "abc_def", after: "AbcDef"},
		testData{before: "abc_DEF", after: "AbcDef"},
		testData{before: "abc_DEF", after: "AbcDef"},
		testData{before: "你好嗎?", after: "你好嗎?"},
	}
	for _, test := range tests {
		got := UnderscoreToUpperCamelCase(test.before)
		if got != test.after {
			t.Errorf("給的值為\"%s\"，想要的值為\"%s\", 但得到的值為\"%s\"", test.before, test.after, got)
		}
	}
}

// 底線單字轉小寫駝峰單字
func TestUnderscoreToLowerCamelCase(t *testing.T) {
	tests := []testData{
		testData{before: "abc", after: "abc"},
		testData{before: "DEF", after: "def"},
		testData{before: "abc_def", after: "abcDef"},
		testData{before: "abc_DEF", after: "abcDef"},
		testData{before: "abc_DEF", after: "abcDef"},
		testData{before: "你好嗎?", after: "你好嗎?"},
	}
	for _, test := range tests {
		got := UnderscoreToLowerCamelCase(test.before)
		if got != test.after {
			t.Errorf("給的值為\"%s\"，想要的值為\"%s\", 但得到的值為\"%s\"", test.before, test.after, got)
		}
	}
}

// 駝峰單字轉底線單字
func TestCamelCaseToUnderscore(t *testing.T) {
	tests := []testData{
		testData{before: "abc", after: "abc"},
		testData{before: "DEF", after: "d_e_f"},
		testData{before: "abcDef", after: "abc_def"},
		testData{before: "abcDEF", after: "abc_d_e_f"},
		testData{before: "AbcDEF", after: "abc_d_e_f"},
		testData{before: "你好嗎?", after: "你好嗎?"},
	}
	for _, test := range tests {
		got := CamelCaseToUnderscore(test.before)
		if got != test.after {
			t.Errorf("給的值為\"%s\"，想要的值為\"%s\", 但得到的值為\"%s\"", test.before, test.after, got)
		}
	}
}
