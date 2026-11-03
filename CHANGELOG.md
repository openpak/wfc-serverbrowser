# Changelog

## v0.1.2 - 2026-11-03

- Internal `/internal/` polls are counted and summarised once a minute instead of logged per call (wfc-common v0.3.1).

## v0.1.1 - 2026-11-03

- Every serverbrowser message now logs one line to stderr, `serverbrowser cmd <name> -> <result> <ms>ms` (verb and result only, never payload values). Internal HTTP logs as above through wfc-common v0.3.0 `httplog`.

