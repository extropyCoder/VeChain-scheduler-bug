package vuln

import "testing"

// PoC: demonstrates vulnerable behavior by asserting the inactive proposer
// is present in the actives set (this is the bug).
func Test_InactiveProposerStillScheduled(t *testing.T) {
    requested := Address("0xAA")
    proposers := []Proposer{
        {Address: Address("0x01"), Active: true},
        {Address: requested, Active: false}, // inactive requested proposer
        {Address: Address("0x02"), Active: true},
    }

    sched, err := NewSchedulerV1(requested, proposers)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    found := false
    for _, p := range sched.Actives {
        if p.Address == requested {
            found = true
            break
        }
    }
    if !found {
        t.Fatalf("inactive proposer not in actives (expected vulnerable behavior)")
    }
    t.Logf("PoC success: inactive proposer present in actives -> bug confirmed")
}
