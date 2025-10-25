# VeChain Scheduler Bug PoC (Minimal)

Minimal, local, and **safe** PoC showing that an inactive proposer can still be scheduled.

## Run PoC (vulnerable behavior)
```bash
go test ./poc/vuln -v
```

## Run fixed version
```bash
go test ./fix -v
```
