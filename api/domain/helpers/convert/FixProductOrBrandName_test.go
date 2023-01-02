package convert_test

import (
	"NandoFR/Webscraping/domain/helpers/convert"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixProductOrBrandName(t *testing.T) {

	c := convert.New()

	word := c.FixProductOrBrandeName("iphone 11 pro max ?")
	assert.Equal(t, "iPhone 11 Pro Max ?", word)
}
