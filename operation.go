package imgi

import (
	"fmt"

	"github.com/h2non/bimg"
)

type Image struct {
	Buf  []byte
	Mime string
}

type Operation struct {
	Name    string
	Options Options
	Action  Action
}

type Options struct {
	Width  int
	Height int
	Mode   Mode
	X      int
	Y      int
}

type Mode int

const (
	Fit  Mode = iota // aspect ratio preserved, width and height attain
	Fill             // aspect ratio ignored, width and height attain
	Flex             // aspect ratio preserved, width or height attain
)

type Action func(Image, Options) (Image, error)

func (a Action) Act(img Image, opts Options) (Image, error) {
	return a(img, opts)
}

var actions = map[string]Action{
	"resize": Resize,
	"crop":   Crop,
}

func Resize(img Image, opts Options) (Image, error) {
	o := bimg.Options{
		Width:   opts.Width,
		Height:  opts.Height,
		Gravity: bimg.GravitySmart,
	}

	if opts.Width > 0 && opts.Height > 0 {
		switch opts.Mode {
		case Fit:
			o.Crop = true
			o.Enlarge = true
		case Flex:
			o.Height = 0
		}
	}

	buf, err := bimg.Resize(img.Buf, o)
	if err != nil {
		return img, err
	}
	return Image{buf, img.Mime}, nil
}

func Crop(img Image, opts Options) (Image, error) {
	meta, err := bimg.Metadata(img.Buf)
	if err != nil {
		return img, err
	}
	w := meta.Size.Width
	h := meta.Size.Height
	if opts.X+opts.Width > w {
		return img, fmt.Errorf("crop out of image width")
	}
	if opts.Y+opts.Height > h {
		return img, fmt.Errorf("crop out of image height")
	}
	if opts.Width == 0 {
		opts.Width = w - opts.X
	}
	if opts.Height == 0 {
		opts.Height = h - opts.Y
	}

	o := bimg.Options{
		AreaWidth:  opts.Width,
		AreaHeight: opts.Height,
		Top:        opts.Y,
		Left:       opts.X,
	}
	buf, err := bimg.Resize(img.Buf, o)

	if err != nil {
		return img, err
	}
	return Image{buf, img.Mime}, nil
}
