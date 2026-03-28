# Custom Mode

The `custom` mode allows the user to define their own generation options by 
providing a `.mus` file:

```text
mode: custom
int: varint
uint: varint
float: unsafe
byte: raw
bool: ord
string: ord
byte_slice: ord
time: raw
time_ser: TimeUnixUTC
```

The generator automatically looks for the `.mus` file using an 
"ascending search" logic:
1. It starts by looking in the current directory.
2. If not found, it moves up to the parent directory.
3. This continues until the file is found or the project root is reached (the 
   directory containing `go.mod` or a `.git` folder).

More specific `.mus` files found deeper in the directory structure 
override more general ones found further up.

Here are all supported keys/values for the `.mus` file:

| Key              | Values                                | Description                                             |
| :--------------- | :------------------------------------ | :------------------------------------------------------ |
| **`mode`**       | `custom`                              | Mode name.                                              |
| **`int`**        | `varint`, `raw`, `unsafe`             | Default package for all integer types.                  |
| **`uint`**       | `varint`, `raw`, `unsafe`             | Default package for all unsigned integer types.         |
| **`float`**      | `varint`, `raw`, `unsafe`             | Default package for all float types.                    |
| **`bool`**       | `ord`, `unsafe`                       | Default package for `bool`.                             |
| **`string`**     | `ord`, `unsafe`                       | Default package for `string`.                           |
| **`byte`**       | `varint`, `raw`, `unsafe`             | Default package for `byte`.                             |
| **`byte_slice`** | `ord`, `unsafe`                       | Default package for `[]byte`.                           |
| **`time`**       | `raw`, `unsafe`                       | Default package for `time.Time`.                        |
| **`time_ser`**   | `TimeUnix`, `TimeUnixMilli`,          | Specific `time.Time` serializer from `raw` or `unsafe`. |
|                  | `TimeUnixMicro`, `TimeUnixNano`,      |                                                         |
|                  | `TimeUnixUTC`, `TimeUnixMilliUTC`,    |                                                         |
|                  | `TimeUnixMicroUTC`, `TimeUnixNanoUTC` |                                                         |
