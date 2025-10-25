Perfect — here’s exactly what to enter in **Immunefi’s submission form** so reviewers can process it quickly and clearly.
This format matches what the triage team expects for high-severity logic or consensus bugs.

---

## 🧾 **Immunefi Submission Example**

### **Title**

> Inactive proposers can still be scheduled due to missing active-check in `NewSchedulerV1`

---

### **Summary (Short Description)**

A logic flaw in the scheduler constructor (`NewSchedulerV1`) allows inactive or banned validators to remain in the proposer sequence.
Because the function does not verify that the proposer is active, it includes them in the active list unconditionally.
This leads to inactive or offline validators receiving block production slots, causing potential liveness stalls or deliberate griefing.

---

### **Severity**

> High (Consensus / Liveness Risk)

**Justification:**
The issue directly impacts chain liveness and consensus reliability.
Inactive validators can occupy proposer slots, preventing timely block production and degrading performance across the network.

---

### **Impact**

* Offline or banned proposers can be scheduled.
* Missed block slots → chain stalls, throughput reduction, delayed finality.
* Attackers could intentionally go inactive while still occupying slots (DoS/griefing vector).
* Breaks the documented behavior: *“If addr is not listed or not active, an error is returned.”*

---

### **Proof of Concept**

Attach your ZIP:
📎 `VeChain-scheduler-bug-min.zip`

Include this note in the description box:

> The ZIP contains a safe, self-contained Go project that demonstrates the bug and fix.
>
> * `poc/vuln` shows the inactive proposer being scheduled (vulnerable behavior).
> * `fix` contains the patched implementation rejecting inactive proposers.
>
> Run locally:
>
> ```bash
> go test ./poc/vuln -v
> go test ./fix -v
> ```
>
> The first test passes (bug present), the second passes after the fix (bug resolved).
> No network, funds, or mainnet interactions occur.

---

### **Steps to Reproduce (inline summary)**

1. Clone or unzip the PoC repository.
2. Run `go test ./poc/vuln -v`.
3. Observe that the inactive proposer is still added to the active set.
4. Apply the fix (`fix/sched_fixed.go`) and re-run `go test ./fix -v`.
5. Observe that the function now rejects inactive proposers.

---

### **Expected Result**

Inactive or banned proposers should **not** be included in the active scheduling sequence.

---

### **Actual Result**

Inactive proposers are still appended to the active list and can be scheduled to produce blocks.

---

### **Suggested Fix**

```go
if !listed || !proposer.Active {
    return nil, errors.New("unauthorized or inactive block proposer")
}
```

---

### **Safety Statement**

This PoC is fully **local and offline**.
It does **not** interact with any mainnet or public systems, nor does it affect real validators or funds.
It is safe to run in any environment.

---

### **Attachments**

✅ `VeChain-scheduler-bug-min.zip` (PoC + fixed version)
✅ `report.md` (full writeup)
✅ Optional: screenshot of passing test logs

---

Would you like me to give you a one-sentence **Impact summary line** (the one Immunefi uses to label submissions like “This bug affects network liveness and proposer fairness”)? It’s useful for your report header.
