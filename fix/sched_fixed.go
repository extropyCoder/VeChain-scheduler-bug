package fix

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

// Fixed: ensure requested proposer is active; do not include inactive in actives.
func NewSchedulerV1(addr Address, proposers []Proposer) (*SchedulerV1, error) {
    actives := make([]Proposer, 0, len(proposers))
    listed := false
    var proposer Proposer
    for _, p := range proposers {
        if p.Address == addr {
            proposer = p
            listed = true
            if p.Active {
                actives = append(actives, p)
            }
        } else if p.Active {
            actives = append(actives, p)
        }
    }
    if !listed {
        return nil, errors.New("unauthorized block proposer")
    }
    if !proposer.Active {
        return nil, errors.New("unauthorized or inactive block proposer")
    }
    return &SchedulerV1{proposer, actives}, nil
}
