package fasta

import (
	"errors"
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
					Sequence: "ATCGATCGATCGATCGATCGATCGATCGATCGCGATCGATCGATCGATCGATCGATCGATCGAT",
				},
				{
					ID:       "Sequence3",
					Sequence: "ATCGATCGATCGATCGATCGATCGATCGATCGCGATCGATCGATCGATCGATCGATCGATCGATTACGGATACAGGTACCGAGCTCGATCGATCG",
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotSeqs, err := NewReader().ReadFile(tt.path)
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
					Sequence: "ATCGATCGATCGATCGATCGATCGATCGATCGCGATCGATCGATCGATCGATCGATCGATCGAT",
				},
				{
					ID:       "Sequence3",
					Sequence: "ATCGATCGATCGATCGATCGATCGATCGATCGCGATCGATCGATCGATCGATCGATCGATCGATTACGGATACAGGTACCGAGCTCGATCGATCG",
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

			reader := NewReader()

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
