# Review Journal

The repository goal stays the same: create a Go reference implementation for sync workflows, centered on diagnostic reporting, negative fixtures, and human-readable error snapshots. This note explains the added review angle.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 231, lane `ship`
- `stress`: `sync drift`, score 211, lane `ship`
- `edge`: `local state`, score 221, lane `ship`
- `recovery`: `conflict cost`, score 197, lane `ship`
- `stale`: `form pressure`, score 199, lane `ship`

## Note

A future change should add new cases before it changes the scoring rule.
