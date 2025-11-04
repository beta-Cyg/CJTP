# Cygnuge Judge system Transfer Protocol (CJTP)

## Protocol Version (1 byte)

$0 ~ 2^8 - 1$

## Content Length (4 bytes)

$0 ~ 2^32 - 1$

- the length includes header
- minimum: 6 bytes

## Response/Request (1 bit)

- Response: 0
- Request: 1

## Response/Request Type (7 bit)

According to CJ

## Data Content (Determined by content length)

```
PV000000 CL000000
CL000000 CL000000
CL000000 R RT00000
DATA...
```

