package normalizers

import (
	"strings"
	"sync"
)

////////////////////////////////////////////////////////////////////////////////

type CodeNormalizer interface {
	Normalize(code string) string
}

type CodeNormalizerConfigFn func(b *CodeNormalizerBuilder)

type CodeNormalizerMapperFn func(ch rune) rune

////////////////////////////////////////////////////////////////////////////////

type CodeNormalizerHolder struct {
	normalizer CodeNormalizer
	mu         sync.Mutex
}

func (inst *CodeNormalizerHolder) GetNormalizer(config CodeNormalizerConfigFn) CodeNormalizer {
	n := inst.normalizer
	if n == nil {
		n = inst.innerGetNormalizer2(config)
	}
	return n
}

func (inst *CodeNormalizerHolder) innerGetNormalizer2(config CodeNormalizerConfigFn) CodeNormalizer {

	mu := &inst.mu
	mu.Lock()
	defer mu.Unlock()

	n := inst.normalizer
	if n == nil {
		n = inst.innerLoadNormalizer(config)
		inst.normalizer = n
	}

	return n
}

func (inst *CodeNormalizerHolder) innerLoadNormalizer(config CodeNormalizerConfigFn) CodeNormalizer {
	b := new(CodeNormalizerBuilder)
	if config != nil {
		config(b)
	}
	return b.Build()
}

////////////////////////////////////////////////////////////////////////////////

type CodeNormalizerBuilder struct {
	useDigit bool
	useUpper bool
	useLower bool

	marks []*innerRuneMapping
}

func (inst *CodeNormalizerBuilder) UseUpperCase() *CodeNormalizerBuilder {
	inst.useUpper = true
	return inst
}

func (inst *CodeNormalizerBuilder) UseLowerCase() *CodeNormalizerBuilder {
	inst.useLower = true
	return inst
}

func (inst *CodeNormalizerBuilder) UseDigit() *CodeNormalizerBuilder {
	inst.useDigit = true
	return inst
}

func (inst *CodeNormalizerBuilder) UseMarks(marks ...rune) *CodeNormalizerBuilder {
	for _, mark := range marks {
		item := &innerRuneMapping{
			from: mark,
			to:   mark,
		}
		inst.marks = append(inst.marks, item)
	}
	return inst
}

func (inst *CodeNormalizerBuilder) ReplaceMark(from, to rune) *CodeNormalizerBuilder {
	item := &innerRuneMapping{
		from: from,
		to:   to,
	}
	inst.marks = append(inst.marks, item)
	return inst
}

func (inst *CodeNormalizerBuilder) Build() CodeNormalizer {

	normalizer := new(innerCommonCodeNormalizer)
	normalizer.mappings = inst.marks

	if inst.useDigit {
		normalizer.mapperForDigit = normalizer.handleRuneBypass
	} else {
		normalizer.mapperForDigit = normalizer.handleRuneIgnore
	}

	inst.configureNormalizer(normalizer)

	return normalizer
}

func (inst *CodeNormalizerBuilder) configureNormalizer(n *innerCommonCodeNormalizer) {

	// config letters
	letterCount := 0
	letterUpper := false
	if inst.useUpper {
		letterCount++
		letterUpper = true
	}
	if inst.useLower {
		letterCount++
	}

	switch letterCount {
	case 0:
		n.mapperForUpperCase = n.handleRuneIgnore
		n.mapperForLowerCase = n.handleRuneIgnore
		return

	case 2:
		n.mapperForUpperCase = n.handleRuneBypass
		n.mapperForLowerCase = n.handleRuneBypass
		return
	}

	if letterUpper {
		fn := n.handleRuneToUpper
		n.mapperForUpperCase = n.handleRuneBypass
		n.mapperForLowerCase = fn
	} else {
		fn := n.handleRuneToLower
		n.mapperForUpperCase = fn
		n.mapperForLowerCase = n.handleRuneBypass
	}
}

////////////////////////////////////////////////////////////////////////////////

type innerCommonCodeNormalizer struct {
	mappings []*innerRuneMapping

	mapperForDigit     CodeNormalizerMapperFn
	mapperForUpperCase CodeNormalizerMapperFn
	mapperForLowerCase CodeNormalizerMapperFn
}

// Normalize implements CodeNormalizer.
func (inst *innerCommonCodeNormalizer) Normalize(code string) string {

	m := new(innerCommonCodeMapping)
	b := new(strings.Builder)
	src := []rune(code)

	m.init(inst)

	for _, ch := range src {
		m.handleRune(ch, b)
	}

	return b.String()
}

func (inst *innerCommonCodeNormalizer) handleRuneBypass(ch rune) rune {
	return ch
}

func (inst *innerCommonCodeNormalizer) handleRuneIgnore(ch rune) rune {
	return 0
}

func (inst *innerCommonCodeNormalizer) handleRuneToUpper(ch rune) rune {
	const (
		upperA rune = 'A'
		lowerA rune = 'a'
	)
	return (ch - lowerA + upperA)
}

func (inst *innerCommonCodeNormalizer) handleRuneToLower(ch rune) rune {
	const (
		upperA rune = 'A'
		lowerA rune = 'a'
	)
	return (ch + lowerA - upperA)
}

func (inst *innerCommonCodeNormalizer) _impl() CodeNormalizer {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerRuneMapping struct {
	from rune
	to   rune
}

////////////////////////////////////////////////////////////////////////////////

type innerCommonCodeMapping struct {
	mappings map[rune]*innerRuneMapping

	mapperForDigit     CodeNormalizerMapperFn
	mapperForUpperCase CodeNormalizerMapperFn
	mapperForLowerCase CodeNormalizerMapperFn
	mapperForOther     CodeNormalizerMapperFn
}

func (inst *innerCommonCodeMapping) init(parent *innerCommonCodeNormalizer) {

	dst := make(map[rune]*innerRuneMapping)
	src := parent.mappings

	for _, it := range src {
		dst[it.from] = it
	}

	inst.mapperForDigit = parent.mapperForDigit
	inst.mapperForUpperCase = parent.mapperForUpperCase
	inst.mapperForLowerCase = parent.mapperForLowerCase
	inst.mapperForOther = inst.handleOtherMarks
	inst.mappings = dst
}

func (inst *innerCommonCodeMapping) handleOtherMarks(ch rune) rune {
	m := inst.mappings[ch]
	if m == nil {
		return 0
	}
	return m.to
}

func (inst *innerCommonCodeMapping) handleRune(ch1 rune, b *strings.Builder) {

	var ch2 rune

	if ('0' <= ch1) && (ch1 <= '9') {
		ch2 = inst.mapperForDigit(ch1)

	} else if ('A' <= ch1) && (ch1 <= 'Z') {
		ch2 = inst.mapperForUpperCase(ch1)

	} else if ('a' <= ch1) && (ch1 <= 'z') {
		ch2 = inst.mapperForLowerCase(ch1)

	} else {
		ch2 = inst.mapperForOther(ch1)
	}

	if ch2 > 0 {
		b.WriteRune(ch2)
	}
}

////////////////////////////////////////////////////////////////////////////////
// EOF
