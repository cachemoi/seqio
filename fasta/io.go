// Package fasta is used to read and write FASTA formatted biological sequences
package fasta

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// IO returns a structure able to perform Input/Output operations for sequences using the FASTA format
type IO struct{}

// NewIO returns a structure able to perform Input/Output operations for sequences using the FASTA format
func NewIO() *IO {
	return &IO{}
}

// Read will extract sequence data from the given file path, expecting a FASTA formatted sequence file.
//
// An example of FASTA formatted sequences looks like this:
//
//	>Sequence1
//	ATCGATCGATCGATCGATCGATCGATCGATCG
//	>Sequence2
//	CGATCGATCGATCGATCGATCGATCGATCGAT
//	>Sequence3
//	TACGGATACAGGTACCGAGCTCGATCGATCG
func (r *IO) ReadFile(path string) ([]*Data, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return r.Read(file)
}

// Read will extract sequence data from the given reader, expecting a FASTA formatted sequence file.
//
// An example of FASTA formatted sequences looks like this:
//
//	>Sequence1
//	ATCGATCGATCGATCGATCGATCGATCGATCG
//	>Sequence2
//	CGATCGATCGATCGATCGATCGATCGATCGAT
//	>Sequence3
//	TACGGATACAGGTACCGAGCTCGATCGATCG
func (r *IO) Read(reader io.Reader) ([]*Data, error) {

	seqs := []*Data{}

	var id ID
	var sequence Sequence
	var err error

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {

		line := scanner.Text()

		// check if the line is an ID

		if strings.HasPrefix(line, ">") {

			// append the previous record and reset the variables

			if id != "" && sequence != "" {
				seqs = append(seqs, NewData(id, sequence))
			}

			id, err = IDFromString(line[1:])
			if err != nil {
				return seqs, err
			}

			// reset Sequence
			sequence = Sequence("")

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

// WriteFile writes the sequences to a file in FASTA format.
func (r *IO) WriteFile(path string, seqs []*Data) error {

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return r.Write(file, seqs)
}

// Write writes the sequences in FASTA format to the provided io.Writer.
func (r *IO) Write(writer io.Writer, seqs []*Data) error {

	for _, seq := range seqs {
		// Write header
		if _, err := fmt.Fprintf(writer, ">%s\n", seq.ID); err != nil {
			return err
		}

		// Write sequence
		if _, err := fmt.Fprintf(writer, "%s\n", seq.Sequence); err != nil {
			return err
		}
	}

	return nil
}
