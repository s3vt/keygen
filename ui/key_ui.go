package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateWindow() {
	myapp := app.New()
	win := myapp.NewWindow("Key Generator Utility")
	win.Resize(fyne.Size{Width: 500, Height: 400})
	addToolbar(win)

	win.ShowAndRun()
}

func addToolbar(myWindow fyne.Window) error {
	// toolbar := widget.NewToolbar(
	// 	widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
	// 		log.Println("New document")
	// 	}),
	// 	widget.NewToolbarSeparator(),
	// 	widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
	// 	widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
	// 	widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
	// 	widget.NewToolbarSpacer(),
	// 	widget.NewToolbarAction(theme.HelpIcon(), func() {
	// 		log.Println("Display help")
	// 	}),
	// )

	// content := container.NewBorder(toolbar, nil, nil, nil, widget.NewLabel("Content"))
	// myWindow.SetContent(content)
	content := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")),
		container.NewTabItem("RSA", createKeyContent()),
		container.NewTabItem("ECDSA", widget.NewLabel("World!")),
		container.NewTabItem("Certificate", widget.NewLabel("World!")),
	)

	content.SetTabLocation(container.TabLocationLeading)
	myWindow.SetContent(content)

	return nil
}

func createKeyContent() fyne.CanvasObject {
	keyContent := container.NewGridWithRows(
		4,
		widget.NewLabel("RSA Keys Generation"),
		widget.NewButton("Generate keys", func() { fmt.Println("ME pressed") }),
		canvas.NewText("Private key is shown here", color.White),
		canvas.NewText("Public key is shown here", color.White),
	)

	return keyContent

}
