# VeChain Scheduler Bug PoC

This repository demonstrates a bug where inactive proposers can still be scheduled
due to a missing active check in `NewSchedulerV1`.

- `poc/vuln`: vulnerable version + failing PoC test (shows inactive proposer scheduled)
- `fix`: patched version + passing test
- `report.md`: Immunefi-ready report

## Run PoC

```bash
go test ./poc/vuln -v
```

## Run fixed version

```bash
go test ./fix -v
```

## Package for submission

```bash
zip -r VeChain-scheduler-bug.zip .
```
