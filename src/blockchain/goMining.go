package blockchain

import (
	"work_queue"
)

type miningWorker struct {
	// TODO. Should implement work_queue.Worker
	block Block
	start uint64
	end   uint64
}

type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	queue := work_queue.Create(uint(workers), uint(chunks))
	chunk_range := end / chunks
	mine_result := new(MiningResult)

	for i := start; i <= end; i = i + chunk_range {
		mine_worker := new(miningWorker)
		mine_worker.start = uint64(i)
		mine_worker.end = uint64(i + chunk_range)
		mine_worker.block = blk

		//puts into the Jobs queue
		queue.Enqueue(mine_worker)
	}

	//what happens if I can't find a value? haha I don't know
	//so this became an issue
	for true {
		result := <-queue.Results
		new_mine_result := result.(MiningResult)

		if new_mine_result.Found {
			queue.Shutdown()
			return new_mine_result
		}

	}

	return *mine_result
}

//Interface for workqueue to use
func (mine *miningWorker) Run() interface{} {
	mine_result := new(MiningResult)

	//Must default because if it returns without finding a result, we don't want it to be true
	mine_result.Found = false

	// checking if proof value is valid
	for i := mine.start; i <= mine.end; i++ {
		mine.block.SetProof(i)

		if mine.block.ValidHash() {
			//If I set the proof above, do I even have to do this? -> Yes, mine_result isn't a block
			mine_result.Found = true
			mine_result.Proof = i
			return *mine_result
		}
	}

	return *mine_result
}

// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << blk.Difficulty) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}
