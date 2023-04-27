package Genesis

import{
	"github.com/vidarsolutions/Ring"
	"github.com/vidarsolutions/Node"

}

var(
	rings = 	Ring.Rings.loadRings()
)
func createFirstNodes() []Node.Node{
	// Create a new Node
	node := Node.Node{}
	node.NodeID =uint64(0)
	node.RingId =uint64(0)
	node.Tor = "Need to Generate and Add tor address here"
	node.I2p = "Need to Generate and Add i2p address here"
	// Add the Node to the Map of all Nodes
	Node.Nodes.AddNode(node)

	// Get the updated Nodes map
	firstNodes := Node.Nodes.GetNodes()
	if(Logging){
		fmt.Println("Node info: "+ strconv.FormatUint(firstNodes[0].NodeID, 10 ))
	}
	nodes := []Node.Node{node}
	newNode := Node.Node{}
	newNode.RingId =uint64(0)
	newNode.NodeID = uint64(1)
	newNode.Tor = "Need to Generate and Add tor address here"
	newNode.I2p = "Need to Generate and Add i2p address here"
	nodes = append(nodes, newNode)
	return nodes
}
func generateRingZero(numRings uint64){
	if numRings == 0{
		say("Generating Ring Zero:", nil, true)
		// Create Vidar Ring 0
		ring := Ring.Ring{}
		ring.RingId =uint64(0)
		ring.LastRing =uint64(0)
		nodes:=createFirstNodes()
		
		ring.Nodes = nodes
		
		// Add the Ring to the Map of all Rings
		rings[0]=ring
		if(Logging){
			for _, node := range rings[0].Nodes {
				fmt.Println("NodeID:", node.NodeID)
				fmt.Println("Tor:", node.Tor)
				fmt.Println("I2p:", node.I2p)
			}
		}
		rings.saveRings()
	}else{
		say("Ring Zero exists joining the Vidar Network", nil, true)
		if Logging{
			for _, node := range rings[0].Nodes {
				say("NodeID:", node.NodeID, Logging)
				say("Tor:", node.Tor, Logging)
				say("I2p:", node.I2p, Logging)
			}
		}
		
	}
}

func say(msg string, data interface{}, logging) {
    if logging {
        var dataStr string
        switch v := data.(type) {
        case string:
            dataStr = v
        case nil:
            dataStr = "<nil>"
        default:
            dataStr = fmt.Sprintf("%v", v)
        }
        fmt.Printf("%s %s\n", msg, dataStr)
    }
}

