package controllers

import (
	"gioui-experiment/apps/geography/table"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions

	DisplayedColumns struct {
		table.Row
		list       layout.List
		checkboxes []checkBox

		loaded      bool
		initialized bool
	}

	checkBox struct {
		name string
		box  widget.Bool
	}
)

// Layout - Lays out the column checkboxes
func (dc *DisplayedColumns) Layout(gtx C, th *material.Theme) D {
	if !dc.loaded {
		dc.list.Axis = layout.Vertical
		dc.checkboxes = make([]checkBox, len(table.ColNames))
		for i := range dc.checkboxes {
			dc.checkboxes[i] = checkBox{
				name: table.ColNames[i],
				box:  widget.Bool{Value: true},
			}
		}
		dc.Row.GenerateColumns()
		dc.loaded = true
	}
	return dc.list.Layout(gtx, len(table.ColNames), func(gtx C, i int) D {
		var cb material.CheckBoxStyle
		cb = material.CheckBox(th, &dc.checkboxes[i].box, dc.checkboxes[i].name)

		if cb.CheckBox.Changed() {
			for j := range dc.Row.Columns {
				if dc.Row.Columns[j].HeadCell == dc.checkboxes[i].name {
					dc.Row.Columns[j].IsEnabled = cb.CheckBox.Value
				}
			}
		}
		op.InvalidateOp{}.Add(gtx.Ops)
		return cb.Layout(gtx)
	})
}
