package main

import (
	"github.com/tadvi/winc"
)

var turn bool

var mainWindow = winc.NewForm(nil)

var tab = [...]string{"", "", "", "", "", "", "", "", ""}

var GagnantX = winc.NewPushButton(mainWindow)

var btn1 = winc.NewPushButton(mainWindow)
var btn2 = winc.NewPushButton(mainWindow)
var btn3 = winc.NewPushButton(mainWindow)
var btn4 = winc.NewPushButton(mainWindow)
var btn5 = winc.NewPushButton(mainWindow)
var btn6 = winc.NewPushButton(mainWindow)
var btn7 = winc.NewPushButton(mainWindow)
var btn8 = winc.NewPushButton(mainWindow)
var btn9 = winc.NewPushButton(mainWindow)

func main() {
	GagnantX.SetText("Recommencer")
	GagnantX.OnClick().Bind(Replay)
	mainWindow.SetSize(400, 400)
	mainWindow.SetText("Tp morpion")
	GagnantX.SetSize(300, 60)
	GagnantX.SetPos(45, 230)
	btn := winc.NewPushButton(mainWindow)
	btn.SetText("Fermer")
	btn.SetPos(145, 300)
	btn.SetSize(100, 40)
	btn.OnClick().Bind(wndOnClose)

	btn1.SetText("")
	btn1.SetSize(60, 60)
	btn1.SetPos(95, 20)
	btn1.OnClick().Bind(func(e *winc.Event) {
		if tab[0] == "" {
			if turn {
				tab[0] = "O"
				turn = false
				btn1.SetText("O")
			} else {
				tab[0] = "X"
				turn = true
				btn1.SetText("X")
			}
		}
		IsWin()
		IsFinish()

	})

	btn2.SetText("")
	btn2.SetSize(60, 60)
	btn2.SetPos(165, 20)
	btn2.OnClick().Bind(func(e *winc.Event) {
		if tab[1] == "" {
			if turn {
				tab[1] = "O"
				turn = false
				btn2.SetText("O")
			} else {
				tab[1] = "X"
				turn = true
				btn2.SetText("X")
			}
		}
		IsWin()
		IsFinish()
	})
	btn3.SetText("")
	btn3.SetSize(60, 60)
	btn3.SetPos(235, 20)
	btn3.OnClick().Bind(func(e *winc.Event) {
		if tab[2] == "" {
			if turn {
				tab[2] = "O"
				turn = false
				btn3.SetText("O")
			} else {
				tab[2] = "X"
				turn = true
				btn3.SetText("X")
			}
		}
		IsWin()
		IsFinish()
	})
	btn4.SetText("")
	btn4.SetSize(60, 60)
	btn4.SetPos(95, 90)
	btn4.OnClick().Bind(func(e *winc.Event) {
		if tab[3] == "" {
			if turn {
				tab[3] = "O"
				turn = false
				btn4.SetText("O")
			} else {
				tab[3] = "X"
				turn = true
				btn4.SetText("X")
			}
		}
		IsWin()
		IsFinish()
	})
	btn5.SetText("")
	btn5.SetSize(60, 60)
	btn5.SetPos(165, 90)
	btn5.OnClick().Bind(func(e *winc.Event) {
		if tab[4] == "" {
			if turn {
				tab[4] = "O"
				turn = false
				btn5.SetText("O")
			} else {
				tab[4] = "X"
				turn = true
				btn5.SetText("X")
			}
		}
		IsWin()
		IsFinish()
	})
	btn6.SetText("")
	btn6.SetSize(60, 60)
	btn6.SetPos(235, 90)
	btn6.OnClick().Bind(func(e *winc.Event) {
		if tab[5] == "" {
			if turn {
				tab[5] = "O"
				turn = false
				btn6.SetText("O")
			} else {
				tab[5] = "X"
				turn = true
				btn6.SetText("X")
			}
		}
		IsWin()
		IsFinish()
	})
	btn7.SetText("")
	btn7.SetSize(60, 60)
	btn7.SetPos(95, 160)
	btn7.OnClick().Bind(func(e *winc.Event) {
		if tab[6] == "" {
			if turn {
				tab[6] = "O"
				turn = false
				btn7.SetText("O")
			} else {
				tab[6] = "X"
				turn = true
				btn7.SetText("X")
			}
		}
		IsWin()
		IsFinish()
	})
	btn8.SetText("")
	btn8.SetSize(60, 60)
	btn8.SetPos(165, 160)
	btn8.OnClick().Bind(func(e *winc.Event) {
		if tab[7] == "" {
			if turn {
				tab[7] = "O"
				turn = false
				btn8.SetText("O")
			} else {
				tab[7] = "X"
				turn = true
				btn8.SetText("X")
			}
		}
		IsWin()
		IsFinish()
	})
	btn9.SetText("")
	btn9.SetSize(60, 60)
	btn9.SetPos(235, 160)
	btn9.OnClick().Bind(func(e *winc.Event) {
		if tab[8] == "" {
			if turn {
				tab[8] = "O"
				turn = false
				btn9.SetText("O")
			} else {
				tab[8] = "X"
				turn = true
				btn9.SetText("X")
			}
		}
		IsWin()
		IsFinish()
	})
	mainWindow.Center()
	mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)
	winc.RunMainLoop()
}
func wndOnClose(arg *winc.Event) {
	winc.Exit()
}

func IsWin() {
	if tab[0] == "X" && tab[1] == "X" && tab[2] == "X" {

		GagnantX.SetText("X a gagné, rejouer ?")

	}
	if tab[3] == "X" && tab[4] == "X" && tab[5] == "X" {

		GagnantX.SetText("X a gagné, rejouer ?")

	}
	if tab[6] == "X" && tab[7] == "X" && tab[8] == "X" {

		GagnantX.SetText("X a gagné, rejouer ?")

	}
	if tab[0] == "X" && tab[3] == "X" && tab[6] == "X" {

		GagnantX.SetText("X a gagné, rejouer ?")

	}
	if tab[1] == "X" && tab[4] == "X" && tab[7] == "X" {

		GagnantX.SetText("X a gagné, rejouer ?")

	}
	if tab[2] == "X" && tab[5] == "X" && tab[8] == "X" {

		GagnantX.SetText("X a gagné, rejouer ?")

	}
	if tab[0] == "X" && tab[4] == "X" && tab[8] == "X" {

		GagnantX.SetText("X a gagné, rejouer ?")

	}
	if tab[2] == "X" && tab[4] == "X" && tab[6] == "X" {

		GagnantX.SetText("X a gagné, rejouer ?")

	}
	if tab[0] == "O" && tab[1] == "O" && tab[2] == "O" {

		GagnantX.SetText("O a gagné, rejouer ?")

	}
	if tab[3] == "O" && tab[4] == "O" && tab[5] == "O" {

		GagnantX.SetText("O a gagné, rejouer ?")

	}
	if tab[6] == "O" && tab[7] == "O" && tab[8] == "O" {

		GagnantX.SetText("O a gagné, rejouer ?")

	}
	if tab[0] == "O" && tab[3] == "O" && tab[6] == "O" {

		GagnantX.SetText("O a gagné, rejouer ?")

	}
	if tab[1] == "O" && tab[4] == "O" && tab[7] == "O" {

		GagnantX.SetText("O a gagné, rejouer ?")

	}
	if tab[2] == "O" && tab[5] == "O" && tab[8] == "O" {

		GagnantX.SetText("O a gagné, rejouer ?")

	}
	if tab[0] == "O" && tab[4] == "O" && tab[8] == "O" {

		GagnantX.SetText("O a gagné, rejouer ?")

	}
	if tab[2] == "O" && tab[4] == "O" && tab[6] == "O" {

		GagnantX.SetText("O a gagné, rejouer ?")

	}
}

func IsFinish() {
	if tab[0] != "" && tab[1] != "" && tab[2] != "" && tab[3] != "" && tab[4] != "" && tab[5] != "" && tab[6] != "" && tab[7] != "" && tab[8] != "" {
		GagnantX.SetText("Egalité, rejouer ?")

	}
}

func Replay(arg *winc.Event) {
	tab[0] = ""
	tab[1] = ""
	tab[2] = ""
	tab[3] = ""
	tab[4] = ""
	tab[5] = ""
	tab[6] = ""
	tab[7] = ""
	tab[8] = ""
	btn1.SetText("")
	btn2.SetText("")
	btn3.SetText("")
	btn4.SetText("")
	btn5.SetText("")
	btn6.SetText("")
	btn7.SetText("")
	btn8.SetText("")
	btn9.SetText("")
	GagnantX.SetText("Recommencer")
	GagnantX.OnClick().Bind(Replay)
}
