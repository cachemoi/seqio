package fasta

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// Reader returns a structure able to parse DNA stored in FASTA format
type Reader struct{}

// NewReader returns a structure able to parse DNA stored in FASTA format
func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) ReadFile(path string) ([]*Data, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return r.Read(file)
}

func (r *Reader) Read(reader io.Reader) ([]*Data, error) {

	seqs := []*Data{}

	var id ID
	var sequence Sequence
	var err error

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {

		line := scanner.Text()
		if strings.HasPrefix(line, ">") {

			if id != "" {
				// we should be done with the previous record
				seqs = append(seqs, NewData(id, sequence))
			}

			id, err = IDFromString(strings.TrimSpace(line[1:]))
			if err != nil {
				return seqs, err
			}

			continue
		}

		// Otherwise append sequence line

		parsed, err := SeqFromString(line)
		if err != nil {
			return seqs, err
		}

		sequence += parsed
	}

	// Append last record
	if id != "" {
		seqs = append(seqs, NewData(id, sequence))
	}

	if err := scanner.Err(); err != nil {
		return seqs, err
	}

	return seqs, nil
}
