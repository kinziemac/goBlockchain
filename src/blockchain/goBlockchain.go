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

func (chain Blockchain) IsValid() bool {
	diff := uint8(0)
	prevHash := make([]byte, 32)
	generation := uint64(0)

	//checking all block values to make sure they are vaild
	for i := 0; i < len(chain.Chain); i++ {
		blk := chain.Chain[i]
		diff = chain.Chain[0].Difficulty

		//First Value
		if i == 0 {
			for j := 0; j < 32; j++ {
				if chain.Chain[0].PrevHash[j] != []byte("\x00")[0] {
					fmt.Println("First Prev did not match all zeros")
					return false
				}
			}
		}

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
			fmt.Println("in CurrentHash")
			return false
		}
	}

	return true
}
