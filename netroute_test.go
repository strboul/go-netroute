package netroute

import (
	"net"
	"testing"
)

func TestRoute(t *testing.T) {
	_, gw, src, err := Route(net.IPv4(127, 0, 0, 1))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Default route is via %v from %v", gw, src)
}