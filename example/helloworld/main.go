package main

import (
	"fmt"
	"runtime"

	"github.com/apoloval/iris"
	"github.com/apoloval/iris/gfx"
)

func main() {
	app, err := iris.NewApp()
	if err != nil {
		panic(err)
	}

	col := gfx.ColorWhite
	for {
		app.BeginFrame()

		app.BeginLayoutH(iris.Padding(10))
		app.BeginLayoutV(iris.Padding(10))
		app.BeginLayoutV(iris.Padding(5))

		if app.Label(1, "Hello World!", iris.FontColor(col)) {
			col = gfx.ColorRed
		} else {
			col = gfx.ColorWhite
		}

		stats := app.Stats()
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)

		app.Label(2, "Performance statistics:")

		app.BeginLayoutH(iris.Padding(5))

		app.BeginLayoutV(iris.Expand(300))
		app.Label(3, "Frames per second :", iris.Align(gfx.AlignRight))
		app.Label(4, "Frame render time :", iris.Align(gfx.AlignRight))
		app.Label(5, "Allocated memory :", iris.Align(gfx.AlignRight))
		app.EndLayout()

		app.BeginLayoutV()
		app.Label(6, fmt.Sprintf("%.2f", stats.FramesPerSecond))
		app.Label(7, fmt.Sprintf("%v", stats.FrameRenderTime))
		app.Label(8, bytesForHuman(mem.Alloc))
		app.EndLayout()

		app.EndLayout()

		app.EndLayout()
		app.EndLayout()
		app.EndLayout()

		if app.EndFrame() {
			break
		}
	}
}

func bytesForHuman(bytes uint64) string {
	if bytes < 10*1024 {
		return fmt.Sprintf("%d B", bytes)
	} else if bytes < 10*1024*1024 {
		return fmt.Sprintf("%d KB", bytes/1024)
	} else if bytes < 10*1024*1024*1024 {
		return fmt.Sprintf("%d MB", bytes/(1024*1024))
	} else if bytes < 10*1024*1024*1024*1024 {
		return fmt.Sprintf("%d GB", bytes/(1024*1024*1024))
	}
	return fmt.Sprintf("%d B", bytes)
}
