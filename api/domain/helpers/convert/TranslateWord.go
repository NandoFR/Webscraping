package convert

import "strings"

func (c Convert) TranslateWord(text string) string {

	text = strings.ToLower(text)

	x := func(old string, replace string) {
		oldLower := strings.ToLower(old)
		replaceLower := strings.ToLower(replace)

		if strings.Contains(text, oldLower) {
			text = strings.Replace(text, oldLower, replaceLower, -1)
		}
	}

	x("computador", "pc")
	x("?iphone", "iphone")

	return strings.Title(text)
}
