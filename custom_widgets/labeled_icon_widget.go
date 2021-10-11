package custom_widgets

import (
	"gioui-experiment/globals"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
	"image/color"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

// LabeledIconBtn - component which returns a custom-made widget, based on its given properties
type LabeledIconBtn struct {
	Theme               *material.Theme
	BgColor, LabelColor color.NRGBA
	Button              *widget.Clickable
	Icon                *widget.Icon
	Label               string
}

// Layout - currently returns a button-like widget, which contains:
// an icon on the left side and a text label n the right side.
// These 2 are separated by 5 device pixels
func (lib LabeledIconBtn) Layout(gtx C) D {

	// Set custom colours for Icon and Label
	// Set the TextSize for the Label
	lib.Theme.Palette.ContrastBg = lib.BgColor
	lib.Theme.TextSize.Scale(14.0 / 16.0)

	// return a ButtonLayout dimension which contains the Icon and Label
	return material.ButtonLayout(lib.Theme, lib.Button).Layout(gtx, func(gtx C) D {
		return layout.UniformInset(globals.DefaultMargin).Layout(gtx, func(gtx C) D {

			// labeledIcon will be used to return all its content dimensions
			labeledIcon := layout.Flex{
				Axis:      layout.Horizontal,
				Alignment: layout.Middle,
			}
			// spacer is used for separating the Icon from the Label by an amount of pixels
			spacer := unit.Dp(5)
			// This is the actual Icon of the Button. It will return the dimensions of the Icon widget
			layIcon := layout.Rigid(func(gtx C) D {
				return layout.Inset{
					Right: spacer,
				}.Layout(gtx, func(gtx C) D {
					var d D
					if lib.Icon != nil {
						size := gtx.Px(unit.Dp(56)) - 2*gtx.Px(unit.Dp(16))
						gtx.Constraints = layout.Exact(image.Pt(size, size))
						d = lib.Icon.Layout(gtx, lib.LabelColor)
					}
					return d
				})
			})

			// This is the actual Label of the Button, treated as a Body1 Typography element
			layLabel := layout.Rigid(func(gtx C) D {
				return layout.Inset{Left: spacer}.Layout(gtx, func(gtx C) D {
					l := material.Body1(lib.Theme, lib.Label)
					l.Color = lib.LabelColor
					return l.Layout(gtx)
				})
			})

			// eventually return labeledIcon with all its children added as parameters in the right order
			return labeledIcon.Layout(gtx, layIcon, layLabel)
		})
	})
}
