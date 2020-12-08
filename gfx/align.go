package gfx

// AlignH is an horizontal alignment policy
type AlignH interface {
	CalculateX(src, dst int) int
}

// AlignV is a vertical alignment policy
type AlignV interface {
	CalculateY(src, dst int) int
}

// Align is a horizontal and vertical alignment policy
type Align interface {
	AlignH
	AlignV
}

// Alignment definitions
var (
	AlignLeft   AlignH = alignLeft{}
	AlignRight  AlignH = alignRight{}
	AlignTop    AlignV = alignTop{}
	AlignBottom AlignV = alignBottom{}
	AlignCenter Align  = alignCenter{}
)

type alignLeft struct{}

func (alignLeft) CalculateX(src, dst int) int {
	return 0
}

type alignRight struct{}

func (alignRight) CalculateX(src, dst int) int {
	return dst - src
}

type alignTop struct{}

func (alignTop) CalculateY(src, dst int) int {
	return 0
}

type alignBottom struct{}

func (alignBottom) CalculateY(src, dst int) int {
	return dst - src
}

type alignCenter struct{}

func (alignCenter) CalculateX(src, dst int) int {
	return (dst - src) / 2
}

func (alignCenter) CalculateY(src, dst int) int {
	return (dst - src) / 2
}

// AlignTo combines the given horizontal and vertical alignments
func AlignTo(h AlignH, v AlignV) Align {
	return alignPair{h: h, v: v}
}

type alignPair struct {
	h AlignH
	v AlignV
}

func (p alignPair) CalculateX(src, dst int) int {
	return p.h.CalculateX(src, dst)
}

func (p alignPair) CalculateY(src, dst int) int {
	return p.v.CalculateY(src, dst)
}
