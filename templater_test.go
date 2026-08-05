package templater

import "testing"

func TestRender(t *testing.T) {
	data := map[string]string{"name": "Ada", "role": "engineer"}
	cases := map[string]string{
		"hi {{name}}":            "hi Ada",
		"{{ name }} the {{role}}": "Ada the engineer",
		"{{missing}}!":           "!", // unknown keys render empty
		"no placeholders":        "no placeholders",
	}
	for in, want := range cases {
		if got := Render(in, data); got != want {
			t.Errorf("Render(%q) = %q, want %q", in, got, want)
		}
	}
}
