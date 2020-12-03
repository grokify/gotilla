package svgutil

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/rustyoz/svg"
)

func ReadFile(filename, name string, scale float64) (*svg.Svg, error) {
	// https://godoc.org/github.com/rustyoz/svg#Svg
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	name = strings.TrimSpace(name)
	if len(name) == 0 {
		name = filename
	}

	return svg.ParseSvgFromReader(file, filename, scale)
}

func AspectRatio(img *svg.Svg) (float64, error) {
	if img == nil {
		return 0, errors.New("E_NO_IMAGE")
	}
	vals, err := img.ViewBoxValues()
	if err != nil {
		return 0, err
	}
	if len(vals) == 2 {
		if vals[0] == 0 {
			return 0, errors.New("E_DENOM_HEIGHT==0")
		}
		return vals[1] / vals[0], nil
	} else if len(vals) == 4 {
		w := vals[2] - vals[0]
		h := vals[3] - vals[1]
		if h == 0 {
			return 0, errors.New("E_DENOM_HEIGHT==0")
		}
		return w / h, nil
	}
	return 0, fmt.Errorf("E_BAD_VIEWBOX_LEN [%v]", len(vals))
}