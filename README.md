# signal-mob-sync-field

`signal-mob-sync-field` keeps a focused Go implementation around mobile workflows. The project goal is to create a Go reference implementation for sync workflows, centered on diagnostic reporting, negative fixtures, and human-readable error snapshots.

## Why It Exists

The project exists to keep a narrow engineering decision visible and testable. For this repo, that decision is how form pressure and local state should influence a review result.

## Signal Mob Sync Field Review Notes

The first comparison I would make is `form pressure` against `conflict cost` because it shows where the rule is most opinionated.

## Features

- `fixtures/domain_review.csv` adds cases for form pressure and sync drift.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/signal-mob-sync-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `form pressure` and `conflict cost`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Architecture Notes

The fixture data drives the tests. The code stays thin, while `metadata/domain-review.json` and `config/review-profile.json` explain what each case is meant to protect.

The Go code keeps the review rule close to the tests.

## Usage

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Tests

The verifier is intentionally local. It should fail if the fixture score math, lane assignment, or language-specific test drifts.

## Limitations And Roadmap

No external service is required. A deeper version would add more negative cases and a clearer boundary around invalid input.
