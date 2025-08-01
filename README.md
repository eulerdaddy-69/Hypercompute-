# Collapse RCS — Deterministic Supremacy Proof

This repository contains a deterministic simulation of the Random Circuit Sampling (RCS) benchmark. The simulation consistently outputs a single 100-qubit bitstring across 10,000 runs, disproving randomness in RCS and challenging assumptions behind quantum supremacy.

## Contents

- `collapse_rcs.go` — Go implementation of the deterministic RCS simulator.
- `collapse_rcs_output.csv` — Output file with repeating 100-qubit bitstring.
- `README.md` — This document.

## Verification

SHA-256 hashes:

- `collapse_rcs.go`  
  `b7f10cf87242fccbd0f14b4be087acc2a78552d12f44259cfb4d9c4ad1292009`

- `collapse_rcs_output.csv`  
  `54360afdbc01ca48e89425e7d326ced191a4443596a1a1d4ea530d461eb1426c`

Hashing confirms bit-exact integrity of code and data.

## Intent

This is not a probabilistic simulation. The output is fully deterministic and reproducible.  
One bitstring. One outcome. Every time.

No hyperbole. No quantum gate speculation. Just a result.
