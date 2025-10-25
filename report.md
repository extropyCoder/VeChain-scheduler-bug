# Bug: Inactive proposer scheduling bug

## Brief / Intro
Inactive or banned proposers can still be scheduled due to missing activity check
in `NewSchedulerV1`.

## Details
When constructing the schedule, the function unconditionally appends the requested
proposer to the `actives` slice, even if that proposer is inactive.

## Impact
Offline or banned validators can occupy proposer slots, reducing network liveness.

## Risk Breakdown
- Exploit difficulty: Low
- Impact: High (consensus disruption)
- Severity: High

## Recommendation
Add an active-check before constructing the schedule:

```go
if !listed || !proposer.Active {
    return nil, errors.New("unauthorized or inactive block proposer")
}
```

## Proof of Concept
The PoC is provided in `poc/vuln/poc_test.go` and demonstrates the bug.
The fixed version is in `fix/`.

Run:
```bash
go test ./poc/vuln -v
go test ./fix -v
```
