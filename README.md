# gostudy

A pair of small Go projects used to study and gain hands-on familiarity with the
language — its type system, testing conventions, and idioms — rather than to
ship a product. Each project is its own module.

## cipher

A toolkit of classical (pre-modern) ciphers, implemented behind a common
`Cipher` interface (`Name`, `Encrypt`, `Decrypt`):

- **Caesar** and **Rot13** shift ciphers
- **Atbash** substitution cipher
- **Rail Fence** transposition cipher
- **Combiner**, which chains two ciphers together (encrypt outer-then-inner,
  decrypt inner-then-outer)

See [cipher/cipher.go](cipher/cipher.go) for the interface and a runnable demo.

## npcgen

A random NPC generator for tabletop D&D-style character creation: race,
class, and ability scores, built around a reusable weighted-roll
`GeneratorTable` that can be loaded from a comma-separated string or a CSV
file and rolled against a seeded `Dice`.

See [npcgen/design_notes.txt](npcgen/design_notes.txt) for the original
project scope and ordering of generation steps.

## Practice

This repo also doubles as a source of small, real bugs for short (~30 minute)
bug-fixing practice sessions — find one, write a failing test that proves it,
fix it, confirm the test passes.
