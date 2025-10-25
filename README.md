> **Disclosure-ready PoC for Immunefi submission**  
> Demonstrates inactive proposer scheduling bug in VeChain Scheduler.


# Inactive proposers can still be scheduled due to missing active-check in NewSchedulerV1

## Summary 

A logic flaw in the scheduler constructor (NewSchedulerV1) allows inactive or banned validators to remain in the proposer sequence.
Because the function does not verify that the proposer is active, it includes them in the active list unconditionally.
This leads to inactive or offline validators receiving block production slots, causing potential liveness stalls or deliberate griefing.

## Severity

High (Consensus / Liveness Risk)

Justification:
The issue directly impacts chain liveness and consensus reliability.
Inactive validators can occupy proposer slots, preventing timely block production and degrading performance across the network.

## Impact

Offline or banned proposers can be scheduled.

Missed block slots → chain stalls, throughput reduction, delayed finality.

Attackers could intentionally go inactive while still occupying slots (DoS/griefing vector).

Breaks the documented behavior: “If addr is not listed or not active, an error is returned.”

## Proof of Concept

The full runnable Proof of Concept is hosted at https://github.com/extropyCoder/VeChain-scheduler-bug
It contains:

poc/vuln/ — the vulnerable version (go test ./poc/vuln -v shows the inactive proposer being scheduled).

fix/ — the corrected version that rejects inactive proposers.

The PoC is local, safe, and fully self-contained. It does not connect to mainnet or modify live systems.
poc/vuln shows the inactive proposer being scheduled (vulnerable behavior).


Run locally:

go test ./poc/vuln -v
go test ./fix -v


The first test passes (bug present), the second passes after the fix (bug resolved).
No network, funds, or mainnet interactions occur.

### Steps to Reproduce 

Clone the PoC repository.

Run go test ./poc/vuln -v.

Observe that the inactive proposer is still added to the active set.

Apply the fix (fix/sched_fixed.go) and re-run go test ./fix -v.

Observe that the function now rejects inactive proposers.

### Expected Result

Inactive or banned proposers should not be included in the active scheduling sequence.

### Actual Result

Inactive proposers are still appended to the active list and can be scheduled to produce blocks.

Suggested Fix
if !listed || !proposer.Active {
    return nil, errors.New("unauthorized or inactive block proposer")
}

#### Safety Statement

This PoC is fully local and offline.
It does not interact with any mainnet or public systems, nor does it affect real validators or funds.
It is safe to run in any environment.