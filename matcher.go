package envflag

import "strings"

type matcher struct {
	ignoreCase bool
	exacts     stringSlice
	prefixes   stringSlice
}

func (m *matcher) addExact(exact string) {
	m.exacts = append(m.exacts, exact)
}

func (m *matcher) addPrefix(prefix string) {
	m.prefixes = append(m.prefixes, prefix)
}

func (m *matcher) hasPattern() bool {
	return !m.exacts.isEmpty() || !m.prefixes.isEmpty()
}

func (m *matcher) match(target string) bool {
	if m.ignoreCase {
		target = strings.ToLower(target)
	}
	matchExact := func(s string) bool {
		if m.ignoreCase {
			s = strings.ToLower(s)
		}
		return target == s
	}
	if m.exacts.exists(matchExact) {
		return true
	}
	matchPrefix := func(s string) bool {
		if m.ignoreCase {
			s = strings.ToLower(s)
		}
		return strings.HasPrefix(target, s)
	}
	if m.prefixes.exists(matchPrefix) {
		return true
	}
	return false
}

type stringSlice []string

func (ss stringSlice) exists(f func(s string) bool) bool {
	for _, a := range ss {
		if f(a) {
			return true
		}
	}
	return false
}

func (ss stringSlice) isEmpty() bool {
	return len(ss) == 0
}
