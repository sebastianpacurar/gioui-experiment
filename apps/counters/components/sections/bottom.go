package sections

import (
	"gioui-experiment/apps/counters/components/controllers"
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Bottom struct {
		ValueHandlers controllers.ValueHandler
		Incrementor   controllers.Incrementor
	}
)

// TODO
// Will probably give up on the Bottom section to give more room for displaying huge and multiple data at the same time
func (b *Bottom) Layout(th *material.Theme, gtx C) D {
	return layout.Flex{
		Axis:      layout.Horizontal,
		Alignment: layout.Middle,
		Spacing:   layout.SpaceEvenly,
		WeightSum: 2,
	}.Layout(gtx,
		layout.Flexed(1, func(gtx C) D {
			return b.Incrementor.Layout(th, gtx)
		}),
		layout.Flexed(1, func(gtx C) D {
			return b.ValueHandlers.Layout(th, gtx)
		}),
	)
}
