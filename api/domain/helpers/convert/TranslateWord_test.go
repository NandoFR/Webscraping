package convert_test

import (
	"NandoFR/Webscraping/domain/helpers/convert"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslatorWordFromGivenText(t *testing.T) {

	c := convert.New()
	word := c.TranslateWord("Computador I5 Lenovo (Barato) - vem nas propostas")
	assert.Equal(t, "Pc I5 Lenovo (Barato) - Vem Nas Propostas", word)

	word = c.TranslateWord("?iphone 11 pro max ?")
	assert.Equal(t, "Iphone 11 Pro Max ?", word)
}
