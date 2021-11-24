package i18n

// TranslationSet is a set of localised strings for a given language
type TranslationSet struct {
	Confirm string
	Remove  string
	Return  string
	No      string
	Yes     string
}

func englishSet() TranslationSet {
	return TranslationSet{
		Confirm: "Confirm",
		Return:  "return",
		Remove:  "remove",
		No:      "no",
		Yes:     "yes",
	}
}
