package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/lxn/walk"
	d "github.com/lxn/walk/declarative"
)


type MyLabel struct{
	d.Label
}

func makeLabel(t string) MyLabel  {
	l:= MyLabel{}
	//StretchFactor:1,AlwaysConsumeSpace:false,MaxSize:d.Size{Height:10,Width:0}
	l.Text = t
	l.Font = defFont
	//l.StretchFactor = 1
	//l.AlwaysConsumeSpace = false
	//l.MaxSize = d.Size{Height:20,Width:0}
	// l.Background = d.SolidColorBrush{
	// 	Color: walk.RGB(128,0,128),
	// }
	return l
}

func main() {
	
	
	var fromModel, toModel unit
	var fdb, tdb *walk.DataBinder
	var fte, tte *walk.TextEdit
	winsz := d.Size{Width: 800, Height: 0}
	sz := d.Size{Width: 0, Height:420}
	sz2 := d.Size{Width: 0, Height:200}
	 
	//noVgrowLayout := walk.NewVBoxLayout()
	//noVgrowLayout.

	if _, err := (d.MainWindow{
		AssignTo: &mw,
		Title:    "Bandwidth Converter",
		Icon: "assets/bconv.ico",
		Layout:   d.VBox{},
		MinSize: winsz,
		Children: []d.Widget{
			d.Composite{
				Layout: d.Grid{},
				Children: []d.Widget{
					d.HSpacer{StretchFactor: 100,Row:1,Column: 1},
					d.HSpacer{StretchFactor: 100,Row:1,Column: 2},
					d.PushButton{
						Font:defFont,
						Text: "Convert",
						Row: 1,
						Column: 3,
						AlwaysConsumeSpace: false,
						OnClicked: func() {
							if err := fdb.Submit(); err != nil {
								log.Print(err)
								return
							}
							if err := tdb.Submit(); err != nil {
								log.Print(err)
								return
							}
							
							c := fromModel.convetTo(toModel)
							
							i, err := strconv.ParseFloat(fte.Text(), 64)
							if err != nil {
								
								errorDialog(mw, err, "Wrong Format")
								fte.SetText("")
							}
							o := c.apply(i)
							tte.SetText(fmt.Sprint(o))
						},
					},

				},
				
			},			
			d.Composite{
				Border: true,
				MaxSize: sz,
				Layout: d.VBox{},
				Children: []d.Widget{
					d.Composite{
						Layout: d.VBox{},
						Children: []d.Widget{
							makeLabel("From"),
							d.Composite{
								MaxSize: sz2,
								Layout: d.HBox{},
								DataBinder: d.DataBinder{
									AssignTo:       &fdb,
									DataSource:     &fromModel,
									ErrorPresenter: d.ToolTipErrorPresenter{},
								},
								Children: []d.Widget{
									d.Composite{
										Layout:   d.VBox{},
										Children: []d.Widget{makeLabel("Value"), d.TextEdit{AssignTo: &fte, Font:defFont}},
									},
									d.Composite{
										Layout:   d.VBox{},
										Children: []d.Widget{makeLabel("Order"), d.ComboBox{ Font:defFont, Value: d.Bind("U"), BindingMember: "V", DisplayMember:"N", Model: M()}},
									},
									d.Composite{
										Layout:   d.VBox{},
										Children: []d.Widget{makeLabel("B/b"), d.ComboBox{ Font:defFont,Value: d.Bind("B"), BindingMember: "V", DisplayMember:"N", Model: B()}},
									},
									d.Composite{
										Layout:   d.VBox{},
										Children: []d.Widget{makeLabel("time"), d.ComboBox{ Font:defFont,Value: d.Bind("S"), BindingMember: "V", DisplayMember:"N", Model: T()}},
									},
								},
							},
						},
					},
					d.VSeparator{},
					d.Composite{
						Layout: d.VBox{},
						Children: []d.Widget{
							makeLabel("To"),
							d.Composite{
								MaxSize: sz2,
								Layout: d.HBox{},
								DataBinder: d.DataBinder{
									AssignTo:       &tdb,
									DataSource:     &toModel,
									ErrorPresenter: d.ToolTipErrorPresenter{},
								},
								Children: []d.Widget{
									d.Composite{
										Layout:   d.VBox{},
										Children: []d.Widget{makeLabel("Value"), d.TextEdit{AssignTo: &tte,ReadOnly: true, Font:defFont}},
									},
									d.Composite{
										Layout:   d.VBox{},
										Children: []d.Widget{makeLabel("Order"), d.ComboBox{ Font:defFont,Value: d.Bind("U"), BindingMember: "V", DisplayMember:"N", Model: M()}},
									},
									d.Composite{
										Layout:   d.VBox{},
										Children: []d.Widget{makeLabel("B/b"), d.ComboBox{ Font:defFont,Value: d.Bind("B"), BindingMember: "V", DisplayMember:"N", Model: B()}},
									},
									d.Composite{
										Layout:   d.VBox{},
										Children: []d.Widget{makeLabel("time"), d.ComboBox{ Font:defFont,Value: d.Bind("S"), BindingMember: "V", DisplayMember:"N", Model: T()}},
									},
								},
							},
						},
					},
				},
			},
			d.VSpacer{},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
	//log.Println("Outside")
	
	

}

var defFont = d.Font{PointSize: 24}

var errDlg *walk.Dialog
var mw *walk.MainWindow

type CType struct{
	V int
	N string
}

func M()[]*CType{
	return []*CType{
		{0," "},
		{3,"K"},
		{6,"M"},
		{9,"G"},
		{12,"T"},
	}
}

func B()[]*CType{
	return []*CType{
		{0,"b"},
		{1,"B"},
	}
}

func T()[]*CType{
	return []*CType{
		{0,"s"},
		{1,"m"},
		{2,"h"},
	}
}

type unit struct {
	U int
	B int
	S int
}

func (f unit) convetTo(t unit) unit {
	return unit{U: f.U - t.U, B: f.B - t.B, S: t.S - f.S}
}

func (f unit) apply(v float64) float64 {
	return v * math.Pow10(f.U) * math.Pow(60, float64(f.S)) * math.Pow(8, float64(f.B))
}

func errorDialog( c *walk.MainWindow, e error, s string)int{
	return walk.MsgBox(c,
		s,
		e.Error(),
		walk.MsgBoxOK|walk.MsgBoxIconError)

}
