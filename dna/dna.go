package dna

type ID string

func IDFromString(data string) (ID, error) {
	// TODO add validation
	return ID(data), nil
}

type Sequence string

func SeqFromString(data string) (Sequence, error) {
	// TODO add validation
	return Sequence(data), nil
}

// Data is a general purpose representation of a DNA sequence
type Data struct {
	ID       ID
	Sequence Sequence
}

// NewData returns a general purpose representation of a DNA sequence
func NewData(ID ID, sequence Sequence) *Data {
	return &Data{
		ID:       ID,
		Sequence: sequence,
	}
}
