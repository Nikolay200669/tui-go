package main

import (
	"fmt"

	tui "github.com/marcusolsson/tui-go"
)

func main() {
	s1 := newScrollArea(10)
	s2 := newScrollArea(10)
	s3 := newScrollArea(10)
	root := tui.NewVBox()

	scrollBox1 := tui.NewVBox(s1)
	scrollBox1.SetBorder(true)

	matilda := tui.NewVBox(s2)
	matilda.SetBorder(true)
	matilda.SetSizePolicy(tui.Expanding, tui.Expanding)

	scrollBox2 := tui.NewHBox(tui.NewSpacer(), matilda, tui.NewSpacer())
	scrollBox2.SetBorder(true)

	scrollBox3 := tui.NewVBox(s3)
	scrollBox3.SetBorder(true)

	root.Append(scrollBox1)
	root.Append(scrollBox2)
	root.Append(scrollBox3)

	ui := tui.New(root)
	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Up", func() { s2.Scroll(0, -1) })
	ui.SetKeybinding("Down", func() { s2.Scroll(0, 1) })
	ui.SetKeybinding("Left", func() { s2.Scroll(-1, 0) })
	ui.SetKeybinding("Right", func() { s2.Scroll(1, 0) })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}

func newScrollArea(n int) *tui.ScrollArea {
	var labels []tui.Widget
	for i := 0; i < n; i++ {
		labels = append(labels, tui.NewLabel(fmt.Sprintf("foo %d", i)))
	}
	b := tui.NewVBox(labels...)
	b.SetBorder(true)

	marcus := tui.NewVBox(tui.NewLabel("bar"))
	marcus.SetBorder(true)

	wrap := tui.NewVBox(b, marcus)
	wrap.SetBorder(true)

	return tui.NewScrollArea(wrap)
}
