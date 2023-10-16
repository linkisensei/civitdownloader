package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/dustin/go-humanize"
)

// WriteCounter counts the number of bytes written to it. It implements to the io.Writer interface
// and we can pass this into io.TeeReader() which will report progress on each write cycle.
type WriteCounter struct {
	Total uint64
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

func DownloadFile(fileUrl string, filePath string) error {

	// Create the file, but give it a tmp file extension, this means we won't overwrite a
	// file until it's downloaded, but we'll remove the tmp extension once downloaded.
	out, err := os.Create(filePath + ".tmp")
	if err != nil {
		return err
	}

	// Get the data
	resp, err := http.Get(fileUrl)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()

	// Create our progress reporter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Print("\n")

	// Close the file without defer so it can happen before Rename()
	out.Close()

	if err = os.Rename(filePath+".tmp", filePath); err != nil {
		return err
	}

	return nil
}

func CreateDirectoryIfDoesntExists(dirpath string) error {
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		// Directory does not exist, so create it
		if err := os.MkdirAll(dirpath, 0755); err != nil {
			return err
		}

		return nil
	} else if err != nil {
		return err
	} else {
		return nil
	}

}
