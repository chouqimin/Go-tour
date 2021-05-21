package word

import "strings"
import "unicode"

// 單字全部轉為大寫
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 單字全部轉為小寫
func ToLower(s string) string {
	return strings.ToLower(s)
}

// 底線單字轉大寫駝峰單字
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.ToLower(s)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// 底線單字轉小寫駝峰單字
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	message := []rune(s)
	return strings.ToLower(string(message[0])) + string(message[1:])
}

// 駝峰單字轉底線單字
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		} else if unicode.IsUpper(r) {
			output = append(output, '_', unicode.ToLower(r))
		} else {
			output = append(output, r)
		}
	}
	return string(output)
}
