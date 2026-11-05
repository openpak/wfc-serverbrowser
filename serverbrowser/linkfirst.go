package serverbrowser

// Link-first (prds repos/wfc-link-first-design.md, decision D12): a server list goes only to an
// address that has a live, linked GPCM login. Every DWC title logs in to GP before it asks for
// a list, so a real console always qualifies; a crafted client that never logged in cannot
// scrape hosts' names and addresses. With WFC_LINK_FIRST off, everyone qualifies.

import (
	"context"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"openpak/wfc-common/database"
	"openpak/wfc-common/gpcmclient"
	"openpak/wfc-common/logging"
	"openpak/wfc-common/wfcapi"
)

var (
	linkedIPsMu    sync.Mutex
	linkedIPs      map[string]bool
	linkedIPsUntil time.Time
	refusedLists   atomic.Int64
	// gpcmSessions is the GPCM session feed; tests swap it.
	gpcmSessions = gpcmclient.Sessions
)

func init() {
	go func() {
		for range time.Tick(time.Minute) {
			if n := refusedLists.Swap(0); n > 0 {
				logging.Notice("SB", "link gate: refused", n, "list requests from addresses with no linked login in the last minute")
			}
		}
	}()
}

func hostOnly(ip string) string {
	if i := strings.LastIndex(ip, ":"); i >= 0 {
		return ip[:i]
	}
	return ip
}

// linkedCaller reports whether a list may go to this caller ("a.b.c.d:port"). The set of
// linked login addresses is refreshed at most every 10 seconds; if GPCM cannot be asked, the
// last known set is used, and with none yet the list is allowed (an outage must not blank
// every console's search).
func linkedCaller(callerPublicIP string) bool {
	if !database.LinkFirst() {
		return true
	}
	ip := hostOnly(callerPublicIP)
	linkedIPsMu.Lock()
	defer linkedIPsMu.Unlock()
	if time.Now().After(linkedIPsUntil) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		rows, err := gpcmSessions(ctx)
		cancel()
		if err == nil {
			linkedIPs = linkedSet(rows)
			linkedIPsUntil = time.Now().Add(10 * time.Second)
		} else {
			logging.Warn("SB", "link gate: gpcm sessions:", err)
			if linkedIPs == nil {
				return true
			}
		}
	}
	if linkedIPs[ip] {
		return true
	}
	refusedLists.Add(1)
	return false
}

func linkedSet(rows []wfcapi.Session) map[string]bool {
	set := map[string]bool{}
	for _, r := range rows {
		if !r.Login.Unlinked && r.Login.GPPublicIP != "" {
			set[hostOnly(r.Login.GPPublicIP)] = true
		}
	}
	return set
}
