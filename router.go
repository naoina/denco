// Package denco provides fast URL router.
package denco

import (
	"fmt"
	"sort"
	"strings"
)

const (
	// A special character for path parameter.
	ParamCharacter = ':'

	// A special character for wildcard path parameter.
	WildcardCharacter = '*'

	// TerminalCharacter is a special character for end of path.
	TerminationCharacter = '#'

	// Block size of array of BASE/CHECK of Double-Array.
	blockSize = 256
)

// Router represents a URL router.
type Router struct {
	static map[string]interface{}
	param  *doubleArray
}

// New returns a new Router.
func New() *Router {
	return &Router{
		static: make(map[string]interface{}),
		param:  newDoubleArray(blockSize),
	}
}

// Lookup returns data and path parameters that associated with path.
// params is a slice of the Param that arranged in the order in which parameters appeared.
// e.g. when built routing path is "/path/:id/:name" and given path is "/path/to/1/alice". params order is [{"id": "1"}, {"name": "alice"}], not [{"name": "alice"}, {"id": "1"}].
func (rt *Router) Lookup(path string) (data interface{}, params []Param, found bool) {
	if data, found := rt.static[path]; found {
		return data, nil, true
	}
	idx, values, found := rt.param.lookup(path, nil, 1)
	if !found {
		return nil, nil, false
	}
	nd := rt.param.node[-rt.param.bc[idx].base]
	if nd == nil {
		return nil, nil, false
	}
	if len(values) > 0 {
		params = make([]Param, len(values))
		for i, v := range values {
			params[i] = Param{Name: nd.paramNames[i], Value: v}
		}
	}
	return nd.data, params, true
}

// Build builds URL router from records.
func (rt *Router) Build(records []Record) error {
	statics, params := makeRecords(records)
	for _, r := range statics {
		rt.static[r.Key] = r.Value
	}
	if err := rt.param.build(params, 1, 0); err != nil {
		return err
	}
	return nil
}

// Param represents name and value of path parameter.
type Param struct {
	Name  string
	Value string
}

type doubleArray struct {
	bc   []baseCheck
	node []*node
}

func newDoubleArray(size int) *doubleArray {
	return &doubleArray{
		bc:   newBaseCheckArray(size + 1),
		node: []*node{nil}, // A start index is adjusting to 1 because 0 will be used as a mark of non-existent node.
	}
}

// newBaseCheckArray returns a new slice of baseCheck with given size.
func newBaseCheckArray(size int) []baseCheck {
	return make([]baseCheck, size)
}

// baseCheck represents a BASE/CHECK node.
type baseCheck struct {
	base      int
	check     byte
	paramType paramType
}

type paramType uint8

func (p paramType) IsParam() bool {
	return p&paramTypeSingle == paramTypeSingle
}

func (p paramType) IsWildcard() bool {
	return p&paramTypeWildcard == paramTypeWildcard
}

func (p paramType) IsAny() bool {
	return p > 0
}

func (p *paramType) SetSingle() {
	*p |= paramTypeSingle
}

func (p *paramType) SetWildcard() {
	*p |= paramTypeWildcard
}

const (
	paramTypeSingle = 1 << iota
	paramTypeWildcard
)

func (da *doubleArray) lookup(path string, params []string, idx int) (int, []string, bool) {
	indices := make([]uint64, 0, 1)
	for i := 0; i < len(path); i++ {
		if da.bc[idx].paramType.IsAny() {
			indices = append(indices, (uint64(i&0xffffffff)<<32)|uint64(idx&0xffffffff))
		}
		c := path[i]
		next := nextIndex(da.bc[idx].base, c)
		if da.bc[next].check != c {
			goto BACKTRACKING
		}
		idx = next
	}
	if next := nextIndex(da.bc[idx].base, TerminationCharacter); da.bc[next].check == TerminationCharacter {
		return next, params, true
	}
	return -1, nil, false
BACKTRACKING:
	for j := len(indices) - 1; j >= 0; j-- {
		i, idx := int((indices[j]>>32)&0xffffffff), int(indices[j]&0xffffffff)
		if da.bc[idx].paramType.IsParam() {
			next := NextSeparator(path, i)
			idx := nextIndex(da.bc[idx].base, ParamCharacter)
			params := append(params, path[i:next])
			path := path[next:]
			if idx, params, found := da.lookup(path, params, idx); found {
				return idx, params, true
			}
		}
		if da.bc[idx].paramType.IsWildcard() {
			idx := nextIndex(da.bc[idx].base, WildcardCharacter)
			params := append(params, path[i:])
			return idx, params, true
		}
	}
	return -1, nil, false
}

// build builds double-array from records.
func (da *doubleArray) build(srcs []*record, idx, depth int) error {
	sort.Sort(recordSlice(srcs))
	base, siblings, leaf, err := da.arrange(srcs, idx, depth)
	if err != nil {
		return err
	}
	if leaf != nil {
		nd, err := makeNode(leaf)
		if err != nil {
			return err
		}
		da.bc[idx].base = -len(da.node)
		da.node = append(da.node, nd)
	}
	for _, sib := range siblings {
		da.setCheck(nextIndex(base, sib.c), sib.c)
	}
	for _, sib := range siblings {
		records := srcs[sib.start:sib.end]
		switch sib.c {
		case ParamCharacter:
			for _, r := range records {
				next := NextSeparator(r.Key, depth+1)
				name := r.Key[depth+1 : next]
				r.paramNames = append(r.paramNames, name)
				r.Key = r.Key[next:]
			}
			da.bc[idx].paramType.SetSingle()
			if err := da.build(records, nextIndex(base, sib.c), 0); err != nil {
				return err
			}
		case WildcardCharacter:
			r := records[0]
			name := r.Key[depth+1 : len(r.Key)-1]
			r.paramNames = append(r.paramNames, name)
			r.Key = ""
			da.bc[idx].paramType.SetWildcard()
			if err := da.build(records, nextIndex(base, sib.c), 0); err != nil {
				return err
			}
		default:
			if err := da.build(records, nextIndex(base, sib.c), depth+1); err != nil {
				return err
			}
		}
	}
	return nil
}

// setBase sets BASE.
func (da *doubleArray) setBase(i, base int) {
	da.bc[i].base = base
}

// setCheck sets CHECK.
func (da *doubleArray) setCheck(i int, check byte) {
	da.bc[i].check = check
}

// extendBaseCheckArray extends array of BASE/CHECK.
func (da *doubleArray) extendBaseCheckArray() {
	da.bc = append(da.bc, newBaseCheckArray(blockSize)...)
}

// findEmptyIndex returns an index of unused BASE/CHECK node.
func (da *doubleArray) findEmptyIndex(start int) int {
	i := start
	for ; i < len(da.bc); i++ {
		if da.bc[i].base == 0 && da.bc[i].check == 0 {
			break
		}
	}
	return i
}

// findBase returns good BASE.
func (da *doubleArray) findBase(siblings []sibling, start int) (base int) {
	idx := start + 1
	firstChar := siblings[0].c
	for ; idx < len(da.bc); idx = da.findEmptyIndex(idx + 1) {
		base = nextIndex(idx, firstChar)
		i := 0
		for ; i < len(siblings); i++ {
			if next := nextIndex(base, siblings[i].c); len(da.bc) <= next || da.bc[next].base != 0 || da.bc[next].check != 0 {
				break
			}
		}
		if i == len(siblings) {
			return base
		}
	}
	da.extendBaseCheckArray()
	return nextIndex(idx, firstChar)
}

func (da *doubleArray) arrange(records []*record, idx, depth int) (base int, siblings []sibling, leaf *record, err error) {
	siblings, leaf, err = makeSiblings(records, depth)
	if err != nil {
		return -1, nil, nil, err
	}
	if len(siblings) < 1 {
		return -1, nil, leaf, nil
	}
	base = da.findBase(siblings, idx)
	da.setBase(idx, base)
	return base, siblings, leaf, err
}

// node represents a node of Double-Array.
type node struct {
	data interface{}

	// Names of path parameters.
	paramNames []string
}

// makeNode returns a new node from record.
func makeNode(r *record) (*node, error) {
	dups := make(map[string]bool)
	for _, name := range r.paramNames {
		if dups[name] {
			return nil, fmt.Errorf("denco: path parameter `%v' is duplicated in the key `%v'", name, r.Key)
		}
		dups[name] = true
	}
	return &node{data: r.Value, paramNames: r.paramNames}, nil
}

// sibling represents an intermediate data of build for Double-Array.
type sibling struct {
	// An index of start of duplicated characters.
	start int

	// An index of end of duplicated characters.
	end int

	// A character of sibling.
	c byte
}

// nextIndex returns a next index of array of BASE/CHECK.
func nextIndex(base int, c byte) int {
	return base ^ int(c)
}

// makeSiblings returns slice of sibling.
func makeSiblings(records []*record, depth int) (sib []sibling, leaf *record, err error) {
	var (
		pc byte
		n  int
	)
	for i, r := range records {
		if len(r.Key) <= depth {
			leaf = r
			continue
		}
		c := r.Key[depth]
		switch {
		case pc < c:
			sib = append(sib, sibling{start: i, c: c})
		case pc == c:
			continue
		default:
			return nil, nil, fmt.Errorf("denco: BUG: routing table hasn't been sorted")
		}
		if n > 0 {
			sib[n-1].end = i
		}
		pc = c
		n++
	}
	if n == 0 {
		return nil, leaf, nil
	}
	sib[n-1].end = len(records)
	return sib, leaf, nil
}

// Record represents a record data for router construction.
type Record struct {
	// Key for router construction.
	Key string

	// Result value for Key.
	Value interface{}
}

// NewRecord returns a new Record.
func NewRecord(key string, value interface{}) Record {
	return Record{
		Key:   key,
		Value: value,
	}
}

// record represents a record that use to build the Double-Array.
type record struct {
	Record
	paramNames []string
}

// makeRecords returns the records that use to build Double-Arrays.
func makeRecords(srcs []Record) (statics, params []*record) {
	spChars := string([]byte{ParamCharacter, WildcardCharacter})
	termChar := string(TerminationCharacter)
	for _, r := range srcs {
		if strings.ContainsAny(r.Key, spChars) {
			r.Key += termChar
			params = append(params, &record{Record: r})
		} else {
			statics = append(statics, &record{Record: r})
		}
	}
	return statics, params
}

// recordSlice represents a slice of Record for sort and implements the sort.Interface.
type recordSlice []*record

// Len implements the sort.Interface.Len.
func (rs recordSlice) Len() int {
	return len(rs)
}

// Less implements the sort.Interface.Less.
func (rs recordSlice) Less(i, j int) bool {
	return rs[i].Key < rs[j].Key
}

// Swap implements the sort.Interface.Swap.
func (rs recordSlice) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}
