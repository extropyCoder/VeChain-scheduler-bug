package fix

import "testing"

func Test_InactiveProposerRejected(t *testing.T) {
    requested := Address("0xAA")
    proposers := []Proposer{
        {Address: Address("0x01"), Active: true},
        {Address: requested, Active: false}, // inactive requested proposer
        {Address: Address("0x02"), Active: true},
    }

    _, err := NewSchedulerV1(requested, proposers)
    if err == nil {
        t.Fatalf("expected error for inactive proposer, got nil")
    }
    t.Logf("Fixed behavior: rejected inactive proposer (%v)", err)
}
