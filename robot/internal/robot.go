package internal

import (
	"fmt"
	"image"
	"strconv"

	"github.com/go-vgo/robotgo"
)

type DIR struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

type SpeechBox struct {
	X      int
	Y      int
	Width  int
	Height int
}

func CaptureScreen(dir *SpeechBox) *robotgo.CBitmap {
	bit := robotgo.CaptureScreen(dir.X, dir.Y, dir.Width, dir.Height)
	return &bit
}

func FreeBitmap(bitmap *robotgo.CBitmap) {
	robotgo.FreeBitmap(*bitmap)
}

func CollectRedPixels(dir *DIR) []image.Point {
	var redPixels []image.Point

	//fmt.Println(dir.X1, dir.Y1, dir.X2, dir.Y2)
	for y := dir.Y1; y < dir.Y2; y++ {
		for x := dir.X1; x < dir.X2; x++ {

			color := robotgo.GetPixelColor(x, y)
			r, g, b := hexToRGB(color)
			if (b > 70 && g > 90 && r > 220) && (b < 90 && g < 110 && r < 240) {
				fmt.Println(r, g, b)
				redPixels = append(redPixels, image.Point{x, y})
			}
		}
	}
	return redPixels
}
func hexToRGB(color string) (int64, int64, int64) {
	r, _ := strconv.ParseInt(color[0:2], 16, 0)
	g, _ := strconv.ParseInt(color[2:4], 16, 0)
	b, _ := strconv.ParseInt(color[4:6], 16, 0)
	return r, g, b
}

func Click(x, y int) {
	robotgo.Move(x, y)
	robotgo.MilliSleep(100)
	robotgo.Click("left", false)
	robotgo.MilliSleep(100)
	robotgo.Click("left", true)
}

func ScrollDown() {
	robotgo.MilliSleep(100)
	robotgo.Scroll(0, -50)
	robotgo.MilliSleep(100)
}

func Input(keys ...string) {
	robotgo.MilliSleep(100)
	if len(keys) == 1 {
		robotgo.KeyTap(keys[0])
	}
}

func EnterInput(res string) {
	str := fmt.Sprint(res)
	robotgo.MilliSleep(100)
	robotgo.TypeStr(str)
	robotgo.MilliSleep(100)
	Input("enter")
}
