package ajson

func Parse(v string) (*JSONDocument, error) {
	return DefaultOptions.DefaultParser.Parse(v)
}
