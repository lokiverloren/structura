// structura -  simple application to allow composition of a document as labelled paragraphs that can be easily rearranged
//              and placed into a tree structure to help organise ideas
// project initiated by Loki Verloren <loki.verloren@gmail.com>

package main

import (
        "gopkg.in/qml.v1"
	"math/rand"
	"time"
)

func main() {
	qml.Init(nil)
	engine := qml.NewEngine()
	component, err := engine.LoadFile("share/structura/main.qml")
	if err != nil {
		panic(err)
	}

	ctrl := Control{Message: "Hello from Go!"}

	context := engine.Context()
	context.SetVar("ctrl", &ctrl)

	window := component.CreateWindow(nil)

	ctrl.Root = window.Root()

	rand.Seed(time.Now().Unix())

	window.Show()
	window.Wait()
}

type Control struct {
	Root    qml.Object
	Message string
}

func (ctrl *Control) Hello() {
    go func() {
            ctrl.Message = "Hello from Go Again!"
            qml.Changed(ctrl, &ctrl.Message)
    }()
}

