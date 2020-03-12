package pbft

import "net"

// All BFT phases are implemented here: pre-prepare, prepare, commit.
// Communication: TCP.

// Replica : defines attrs for a replicated nodes in the network.
type Replica struct {
	listener *net.TCPListener
	ID       int64
	addr     string
	// Every node has a "list" of all other node in the network. map[ID]addr.
	replicas       map[int64]string
	currSequence   int64
	view           int64
	isInActiveView bool
	checkpoints    []*Checkpoint
}

// Server :
type Server struct {
	Node *Replica
	Port int
}

func electPrimary(currentView int) {
	// The ID of the new primary node is determined by : v mod N (current view number, v and total number of servers, N).
}
