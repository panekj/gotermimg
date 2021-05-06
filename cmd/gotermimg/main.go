package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/term"

	"github.com/panekj/gotermimg"
)

var (
	isUTF8 bool
	width  uint
	height uint
)

func init() {
	flag.BoolVar(&isUTF8, "u", false, "Enable UTF8 output")
	flag.UintVar(&width, "x", 0, `Scale to n*2 columns wide in ANSI mode, n columns wide in UTF8 mode.
        When -x=0 (the default), aspect ratio is maintained.
        For example if -y is provided without -x, width is scaled to
        maintain aspect ratio`)
	flag.UintVar(&height, "y", 0, `Scale to n rows high in ANSI mode, n/2 rows high in UTF8 mode.
        When -y=0 (the default), aspect ratio is maintained.
        For example if -x is provided without -y, height is scaled to
        maintain aspect ratio`)

	flag.Usage = func() {
		_, err := fmt.Fprint(os.Stderr, `Usage: gotermimg [-u] [-x=n] [-y=n] [IMAGEFILE]
  IMAGEFILE - png or jpg.
  Image data can be piped to stdin instead of providing IMAGEFILE.

  If neither -x or -y are provided, and the image is larger than your current
  terminal, it will be automatically scaled to fit.

`)
		if err != nil {
			return
		}
		flag.PrintDefaults()
	}

	flag.Parse()
}

func main() {
	var buf *bytes.Reader
	switch {
	case !term.IsTerminal(int(os.Stdin.Fd())):
		bufData, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		buf = bytes.NewReader(bufData)
	case len(flag.Args()) < 1:
		flag.Usage()
		os.Exit(1)
	default:
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		bufData, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		err = file.Close()
		if err != nil {
			return
		}
		buf = bytes.NewReader(bufData)
	}

	conf, imgFormat, err := image.DecodeConfig(buf)
	if err != nil {
		log.Fatal(err)
	}
	_, err = buf.Seek(0, 0)
	if err != nil {
		return
	}

	var conv gotermimg.Converter
	if isUTF8 {
		conv = gotermimg.UTF8
	} else {
		conv = gotermimg.ANSI
	}

	var trans gotermimg.Transformer
	if width != 0 || height != 0 {
		trans = gotermimg.Resize(width, height)
	} else if term.IsTerminal(int(os.Stdout.Fd())) {
		x, y, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			log.Fatal(err)
		}

		y--

		// Convert the actual terminal dimensions into effective dimensions
		switch {
		case isUTF8:
			y *= 2
		case x%2 == 0:
			x /= 2
		default:
			x = (x - 1) / 2
		}

		if conf.Width > x || conf.Height > y {
			aspectTerm := float32(x) / float32(y)
			aspectImg := float32(conf.Width) / float32(conf.Height)

			if aspectImg > aspectTerm {
				trans = gotermimg.Resize(uint(x), 0)
			} else {
				trans = gotermimg.Resize(0, uint(y))
			}
		}
	}

	if imgFormat != "gif" {
		img, _, err := image.Decode(buf)
		if err != nil {
			log.Fatal(err)
		}
		gotermimg.PrintImage(img, conv, trans)
	}
}
