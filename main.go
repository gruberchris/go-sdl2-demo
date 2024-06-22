package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		1280, 768, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	err = surface.FillRect(nil, 0)
	if err != nil {
		panic(err)
	}

	rect := sdl.Rect{100, 0, 200, 200}
	rect2 := sdl.Rect{200, 100, 200, 200}
	rect3 := sdl.Rect{300, 200, 200, 200}
	rect4 := sdl.Rect{400, 300, 200, 200}
	rect5 := sdl.Rect{500, 400, 200, 200}
	purple := sdl.Color{R: 255, G: 0, B: 255, A: 255}
	cyan := sdl.Color{R: 0, G: 255, B: 255, A: 255}
	yellow := sdl.Color{R: 255, G: 255, B: 0, A: 255}
	blue := sdl.Color{R: 0, G: 0, B: 255, A: 255}
	green := sdl.Color{R: 0, G: 255, B: 0, A: 255}
	pixel := sdl.MapRGBA(surface.Format, purple.R, purple.G, purple.B, purple.A)
	pixel2 := sdl.MapRGBA(surface.Format, cyan.R, cyan.G, cyan.B, cyan.A)
	pixel3 := sdl.MapRGBA(surface.Format, yellow.R, yellow.G, yellow.B, yellow.A)
	pixel4 := sdl.MapRGBA(surface.Format, blue.R, blue.G, blue.B, blue.A)
	pixel5 := sdl.MapRGBA(surface.Format, green.R, green.G, green.B, green.A)

	err = surface.FillRect(&rect, pixel)
	if err != nil {
		panic(err)
	}

	err = surface.FillRect(&rect2, pixel2)
	if err != nil {
		panic(err)
	}

	err = surface.FillRect(&rect3, pixel3)
	if err != nil {
		panic(err)
	}

	err = surface.FillRect(&rect4, pixel4)
	if err != nil {
		panic(err)
	}

	err = surface.FillRect(&rect5, pixel5)
	if err != nil {
		panic(err)
	}

	err = window.UpdateSurface()
	if err != nil {
		panic(err)
	}

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.MouseButtonEvent:
				println("MouseButtonEvent")
				break
			}
		}
	}
}
