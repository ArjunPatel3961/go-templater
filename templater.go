// Package templater renders {{key}} placeholders, dependency-free.
package templater

import "regexp"

var varRe = regexp.MustCompile(`{{\s*(\w+)\s*}}`)

func Render(tmpl string, data map[string]string) string {
	return varRe.ReplaceAllStringFunc(tmpl, func(m string) string {
		key := varRe.FindStringSubmatch(m)[1]
		return data[key]
	})
}
