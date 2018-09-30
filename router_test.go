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
		{Key: "/", Value: "testroute0"},
		{Key: "/path/to/route", Value: "testroute1"},
		{Key: "/path/to/other", Value: "testroute2"},
		{Key: "/path/to/route/a", Value: "testroute3"},
		{Key: "/path/to/:param", Value: "testroute4"},
		{Key: "/path/to/wildcard/*routepath", Value: "testroute5"},
		{Key: "/path/to/:param1/:param2", Value: "testroute6"},
		{Key: "/path/to/:param1/sep/:param2", Value: "testroute7"},
		{Key: "/:year/:month/:day", Value: "testroute8"},
		{Key: "/user/:id", Value: "testroute9"},
		{Key: "/a/to/b/:param/*routepath", Value: "testroute10"},
	}
}

var realURIs = []denco.Record{
	{Key: "/authorizations", Value: "/authorizations"},
	{Key: "/authorizations/:id", Value: "/authorizations/:id"},
	{Key: "/applications/:client_id/tokens/:access_token", Value: "/applications/:client_id/tokens/:access_token"},
	{Key: "/events", Value: "/events"},
	{Key: "/repos/:owner/:repo/events", Value: "/repos/:owner/:repo/events"},
	{Key: "/networks/:owner/:repo/events", Value: "/networks/:owner/:repo/events"},
	{Key: "/orgs/:org/events", Value: "/orgs/:org/events"},
	{Key: "/users/:user/received_events", Value: "/users/:user/received_events"},
	{Key: "/users/:user/received_events/public", Value: "/users/:user/received_events/public"},
	{Key: "/users/:user/events", Value: "/users/:user/events"},
	{Key: "/users/:user/events/public", Value: "/users/:user/events/public"},
	{Key: "/users/:user/events/orgs/:org", Value: "/users/:user/events/orgs/:org"},
	{Key: "/feeds", Value: "/feeds"},
	{Key: "/notifications", Value: "/notifications"},
	{Key: "/repos/:owner/:repo/notifications", Value: "/repos/:owner/:repo/notifications"},
	{Key: "/notifications/threads/:id", Value: "/notifications/threads/:id"},
	{Key: "/notifications/threads/:id/subscription", Value: "/notifications/threads/:id/subscription"},
	{Key: "/repos/:owner/:repo/stargazers", Value: "/repos/:owner/:repo/stargazers"},
	{Key: "/users/:user/starred", Value: "/users/:user/starred"},
	{Key: "/user/starred", Value: "/user/starred"},
	{Key: "/user/starred/:owner/:repo", Value: "/user/starred/:owner/:repo"},
	{Key: "/repos/:owner/:repo/subscribers", Value: "/repos/:owner/:repo/subscribers"},
	{Key: "/users/:user/subscriptions", Value: "/users/:user/subscriptions"},
	{Key: "/user/subscriptions", Value: "/user/subscriptions"},
	{Key: "/repos/:owner/:repo/subscription", Value: "/repos/:owner/:repo/subscription"},
	{Key: "/user/subscriptions/:owner/:repo", Value: "/user/subscriptions/:owner/:repo"},
	{Key: "/users/:user/gists", Value: "/users/:user/gists"},
	{Key: "/gists", Value: "/gists"},
	{Key: "/gists/:id", Value: "/gists/:id"},
	{Key: "/gists/:id/star", Value: "/gists/:id/star"},
	{Key: "/repos/:owner/:repo/git/blobs/:sha", Value: "/repos/:owner/:repo/git/blobs/:sha"},
	{Key: "/repos/:owner/:repo/git/commits/:sha", Value: "/repos/:owner/:repo/git/commits/:sha"},
	{Key: "/repos/:owner/:repo/git/refs", Value: "/repos/:owner/:repo/git/refs"},
	{Key: "/repos/:owner/:repo/git/tags/:sha", Value: "/repos/:owner/:repo/git/tags/:sha"},
	{Key: "/repos/:owner/:repo/git/trees/:sha", Value: "/repos/:owner/:repo/git/trees/:sha"},
	{Key: "/issues", Value: "/issues"},
	{Key: "/user/issues", Value: "/user/issues"},
	{Key: "/orgs/:org/issues", Value: "/orgs/:org/issues"},
	{Key: "/repos/:owner/:repo/issues", Value: "/repos/:owner/:repo/issues"},
	{Key: "/repos/:owner/:repo/issues/:number", Value: "/repos/:owner/:repo/issues/:number"},
	{Key: "/repos/:owner/:repo/assignees", Value: "/repos/:owner/:repo/assignees"},
	{Key: "/repos/:owner/:repo/assignees/:assignee", Value: "/repos/:owner/:repo/assignees/:assignee"},
	{Key: "/repos/:owner/:repo/issues/:number/comments", Value: "/repos/:owner/:repo/issues/:number/comments"},
	{Key: "/repos/:owner/:repo/issues/:number/events", Value: "/repos/:owner/:repo/issues/:number/events"},
	{Key: "/repos/:owner/:repo/labels", Value: "/repos/:owner/:repo/labels"},
	{Key: "/repos/:owner/:repo/labels/:name", Value: "/repos/:owner/:repo/labels/:name"},
	{Key: "/repos/:owner/:repo/issues/:number/labels", Value: "/repos/:owner/:repo/issues/:number/labels"},
	{Key: "/repos/:owner/:repo/milestones/:number/labels", Value: "/repos/:owner/:repo/milestones/:number/labels"},
	{Key: "/repos/:owner/:repo/milestones", Value: "/repos/:owner/:repo/milestones"},
	{Key: "/repos/:owner/:repo/milestones/:number", Value: "/repos/:owner/:repo/milestones/:number"},
	{Key: "/emojis", Value: "/emojis"},
	{Key: "/gitignore/templates", Value: "/gitignore/templates"},
	{Key: "/gitignore/templates/:name", Value: "/gitignore/templates/:name"},
	{Key: "/meta", Value: "/meta"},
	{Key: "/rate_limit", Value: "/rate_limit"},
	{Key: "/users/:user/orgs", Value: "/users/:user/orgs"},
	{Key: "/user/orgs", Value: "/user/orgs"},
	{Key: "/orgs/:org", Value: "/orgs/:org"},
	{Key: "/orgs/:org/members", Value: "/orgs/:org/members"},
	{Key: "/orgs/:org/members/:user", Value: "/orgs/:org/members/:user"},
	{Key: "/orgs/:org/public_members", Value: "/orgs/:org/public_members"},
	{Key: "/orgs/:org/public_members/:user", Value: "/orgs/:org/public_members/:user"},
	{Key: "/orgs/:org/teams", Value: "/orgs/:org/teams"},
	{Key: "/teams/:id", Value: "/teams/:id"},
	{Key: "/teams/:id/members", Value: "/teams/:id/members"},
	{Key: "/teams/:id/members/:user", Value: "/teams/:id/members/:user"},
	{Key: "/teams/:id/repos", Value: "/teams/:id/repos"},
	{Key: "/teams/:id/repos/:owner/:repo", Value: "/teams/:id/repos/:owner/:repo"},
	{Key: "/user/teams", Value: "/user/teams"},
	{Key: "/repos/:owner/:repo/pulls", Value: "/repos/:owner/:repo/pulls"},
	{Key: "/repos/:owner/:repo/pulls/:number", Value: "/repos/:owner/:repo/pulls/:number"},
	{Key: "/repos/:owner/:repo/pulls/:number/commits", Value: "/repos/:owner/:repo/pulls/:number/commits"},
	{Key: "/repos/:owner/:repo/pulls/:number/files", Value: "/repos/:owner/:repo/pulls/:number/files"},
	{Key: "/repos/:owner/:repo/pulls/:number/merge", Value: "/repos/:owner/:repo/pulls/:number/merge"},
	{Key: "/repos/:owner/:repo/pulls/:number/comments", Value: "/repos/:owner/:repo/pulls/:number/comments"},
	{Key: "/user/repos", Value: "/user/repos"},
	{Key: "/users/:user/repos", Value: "/users/:user/repos"},
	{Key: "/orgs/:org/repos", Value: "/orgs/:org/repos"},
	{Key: "/repositories", Value: "/repositories"},
	{Key: "/repos/:owner/:repo", Value: "/repos/:owner/:repo"},
	{Key: "/repos/:owner/:repo/contributors", Value: "/repos/:owner/:repo/contributors"},
	{Key: "/repos/:owner/:repo/languages", Value: "/repos/:owner/:repo/languages"},
	{Key: "/repos/:owner/:repo/teams", Value: "/repos/:owner/:repo/teams"},
	{Key: "/repos/:owner/:repo/tags", Value: "/repos/:owner/:repo/tags"},
	{Key: "/repos/:owner/:repo/branches", Value: "/repos/:owner/:repo/branches"},
	{Key: "/repos/:owner/:repo/branches/:branch", Value: "/repos/:owner/:repo/branches/:branch"},
	{Key: "/repos/:owner/:repo/collaborators", Value: "/repos/:owner/:repo/collaborators"},
	{Key: "/repos/:owner/:repo/collaborators/:user", Value: "/repos/:owner/:repo/collaborators/:user"},
	{Key: "/repos/:owner/:repo/comments", Value: "/repos/:owner/:repo/comments"},
	{Key: "/repos/:owner/:repo/commits/:sha/comments", Value: "/repos/:owner/:repo/commits/:sha/comments"},
	{Key: "/repos/:owner/:repo/comments/:id", Value: "/repos/:owner/:repo/comments/:id"},
	{Key: "/repos/:owner/:repo/commits", Value: "/repos/:owner/:repo/commits"},
	{Key: "/repos/:owner/:repo/commits/:sha", Value: "/repos/:owner/:repo/commits/:sha"},
	{Key: "/repos/:owner/:repo/readme", Value: "/repos/:owner/:repo/readme"},
	{Key: "/repos/:owner/:repo/keys", Value: "/repos/:owner/:repo/keys"},
	{Key: "/repos/:owner/:repo/keys/:id", Value: "/repos/:owner/:repo/keys/:id"},
	{Key: "/repos/:owner/:repo/downloads", Value: "/repos/:owner/:repo/downloads"},
	{Key: "/repos/:owner/:repo/downloads/:id", Value: "/repos/:owner/:repo/downloads/:id"},
	{Key: "/repos/:owner/:repo/forks", Value: "/repos/:owner/:repo/forks"},
	{Key: "/repos/:owner/:repo/hooks", Value: "/repos/:owner/:repo/hooks"},
	{Key: "/repos/:owner/:repo/hooks/:id", Value: "/repos/:owner/:repo/hooks/:id"},
	{Key: "/repos/:owner/:repo/releases", Value: "/repos/:owner/:repo/releases"},
	{Key: "/repos/:owner/:repo/releases/:id", Value: "/repos/:owner/:repo/releases/:id"},
	{Key: "/repos/:owner/:repo/releases/:id/assets", Value: "/repos/:owner/:repo/releases/:id/assets"},
	{Key: "/repos/:owner/:repo/stats/contributors", Value: "/repos/:owner/:repo/stats/contributors"},
	{Key: "/repos/:owner/:repo/stats/commit_activity", Value: "/repos/:owner/:repo/stats/commit_activity"},
	{Key: "/repos/:owner/:repo/stats/code_frequency", Value: "/repos/:owner/:repo/stats/code_frequency"},
	{Key: "/repos/:owner/:repo/stats/participation", Value: "/repos/:owner/:repo/stats/participation"},
	{Key: "/repos/:owner/:repo/stats/punch_card", Value: "/repos/:owner/:repo/stats/punch_card"},
	{Key: "/repos/:owner/:repo/statuses/:ref", Value: "/repos/:owner/:repo/statuses/:ref"},
	{Key: "/search/repositories", Value: "/search/repositories"},
	{Key: "/search/code", Value: "/search/code"},
	{Key: "/search/issues", Value: "/search/issues"},
	{Key: "/search/users", Value: "/search/users"},
	{Key: "/legacy/issues/search/:owner/:repository/:state/:keyword", Value: "/legacy/issues/search/:owner/:repository/:state/:keyword"},
	{Key: "/legacy/repos/search/:keyword", Value: "/legacy/repos/search/:keyword"},
	{Key: "/legacy/user/search/:keyword", Value: "/legacy/user/search/:keyword"},
	{Key: "/legacy/user/email/:email", Value: "/legacy/user/email/:email"},
	{Key: "/users/:user", Value: "/users/:user"},
	{Key: "/user", Value: "/user"},
	{Key: "/users", Value: "/users"},
	{Key: "/user/emails", Value: "/user/emails"},
	{Key: "/users/:user/followers", Value: "/users/:user/followers"},
	{Key: "/user/followers", Value: "/user/followers"},
	{Key: "/users/:user/following", Value: "/users/:user/following"},
	{Key: "/user/following", Value: "/user/following"},
	{Key: "/user/following/:user", Value: "/user/following/:user"},
	{Key: "/users/:user/following/:target_user", Value: "/users/:user/following/:target_user"},
	{Key: "/users/:user/keys", Value: "/users/:user/keys"},
	{Key: "/user/keys", Value: "/user/keys"},
	{Key: "/user/keys/:id", Value: "/user/keys/:id"},
	{Key: "/people/:userId", Value: "/people/:userId"},
	{Key: "/people", Value: "/people"},
	{Key: "/activities/:activityId/people/:collection", Value: "/activities/:activityId/people/:collection"},
	{Key: "/people/:userId/people/:collection", Value: "/people/:userId/people/:collection"},
	{Key: "/people/:userId/openIdConnect", Value: "/people/:userId/openIdConnect"},
	{Key: "/people/:userId/activities/:collection", Value: "/people/:userId/activities/:collection"},
	{Key: "/activities/:activityId", Value: "/activities/:activityId"},
	{Key: "/activities", Value: "/activities"},
	{Key: "/activities/:activityId/comments", Value: "/activities/:activityId/comments"},
	{Key: "/comments/:commentId", Value: "/comments/:commentId"},
	{Key: "/people/:userId/moments/:collection", Value: "/people/:userId/moments/:collection"},
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
		if !reflect.DeepEqual(data, testcase.value) || !reflect.DeepEqual(params, denco.Params(testcase.params)) || !reflect.DeepEqual(found, testcase.found) {
			t.Errorf("Router.Lookup(%q) => (%#v, %#v, %#v), want (%#v, %#v, %#v)", testcase.path, data, params, found, testcase.value, denco.Params(testcase.params), testcase.found)
		}
	}
}

func TestRouter_Lookup(t *testing.T) {
	testcases := []testcase{
		{"/", "testroute0", nil, true},
		{"/path/to/route", "testroute1", nil, true},
		{"/path/to/other", "testroute2", nil, true},
		{"/path/to/route/a", "testroute3", nil, true},
		{"/path/to/hoge", "testroute4", []denco.Param{{Name: "param", Value: "hoge"}}, true},
		{"/path/to/wildcard/some/params", "testroute5", []denco.Param{{Name: "routepath", Value: "some/params"}}, true},
		{"/path/to/o1/o2", "testroute6", []denco.Param{{Name: "param1", Value: "o1"}, {Name: "param2", Value: "o2"}}, true},
		{"/path/to/p1/sep/p2", "testroute7", []denco.Param{{Name: "param1", Value: "p1"}, {Name: "param2", Value: "p2"}}, true},
		{"/2014/01/06", "testroute8", []denco.Param{{Name: "year", Value: "2014"}, {Name: "month", Value: "01"}, {Name: "day", Value: "06"}}, true},
		{"/user/777", "testroute9", []denco.Param{{Name: "id", Value: "777"}}, true},
		{"/a/to/b/p1/some/wildcard/params", "testroute10", []denco.Param{{Name: "param", Value: "p1"}, {Name: "routepath", Value: "some/wildcard/params"}}, true},
		{"/missing", nil, nil, false},
	}
	runLookupTest(t, routes(), testcases)

	records := []denco.Record{
		{Key: "/", Value: "testroute0"},
		{Key: "/:b", Value: "testroute1"},
		{Key: "/*wildcard", Value: "testroute2"},
	}
	testcases = []testcase{
		{"/", "testroute0", nil, true},
		{"/true", "testroute1", []denco.Param{{Name: "b", Value: "true"}}, true},
		{"/foo/bar", "testroute2", []denco.Param{{Name: "wildcard", Value: "foo/bar"}}, true},
	}
	runLookupTest(t, records, testcases)

	records = []denco.Record{
		{Key: "/networks/:owner/:repo/events", Value: "testroute0"},
		{Key: "/orgs/:org/events", Value: "testroute1"},
		{Key: "/notifications/threads/:id", Value: "testroute2"},
	}
	testcases = []testcase{
		{"/networks/:owner/:repo/events", "testroute0", []denco.Param{{Name: "owner", Value: ":owner"}, {Name: "repo", Value: ":repo"}}, true},
		{"/orgs/:org/events", "testroute1", []denco.Param{{Name: "org", Value: ":org"}}, true},
		{"/notifications/threads/:id", "testroute2", []denco.Param{{Name: "id", Value: ":id"}}, true},
	}
	runLookupTest(t, records, testcases)

	runLookupTest(t, []denco.Record{
		{Key: "/", Value: "route2"},
	}, []testcase{
		{"/user/alice", nil, nil, false},
	})

	runLookupTest(t, []denco.Record{
		{Key: "/user/:name", Value: "route1"},
	}, []testcase{
		{"/", nil, nil, false},
	})

	runLookupTest(t, []denco.Record{
		{Key: "/*wildcard", Value: "testroute0"},
		{Key: "/a/:b", Value: "testroute1"},
	}, []testcase{
		{"/a", "testroute0", []denco.Param{{Name: "wildcard", Value: "a"}}, true},
	})

	runLookupTest(t, []denco.Record{
		{Key: "/stock-adjustments/:stock-adjustment-id/lines/upload", Value: "lines-up"},
		{Key: "/stock-adjustments/:stock-adjustment-id/lines/:id", Value: "lines-id"},
	}, []testcase{
		{"/stock-adjustments/123/lines/upload", "lines-up", []denco.Param{{Name: "stock-adjustment-id", Value: "123"}}, true},
	})
}

func TestRouter_Lookup_withManyRoutes(t *testing.T) {
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

func TestRouter_Lookup_realURIs(t *testing.T) {
	testcases := []testcase{
		{"/authorizations", "/authorizations", nil, true},
		{"/authorizations/1", "/authorizations/:id", []denco.Param{{Name: "id", Value: "1"}}, true},
		{"/applications/1/tokens/zohRoo7e", "/applications/:client_id/tokens/:access_token", []denco.Param{{Name: "client_id", Value: "1"}, {Name: "access_token", Value: "zohRoo7e"}}, true},
		{"/events", "/events", nil, true},
		{"/repos/naoina/denco/events", "/repos/:owner/:repo/events", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/networks/naoina/denco/events", "/networks/:owner/:repo/events", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/orgs/something/events", "/orgs/:org/events", []denco.Param{{Name: "org", Value: "something"}}, true},
		{"/users/naoina/received_events", "/users/:user/received_events", []denco.Param{{Name: "user", Value: "naoina"}}, true},
		{"/users/naoina/received_events/public", "/users/:user/received_events/public", []denco.Param{{Name: "user", Value: "naoina"}}, true},
		{"/users/naoina/events", "/users/:user/events", []denco.Param{{Name: "user", Value: "naoina"}}, true},
		{"/users/naoina/events/public", "/users/:user/events/public", []denco.Param{{Name: "user", Value: "naoina"}}, true},
		{"/users/naoina/events/orgs/something", "/users/:user/events/orgs/:org", []denco.Param{{Name: "user", Value: "naoina"}, {Name: "org", Value: "something"}}, true},
		{"/feeds", "/feeds", nil, true},
		{"/notifications", "/notifications", nil, true},
		{"/repos/naoina/denco/notifications", "/repos/:owner/:repo/notifications", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/notifications/threads/1", "/notifications/threads/:id", []denco.Param{{Name: "id", Value: "1"}}, true},
		{"/notifications/threads/2/subscription", "/notifications/threads/:id/subscription", []denco.Param{{Name: "id", Value: "2"}}, true},
		{"/repos/naoina/denco/stargazers", "/repos/:owner/:repo/stargazers", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/users/naoina/starred", "/users/:user/starred", []denco.Param{{Name: "user", Value: "naoina"}}, true},
		{"/user/starred", "/user/starred", nil, true},
		{"/user/starred/naoina/denco", "/user/starred/:owner/:repo", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/subscribers", "/repos/:owner/:repo/subscribers", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/users/naoina/subscriptions", "/users/:user/subscriptions", []denco.Param{{Name: "user", Value: "naoina"}}, true},
		{"/user/subscriptions", "/user/subscriptions", nil, true},
		{"/repos/naoina/denco/subscription", "/repos/:owner/:repo/subscription", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/user/subscriptions/naoina/denco", "/user/subscriptions/:owner/:repo", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/users/naoina/gists", "/users/:user/gists", []denco.Param{{Name: "user", Value: "naoina"}}, true},
		{"/gists", "/gists", nil, true},
		{"/gists/1", "/gists/:id", []denco.Param{{Name: "id", Value: "1"}}, true},
		{"/gists/2/star", "/gists/:id/star", []denco.Param{{Name: "id", Value: "2"}}, true},
		{"/repos/naoina/denco/git/blobs/03c3bbc7f0d12268b9ca53d4fbfd8dc5ae5697b9", "/repos/:owner/:repo/git/blobs/:sha", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "sha", Value: "03c3bbc7f0d12268b9ca53d4fbfd8dc5ae5697b9"}}, true},
		{"/repos/naoina/denco/git/commits/03c3bbc7f0d12268b9ca53d4fbfd8dc5ae5697b9", "/repos/:owner/:repo/git/commits/:sha", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "sha", Value: "03c3bbc7f0d12268b9ca53d4fbfd8dc5ae5697b9"}}, true},
		{"/repos/naoina/denco/git/refs", "/repos/:owner/:repo/git/refs", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/git/tags/03c3bbc7f0d12268b9ca53d4fbfd8dc5ae5697b9", "/repos/:owner/:repo/git/tags/:sha", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "sha", Value: "03c3bbc7f0d12268b9ca53d4fbfd8dc5ae5697b9"}}, true},
		{"/repos/naoina/denco/git/trees/03c3bbc7f0d12268b9ca53d4fbfd8dc5ae5697b9", "/repos/:owner/:repo/git/trees/:sha", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "sha", Value: "03c3bbc7f0d12268b9ca53d4fbfd8dc5ae5697b9"}}, true},
		{"/issues", "/issues", nil, true},
		{"/user/issues", "/user/issues", nil, true},
		{"/orgs/something/issues", "/orgs/:org/issues", []denco.Param{{Name: "org", Value: "something"}}, true},
		{"/repos/naoina/denco/issues", "/repos/:owner/:repo/issues", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/issues/1", "/repos/:owner/:repo/issues/:number", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "number", Value: "1"}}, true},
		{"/repos/naoina/denco/assignees", "/repos/:owner/:repo/assignees", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/assignees/foo", "/repos/:owner/:repo/assignees/:assignee", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "assignee", Value: "foo"}}, true},
		{"/repos/naoina/denco/issues/1/comments", "/repos/:owner/:repo/issues/:number/comments", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "number", Value: "1"}}, true},
		{"/repos/naoina/denco/issues/1/events", "/repos/:owner/:repo/issues/:number/events", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "number", Value: "1"}}, true},
		{"/repos/naoina/denco/labels", "/repos/:owner/:repo/labels", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/labels/bug", "/repos/:owner/:repo/labels/:name", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "name", Value: "bug"}}, true},
		{"/repos/naoina/denco/issues/1/labels", "/repos/:owner/:repo/issues/:number/labels", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "number", Value: "1"}}, true},
		{"/repos/naoina/denco/milestones/1/labels", "/repos/:owner/:repo/milestones/:number/labels", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "number", Value: "1"}}, true},
		{"/repos/naoina/denco/milestones", "/repos/:owner/:repo/milestones", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/milestones/1", "/repos/:owner/:repo/milestones/:number", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "number", Value: "1"}}, true},
		{"/emojis", "/emojis", nil, true},
		{"/gitignore/templates", "/gitignore/templates", nil, true},
		{"/gitignore/templates/Go", "/gitignore/templates/:name", []denco.Param{{Name: "name", Value: "Go"}}, true},
		{"/meta", "/meta", nil, true},
		{"/rate_limit", "/rate_limit", nil, true},
		{"/users/naoina/orgs", "/users/:user/orgs", []denco.Param{{Name: "user", Value: "naoina"}}, true},
		{"/user/orgs", "/user/orgs", nil, true},
		{"/orgs/something", "/orgs/:org", []denco.Param{{Name: "org", Value: "something"}}, true},
		{"/orgs/something/members", "/orgs/:org/members", []denco.Param{{Name: "org", Value: "something"}}, true},
		{"/orgs/something/members/naoina", "/orgs/:org/members/:user", []denco.Param{{Name: "org", Value: "something"}, {Name: "user", Value: "naoina"}}, true},
		{"/orgs/something/public_members", "/orgs/:org/public_members", []denco.Param{{Name: "org", Value: "something"}}, true},
		{"/orgs/something/public_members/naoina", "/orgs/:org/public_members/:user", []denco.Param{{Name: "org", Value: "something"}, {Name: "user", Value: "naoina"}}, true},
		{"/orgs/something/teams", "/orgs/:org/teams", []denco.Param{{Name: "org", Value: "something"}}, true},
		{"/teams/1", "/teams/:id", []denco.Param{{Name: "id", Value: "1"}}, true},
		{"/teams/2/members", "/teams/:id/members", []denco.Param{{Name: "id", Value: "2"}}, true},
		{"/teams/3/members/naoina", "/teams/:id/members/:user", []denco.Param{{Name: "id", Value: "3"}, {Name: "user", Value: "naoina"}}, true},
		{"/teams/4/repos", "/teams/:id/repos", []denco.Param{{Name: "id", Value: "4"}}, true},
		{"/teams/5/repos/naoina/denco", "/teams/:id/repos/:owner/:repo", []denco.Param{{Name: "id", Value: "5"}, {Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/user/teams", "/user/teams", nil, true},
		{"/repos/naoina/denco/pulls", "/repos/:owner/:repo/pulls", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/pulls/1", "/repos/:owner/:repo/pulls/:number", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "number", Value: "1"}}, true},
		{"/repos/naoina/denco/pulls/1/commits", "/repos/:owner/:repo/pulls/:number/commits", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "number", Value: "1"}}, true},
		{"/repos/naoina/denco/pulls/1/files", "/repos/:owner/:repo/pulls/:number/files", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "number", Value: "1"}}, true},
		{"/repos/naoina/denco/pulls/1/merge", "/repos/:owner/:repo/pulls/:number/merge", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "number", Value: "1"}}, true},
		{"/repos/naoina/denco/pulls/1/comments", "/repos/:owner/:repo/pulls/:number/comments", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "number", Value: "1"}}, true},
		{"/user/repos", "/user/repos", nil, true},
		{"/users/naoina/repos", "/users/:user/repos", []denco.Param{{Name: "user", Value: "naoina"}}, true},
		{"/orgs/something/repos", "/orgs/:org/repos", []denco.Param{{Name: "org", Value: "something"}}, true},
		{"/repositories", "/repositories", nil, true},
		{"/repos/naoina/denco", "/repos/:owner/:repo", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/contributors", "/repos/:owner/:repo/contributors", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/languages", "/repos/:owner/:repo/languages", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/teams", "/repos/:owner/:repo/teams", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/tags", "/repos/:owner/:repo/tags", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/branches", "/repos/:owner/:repo/branches", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/branches/master", "/repos/:owner/:repo/branches/:branch", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "branch", Value: "master"}}, true},
		{"/repos/naoina/denco/collaborators", "/repos/:owner/:repo/collaborators", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/collaborators/something", "/repos/:owner/:repo/collaborators/:user", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "user", Value: "something"}}, true},
		{"/repos/naoina/denco/comments", "/repos/:owner/:repo/comments", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/commits/03c3bbc7f0d12268b9ca53d4fbfd8dc5ae5697b9/comments", "/repos/:owner/:repo/commits/:sha/comments", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "sha", Value: "03c3bbc7f0d12268b9ca53d4fbfd8dc5ae5697b9"}}, true},
		{"/repos/naoina/denco/comments/1", "/repos/:owner/:repo/comments/:id", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "id", Value: "1"}}, true},
		{"/repos/naoina/denco/commits", "/repos/:owner/:repo/commits", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/commits/03c3bbc7f0d12268b9ca53d4fbfd8dc5ae5697b9", "/repos/:owner/:repo/commits/:sha", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "sha", Value: "03c3bbc7f0d12268b9ca53d4fbfd8dc5ae5697b9"}}, true},
		{"/repos/naoina/denco/readme", "/repos/:owner/:repo/readme", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/keys", "/repos/:owner/:repo/keys", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/keys/1", "/repos/:owner/:repo/keys/:id", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "id", Value: "1"}}, true},
		{"/repos/naoina/denco/downloads", "/repos/:owner/:repo/downloads", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/downloads/2", "/repos/:owner/:repo/downloads/:id", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "id", Value: "2"}}, true},
		{"/repos/naoina/denco/forks", "/repos/:owner/:repo/forks", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/hooks", "/repos/:owner/:repo/hooks", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/hooks/2", "/repos/:owner/:repo/hooks/:id", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "id", Value: "2"}}, true},
		{"/repos/naoina/denco/releases", "/repos/:owner/:repo/releases", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/releases/1", "/repos/:owner/:repo/releases/:id", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "id", Value: "1"}}, true},
		{"/repos/naoina/denco/releases/1/assets", "/repos/:owner/:repo/releases/:id/assets", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "id", Value: "1"}}, true},
		{"/repos/naoina/denco/stats/contributors", "/repos/:owner/:repo/stats/contributors", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/stats/commit_activity", "/repos/:owner/:repo/stats/commit_activity", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/stats/code_frequency", "/repos/:owner/:repo/stats/code_frequency", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/stats/participation", "/repos/:owner/:repo/stats/participation", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/stats/punch_card", "/repos/:owner/:repo/stats/punch_card", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}}, true},
		{"/repos/naoina/denco/statuses/master", "/repos/:owner/:repo/statuses/:ref", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repo", Value: "denco"}, {Name: "ref", Value: "master"}}, true},
		{"/search/repositories", "/search/repositories", nil, true},
		{"/search/code", "/search/code", nil, true},
		{"/search/issues", "/search/issues", nil, true},
		{"/search/users", "/search/users", nil, true},
		{"/legacy/issues/search/naoina/denco/closed/test", "/legacy/issues/search/:owner/:repository/:state/:keyword", []denco.Param{{Name: "owner", Value: "naoina"}, {Name: "repository", Value: "denco"}, {Name: "state", Value: "closed"}, {Name: "keyword", Value: "test"}}, true},
		{"/legacy/repos/search/test", "/legacy/repos/search/:keyword", []denco.Param{{Name: "keyword", Value: "test"}}, true},
		{"/legacy/user/search/test", "/legacy/user/search/:keyword", []denco.Param{{Name: "keyword", Value: "test"}}, true},
		{"/legacy/user/email/naoina@kuune.org", "/legacy/user/email/:email", []denco.Param{{Name: "email", Value: "naoina@kuune.org"}}, true},
		{"/users/naoina", "/users/:user", []denco.Param{{Name: "user", Value: "naoina"}}, true},
		{"/user", "/user", nil, true},
		{"/users", "/users", nil, true},
		{"/user/emails", "/user/emails", nil, true},
		{"/users/naoina/followers", "/users/:user/followers", []denco.Param{{Name: "user", Value: "naoina"}}, true},
		{"/user/followers", "/user/followers", nil, true},
		{"/users/naoina/following", "/users/:user/following", []denco.Param{{Name: "user", Value: "naoina"}}, true},
		{"/user/following", "/user/following", nil, true},
		{"/user/following/naoina", "/user/following/:user", []denco.Param{{Name: "user", Value: "naoina"}}, true},
		{"/users/naoina/following/target", "/users/:user/following/:target_user", []denco.Param{{Name: "user", Value: "naoina"}, {Name: "target_user", Value: "target"}}, true},
		{"/users/naoina/keys", "/users/:user/keys", []denco.Param{{Name: "user", Value: "naoina"}}, true},
		{"/user/keys", "/user/keys", nil, true},
		{"/user/keys/1", "/user/keys/:id", []denco.Param{{Name: "id", Value: "1"}}, true},
		{"/people/me", "/people/:userId", []denco.Param{{Name: "userId", Value: "me"}}, true},
		{"/people", "/people", nil, true},
		{"/activities/foo/people/vault", "/activities/:activityId/people/:collection", []denco.Param{{Name: "activityId", Value: "foo"}, {Name: "collection", Value: "vault"}}, true},
		{"/people/me/people/vault", "/people/:userId/people/:collection", []denco.Param{{Name: "userId", Value: "me"}, {Name: "collection", Value: "vault"}}, true},
		{"/people/me/openIdConnect", "/people/:userId/openIdConnect", []denco.Param{{Name: "userId", Value: "me"}}, true},
		{"/people/me/activities/vault", "/people/:userId/activities/:collection", []denco.Param{{Name: "userId", Value: "me"}, {Name: "collection", Value: "vault"}}, true},
		{"/activities/foo", "/activities/:activityId", []denco.Param{{Name: "activityId", Value: "foo"}}, true},
		{"/activities", "/activities", nil, true},
		{"/activities/foo/comments", "/activities/:activityId/comments", []denco.Param{{Name: "activityId", Value: "foo"}}, true},
		{"/comments/hoge", "/comments/:commentId", []denco.Param{{Name: "commentId", Value: "hoge"}}, true},
		{"/people/me/moments/vault", "/people/:userId/moments/:collection", []denco.Param{{Name: "userId", Value: "me"}, {Name: "collection", Value: "vault"}}, true},
	}
	runLookupTest(t, realURIs, testcases)
}

func TestRouter_Build(t *testing.T) {
	// test for duplicate name of path parameters.
	func() {
		r := denco.New()
		if err := r.Build([]denco.Record{
			{Key: "/:user/:id/:id", Value: "testroute0"},
			{Key: "/:user/:user/:id", Value: "testroute0"},
		}); err == nil {
			t.Errorf("no error returned by duplicate name of path parameters")
		}
	}()
}

func TestRouter_Build_withoutSizeHint(t *testing.T) {
	for _, v := range []struct {
		keys     []string
		sizeHint int
	}{
		{[]string{"/user"}, 0},
		{[]string{"/user/:id"}, 1},
		{[]string{"/user/:id/post"}, 1},
		{[]string{"/user/:id/:group"}, 2},
		{[]string{"/user/:id/post/:cid"}, 2},
		{[]string{"/user/:id/post/:cid", "/admin/:id/post/:cid"}, 2},
		{[]string{"/user/:id", "/admin/:id/post/:cid"}, 2},
		{[]string{"/user/:id/post/:cid", "/admin/:id/post/:cid/:type"}, 3},
	} {
		r := denco.New()
		actual := r.SizeHint
		expect := -1
		if !reflect.DeepEqual(actual, expect) {
			t.Errorf(`before Build; Router.SizeHint => (%[1]T=%#[1]v); want (%[2]T=%#[2]v)`, actual, expect)
		}
		records := make([]denco.Record, len(v.keys))
		for i, k := range v.keys {
			records[i] = denco.Record{Key: k, Value: "value"}
		}
		if err := r.Build(records); err != nil {
			t.Fatal(err)
		}
		actual = r.SizeHint
		expect = v.sizeHint
		if !reflect.DeepEqual(actual, expect) {
			t.Errorf(`Router.Build(%#v); Router.SizeHint => (%[2]T=%#[2]v); want (%[3]T=%#[3]v)`, records, actual, expect)
		}
	}
}

func TestRouter_Build_withSizeHint(t *testing.T) {
	for _, v := range []struct {
		key      string
		sizeHint int
		expect   int
	}{
		{"/user", 0, 0},
		{"/user", 1, 1},
		{"/user", 2, 2},
		{"/user/:id", 3, 3},
		{"/user/:id/:group", 0, 0},
		{"/user/:id/:group", 1, 1},
	} {
		r := denco.New()
		r.SizeHint = v.sizeHint
		records := []denco.Record{
			{Key: v.key, Value: "value"},
		}
		if err := r.Build(records); err != nil {
			t.Fatal(err)
		}
		actual := r.SizeHint
		expect := v.expect
		if !reflect.DeepEqual(actual, expect) {
			t.Errorf(`Router.Build(%#v); Router.SizeHint => (%[2]T=%#[2]v); want (%[3]T=%#[3]v)`, records, actual, expect)
		}
	}
}

func TestParams_Get(t *testing.T) {
	params := denco.Params([]denco.Param{
		{Name: "name1", Value: "value1"},
		{Name: "name2", Value: "value2"},
		{Name: "name3", Value: "value3"},
		{Name: "name1", Value: "value4"},
	})
	for _, v := range []struct{ value, expected string }{
		{"name1", "value1"},
		{"name2", "value2"},
		{"name3", "value3"},
		{"name4", ""},
	} {
		actual := params.Get(v.value)
		expected := v.expected
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Params.Get(%q) => %#v, want %#v", v.value, actual, expected)
		}
	}
}
