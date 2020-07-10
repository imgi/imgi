package imgi

import "github.com/h2non/bimg"

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
}

func Resize(img Image, opts Options) (Image, error) {
	// meta, err := bimg.Metadata(img.Buf)
	// if err != nil {
	// 	return img, err
	// }

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
