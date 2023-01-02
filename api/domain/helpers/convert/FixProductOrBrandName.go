package convert

import "strings"

func (c Convert) FixProductOrBrandeName(text string) string {

	text = strings.Title(text)
	x := func(old string, replace string) {
		oldTitle := strings.Title(old)
		if strings.Contains(text, oldTitle) {
			text = strings.Replace(text, oldTitle, replace, -1)
		}
	}

	x("iphone", "iPhone")
	x("pc", "PC")

	return text
}
