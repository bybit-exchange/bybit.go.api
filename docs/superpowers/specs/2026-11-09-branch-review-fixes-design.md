# Branch Review Fixes Design

## Goal

Fix the defects found while reviewing the `doris-sdk-go` branch without broadening the public API change beyond the newly added endpoints.

## Scope

- Correct the request fields and JSON tags of the new Prediction API DTOs so they match the official V5 API contract.
- Give Crypto Loan fixed inventory its own response model because testnet returns `term` as a JSON number. Keep the Spot Margin response model independent.
- Change newly added millisecond timestamps, including tax report request boundaries, from `int` to `int64`.
- Remove trailing whitespace reported by `git diff --check`.
- Remove the obsolete RWA stocks removal design document from this branch.
- Do not change `InsLoanDelayLiqStatus`: its HTTP method and path match the official documentation, while its non-JSON testnet 404 is an environment-side limitation.

## Compatibility Strategy

Use the smallest type changes that make the wire contract accurate. Crypto Loan and Spot Margin will not share the fixed-inventory response type, so correcting Crypto Loan's numeric `term` will not silently change Spot Margin consumers. Prediction request structs will expose only documented fields and tags. Millisecond values will use `int64` so the SDK is safe on 32-bit targets.

## Testing

Tests will be added before production-code changes and will cover:

- JSON encoding of each corrected Prediction request DTO.
- Decoding a representative testnet Crypto Loan fixed-inventory response containing a numeric `term`.
- Compile-time or architecture-targeted checks that time fields use `int64` and remain safe on 32-bit platforms.
- Existing unit tests, `go vet ./...`, and `git diff --check main...HEAD`.

No state-changing testnet endpoint will be invoked. Read-only testnet behavior already observed during review will be used as validation evidence; credentials will not be stored in source, tests, logs, or documentation.

## Acceptance Criteria

- The affected request and response models match the documented wire format.
- Numeric Crypto Loan `term` values deserialize successfully.
- Newly introduced millisecond timestamps do not depend on platform `int` width.
- All tests and static checks pass.
- The branch contains no unrelated RWA stocks design artifact or whitespace error.
