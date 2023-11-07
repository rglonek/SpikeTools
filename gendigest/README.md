# GenDigest

Generate Aerospike Key Digest from Set Name and Key

## Installation and running

Two methods are available:
1. Head to the releases page, download the single binary and run it on your machine
2. Using docker: `docker run -it --rm robertglonek/gendigest`
   * example print help: `docker run -it --rm robertglonek/gendigest --help`
   * example use: `docker run -it --rm robertglonek/gendigest -key "bob" -set "test"`

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
