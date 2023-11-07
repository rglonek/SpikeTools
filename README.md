# SpikeTools

Download from the releases page

## GenDigest

Generate Aerospike digest from Set Name and Primary Key

Two methods of running it:
1. Head to the releases page, download the single binary and run it on your machine
2. Using docker: `docker run -it --rm robertglonek/gendigest`
   * example print help: `docker run -it --rm robertglonek/gendigest --help`
   * example use: `docker run -it --rm robertglonek/gendigest -key "bob" -set "test"`
