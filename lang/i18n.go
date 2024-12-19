package lang

import (
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var Localiser *i18n.Localizer
var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("lang/en.json")
	bundle.LoadMessageFile("lang/ja.json")

	Localiser = i18n.NewLocalizer(bundle, language.English.String(), language.Japanese.String())
}
