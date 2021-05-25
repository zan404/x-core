package grpc_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/metadata"

	"github.com/xtls/xray-core/common/net"
	. "github.com/xtls/xray-core/common/protocol/grpc"
)

func TestParseXRealIP(t *testing.T) {
	md := metadata.New(map[string]string{"X-Real-IP": "129.78.138.66"})
	addr := ParseXRealIP(md)
	if r := cmp.Diff(addr, net.ParseAddress("129.78.138.66")); r != "" {
		t.Error(r)
	}
}

func TestParseXForwardedFor(t *testing.T) {
	md := metadata.Pairs("X-Forwarded-For", "129.78.138.66, 129.78.64.103")
	addrs := ParseXForwardedFor(md)
	if r := cmp.Diff(addrs, []net.Address{net.ParseAddress("129.78.138.66"), net.ParseAddress("129.78.64.103")}); r != "" {
		t.Error(r)
	}
}
