package convert

import "strings"

func (c Convert) TranslateWord(text string) string {

	text = strings.ToLower(text)

	x := func(olds []string, replace string) {
		replaceLower := strings.ToLower(replace)
		for _, v := range olds {
			oldLower := strings.ToLower(v)
			if strings.Contains(text, oldLower) {
				text = strings.Replace(text, oldLower, replaceLower, -1)
			}
		}
	}

	x([]string{"computador"}, "pc")
	x([]string{"?iphone"}, "iphone")
	x([]string{"televisao", "televisão"}, "tv")

	return strings.Title(text)
}
