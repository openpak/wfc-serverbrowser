package serverbrowser

import (
	"context"
	"testing"
	"time"

	"openpak/wfc-common/wfcapi"
)

func TestLinkedCaller(t *testing.T) {
	t.Setenv("WFC_LINK_FIRST", "1")
	before := gpcmSessions
	defer func() { gpcmSessions = before; linkedIPs = nil; linkedIPsUntil = time.Time{} }()
	gpcmSessions = func(ctx context.Context) ([]wfcapi.Session, error) {
		return []wfcapi.Session{
			{Login: wfcapi.Login{GPPublicIP: "198.51.100.1:5000"}},
			{Login: wfcapi.Login{GPPublicIP: "198.51.100.2:5000", Unlinked: true}},
		}, nil
	}
	linkedIPs, linkedIPsUntil = nil, time.Time{}
	if !linkedCaller("198.51.100.1:31000") {
		t.Fatal("linked address refused")
	}
	if linkedCaller("198.51.100.2:31000") || linkedCaller("203.0.113.9:1") {
		t.Fatal("unlinked or unknown address allowed")
	}
	t.Setenv("WFC_LINK_FIRST", "")
	if !linkedCaller("203.0.113.9:1") {
		t.Fatal("gate off must allow everyone")
	}
}
