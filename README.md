# gocached
==========

## Commands

`get`
returns a value for the key, or null

`set`
sets a value for the key

Commands are communicated via JSON with required attribuets `command`, `key`, `value`, and an optional `expiration` attribute.  Example:

```json
{
  command:    "get",
  key:        "user:name",
  value:      "Sean Yu",
  expiration: "1393918020" //unix timestamp
}
```
