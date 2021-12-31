package table

import (
	"fmt"
	"gioui-experiment/apps/geography/components/countries/data"
	g "gioui-experiment/globals"
	"gioui-experiment/themes/colors"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"image"
)

type (
	C = layout.Context
	D = layout.Dimensions

	Table struct {
		rows   []row
		list   widget.List
		loaded bool
	}

	row struct {
		click    widget.Clickable
		name     string
		cca2     string
		cca3     string
		ccn3     string
		area     float64
		active   bool
		selected bool
	}
)

func (t *Table) Layout(gtx C, th *material.Theme) D {
	t.list.Axis = layout.Vertical
	t.list.Alignment = layout.Middle

	if !t.loaded {
		for i := range data.Data {
			t.rows = append(t.rows, row{
				name:   data.Data[i].Name.Common,
				cca2:   data.Data[i].Cca2,
				cca3:   data.Data[i].Cca3,
				ccn3:   data.Data[i].Ccn3,
				area:   data.Data[i].Area,
				active: data.Data[i].Active,
			})
		}
		t.loaded = true
	} else {
		for i := range data.Data {
			t.rows[i].active = data.Data[i].Active
		}
	}
	border := widget.Border{
		Color: g.Colours[colors.GREY],
		Width: unit.Px(1),
	}

	return material.List(th, &t.list).Layout(gtx, len(data.Data), func(gtx C, i int) D {
		return material.Clickable(gtx, &t.rows[i].click, func(gtx C) D {
			var content D
			rowColor := g.Colours[colors.ANTIQUE_WHITE]

			if t.rows[i].selected {
				rowColor = g.Colours[colors.AERO_BLUE]
			}

			if t.rows[i].active {
				if t.rows[i].click.Clicked() {
					if t.rows[i].selected {
						t.rows[i].selected = false
					} else {
						t.rows[i].selected = true
					}
				}

				if t.rows[i].click.Hovered() {
					data.Data[i].Hovered = true
					if !t.rows[i].selected {
						rowColor = g.Colours[colors.NYANZA]
					} else {
						rowColor = g.Colours[colors.LIGHT_SALMON]
					}
				} else {
					data.Data[i].Hovered = false
				}

				content = layout.Flex{}.Layout(gtx,
					layout.Flexed(4, func(gtx C) D {
						return border.Layout(gtx, func(gtx C) D {
							return layout.Inset{
								Top:    unit.Dp(2),
								Bottom: unit.Dp(2),
								Left:   unit.Dp(2),
							}.Layout(gtx, func(gtx C) D {
								return layout.Stack{Alignment: layout.Center}.Layout(gtx,

									layout.Expanded(func(gtx C) D {
										return g.ColoredArea(gtx, image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y), rowColor)
									}),

									layout.Stacked(func(gtx C) D {
										return material.Body1(th, t.rows[i].name).Layout(gtx)
									}))
							})
						})
					}),

					layout.Flexed(1, func(gtx C) D {
						return border.Layout(gtx, func(gtx C) D {
							return layout.Inset{
								Top:    unit.Dp(2),
								Bottom: unit.Dp(2),
								Left:   unit.Dp(2),
							}.Layout(gtx, func(gtx C) D {
								return layout.Stack{Alignment: layout.Center}.Layout(gtx,

									layout.Expanded(func(gtx C) D {
										return g.ColoredArea(gtx, image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y), rowColor)
									}),

									layout.Stacked(func(gtx C) D {
										return material.Body1(th, t.rows[i].cca2).Layout(gtx)
									}))
							})
						})
					}),
					layout.Flexed(1, func(gtx C) D {
						return border.Layout(gtx, func(gtx C) D {
							return layout.Inset{
								Top:    unit.Dp(2),
								Bottom: unit.Dp(2),
								Left:   unit.Dp(2),
							}.Layout(gtx, func(gtx C) D {
								return layout.Stack{Alignment: layout.Center}.Layout(gtx,

									layout.Expanded(func(gtx C) D {
										return g.ColoredArea(gtx, image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y), rowColor)
									}),

									layout.Stacked(func(gtx C) D {
										return material.Body1(th, t.rows[i].cca3).Layout(gtx)
									}))
							})
						})
					}),
					layout.Flexed(1, func(gtx C) D {
						return border.Layout(gtx, func(gtx C) D {
							return layout.Inset{
								Top:    unit.Dp(2),
								Bottom: unit.Dp(2),
								Left:   unit.Dp(2),
							}.Layout(gtx, func(gtx C) D {
								return layout.Stack{Alignment: layout.Center}.Layout(gtx,

									layout.Expanded(func(gtx C) D {
										return g.ColoredArea(gtx, image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y), rowColor)
									}),

									layout.Stacked(func(gtx C) D {
										return material.Body1(th, t.rows[i].ccn3).Layout(gtx)
									}))
							})
						})
					}),
					layout.Flexed(3, func(gtx C) D {
						return border.Layout(gtx, func(gtx C) D {
							return layout.Inset{
								Top:    unit.Dp(2),
								Bottom: unit.Dp(2),
								Left:   unit.Dp(2),
							}.Layout(gtx, func(gtx C) D {
								return layout.Stack{Alignment: layout.N}.Layout(gtx,

									layout.Expanded(func(gtx C) D {
										return g.ColoredArea(gtx, image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y), rowColor)
									}),

									layout.Stacked(func(gtx C) D {
										return material.Body1(th, fmt.Sprintf("%.2f", t.rows[i].area)).Layout(gtx)
									}))
							})
						})
					}),
				)
			}

			return content
		})
	})
}
