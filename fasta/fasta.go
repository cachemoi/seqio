package fasta

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/cachemoi/seqio/dna"
)

// Reader returns a structure able to parse DNA stored in FASTA format
type Reader struct{}

// NewReader returns a structure able to parse DNA stored in FASTA format
func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) ReadFile(path string) ([]*dna.Data, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return r.Read(file)
}

func (r *Reader) Read(reader io.Reader) ([]*dna.Data, error) {

	dnas := []*dna.Data{}

	var id dna.ID
	var sequence dna.Sequence
	var err error

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {

		line := scanner.Text()
		if strings.HasPrefix(line, ">") {

			if id != "" {
				// we should be done with the previous record
				dnas = append(dnas, dna.NewData(id, sequence))
			}

			id, err = dna.IDFromString(strings.TrimSpace(line[1:]))
			if err != nil {
				return nil, err
			}

			continue
		}

		// Otherwise append sequence line

		parsed, err := dna.SeqFromString(line)
		if err != nil {
			return nil, err
		}

		sequence += parsed
	}

	// Append last record
	if id != "" {
		dnas = append(dnas, dna.NewData(id, sequence))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return dnas, nil
}
