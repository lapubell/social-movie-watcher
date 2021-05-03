package main

import "testing"

func TestParseNewNameFromMessage(t *testing.T) {
	cases := map[string]string{
		"!!changeName:OLD::fun times!!::NEW::fun times!!":   "fun times",
		"!!changeName:OLD::fun times!!::NEW::fun times!!!!": "fun times!!",
		"!!changeName:NEW::fun times!!!!":                   "fun times!!",
		"":                                                  "(Lurker)",
		"NEW::":                                             "(Lurker)",
		"NEW:: ":                                            "(Lurker)",
		"NEW::    ":                                         "(Lurker)",
		"NEW::   !!   a!!":                                  "!!   a",
	}

	for msg, expected := range cases {
		if expected != parseNewNameFromMessage(msg) {
			t.Errorf("Expected `%v`, got `%v`", expected, parseNewNameFromMessage(msg))
		}
	}
}
