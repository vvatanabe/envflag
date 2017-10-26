package envflag

import (
	"testing"
	"os"
)

func TestGetEnviron(t *testing.T) {

	tests := []struct {
		key  string
		value          string
	}{
		{"ENV_FLAG_INT", "1"},
		{"ENV_FLAG_STR", "hello"},
		{"ENV_FLAG_BOOL", "true"},
	}

	for _, tc := range tests {
		os.Setenv(tc.key, tc.value)
	}

	environ := getEnviron()

	for _, tc := range tests {
		if v, ok := environ[tc.key]; !ok {
			t.Errorf("not exists env key: %s", tc.key)
		} else if v != tc.value {
			t.Errorf("not match env value: %s", tc.value)
		}
	}

}
