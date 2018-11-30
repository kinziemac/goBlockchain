package blockchain

import (
	"blockchain"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO: some useful tests of Blocks
func TestValidHash(t *testing.T) {
	b0 := blockchain.Initial(19)
  b0.SetProof(87745)

  b1 := b0.Next("hash example 1234")
  b1.SetProof(1407891)

	assert.Equal(t, b0.ValidHash(), true)
	assert.Equal(t, b1.ValidHash(), true)
}

func TestNotValidHash(t *testing.T) {
	b0 := blockchain.Initial(19)
  b0.SetProof(87745)
  b1 := b0.Next("hash example 1234")
  b1.SetProof(140789)

	assert.Equal(t, b0.ValidHash(), true)
	assert.Equal(t, b1.ValidHash(), false)
}

func TestValidMine(t *testing.T) {
	b0 := blockchain.Initial(19)
  b0.SetProof(87745)
  b1 := b0.Next("hash example 1234")
  b1.SetProof(1407891)

	b2 := b1.Next("hello")
	b2.Mine(13)

	t.Log("Proof Value: ", b2.Proof)
	assert.Equal(t, b2.Proof, uint64(128582))
}
