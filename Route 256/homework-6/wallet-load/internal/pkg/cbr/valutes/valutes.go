package valutes

import (
	"encoding/xml"
	"io"

	"golang.org/x/text/encoding/charmap"
)

type ValCurs struct {
	ValCurs []Valute `xml:"Valute"`
}

type Valute struct {
	CharCode string `xml:"CharCode"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

func (v *ValCurs) Decode(r io.Reader) error {
	d := xml.NewDecoder(r)
	d.CharsetReader = func(_ string, input io.Reader) (io.Reader, error) {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}
	return d.Decode(v)
}

func Unmarshal(r io.Reader) (*ValCurs, error) {
	v := &ValCurs{}
	return v, v.Decode(r)
}
