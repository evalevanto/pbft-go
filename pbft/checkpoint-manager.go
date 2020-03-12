package pbft

// Checkpoint describes a record of current state
type Checkpoint struct{}

// MakeCheckpoint is an upcall function that creates a (volatile) checkpoint record.
func makeCheckpoint(currentSeqNumber int, modifiedBlocks []int) {
	// Create digest of the current state using AdHash algorithm.
}

func deleteCheckpoint() {}

func getCheckpoint() {}
