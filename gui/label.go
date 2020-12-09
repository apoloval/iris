package gui

import "github.com/apoloval/karen/gfx"

// Label is a widget to represent a single line of text
type Label struct {
	LabelProps

	rendered gfx.RenderedText
}

// LabelProps are the properties of a label widget
type LabelProps struct {
	WidgetProps
	TextProps

	Text string
}

// Apply the label properties
func (p *LabelProps) Apply(props []LabelProp) {
	for _, prop := range props {
		prop(p)
	}
}

// LabelProp is a single property of a label widget
type LabelProp func(*LabelProps)

// LabelDefaultProps are the default properties for a label
var LabelDefaultProps = struct {
	WidgetProps []WidgetProp
	TextProps   []TextProp
	LabelProps  []LabelProp
}{
	WidgetProps: []WidgetProp{
		WithAlign(gfx.AlignTopLeft),
	},
	TextProps: []TextProp{
		WithFontSize(12),
		WithFontType(gfx.TextFontTypeArial),
		WithFontColor(gfx.ColorWhite),
	},
}

// NewLabel instantiates a new label widget from the given text
func NewLabel(text string, props ...LabelProp) *Label {
	label := &Label{}
	label = label.WithWidgetProps(LabelDefaultProps.WidgetProps...)
	label = label.WithTextProps(LabelDefaultProps.TextProps...)
	label.Apply(props)
	label.Text = text
	return label
}

// WithWidgetProps applies the given widget properties
func (l *Label) WithWidgetProps(props ...WidgetProp) *Label {
	l.WidgetProps.Apply(props)
	return l
}

// WithTextProps applies the given text properties
func (l *Label) WithTextProps(props ...TextProp) *Label {
	l.TextProps.Apply(props)
	return l
}

// Draw this label in the given canvas
func (l *Label) Draw(canvas gfx.Canvas) {
	if l.rendered == nil {
		textParams := gfx.RenderTextParams{
			Font:  l.FontType,
			Size:  l.FontSize,
			Color: l.FontColor,
		}
		rt, err := canvas.Engine().RenderText(l.Text, textParams)
		if err != nil {
			panic(err)
		}
		l.rendered = rt
	}
	canvas = gfx.WithPadding(canvas, l.Padding)
	dst := l.Align(l.rendered.Size().ToRect(), canvas.Size().ToRect())
	canvas.DrawText(dst, l.rendered)
}
