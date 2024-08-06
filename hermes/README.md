# Hermes Client

More information about the Hermes API can be found [here](https://docs.pyth.network/price-feeds/api-instances-and-providers/hermes).

## Current Supported Endpoints

- `/v2/updates/price/latest` ([documentation](https://hermes.pyth.network/docs/#/rest/-latest_price_updates))
  - For batch requests, both single-batch (sync) and multiple concurrent-single (async) requests are supported.
- `/v2/updates/price/stream` ([documentation](https://hermes.pyth.network/docs/#/rest/price_stream_sse_handler))
  - SSE streaming is supported.
