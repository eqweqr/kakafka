package translation_test

import (
	"testing"

	"github.com/eqweqr/kafka/translation"
)

func TestTranslate(t *testing.T) {
	// // Arrange
	// word := "hello"
	// language := "english"

	// // Act
	// res := translation.Translate(word, language)

	// // Assert
	// if res != "hello" {
	// 	t.Errorf(`expected "hello" but received "%s"`, res)
	// }

	// Arrange
	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{
			Word:        "hello",
			Language:    "english",
			Translation: "hello",
		},
		{
			Word:        "hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "finnish",
			Translation: "hei",
		},
		{
			Word:        "hello",
			Language:    "dutch",
			Translation: "",
		},
		{
			Word:        "bye",
			Language:    "german",
			Translation: "",
		},
		{
			Word:        "Hello",
			Language:    "german",
			Translation: "hallo",
		},
	}

	for _, test := range tt {
		// Act
		res := translation.Translate(test.Word, test.Language)

		// Assert
		if res != test.Translation {
			t.Errorf(
				`exptected "%s" to be "%s" from "%s" but received "%s"`,
				test.Word, test.Language, test.Translation, res)
		}
	}

}
