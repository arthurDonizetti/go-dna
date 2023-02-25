package domain_test

import (
	"testing"

	"github.com/arthurDonizetti/go-dna/internals/domain"
	"github.com/stretchr/testify/assert"
)

type mockSequence struct {
	sequence []string
	expected bool
}

func newSequence() []mockSequence {
	return []mockSequence{
		{
			sequence: []string{
				"CTGAGA",
				"CTATGC",
				"TATTGT",
				"AGAGGG",
				"CCCCTA",
				"TCACTG",
			},
			expected: true,
		},
		{
			sequence: []string{
				"GGGGAT",
			},
			expected: false,
		},
		{
			sequence: []string{
				"GGGGAT",
				"CCCCTG",
			},
			expected: true,
		},
		{
			sequence: []string{
				"ATGAGA",
				"CAATGC",
				"TAATGT",
				"AGAATG",
				"CACTTA",
				"TCACTG",
			},
			expected: true,
		},
		{
			sequence: []string{
				"AACGCTACAAGCTCGAAGGA",
				"GGGAGGCAACTCTGCCGCTC",
				"GATCCTTTCAACGGGGAAAA",
				"GCGCTCGTTAAATACTGACT",
				"ATGATGAATATCCGCTGTCA",
				"GGGGCATCTCGTGAAGCTTT",
				"ACCCCTGGCCGCTTTTTTCT",
				"CAAAATATGATACTCAGGCA",
				"AACACGCCGCAGTGGGGTTG",
				"GGTACCCGGCCCTACGAGAA",
				"CGTACCTCGGGCCGTATCCT",
				"GGTGGTTTATAACACCACTG",
				"CCACGGCGTCGTAGTAACGT",
				"GCTATTCTCCGATGGAAGCG",
				"CGTCCACAAAGGGTCCAACA",
				"CAGTACCCGTTTCCCTAGAG",
				"CTTTAATACGAGGGGCTATC",
				"GTGTTTGTGACATCTCTTCA",
				"GTAACAGAGTCCTCCCGGTC",
				"CAAAATATGATACTCAGGCA",
			},
			expected: true,
		},
		{
			sequence: []string {
				"ACGT",
				"TGCA",
				"GCAT",
				"CATG",
			},
			expected: false,
		},
		{
			sequence: []string {
				"ACGT",
				"AGCT",
				"ACAT",
				"AATT",
			},
			expected: true,
		},
		{
			sequence: []string {
				"ACGA",
				"AGAA",
				"AAAT",
				"AATG",
			},
			expected: true,
		},
		{
			sequence: []string {
				"ACGA",
				"TAAA",
				"GAAT",
				"AATA",
			},
			expected: true,
		},
	}
}

func Test_DnaIsSimian(t *testing.T) {
	for _, seq := range newSequence() {
		dna := domain.NewDna(seq.sequence)
		result := dna.IsSimian()
		assert.Equal(t, seq.expected, result)
	}
}
