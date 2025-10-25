package vuln

import "errors"

type Address string

type Proposer struct {
    Address Address
    Active  bool
}

type SchedulerV1 struct {
    Proposer Proposer
    Actives  []Proposer
}

// Vulnerable: always appends requested proposer, even if inactive.
func NewSchedulerV1(addr Address, proposers []Proposer) (*SchedulerV1, error) {
    actives := make([]Proposer, 0, len(proposers))
    listed := false
    var proposer Proposer
    for _, p := range proposers {
        if p.Address == addr {
            proposer = p
            actives = append(actives, p) // BUG: added even if inactive
            listed = true
        } else if p.Active {
            actives = append(actives, p)
        }
    }
    if !listed {
        return nil, errors.New("unauthorized block proposer")
    }
    return &SchedulerV1{proposer, actives}, nil
}
