# Signal Mob Sync Field Walkthrough

I use this file as a small checklist before changing the Go implementation.

| Case | Focus | Score | Lane |
| --- | --- | ---: | --- |
| baseline | form pressure | 231 | ship |
| stress | sync drift | 211 | ship |
| edge | local state | 221 | ship |
| recovery | conflict cost | 197 | ship |
| stale | form pressure | 199 | ship |

Start with `baseline` and `recovery`. They create the widest contrast in this repository's fixture set, which makes them better review anchors than the middle cases.

The next useful expansion would be a malformed fixture around sync drift and conflict cost.
