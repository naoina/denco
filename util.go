package denco

// NextSeparator returns an index of next separator in path.
func NextSeparator(path string, start int) int {
	for start < len(path) && path[start] != '/' && path[start] != '.' {
		start++
	}
	return start
}

// isMetaChar returns whether the meta character.
func IsMetaChar(c byte) bool {
	return c == ParamCharacter || c == WildcardCharacter
}

// ParamNames returns parameter names in given path.
// It returns names which meta character is prefixed.
func ParamNames(path string) (names []string) {
	for i := 0; i < len(path); i++ {
		if IsMetaChar(path[i]) {
			next := NextSeparator(path, i+1)
			names = append(names, path[i:next])
			i = next
		}
	}
	return names
}
