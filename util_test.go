package denco_test

import (
	"reflect"
	"testing"

	"github.com/naoina/denco"
)

func TestNextSeparator(t *testing.T) {
	for _, testcase := range []struct {
		path     string
		start    int
		expected interface{}
	}{
		{"/path/to/route", 0, 0},
		{"/path/to/route", 1, 5},
		{"/path/to/route", 9, 14},
		{"/path.html", 1, 5},
		{"/foo/bar.html", 1, 4},
		{"/foo/bar.html/baz.png", 5, 8},
		{"/foo/bar.html/baz.png", 10, 13},
	} {
		actual := denco.NextSeparator(testcase.path, testcase.start)
		expected := testcase.expected
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("path = %q, start = %v expect %v, but %v", testcase.path, testcase.start, expected, actual)
		}
	}
}

func TestIsMetaChar(t *testing.T) {
	for _, c := range []byte{':', '*'} {
		if !denco.IsMetaChar(c) {
			t.Errorf("Expect %q is meta charcter, but isn't", c)
		}
	}
	for c := byte(0); c < 0xff && c != ':' && c != '*'; c++ {
		if denco.IsMetaChar(c) {
			t.Errorf("Expect %q is not meta character, but isn't", c)
		}
	}
}

func TestParamNames(t *testing.T) {
	for path, expected := range map[string][]string{
		"/:a":    {":a"},
		"/:b":    {":b"},
		"/:a/:b": {":a", ":b"},
		"/:ab":   {":ab"},
		"/*w":    {"*w"},
		"/*w/:p": {"*w", ":p"},
	} {
		actual := denco.ParamNames(path)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%q expects %q, but %q", path, expected, actual)
		}
	}
}
