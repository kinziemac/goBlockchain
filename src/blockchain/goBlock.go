package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	PrevHash   []byte
	Generation uint64
	Difficulty uint8
	Data       string
	Proof      uint64
	Hash       []byte
}

// Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	hash := make([]byte, 32)
	b := Block{
		hash,
		0,
		difficulty,
		"",
		0,
		hash}

	return b
}

// Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	// TODO
	b := Block{
		prev_block.Hash,
		prev_block.Generation + 1,
		prev_block.Difficulty,
		data,
		0,
		prev_block.Hash}

	return b
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {

	src := blk.PrevHash
	encodedStr := hex.EncodeToString(src)
	encodedStr += ":" + fmt.Sprint(blk.Generation)
	encodedStr += ":" + fmt.Sprint(blk.Difficulty)
	encodedStr += ":" + blk.Data
	encodedStr += ":" + fmt.Sprint(blk.Proof)

	h := sha256.New()
	h.Write([]byte(encodedStr))

	return h.Sum(nil)
}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	nBytes := blk.Difficulty / 8
	nBits := blk.Difficulty % 8
	zeroValues := len(blk.Hash) - int(nBytes)

	//check last bytes to make sure they're 0
	for i := zeroValues; i < len(blk.Hash); i++ {
		if blk.Hash[i] != []byte("\x00")[0] {
			return false
		}
	}

	if blk.Hash[zeroValues-1]%(1<<nBits) != 0 {
		return false
	}

	return true
}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
