package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"io"
	"net/http"
	"os"
	"strings"
)

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
	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

func main4() {
	fmt.Println("Download Started")

	fileUrl := "https://download.jetbrains.com/idea/ideaIU-2024.1.4.exe?_gl=1*dwjzzh*_gcl_au*MzAxOTMzNjQwLjE3MTM0OTM4ODc.*_ga*MTkyNjcyMzAxNC4xNjgzMTg0MzQ5*_ga_9J976DJZ68*MTcxOTQ2OTg4OS4xMy4xLjE3MTk0Njk5OTYuMTkuMC4w&_ga=2.162534655.2112090742.1719469889-1926723014.1683184349"
	err := DownloadFile("9.png", fileUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Download Finished")
}
func DownloadFile(filepath string, url string) error {
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	resp, err := http.Get(url)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()
	counter := &WriteCounter{}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}
	fmt.Print("\n")
	out.Close()
	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}
