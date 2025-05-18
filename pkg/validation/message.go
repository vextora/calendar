package validation

import (
	"fmt"
)

var errorMessages = map[string]string{
	"required":  "%s wajib diisi",
	"minLength": "%s belum memenuhi panjang minimal %s karakter",
	"minWord":   "%s belum memenuhi panjang minimal %s kata",
	"email":     "Pastikan format %s sudah benar",
}

func GetErrorMessage(field, tag, param string) string {
	if template, ok := errorMessages[tag]; ok {
		if param != "" && countPlaceholders(template) == 2 {
			return fmt.Sprintf(template, field, param)
		}
		return fmt.Sprintf(template, field)
	}
	return ""
}

func countPlaceholders(s string) int {
	count := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '%' && s[i+1] == 's' {
			count++
		}
	}
	return count
}
