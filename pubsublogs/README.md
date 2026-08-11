# `pubsublogs`

This package defines the legacy `PubsubLog` projection for `pubsub_logs`. It
stores message/topic/payload and processing metadata and implements the normal
generic model contract.

No direct import was found in the checked sibling services. Current event
consumers use [`pubsubactivities`](../pubsubactivities/README.md) for durable
deduplication, leases, and domain-reaction audit state. Do not use `PubsubLog`
as a substitute for that concurrency protocol.

Retain this model for compatibility/reporting until its database table and
historical data are formally retired.

See the [root guide](../README.md) for package lifecycle conventions.
