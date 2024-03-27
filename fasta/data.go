package fasta

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidSequence = errors.New("invalid sequence")

	// Allowed characters for nucleotide sequences
	nucleotideChars = map[rune]struct{}{
		'A': {},
		'C': {},
		'G': {},
		'T': {},
		'U': {},
		'R': {},
		'Y': {},
		'S': {},
		'W': {},
		'K': {},
		'M': {},
		'B': {},
		'D': {},
		'H': {},
		'V': {},
		'N': {},
	}

	// Allowed characters for protein sequences
	aminoAcidChars = map[rune]struct{}{
		'A': {},
		'R': {},
		'N': {},
		'D': {},
		'C': {},
		'Q': {},
		'E': {},
		'G': {},
		'H': {},
		'I': {},
		'L': {},
		'K': {},
		'M': {},
		'F': {},
		'P': {},
		'S': {},
		'T': {},
		'W': {},
		'Y': {},
		'V': {},
	}
)

type ID string

func IDFromString(data string) (ID, error) {
	// TODO add validation
	return ID(data), nil
}

type Sequence string

// SeqFromString seeks to parse a valid biological sequence from the input. FASTA is barely a standard so we could get
// pretty much anything here, but we try to filter out anything really unexpected
//
// we could expect the following valid letters:
//
// For nucleotide sequences (DNA or RNA), the valid letters typically include:
//
//   - Adenine (A)
//   - Cytosine (C)
//   - Guanine (G)
//   - Thymine (T) (for DNA) or Uracil (U) (for RNA)
//   - Additionally, some formats may include ambiguity codes such as R (A or G), Y (C or T), S (G or C), W (A or T), K (G or T), M (A or C), B (C or G or T), D (A or G or T), H (A or C or T), V (A or C or G), N (any base)
//
// For protein sequences, the valid letters typically include the standard amino acid abbreviations:
//
//   - Alanine (A)
//   - Arginine (R)
//   - Asparagine (N)
//   - Aspartic acid (D)
//   - Cysteine (C)
//   - Glutamine (Q)
//   - Glutamic acid (E)
//   - Glycine (G)
//   - Histidine (H)
//   - Isoleucine (I)
//   - Leucine (L)
//   - Lysine (K)
//   - Methionine (M)
//   - Phenylalanine (F)
//   - Proline (P)
//   - Serine (S)
//   - Threonine (T)
//   - Tryptophan (W)
//   - Tyrosine (Y)
//   - Valine (V)
func SeqFromString(data string) (Sequence, error) {

	for _, char := range data {
		_, isValidNucleotide := nucleotideChars[char]
		_, isValidAminoAcid := aminoAcidChars[char]

		if !isValidAminoAcid && !isValidNucleotide {
			return Sequence(""), fmt.Errorf("%w: got %c", ErrInvalidSequence, char)
		}
	}

	return Sequence(data), nil
}

// Data is a general purpose representation of a FASTA parsed sequence
type Data struct {
	ID       ID
	Sequence Sequence
}

// NewData returns a general purpose representation of a FASTA sequence
func NewData(ID ID, sequence Sequence) *Data {
	return &Data{
		ID:       ID,
		Sequence: sequence,
	}
}
