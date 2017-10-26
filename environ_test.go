package envflag

import "testing"

func TestStrToBool(t *testing.T) {

	tests := []struct {
		value string
		want  bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tc := range tests {
		if boolean, err := strToBool(tc.value); err != nil {
			t.Errorf("err: %s", err.Error())
		} else if boolean != tc.want {
			t.Fatalf("fatal want: %v, result: %v", tc.want, boolean)
		}
	}
}
