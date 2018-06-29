package theme

import (
	"fmt"

	"github.com/ungerik/go-cairo"
)

func generateBoard(outfile string, boxes []Box) {

	surface, _ := cairo.NewSurfaceFromPNG("backgrounds/back1.png")

	for _, box := range boxes {
		box.Paint(surface)
	}

}

func (b *PhantomBox) ComputeBoxParams() (x, y, w, h float64) {

	// specified info
	x = float64(b.PosX)
	y = float64(b.PosY)
	w = float64(b.Width)
	h = float64(b.Height)

	// change in width due to margin padding
	w += b.Margin.Left + b.Margin.Right
	w += b.Padding.Left + b.Padding.Right

	h += b.Margin.Top + b.Margin.Bottom
	h += b.Padding.Top + b.Padding.Bottom

	// adding border width
	w += b.Border.Width
	h += b.Border.Width

	return
}

func (b *PhantomBox) Paint(surface *cairo.Surface) {

	x, y, w, h := b.ComputeBoxParams()

	surface.MoveTo(x, y)
	surface.SetSourceRGB(
		float64(b.BgColor.Red),
		float64(b.BgColor.Green),
		float64(b.BgColor.Blue),
	)
	surface.Rectangle(x, y, w, h)
	surface.Fill()

}
