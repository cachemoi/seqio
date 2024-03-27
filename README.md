# Seqio

This library is intended to be a rosetta stone to convert between the vast array of formats used to store biological
sequence data.

A [go library](https://pkg.go.dev/github.com/biogo/biogo/io/seqio) already exists to do a similar job, and there's also
similar libraries in other languages (e.g [biopython](https://biopython.org/wiki/SeqIO))

This is intended to be a tool with a much narrower scope than those generic bioinformatics toolchains, and perhaps gain
some clarity/ergonomics through this clarity of purpose.

## Example usage

```go
io := fasta.NewIO()

// the parsed sequences data here will look something like this:
//
// &fasta.Data{ID:"Sequence1", Sequence:"ATCGATCGATCGATCGATCGATCGATCGATCG"},
//
parsedSequences, err := io.ReadFile("mysequence.fasta")
if err != nil {
  return err
}

// writing the sequences we just parsed to another file
err := io.WriteFile("somesequences.fasta", parsedSequences)
if err != nil {
  return err
}
```

You could also use a generic io.Reader/io.Writer interfaces rather than a file

```go
var writer io.Writer
var reader io.Writer

io := fasta.NewIO()

// the parsed sequences data here will look something like this:
//
// &fasta.Data{ID:"Sequence1", Sequence:"ATCGATCGATCGATCGATCGATCGATCGATCG"},
//
parsedSequences, err := io.Read(reader)
if err != nil {
  return err
}

// writing the sequences we just parsed to another file
err := io.Write(writer, parsedSequences)
if err != nil {
  return err
}
```
