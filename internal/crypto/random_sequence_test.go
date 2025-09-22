package crypto

import (
	"testing"

	"github.com/Sorrowful-free/short-url-service/internal/consts"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomSequence(t *testing.T) {
	sequence, err := GenerateRandomSequence(consts.TestUIDLength)
	if err != nil {
		t.Fatalf("failed to generate random sequence: %v", err)
	}
	assert.NotEmpty(t, sequence, "sequence must be not empty")
}

func TestGenerateRandomSequenceString(t *testing.T) {
	sequence, err := GenerateRandomSequenceString(consts.TestUIDLength)
	if err != nil {
		t.Fatalf("failed to generate random sequence string: %v", err)
	}
	assert.NotEmpty(t, sequence, "sequence must be not empty")
}
