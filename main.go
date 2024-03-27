package main

import (
	"github.com/cachemoi/seqio/fasta"
)

func main() {

	// TODO writer to bin format (parquet?)

	reader := fasta.NewReader()

	_, err := reader.ReadFile("fasta/testdata/example.fasta")
	if err != nil {
		panic(err)
	}

}
