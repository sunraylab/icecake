package lib

import "testing"

func TestParseAttributes(t *testing.T) {

	tests := []struct {
		in   string
		want string
	}{
		{in: "a", want: "a"},
		{in: "attr", want: "attr"},
		{in: "  attr1  attr2  ", want: "attr1 attr2"},
		{in: "attr1='val1'", want: "attr1='val1'"},
		{in: "attr1 = val1", want: "attr1='val1'"},
		{in: "attr1  =  ' val1 ' ", want: "attr1=' val1 '"},
		{in: "attr1=''", want: "attr1"},
		{in: "attr1='val\"ue'", want: "attr1='val\"ue'"},
		{in: "attr1='val\"ue' attr2", want: "attr1='val\"ue' attr2"},
		{in: "attr1 attr2='val2' attr3 attr4='val4'", want: "attr1 attr2='val2' attr3 attr4='val4'"},

		{in: "", want: ""},
		{in: "=18", want: "unexpected starting char '=' at pos 0"},
		{in: "attr1='va'lue", want: "quoted value must be separated with a space with the next attribute"},
		{in: "attr1=", want: "missing value after '='"},
		{in: "attr1='value", want: "missing ending quote in value starting at pos 7"},
		{in: "attr1='value attr2", want: "missing ending quote in value starting at pos 7"},
	}

	for i, tst := range tests {
		tks, err := ParseAttributes(tst.in)
		if err == nil {
			if tks.String() != tst.want {
				t.Errorf("%d failed. out: %q, want:%s", i, tks.String(), tst.want)
			}
		} else {
			if err.Error() != tst.want {
				t.Errorf("%d failed. out: %q, want:%s", i, err.Error(), tst.want)
			}
		}
	}

}
