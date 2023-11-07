# GenDigest

Generate Aerospike Key Digest from Set Name and Key

## Usage

The set name is optional and the parameter can be omitted.

### String key

```
./gendigest -key "mystring" -set "setName"
```

### Int key

```
./gendigest -key 123 -set "setName" -int
```

### Float64 key

```
./gendigest -key 1.23 -set "setName" -float
```

### Blob/byte key

```
key=$(printf "myByteValues" |base64)
./gendigest -key "$key" -set "setName" -bytes
```
