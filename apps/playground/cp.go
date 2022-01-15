package playground

import (
	"gioui-experiment/apps/playground/controllers/counter"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type (
	ControlPanel struct {
		controllers []Controller
		list        widget.List

		vh       counter.ValueHandler
		inc      counter.Incrementor
		sequence counter.Sequence
		status   counter.Status

		// hardcoded in order to keep track of the specific current state
		incState     component.DiscloserState
		vhState      component.DiscloserState
		displayState component.DiscloserState
		statusState  component.DiscloserState
	}

	Controller struct {
		name   string
		layout func(C, *Controller) D
	}
)

var controllerInset = layout.Inset{
	Top:    unit.Dp(10),
	Right:  unit.Dp(25),
	Bottom: unit.Dp(10),
}

func (cp *ControlPanel) Layout(gtx C, th *material.Theme) D {
	cp.list.Axis = layout.Vertical

	divider := layout.Rigid(func(gtx C) D {
		div := component.Divider(th)
		div.Left = unit.Dp(5)
		div.Right = unit.Dp(5)
		return div.Layout(gtx)
	})

	// every controller is a vertical flex which contains 2 rigids - discloser and the divider
	cp.controllers = []Controller{
		{
			name: "Sequence",
			layout: func(gtx C, c *Controller) D {
				content := layout.Rigid(func(gtx C) D {
					return component.SimpleDiscloser(th, &cp.displayState).Layout(gtx,
						material.Body1(th, c.name).Layout,
						func(gtx C) D {
							return controllerInset.Layout(gtx, func(gtx C) D {
								return cp.sequence.Layout(gtx, th)
							})
						})
				})
				return cp.LayOutset(gtx, content, divider)
			},
		},

		//TODO: heads up on the Start and Step Values layout
		// after opening the discloser the right side of the border goes out of frame
		// closing the discloser causes the border to reposition accordingly
		{
			name: "Start and Step Values",
			layout: func(gtx C, c *Controller) D {
				content := layout.Rigid(func(gtx C) D {
					return component.SimpleDiscloser(th, &cp.vhState).Layout(gtx,
						material.Body1(th, c.name).Layout,
						func(gtx C) D {
							return controllerInset.Layout(gtx, func(gtx C) D {
								return cp.vh.Layout(gtx, th)
							})
						})
				})
				return cp.LayOutset(gtx, content, divider)
			},
		},
		{
			name: "Manual Incrementors",
			layout: func(gtx C, c *Controller) D {
				content := layout.Rigid(func(gtx C) D {
					return component.SimpleDiscloser(th, &cp.incState).Layout(gtx,
						material.Body1(th, c.name).Layout,
						func(gtx C) D {
							return controllerInset.Layout(gtx, func(gtx C) D {
								return cp.inc.Layout(gtx, th)
							})
						})
				})
				return cp.LayOutset(gtx, content, divider)
			},
		},
		{
			name: "Stats",
			layout: func(gtx C, c *Controller) D {
				content := layout.Rigid(func(gtx C) D {
					return component.SimpleDiscloser(th, &cp.statusState).Layout(gtx,
						material.Body1(th, c.name).Layout,
						func(gtx C) D {
							return controllerInset.Layout(gtx, func(gtx C) D {
								return cp.status.Layout(gtx, th)
							})
						})
				})
				return cp.LayOutset(gtx, content, divider)
			},
		},
	}

	// return a vertical list of (discloser, divider) groups, as ListElements
	return material.List(th, &cp.list).Layout(gtx, len(cp.controllers), func(gtx C, i int) D {
		return cp.controllers[i].layout(gtx, &cp.controllers[i])
	})
}

// LayOutset - wraps the discloser and divider in a vertical flex layout
func (cp *ControlPanel) LayOutset(gtx C, discloser, divider layout.FlexChild) D {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, discloser, divider)
}
