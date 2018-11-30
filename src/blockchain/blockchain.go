package blockchain

import (
	"reflect"
	"fmt"
)

type Blockchain struct {
	Chain []Block
}

func (chain *Blockchain) Add(blk Block) {
	// You can remove the panic() here if you wish.
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}

	//adds block to blockchain
	chain.Chain = append(chain.Chain, blk)
}
//
// type Block struct {
// 	PrevHash   []byte
// 	Generation uint64
// 	Difficulty uint8
// 	Data       string
// 	Proof      uint64
// 	Hash       []byte
// }

func (chain Blockchain) IsValid() bool {
	diff := uint8(0)
	prevHash := make([]byte, 32)
	generation := uint64(0)

	if len(chain.Chain) > 0 {
		//this way, if the blockchain is empty - it won't seg fault
		diff = chain.Chain[0].Difficulty
	}

	//not sure why first value doesn't work..
	for i := 0; i < 32; i++ {
	  if chain.Chain[0].PrevHash[i] != []byte("\x00")[0] {
			fmt.Println("First Prev did not match all zeros")
			return false
		}
	}


	//checking generation and difficulty
	for i := 0; i < len(chain.Chain); i++ {
		blk := chain.Chain[i]

		//Difficulty
		if diff != blk.Difficulty {
			fmt.Println("in Difficulty")
			return false
		}

		//Generation
		if generation != blk.Generation {
			fmt.Println("in Generation")
			return false
		}
		generation++

		//Previous Hash
		//something about to zeros doesn't work
		if i > 0 && !reflect.DeepEqual(blk.PrevHash, prevHash) {
			fmt.Println("in PrevHash")
			return false
		}
		prevHash = blk.Hash

		//ValidHash
		if !blk.ValidHash() {
			fmt.Println("in ValidHash")
			return false
		}

		//make function for calculating Hash
		currentHash := blk.CalcHash()
		if !reflect.DeepEqual(currentHash, blk.Hash) {
			fmt.Println(currentHash)
			fmt.Println(blk.Hash)
			fmt.Println("in CurrentHash")
			return false
		}
	}

	return true
}
