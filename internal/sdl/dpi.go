package sdl

import (
	"github.com/veandco/go-sdl2/sdl"
)

type dpi struct {
	x int
	y int
}

func initDPI() (*dpi, error) {
	prepareDPI()
	_, h, v, err := sdl.GetDisplayDPI(0)
	if err != nil {
		return nil, err
	}

	return &dpi{
		x: int(h),
		y: int(v),
	}, nil
}

func (d *dpi) scaleX(size int) int {
	return d.scale(d.x, size)
}

func (d *dpi) scaleY(size int) int {
	return d.scale(d.y, size)
}

func (d *dpi) scale(dpi, size int) int {
	switch dpi {
	case 96:
		return size
	case 120:
		return size * 5 / 2
	case 144:
		return size * 3 / 2
	case 192:
		return size * 2
	default:
		return size
	}
}
