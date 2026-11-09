# Remove RWA Stocks APIs

## Scope

Remove every Go client method whose endpoint starts with `/v5/rwa/stocks/`, together with the response/request types used only by those methods and their entries in the 2026-11-09 release notes.

The existing `/v5/earn/rwa/` APIs are unrelated and remain unchanged.

## Changes

- Delete `stocks.go`, which contains all nine `/v5/rwa/stocks/` client methods.
- Delete `models/stocksResponse.go`, whose types are dedicated to those methods.
- Remove the nine function entries and sixteen struct entries from `RELEASE_NOTE.md`.
- Update the release-note totals from 48 to 39 functions and from 110 to 94 structs.

## Verification

- Search the repository to confirm no `/v5/rwa/stocks/` endpoint, stock API method, or dedicated stock model remains.
- Run formatting/checks applicable to the edited Markdown and Go repository.
- Run the full Go test suite.
