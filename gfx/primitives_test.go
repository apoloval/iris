package gfx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPosAdd(t *testing.T) {
	assert.Equal(t, Pos{47, 66}, Pos{13, 54}.Add(Pos{34, 12}))
}

func TestPosSub(t *testing.T) {
	assert.Equal(t, Pos{-21, 42}, Pos{13, 54}.Sub(Pos{34, 12}))
}

func TestRectTopLeft(t *testing.T) {
	assert.Equal(t, Pos{13, 54}, Rect{Pos{13, 54}, Size{39, 80}}.TopLeft())
}

func TestRectTopRight(t *testing.T) {
	assert.Equal(t, Pos{52, 54}, Rect{Pos{13, 54}, Size{39, 80}}.TopRight())
}

func TestRectBottomLeft(t *testing.T) {
	assert.Equal(t, Pos{13, 134}, Rect{Pos{13, 54}, Size{39, 80}}.BottomLeft())
}

func TestRectBottomRight(t *testing.T) {
	assert.Equal(t, Pos{52, 134}, Rect{Pos{13, 54}, Size{39, 80}}.BottomRight())
}
