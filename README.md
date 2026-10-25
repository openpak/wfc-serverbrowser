# Wii/DS Server Browser Server

OpenPak's Wii/DS Server Browser Server.

| | |
| --- | --- |
| Network | Wii, DS |
| Role | GameSpy server-browser TCP queries and client messages; QR2 owns the directory |
| Depends on | `WFC_QR2_URL` (QR2 service), existing WFC database |
| Used by | wfc-nas (console HTTP) |
| Image | `ghcr.io/openpak/wfc-serverbrowser` |
| Based on | WiiLink WFC (AGPL-3.0) |
| Licence | AGPL-3.0-only |

## What it serves

| Endpoint | What |
| --- | --- |
| TCP 28910 | GameSpy server-browser TCP queries and client messages |
| HTTP :20202 `GET /healthz` | Unauthenticated health check; protocol/internal paths need `WFC_INTERNAL_KEY` |

## Build

```sh
docker build -t wfc-serverbrowser .
```

## Docs

- [Protocol](docs/protocol.md)
- [Evidence](docs/evidence.md)
- [Deploy](docs/deploy.md)
- [Notes](docs/notes.md)

## Licence

AGPL-3.0-only. See [LICENSE](LICENSE).
