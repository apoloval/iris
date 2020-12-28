package sdl

import (
	"github.com/JamesHovious/w32"
)

func prepareDPI() {
	w32.SetProcessDPIAware()
}
