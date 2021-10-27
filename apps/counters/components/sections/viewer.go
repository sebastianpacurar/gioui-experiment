package sections

import (
	"fmt"
	"gioui-experiment/apps/counters/components/utils"
	g "gioui-experiment/globals"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"image"
	"strconv"
)

type View struct{}

func (v *View) Layout(th *material.Theme, gtx C) D {
	cv := utils.CounterVals
	size := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)
	return layout.Stack{
		Alignment: layout.NW,
	}.Layout(
		gtx,
		layout.Expanded(func(gtx C) D {
			view := g.ColoredArea(
				gtx,
				gtx.Constraints.Constrain(size),
				g.Colours["antique-white"],
			)
			return view
		}),

		layout.Stacked(func(gtx C) D {
			title := func(gtx C) D {
				return layout.Flex{
					Axis: layout.Horizontal,
				}.Layout(gtx,
					layout.Flexed(1, func(gtx C) D {
						text := material.H6(th, fmt.Sprintf("%s", cv.CurrVal))
						return layout.Inset{
							Top: unit.Dp(20),
						}.Layout(gtx, func(gtx C) D {
							return layout.Center.Layout(gtx, text.Layout)
						})
					}),
				)
			}
			return layout.Stack{}.Layout(gtx,
				layout.Expanded(func(gtx C) D {
					clip.UniformRRect(f32.Rectangle{
						Max: layout.FPt(image.Pt(gtx.Constraints.Max.X, 100)),
					}, 0).Add(gtx.Ops)
					return D{}
				}),
				layout.Stacked(title),
			)
		}),

		layout.Stacked(func(gtx C) D {
			var label material.LabelStyle
			if cv.PEnabled || cv.FEnabled || cv.NEnabled {
				label = material.H3(th, strconv.FormatUint(cv.UCount, 10))
			} else if cv.WEnabled {
				label = material.H3(th, strconv.FormatInt(cv.Count, 10))
			}

			return layout.Inset{
				Top:    unit.Dp(80),
				Right:  unit.Dp(50),
				Bottom: unit.Dp(20),
				Left:   unit.Dp(50),
			}.Layout(gtx, func(gtx C) D {
				return label.Layout(gtx)
			})
		}))
}
