# Changelog

## v0.2.0 - 2026-11-05

- Link-first (prds#139), behind `WFC_LINK_FIRST=1`: a server list goes only to an address with a live, linked GPCM login (from wfc-gpcm's session feed, refreshed every 10 s; an outage keeps the last set, and none yet allows). Refusals are summarised once a minute. Vendors wfc-common v0.8.0.

## v0.1.2 - 2026-11-03

- Internal `/internal/` polls are counted and summarised once a minute instead of logged per call (wfc-common v0.3.1).

## v0.1.1 - 2026-11-03

- Every serverbrowser message now logs one line to stderr, `serverbrowser cmd <name> -> <result> <ms>ms` (verb and result only, never payload values). Internal HTTP logs as above through wfc-common v0.3.0 `httplog`.

