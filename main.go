package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/cheggaaa/pb/v3"
	"github.com/ofman/filesharego"
)

// Mock function to simulate downloading data with progress updates
func DownloadWithProgress(w io.Writer, cid string, progress chan<- int64) error {
	totalSize := int64(100) // Replace with actual total size
	bar := pb.Full.Start64(totalSize)
	bar.SetWriter(ioutil.Discard) // Discard the default output

	for i := int64(0); i < totalSize; i++ {
		time.Sleep(50 * time.Millisecond) // Simulate time taken to download data
		w.Write([]byte("data"))           // Write data to provided writer
		bar.Increment()
		progress <- i + 1 // Send progress update
	}

	bar.Finish()
	close(progress)
	return nil
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("FileShareGo UI")

	label := widget.NewLabel("")

	input := widget.NewEntry()
	input.SetPlaceHolder("Paste your CID address here...")

	progress := widget.NewProgressBar()
	infinite := widget.NewProgressBarInfinite()

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
		// Create channels to receive the results
		msgReturnedChan := make(chan string)
		errChan := make(chan error)
		progressChan := make(chan int64)

		// Run the function in a goroutine
		go func() {
			msgReturned, err, progressReturned := filesharego.DownloadFromCid(input.Text, false)
			msgReturnedChan <- msgReturned
			errChan <- err
			progressChan <- progressReturned
		}()

		progress.Show()

		go func() {
			for p := range progressChan {
				fmt.Printf("\rDownload progress: %d%%", p)
				progress.SetValue(float64(p))
			}
			fmt.Println("\nDownload finished")
		}()

		// Use the results
		msgReturned := <-msgReturnedChan
		err := <-errChan
		filesharego.ErrorCheck(err, false)
		label.SetText("Content successfully downloaded:\n" + msgReturned)

	})

	input.OnChanged = func(input string) {
		buttonCopy.Show()
		buttonSave.Show()
		buttonDownload.Show()
	}

	// fileInfo := fyne.URIReadCloser()
	dialogFile := dialog.NewFileOpen(func(readThis fyne.URIReadCloser, err error) {

		go func() {
			// Do the background work.
			cidStr, err := filesharego.UploadFiles(readThis.URI().Path(), false)

			if err != nil {
				// Handle error.
				filesharego.ErrorCheck(err, false)
				return
			}

			// Update the UI with the result.
			fmt.Printf("Content successfully uploaded and seeding: %s\n", readThis.URI().Path()+"\nCopy and share this CID address below:\n"+cidStr)
			label.SetText("Content successfully uploaded and seeding:\n" + readThis.URI().Path() + "\nCopy and share this CID address below:")
			input.SetText(cidStr)

			buttonCopy.Show()
			buttonSave.Show()
			buttonDownload.Show()
			infinite.Show()
		}()

	}, myWindow)

	dialogFolder := dialog.NewFolderOpen(func(readThis fyne.ListableURI, err error) {
		go func() {
			// Do the background work.
			cidStr, err := filesharego.UploadFiles(readThis.Path(), false)

			if err != nil {
				// Handle error.
				filesharego.ErrorCheck(err, false)
				return
			}

			// Update the UI with the result.
			fmt.Printf("Content successfully uploaded and seeding: %s\n", readThis.Path()+"\nCopy and share this CID address below:\n"+cidStr)
			label.SetText("Content successfully uploaded and seeding:\n" + readThis.Path() + "\nCopy and share this CID address below:")
			input.SetText(cidStr)

			buttonCopy.Show()
			buttonSave.Show()
			buttonDownload.Show()
			infinite.Show()
		}()
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
	progress.Hide()
	infinite.Hide()

	content2 := container.NewVBox(label, input, buttonPaste, buttonCopy, buttonSave, buttonDownload, buttonsUpload, progress, infinite)
	myWindow.SetContent(content2)
	myWindow.Resize(fyne.NewSize(700, 500))
	myWindow.ShowAndRun()
}
