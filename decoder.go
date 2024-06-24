package ajson

import (
	"errors"
)

type Decoder interface {
	CanDecode(Decoder, JSONNode, *Reader, Options) bool
	Decode(Decoder, JSONNode, *Reader, Options) (JSONNode, error)
}

type JSONDecoder struct {
	decoders []Decoder
}

func NewJSONDecoder(decoders ...Decoder) *JSONDecoder {
	return &JSONDecoder{decoders: decoders}
}

func (decoder *JSONDecoder) AddDecoder(d Decoder) {
	decoder.decoders = append(decoder.decoders, d)
}

func (decoder *JSONDecoder) CanDecode(Decoder, JSONNode, *Reader, Options) bool {
	return true
}

func (decoder *JSONDecoder) Decode(_ Decoder, n JSONNode, r *Reader, options Options) (JSONNode, error) {
	for _, d := range decoder.decoders {
		if d.CanDecode(decoder, n, r, options) {
			return d.Decode(decoder, n, r, options)
		}
	}
	return nil, errors.New("unsupported encoding type")
}
