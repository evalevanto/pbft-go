package pbft

import (
	"log"
	"net"
	"os"

	"github.com/spf13/viper"
)

var config *viper.Viper
var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "Network Logs", log.Ltime)
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")
}

type pbftNode struct {
	net.Listener
	nodeID    uint64 // backup identifier, i
	nodeCount uint64 // Nodes in the network, |R|

	sequenceNumber uint64 // Last stable request executed, n
	activeView     bool
	currentView    uint64 // View the node is in, v

	h           uint64            // sequence number of the last stable checkpoint. Lower watermark, h. Used in pre-prep and prep.
	H           uint64            // Water marks limit what messages will be accepted. High water mark, H. H = h + k.
	k           uint64            // Static number that sets the high mark advance steps.
	checkpoints map[uint64]string // A map of sequenceNo of checkpoints and digest.

	currExecReq uint64 // Current request being executed.
	requestChan chan *Request
}

// Constructor
func newPbftNode(replID uint64) *pbftNode {
	node := &pbftNode{}

	node.nodeID = replID
	node.nodeCount = config.GetUint64("network.node_count")
	node.k = config.GetUint64("network.k")
	node.checkpoints = make(map[uint64]string)

	initialCheckpoint := "THE GENESIS XXXX"
	node.checkpoints[0] = initialCheckpoint

	logger.Printf("New node %d created", node.nodeID)

	return node
}

// ------ HELPER FUNCTIONS ------

// whoPrimary function determines the primary node in a view, v.
// The primary of view v is the replica p such that p = mod |R|.
func (node *pbftNode) whoPrimary(view uint64) uint64 {
	return view % node.nodeCount
}

// isSeqInSpace checks that the sequence number, n in the pre-prepare or prepare
// message is between the water marks; before being accepted.
func (node *pbftNode) isSeqInSpace(sequenceNumber uint64) bool {
	if sequenceNumber < node.h || sequenceNumber > node.H {
		return false
	}
	return true
}

// ------ REPLICA COUNT FUNCTIONS ------
func (node *pbftNode) countForPrePrep() {

}
