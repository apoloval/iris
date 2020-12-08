package gui

import "github.com/apoloval/karen/gfx"

// LabelDefaultTextParams are the default render text parameters for labels
var LabelDefaultTextParams = gfx.RenderTextParams{
	Font:  gfx.TextFontTypeArial,
	Size:  12,
	Color: gfx.ColorWhite,
}

// Label is a widget to represent a single line of text
type Label struct {
	Text       string
	TextParams gfx.RenderTextParams
	Align      gfx.Align

	rendered gfx.RenderedText
}

// NewLabel instantiates a new label widget from the given text
func NewLabel(text string) *Label {
	return &Label{
		Text:       text,
		TextParams: LabelDefaultTextParams,
		Align:      gfx.AlignTo(gfx.AlignLeft, gfx.AlignTop),
	}
}

// Draw this label in the given canvas
func (l *Label) Draw(canvas gfx.Canvas) {
	if l.rendered == nil {
		rt, err := canvas.Engine().RenderText(l.Text, l.TextParams)
		if err != nil {
			panic(err)
		}
		l.rendered = rt
	}
	canvas.DrawText(l.rendered, l.Align)
}
