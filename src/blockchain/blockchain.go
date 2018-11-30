package blockchain

import (
	"reflect"
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
	chain.IsValid()

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


	//checking generation and difficulty
	for i := 0; i < len(chain.Chain); i++ {
		blk := chain.Chain[i]

		//Difficulty
		if diff != blk.Difficulty {
			return false
		}

		//Generation
		if generation != blk.Generation {
			return false
		}
		generation++

		//Previous Hash
		if reflect.DeepEqual(blk.PrevHash, prevHash) {
			return false
		}
		prevHash = blk.Hash

		//ValidHash
		blk.ValidHash()

		//make function for calculating Hash
		currentHash := blk.CalcHash()
		if reflect.DeepEqual(currentHash, blk.Hash) {
			return false
		}
	}




	return true
}
