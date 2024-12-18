package adapter

import (
	"net"

	st "gvisor.dev/gvisor/pkg/tcpip/stack"
)

// TCPConn implements the net.Conn interface.
type TCPConn interface {
	net.Conn

	// ID returns the transport endpoint id of TCPConn.
	ID() *st.TransportEndpointID
}

// UDPConn implements net.Conn and net.PacketConn.
type UDPConn interface {
	net.Conn
	net.PacketConn

	// ID returns the transport endpoint id of UDPConn.
	ID() *st.TransportEndpointID
}
