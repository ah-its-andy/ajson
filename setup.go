package ajson

type Options struct {
	Decoders []Decoder

	IntentChar string
}

var DefaultOptions *Options = &Options{
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

func Setup(fn func(*Options)) {
	if fn != nil {
		fn(DefaultOptions)
	}
}
