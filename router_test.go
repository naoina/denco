package denco_test

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/naoina/denco"
)

func routes() []denco.Record {
	return []denco.Record{
		{"/", "testroute0"},
		{"/path/to/route", "testroute1"},
		{"/path/to/other", "testroute2"},
		{"/path/to/route/a", "testroute3"},
		{"/path/to/:param", "testroute4"},
		{"/path/to/wildcard/*routepath", "testroute5"},
		{"/path/to/:param1/:param2", "testroute6"},
		{"/path/to/:param1/sep/:param2", "testroute7"},
		{"/:year/:month/:day", "testroute8"},
		{"/user/:id", "testroute9"},
		{"/a/to/b/:param/*routepath", "testroute10"},
	}
}

type testcase struct {
	path   string
	value  interface{}
	params []denco.Param
	found  bool
}

func runLookupTest(t *testing.T, records []denco.Record, testcases []testcase) {
	r := denco.New()
	if err := r.Build(records); err != nil {
		t.Fatal(err)
	}
	for _, testcase := range testcases {
		data, params, found := r.Lookup(testcase.path)
		if !reflect.DeepEqual(data, testcase.value) || !reflect.DeepEqual(params, testcase.params) || !reflect.DeepEqual(found, testcase.found) {
			t.Errorf("Router.Lookup(%q) => (%#v, %#v, %#v), want (%#v, %#v, %#v)", testcase.path, data, params, found, testcase.value, testcase.params, testcase.found)
		}
	}
}

func TestRouter_Lookup(t *testing.T) {
	testcases := []testcase{
		{"/", "testroute0", nil, true},
		{"/path/to/route", "testroute1", nil, true},
		{"/path/to/other", "testroute2", nil, true},
		{"/path/to/route/a", "testroute3", nil, true},
		{"/path/to/hoge", "testroute4", []denco.Param{{"param", "hoge"}}, true},
		{"/path/to/wildcard/some/params", "testroute5", []denco.Param{{"routepath", "some/params"}}, true},
		{"/path/to/o1/o2", "testroute6", []denco.Param{{"param1", "o1"}, {"param2", "o2"}}, true},
		{"/path/to/p1/sep/p2", "testroute7", []denco.Param{{"param1", "p1"}, {"param2", "p2"}}, true},
		{"/2014/01/06", "testroute8", []denco.Param{{"year", "2014"}, {"month", "01"}, {"day", "06"}}, true},
		{"/user/777", "testroute9", []denco.Param{{"id", "777"}}, true},
		{"/a/to/b/p1/some/wildcard/params", "testroute10", []denco.Param{{"param", "p1"}, {"routepath", "some/wildcard/params"}}, true},
		{"/missing", nil, nil, false},
	}
	runLookupTest(t, routes(), testcases)

	records := []denco.Record{
		{"/", "testroute0"},
		{"/:b", "testroute1"},
		{"/*wildcard", "testroute2"},
	}
	testcases = []testcase{
		{"/", "testroute0", nil, true},
		{"/true", "testroute1", []denco.Param{{"b", "true"}}, true},
		{"/foo/bar", "testroute2", []denco.Param{{"wildcard", "foo/bar"}}, true},
	}
	runLookupTest(t, records, testcases)

	records = []denco.Record{
		{"/networks/:owner/:repo/events", "testroute0"},
		{"/orgs/:org/events", "testroute1"},
		{"/notifications/threads/:id", "testroute2"},
	}
	testcases = []testcase{
		{"/networks/:owner/:repo/events", "testroute0", []denco.Param{{"owner", ":owner"}, {"repo", ":repo"}}, true},
		{"/orgs/:org/events", "testroute1", []denco.Param{{"org", ":org"}}, true},
		{"/notifications/threads/:id", "testroute2", []denco.Param{{"id", ":id"}}, true},
	}
}

func Testdenco_Lookup_withManyRoutes(t *testing.T) {
	n := 1000
	rand.Seed(time.Now().UnixNano())
	records := make([]denco.Record, n)
	for i := 0; i < n; i++ {
		records[i] = denco.Record{Key: "/" + randomString(rand.Intn(50)+10), Value: fmt.Sprintf("route%d", i)}
	}
	router := denco.New()
	if err := router.Build(records); err != nil {
		t.Fatal(err)
	}
	for _, r := range records {
		data, params, found := router.Lookup(r.Key)
		if !reflect.DeepEqual(data, r.Value) || len(params) != 0 || !reflect.DeepEqual(found, true) {
			t.Errorf("Router.Lookup(%q) => (%#v, %#v, %#v), want (%#v, %#v, %#v)", r.Key, data, len(params), found, r.Value, 0, true)
		}
	}
}

func TestRouter_Build(t *testing.T) {
	// test for duplicate name of path parameters.
	func() {
		r := denco.New()
		if err := r.Build([]denco.Record{
			{"/:user/:id/:id", "testroute0"},
			{"/:user/:user/:id", "testroute0"},
		}); err == nil {
			t.Errorf("no error returned by duplicate name of path parameters")
		}
	}()
}
