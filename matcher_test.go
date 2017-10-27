package envflag

import "testing"

func TestMatcherMatchPrefix(t *testing.T) {
	matcher := &matcher{
		false,
		[]string{},
		[]string{},
	}

	prefixes := []string{"pre_1", "pre_2"}

	for _, v := range prefixes {
		matcher.addPrefix(v)
	}

	if !matcher.match("pre_1_value") {
		t.Fatalf("not match: %s", "pre_1_value")
	}
	if !matcher.match("pre_2_value") {
		t.Fatalf("not match: %s", "pre_2_value")
	}
}

func TestMatcherMatchExact(t *testing.T) {
	matcher := &matcher{
		false,
		[]string{},
		[]string{},
	}

	exacts := []string{"exact_1", "exact_2"}

	for _, v := range exacts {
		matcher.addExact(v)
	}

	if !matcher.match("exact_1") {
		t.Fatalf("not match: %s", "exact_1")
	}
	if !matcher.match("exact_2") {
		t.Fatalf("not match: %s", "exact_2")
	}
}

func TestMatcherHasPattern(t *testing.T) {
	matcher := &matcher{
		false,
		[]string{},
		[]string{},
	}
	matcher.addExact("test_1")
	if !matcher.hasPattern() {
		t.Fatal("not has match pattern")
	}
}
