package fasta

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReadFile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string

		path string

		wantSeqs []*Data

		wantErr error
	}{
		{
			name: "ok",

			path: filepath.Join("testdata", "example.fasta"),

			wantSeqs: []*Data{
				{
					ID:       "Sequence1",
					Sequence: "ATCGATCGATCGATCGATCGATCGATCGATCG"},
				{
					ID:       "Sequence2",
					Sequence: "CGATCGATCGATCGATCGATCGATCGATCGAT",
				},
				{
					ID:       "Sequence3",
					Sequence: "TACGGATACAGGTACCGAGCTCGATCGATCG",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotSeqs, err := NewIO().ReadFile(tt.path)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got %v, want %v", err, tt.wantErr)
			}

			if !cmp.Equal(gotSeqs, tt.wantSeqs) {
				t.Errorf("gotSeqs and tt.wantSeqs didn't match %s", cmp.Diff(gotSeqs, tt.wantSeqs))
			}

		})
	}
}
func TestRead(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string

		input string

		wantSeqs []*Data
		wantErr  error
	}{
		{
			name: "ok",
			input: `>Sequence1
ATCGATCGATCGATCGATCGATCGATCGATCG
>Sequence2
CGATCGATCGATCGATCGATCGATCGATCGAT
>Sequence3
TACGGATACAGGTACCGAGCTCGATCGATCG
`,
			wantErr: nil,
			wantSeqs: []*Data{
				{
					ID:       "Sequence1",
					Sequence: "ATCGATCGATCGATCGATCGATCGATCGATCG"},
				{
					ID:       "Sequence2",
					Sequence: "CGATCGATCGATCGATCGATCGATCGATCGAT",
				},
				{
					ID:       "Sequence3",
					Sequence: "TACGGATACAGGTACCGAGCTCGATCGATCG",
				},
			},
		},
		{
			name:     "empty",
			input:    ``,
			wantErr:  nil,
			wantSeqs: []*Data{},
		},
		{
			name: "invalid sequence",
			input: `>Sequence1
ATCGATCGATCGATCGATCGATCGATCGATCG
>Sequence2Bad
CGATCGATCGATCGATCGA🍌CGATCGATCGAT
>Sequence3
TACGGATACAGGTACCGAGCTCGATCGATCG
`,
			wantErr: ErrInvalidSequence,
			wantSeqs: []*Data{
				{
					ID:       "Sequence1",
					Sequence: "ATCGATCGATCGATCGATCGATCGATCGATCG",
				},
			},
		},
	}

	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			testData := strings.NewReader(tt.input)

			reader := NewIO()

			got, err := reader.Read(testData)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got %v, want %v", err, tt.wantErr)
			}

			if !cmp.Equal(got, tt.wantSeqs) {
				t.Errorf("got and tt.wantSeqs didn't match %s", cmp.Diff(got, tt.wantSeqs))
			}

		})
	}
}

func TestWriteFile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string

		sequences []*Data

		wantWritten []byte
		wantErr     error
	}{
		{
			name: "ok",

			sequences: []*Data{
				{
					ID:       "Sequence1",
					Sequence: "ATCGATCGATCGATCGATCGATCGATCGATCG"},
				{
					ID:       "Sequence2",
					Sequence: "CGATCGATCGATCGATCGATCGATCGATCGAT",
				},
				{
					ID:       "Sequence3",
					Sequence: "TACGGATACAGGTACCGAGCTCGATCGATCG",
				},
			},

			wantWritten: []byte(`>Sequence1
ATCGATCGATCGATCGATCGATCGATCGATCG
>Sequence2
CGATCGATCGATCGATCGATCGATCGATCGAT
>Sequence3
TACGGATACAGGTACCGAGCTCGATCGATCG
`),
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Create a temporary directory for testing
			tempDir, err := os.MkdirTemp("", "test")
			if err != nil {
				t.Fatalf("Error creating temporary directory: %v", err)
			}
			defer os.RemoveAll(tempDir) // Clean up the temporary directory

			testPath := filepath.Join(tempDir, "output.fasta")

			err = NewIO().WriteFile(testPath, tt.sequences)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got %v, want %v", err, tt.wantErr)
			}

			gotData, err := os.ReadFile(testPath)
			if err != nil {
				t.Fatalf("Error reading test file: %v", err)
			}

			if !cmp.Equal(gotData, tt.wantWritten) {
				t.Errorf("gotData and tt.wantWritten didn't match %s", cmp.Diff(gotData, tt.wantWritten))
			}

		})
	}
}

func TestWrite(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string

		sequences []*Data

		wantWritten string
		wantErr     error
	}{
		{
			name: "ok",

			sequences: []*Data{
				{
					ID:       "Sequence1",
					Sequence: "ATCGATCGATCGATCGATCGATCGATCGATCG"},
				{
					ID:       "Sequence2",
					Sequence: "CGATCGATCGATCGATCGATCGATCGATCGAT",
				},
				{
					ID:       "Sequence3",
					Sequence: "TACGGATACAGGTACCGAGCTCGATCGATCG",
				},
			},

			wantWritten: `>Sequence1
ATCGATCGATCGATCGATCGATCGATCGATCG
>Sequence2
CGATCGATCGATCGATCGATCGATCGATCGAT
>Sequence3
TACGGATACAGGTACCGAGCTCGATCGATCG
`,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			writer := &strings.Builder{}

			err := NewIO().Write(writer, tt.sequences)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got %v, want %v", err, tt.wantErr)
			}

			if !cmp.Equal(writer.String(), tt.wantWritten) {
				t.Errorf("writer.String() and tt.wantWritten didn't match %s", cmp.Diff(writer.String(), tt.wantWritten))
			}
		})
	}
}

func TestMirror(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string

		sequences []*Data

		wantSeqs []*Data
		wantErr  error
	}{
		{
			name: "ok",

			sequences: []*Data{
				{
					ID:       "Sequence1",
					Sequence: "ATCGATCGATCGATCGATCGATCGATCGATCG"},
				{
					ID:       "Sequence2",
					Sequence: "CGATCGATCGATCGATCGATCGATCGATCGAT",
				},
				{
					ID:       "Sequence3",
					Sequence: "TACGGATACAGGTACCGAGCTCGATCGATCG",
				},
			},

			wantSeqs: []*Data{
				{
					ID:       "Sequence1",
					Sequence: "ATCGATCGATCGATCGATCGATCGATCGATCG"},
				{
					ID:       "Sequence2",
					Sequence: "CGATCGATCGATCGATCGATCGATCGATCGAT",
				},
				{
					ID:       "Sequence3",
					Sequence: "TACGGATACAGGTACCGAGCTCGATCGATCG",
				},
			},

			wantErr: nil,
		},
	}

	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Create a temporary directory for testing
			tempDir, err := os.MkdirTemp("", "test")
			if err != nil {
				t.Fatalf("Error creating temporary directory: %v", err)
			}
			defer os.RemoveAll(tempDir) // Clean up the temporary directory

			testPath := filepath.Join(tempDir, "output.fasta")

			io := NewIO()

			err = io.WriteFile(testPath, tt.sequences)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got %v, want %v", err, tt.wantErr)
			}

			gotSeqs, err := io.ReadFile(testPath)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("got %v, want %v", err, tt.wantErr)
			}

			if !cmp.Equal(gotSeqs, tt.wantSeqs) {
				t.Errorf("gotSeqs and tt.wantSeqs didn't match %s", cmp.Diff(gotSeqs, tt.wantSeqs))
			}

		})
	}
}
