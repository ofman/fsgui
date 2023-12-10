package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/ofman/filesharego"
)

func bgDownloadWorker(inputStr string) (string, error) {
	c1 := make(chan string)
	c2 := make(chan error)
	go filesharego.DownloadFromCid(inputStr, false)

	select {
	// will have to fix timer in the future to check for seeders
	// case <-time.After(10 * time.Second):
	// 	// if c2 == nil {
	// 	fmt.Println("Download taking more than 10 seconds. No seeders?")
	// 	return "Download taking more than 10 seconds. No seeders?", nil
	// 	// } else {
	// 	// 	return "", nil
	// 	// }
	case outputPath := <-c1:
		fmt.Println("Content successfully downloaded to:\n", outputPath)
		return "Content successfully downloaded to:\n" + outputPath, nil
	case err := <-c2:
		fmt.Println("error: ", err)
		return err.Error(), err
	}
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("FileShareGo UI")

	label := widget.NewLabel("")

	input := widget.NewEntry()
	input.SetPlaceHolder("Paste your CID address here...")

	// label2 := widget.NewLabel("or")

	// centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), label2, layout.NewSpacer())
	// content := container.NewVBox(label, input, buttonPaste, centered, buttonsUpload)
	// myWindow.SetContent(content)
	buttonCopy := widget.NewButtonWithIcon("Copy", theme.ContentCopyIcon(), func() {
		myWindow.Clipboard().SetContent(input.Text)
		fmt.Printf("Content was copied: %s\n", input.Text)
		label.SetText("Content successfully copied!")
	})
	buttonSave := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() {
		// in case of /ipfs/exampleCid we strip string and work only on exampleCid
		cidStr := filesharego.GetCidStrFromString(input.Text)
		d1 := []byte(cidStr)
		err := os.WriteFile(cidStr+".txt", d1, 0777)
		filesharego.ErrorCheck(err, false)
		fmt.Printf("Content was saved as .txt file: %s\n", cidStr)
		label.SetText("Content successfully saved:\n" + cidStr + ".txt")
	})
	buttonDownload := widget.NewButtonWithIcon("Download content", theme.DownloadIcon(), func() {
		msgReturned, err := bgDownloadWorker(input.Text)
		filesharego.ErrorCheck(err, false)
		label.SetText(msgReturned)

	})

	input.OnChanged = func(input string) {
		buttonCopy.Show()
		buttonSave.Show()
		buttonDownload.Show()
	}

	// fileInfo := fyne.URIReadCloser()
	dialogFile := dialog.NewFileOpen(func(readThis fyne.URIReadCloser, err error) {
		cidStr, err := filesharego.UploadFiles(readThis.URI().Path(), false)
		filesharego.ErrorCheck(err, false)

		myWindow.Clipboard().SetContent(cidStr)
		fmt.Printf("Content successfully seeding: %s\n", readThis.URI().Path())
		label.SetText("Content successfully seeding:\n" + readThis.URI().Path() + "\nCopy and share this CID address below:")
		input.SetText(cidStr)

		buttonCopy.Show()
		buttonSave.Show()
		buttonDownload.Show()
	}, myWindow)

	dialogFolder := dialog.NewFolderOpen(func(readThis fyne.ListableURI, err error) {
		cidStr, err := filesharego.UploadFiles(readThis.Path(), false)
		filesharego.ErrorCheck(err, false)

		myWindow.Clipboard().SetContent(cidStr)
		fmt.Printf("Content successfully seeding: %s\n", readThis.Path())
		label.SetText("Content successfully seeding:\n" + readThis.Path() + "\nCopy and share this CID address below:")
		input.SetText(cidStr)

		buttonCopy.Show()
		buttonSave.Show()
		buttonDownload.Show()
	}, myWindow)

	buttonPaste := widget.NewButtonWithIcon("Paste", theme.ContentPasteIcon(), func() {
		if myWindow.Clipboard().Content() != "" {
			input.SetText(myWindow.Clipboard().Content())
			fmt.Printf("Content was pasted: %s\n", input.Text)
			label.SetText("Content successfully pasted!")

			buttonCopy.Show()
			buttonSave.Show()
			buttonDownload.Show()
		} else {
			label.SetText("Paste clipboard is empty!")
		}

	})

	buttonUploadFile := widget.NewButtonWithIcon("Upload file", theme.UploadIcon(), func() {
		dialogFile.Show()
	})

	buttonUploadFolder := widget.NewButtonWithIcon("Upload folder", theme.FolderIcon(), func() {
		dialogFolder.Show()
	})

	buttonsUpload := container.New(layout.NewGridLayout(2), buttonUploadFile, buttonUploadFolder)

	buttonCopy.Hide()
	buttonSave.Hide()
	buttonDownload.Hide()

	content2 := container.NewVBox(label, input, buttonPaste, buttonCopy, buttonSave, buttonDownload, buttonsUpload)
	myWindow.SetContent(content2)
	myWindow.Resize(fyne.NewSize(700, 500))
	myWindow.ShowAndRun()
}
