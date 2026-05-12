package normalizers

import (
	"strings"
	"sync"
)

func NormalizeRegularName(name string) string {

	builder := new(strings.Builder)
	dict := innerGetRegularNameDict()
	src := []rune(name)

	dict.handleString(src, builder)

	dst := builder.String()
	return dst
}

////////////////////////////////////////////////////////////////////////////////

type innerRegularNameDict struct {
	mu sync.Mutex

	table map[rune]rune

	bypass rune // = 0
	skip   rune // = -1

}

var theRegularNameDict *innerRegularNameDict

func innerGetRegularNameDict() *innerRegularNameDict {

	dict := theRegularNameDict
	if dict == nil {
		dict = new(innerRegularNameDict)
		dict.init()
		theRegularNameDict = dict
	}
	return dict
}

func (inst *innerRegularNameDict) init() {

	const (
		skip   = -1
		bypass = 0
	)
	t := make(map[rune]rune)

	t[' '] = skip
	t['\t'] = skip
	t['('] = bypass
	t[')'] = bypass
	t['（'] = '('
	t['）'] = ')'

	inst.table = t
	inst.skip = skip
}

func (inst *innerRegularNameDict) handleRune(r rune, dst *strings.Builder) {

	r2 := inst.table[r]

	if r2 > 0 {
		dst.WriteRune(r2)
	} else if r2 == 0 {
		dst.WriteRune(r)
	}

}

func (inst *innerRegularNameDict) handleString(str []rune, dst *strings.Builder) {

	mu := &inst.mu
	mu.Lock()
	defer mu.Unlock()

	for _, r := range str {
		inst.handleRune(r, dst)
	}
}
