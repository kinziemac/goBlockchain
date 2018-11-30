package blockchain

import (
	"blockchain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstHash(t *testing.T) {
	b0 := blockchain.Initial(19)
	b0.Mine(13)

	arr := make([]byte, 32)
	hash := uint64(87745)

	assert.Equal(t, b0.PrevHash, arr)
	assert.Equal(t, b0.Proof, hash)
	assert.Equal(t, b0.Generation, uint64(0))
}

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

	assert.Equal(t, b2.Proof, uint64(128582))
}

func TestAddingToBlockChain(t *testing.T) {
	b0 := blockchain.Initial(19)
	b0.Mine(13)
	assert.Equal(t, b0.ValidHash(), true)

	bChain := new(blockchain.Blockchain)
	assert.Equal(t, len(bChain.Chain), 0)

	bChain.Add(b0)
	assert.Equal(t, len(bChain.Chain), 1)
}
