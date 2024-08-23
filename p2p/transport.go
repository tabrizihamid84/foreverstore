package p2p

// Peer is an interface that represents the remote node
type Peer interface {
}

// Transport is anything that can handles communication between the nodes in the network
// This can be of the form (TCP, Websockets, ...)
type Transport interface {
	ListenAndAccept() error
}
