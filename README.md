
# gotermimg

[![.github/workflows/main.yml](https://github.com/panekj/gotermimg/actions/workflows/main.yml/badge.svg)](https://github.com/panekj/gotermimg/actions/workflows/main.yml)

Update of [github.com/moshen/gotermimg](https://github.com/moshen/gotermimg)

## Installation

Install using `go install`:

```shell
go install github.com/panekj/gotermimg/cmd/gotermimg@latest
```

Installs the `gotermimg` command line application.

## Usage

```none
Usage: gotermimg [-u] [-x=n] [-y=n] [IMAGEFILE]
IMAGEFILE - png or jpg.
Image data can be piped to stdin instead of providing IMAGEFILE.

If neither -x or -y are provided, and the image is larger than your current
terminal, it will be automatically scaled to fit.

-u=false: Enable UTF8 output
-x=0: Scale to n*2 columns wide in ANSI mode, n columns wide in UTF8 mode.
      When -x=0 (the default), aspect ratio is maintained.
      For example if -y is provided without -x, width is scaled to
      maintain aspect ratio
-y=0: Scale to n rows high in ANSI mode, n/2 rows high in UTF8 mode.
      When -y=0 (the default), aspect ratio is maintained.
      For example if -x is provided without -y, height is scaled to
      maintain aspect ratio
```

[![gotermimg on a png with transparency](https://media.giphy.com/media/vpYeVwn2cRxstBp5hS/giphy.gif)](https://media.giphy.com/media/vpYeVwn2cRxstBp5hS/giphy.gif)

[![gotermimg on an animated gif with transparency](https://media.giphy.com/media/b9sXmD1dWBUvbgr87r/giphy.gif)](https://media.giphy.com/media/b9sXmD1dWBUvbgr87r/giphy.gif)

While the render speed on some slower terminals might not look very good, urxvt
looks amazing (click through for HQ).

[![gotermimg on urxvt](https://media.giphy.com/media/Jsg9KArYyntBPgoH4o/giphy.gif)](https://media.giphy.com/media/Jsg9KArYyntBPgoH4o/giphy.gif)

## Author

[Colin Kennedy](https://github.com/moshen)

## Libraries used

- [github.com/nfnt/resize](https://github.com/nfnt/resize) - [MIT Style](https://github.com/nfnt/resize/blob/master/LICENSE)
- [golang.org/x/term](https://github.com/golang/term) - [BSD-3-Clause](https://github.com/golang/term/blob/master/LICENSE)
