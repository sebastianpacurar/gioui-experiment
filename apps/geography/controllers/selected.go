package controllers

import (
	"gioui-experiment/apps/geography/data"
	"gioui-experiment/globals"
	"gioui-experiment/themes/colours"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/outlay"
)

type (
	SelectedCountries struct {
		pills []pill
		api   data.Countries
		wrap  outlay.GridWrap
		count int
	}

	pill struct {
		name      string
		content   data.Country
		btn       widget.Clickable
		removeBtn widget.Clickable
	}
)

func (sc *SelectedCountries) Layout(gtx C, th *material.Theme) D {
	var dims D

	// avoid continuous reiteration
	selectedCount := data.GetSelectedCount()
	if sc.count != selectedCount {
		selected := data.GetSelected()
		sc.pills = make([]pill, selectedCount)
		for i := range sc.pills {
			sc.pills[i].name = selected[i].Name.Common
			sc.pills[i].content = selected[i]
		}
		sc.count = selectedCount
	}

	if selectedCount > 0 {
		dims = layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				return sc.wrap.Layout(gtx, sc.count, func(gtx C, i int) D {

					if sc.pills[i].removeBtn.Clicked() {
						name := sc.pills[i].name
						for i := range data.Cached {
							if data.Cached[i].Name.Common == name {
								data.Cached[i].Selected = false
							}
						}
					}

					if sc.pills[i].btn.Clicked() {
						name := sc.pills[i].name
						for i := range data.Cached {
							if data.Cached[i].Name.Common == name {
								data.Cached[i].IsCPViewed = true
							} else {
								data.Cached[i].IsCPViewed = false
							}
						}
					}

					var area D
					area = layout.Inset{
						Top:    unit.Dp(4),
						Right:  unit.Dp(2),
						Bottom: unit.Dp(4),
						Left:   unit.Dp(2),
					}.Layout(gtx, func(gtx C) D {
						border := widget.Border{
							Color:        globals.Colours[colours.GREY],
							CornerRadius: unit.Dp(1),
							Width:        unit.Dp(1),
						}
						return border.Layout(gtx, func(gtx C) D {
							return layout.Flex{Spacing: layout.SpaceBetween}.Layout(gtx,
								layout.Rigid(func(gtx C) D {
									var btn material.ButtonStyle
									btn = material.Button(th, &sc.pills[i].btn, sc.pills[i].content.Name.Common)
									btn.CornerRadius = unit.Dp(10)
									btn.Background = globals.Colours[colours.AERO_BLUE]
									btn.Color = globals.Colours[colours.BLACK]
									btn.TextSize = th.TextSize.Scale(14.0 / 16.0)
									return btn.Layout(gtx)
								}),
								layout.Rigid(func(gtx C) D {
									var btn material.ButtonStyle
									btn = material.Button(th, &sc.pills[i].removeBtn, "X")
									btn.CornerRadius = unit.Dp(0)
									btn.Background = globals.Colours[colours.FLAME_RED]
									btn.Color = globals.Colours[colours.WHITE]
									btn.TextSize = th.TextSize.Scale(14.0 / 16.0)
									return btn.Layout(gtx)
								}))
						})
					})
					return area
				})
			}))
	} else {
		dims = material.Body1(th, "Nothing selected").Layout(gtx)
	}
	return dims
}
