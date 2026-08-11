# `logs`

This package models provider/API audit records in `log_logs`. `LogLog` is the
persistent row, `ServicesT` identifies the external integration, `LogTypeT`
classifies request direction, and `ObjectType` identifies the related Django
content type.

The main APIs are:

- `LogHttpRequest` for an incoming HTTP request record;
- `(*LogLog).LogRequest` / `LogResponse` hooks for HTTP clients;
- `(*LogLog).UpdateLog` to attach the response;
- `CreateRawEntry` for already-normalized provider request/response data;
- `GetContentID` to resolve an app-label/model pair through Django
  `django_content_type`.

`evm-api`, `wallet-api`, `payment-api`, `escrow-api`, `kyc-api`, `esign-api`,
and `email-worker` import this package for integration auditing or content-type
resolution.

Request/response bodies and headers can contain authorization data, personal
information, signatures, or provider secrets. Callers must redact before
passing them here and apply an appropriate retention/access policy; this
package normalizes JSON but does not perform general secret redaction.

See the [root guide](../README.md) for shared-model boundaries.
