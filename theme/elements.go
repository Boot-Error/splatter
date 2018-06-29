package theme

import (
	"fmt"
	"strconv"
)

type Box interface {
	FromThemeElement(e Element)
	ComputeBoxParams() (x, y, w, h float64)
	Paint(surface *cairo.Surface)
}

type Color struct {
	Red   uint
	Blue  uint
	Green uint
}

type BoxEdge struct {
	Left   int64
	Bottom int64
	Top    int64
	Right  int64
}

type PhantomBox struct {
	Width  int64
	Height int64
	PosX   int64
	PosY   int64
	Border struct {
		Width int64
		Color Color
	}
	Padding BoxEdge
	Margin  BoxEdge
	BgColor Color
}

type TextBox struct {
	BoxProp  PhantomBox
	Text     string
	Fontface string
	Fontsize int64
}

type ImageBox struct {
	BoxProp       PhantomBox
	ImageFileName string
}

type GridBox struct {
	BoxProp PhantomBox
	Flow    string
	Boxes   []Box
}

// convert the properties from the theme file to box structs
func (b *PhantomBox) fromProperties(props map[string]string) {

	if val, ok := props["height"]; ok {
		b.Width, _ = strconv.ParseInt(val, 10, 64)
	} else {
		b.Width = 0
	}

	if val, ok := props["width"]; ok {
		b.Height, _ = strconv.ParseInt(val, 10, 64)
	} else {
		b.Height = 0
	}

	if val, ok := props["left"]; ok {
		b.PosX, _ = val
	} else {
		b.PosX = 0
	}

	if val, ok := props["top"]; ok {
		b.PosY, _ = val
	} else {
		b.PosY = 0
	}

	if val, ok := props["border"]; ok {

		var borderWidth int64
		var red, blue, green uint

		fmt.Sscanf(
			val,
			"%dpx rgb(%d, %d, %d)",
			&borderWidth,
			&red, &blue, &green,
		)
		b.Border.Width = borderWidth
		b.Border.Color = Color{
			Red:   red,
			Green: green,
			Blue:  blue,
		}

	} else {

		b.Border.width = 0
		b.Border.color = Color{
			Red: 0, Green: 0, Blue: 0,
		}
	}

	if val, ok := props["margin"]; ok {

		var left, bottom, top, right int64
		fmt.Sscanf(
			val,
			"%d %d %d %d",
			&left, &bottom, &top, &right,
		)
		b.Margin = BoxEdge{
			Left: left, Bottom: bottom, Top: top, Right: right,
		}

	} else {

		b.Margin = BoxEdge{
			Left: 0, Bottom: 0, Top: 0, Right: 0,
		}
	}

	if val, ok := props["padding"]; ok {

		var left, bottom, top, right int64
		fmt.Sscanf(
			val,
			"%d %d %d %d",
			&left, &bottom, &top, &right,
		)
		b.Margin = BoxEdge{
			Left: left, Bottom: bottom, Top: top, Right: right,
		}

	} else {

		b.Margin = BoxEdge{
			Left: 0, Bottom: 0, Top: 0, Right: 0,
		}
	}

	if val, ok := props["color"]; ok {
		var red, green, blue uint
		fmt.Sscanf(val, "rgb(%d, %d, %d)", &red, &green, &blue)
		b.BgColor = Color{Red: red, Green: green, Blue: blue}

	} else {
		b.BgColor = Color{Red: 255, Green: 255, Blue: 255}
	}
}

func (b *PhantomBox) FromThemeElement(e Element) {
	boxProp.fromProperties(
		e.(map[string]interface{})["properties"].(map[string]string),
	)
}

func (tb *TextBox) FromThemeElement(e Element) {

	var boxProp PhantomBox
	boxProp.fromProperties(
		e.(map[string]interface{})["properties"].(map[string]string),
	)
	tb.BoxProp = boxProp

	if features, ok := e["feature"]; ok {
		if tb.Text, exists = features["text"]; !exists {
			fmt.Println("Error: Text not ")
		}
	} else {
		fmt.Println("Error: Feature not found for element")
	}

}
