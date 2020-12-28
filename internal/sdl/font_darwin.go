package sdl

import (
	"fmt"

	"github.com/apoloval/iris/gfx"
)

func fontPath(fontType gfx.TextFontType) string {
	return fmt.Sprintf("/Library/Fonts/%s.ttf", fontType)
}
