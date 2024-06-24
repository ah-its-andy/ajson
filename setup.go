package ajson

type Options struct {
	Decoders      []Decoder
	DefaultParser *JSONParser

	IntentChar string
}

var DefaultOptions *Options

func init() {
	DefaultOptions = &Options{
		Decoders: []Decoder{
			&ObjectDecoder{
				FieldDecoder: &StringDecoder{},
			},
			&ArrayDecoder{},
			&StringDecoder{},
			&NumberDecoder{},
			&BooleanDecoder{},
			&NullDecoder{},
		},
		IntentChar: "\t",
	}
	DefaultOptions.DefaultParser = NewJSONParser(*DefaultOptions)
}

func Setup(fn func(*Options)) {
	if fn != nil {
		fn(DefaultOptions)
	}
}
