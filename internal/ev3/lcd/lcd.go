package lcd

import (
	"fmt"
	"github.com/ev3go/ev3"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	loadedFont *truetype.Font
)

type Pixel struct {
	X     int
	Y     int
	Color color.Color
}

func Init() (err error) {
	loadedFont, err = loadFont()
	if err != nil {
		log.Printf("ERROR - error while loading font: %v", err)
		return
	}

	return ev3.LCD.Init(true)
}

func SetPixel(x, y int, color color.Color) {
	ev3.LCD.Set(x, y, color)
}

func SetPixels(pixels []Pixel) {
	for _, pixel := range pixels {
		ev3.LCD.Set(pixel.X, pixel.Y, pixel.Color)
	}
}

func SetColor2DArray(xycolors [][]color.Color) {
	for x, ycolors := range xycolors {
		for y, color := range ycolors {
			ev3.LCD.Set(x, y, color)
		}
	}
}

// FillColor fills the LCD with the given color
func FillColor(color color.Color) {
	for x := 0; x < ev3.LCDWidth; x++ {
		for y := 0; y < ev3.LCDHeight; y++ {
			ev3.LCD.Set(x, y, color)
		}
	}
}

func SetImage(image *image.RGBA) {
	dx := image.Bounds().Dx()
	dy := image.Bounds().Dy()

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			ev3.LCD.Set(x, y, image.At(x, y))
		}
	}
}

// Write the given text to the EV3 LCD (display) after clearing the screen
func Write(textContent string) (err error) {
	bgColor := color.White
	bg := image.NewUniform(bgColor)
	draw.Draw(ev3.LCD, ev3.LCD.Bounds(), bg, image.Pt(0, 0), draw.Src)

	return FastWrite(textContent)
}

// FastWrite the given text to the EV3 LCD (display) without clearing the screen first
func FastWrite(textContent string) (err error) {
	fgColor := color.Black
	fontSize := float64(13)

	code := strings.Replace(textContent, "\t", "    ", -1) // convert tabs into spaces
	text := strings.Split(code, "\n")                      // split newlines into arrays

	fg := image.NewUniform(fgColor)

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(loadedFont)
	c.SetFontSize(fontSize)
	c.SetClip(ev3.LCD.Bounds())
	c.SetDst(ev3.LCD)
	c.SetSrc(fg)
	c.SetHinting(font.HintingNone)

	textXOffset := 0
	textYOffset := 0 + int(c.PointToFixed(fontSize)>>6) // Note shift/truncate 6 bits first

	pt := freetype.Pt(textXOffset, textYOffset)
	for _, s := range text {
		_, err = c.DrawString(strings.Replace(s, "\r", "", -1), pt)
		if err != nil {
			return
		}
		pt.Y += c.PointToFixed(fontSize * 1.2)
	}

	return
}

func ShowStatusTTY(b bool) {
	ShowSystemTTY(!b)
}

func ShowSystemTTY(b bool) {
	tty := 5
	if b {
		tty = 2
	}
	exec.Command("/bin/sh", "-c", fmt.Sprintf("echo %s | sudo -S chvt %d", "maker", tty))
}

func loadFont() (f *truetype.Font, err error) {
	fontFile := "/usr/share/fonts/truetype/dejavu/DejaVuSansMono.ttf"
	fontBytes, err := os.ReadFile(fontFile)
	if err != nil {
		return
	}
	f, err = freetype.ParseFont(fontBytes)
	if err != nil {
		return
	}
	return
}
