package main

import (
	"encoding/base64"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/crypto/ripemd160"
)

const (
	NULL    = 0
	INTEGER = 1
	FLOAT   = 2
	STRING  = 3
	BLOB    = 4
	DIGEST  = 6
	BOOL    = 17
	HLL     = 18
	MAP     = 19
	LIST    = 20
	LDT     = 21
	GEOJSON = 23
)

func main() {
	set := flag.String("set", "", "(optional) set name")
	key := flag.String("key", "", "key value")
	isInt := flag.Bool("int", false, "set this parameter to specify the key is an integer, not a string")
	isFloat := flag.Bool("float", false, "set this parameter to indicate that the key is a float64")
	isBytes := flag.Bool("bytes", false, "indicate that the key is a bytes value; if using this flag, key must be a base64-encoded string")
	flag.Parse()
	if *key == "" {
		flag.Usage()
		os.Exit(1)
	}
	var pk interface{}
	var err error
	pk = *key
	if *isInt {
		pk, err = strconv.Atoi(*key)
		if err != nil {
			fmt.Printf("ERROR: key '%s' is not an int\n", *key)
			os.Exit(1)
		}
	} else if *isFloat {
		pk, err = strconv.ParseFloat(*key, 64)
		if err != nil {
			fmt.Printf("ERROR: key '%s' is not a float64\n", *key)
			os.Exit(1)
		}
	} else if *isBytes {
		val, err := base64.StdEncoding.DecodeString(*key)
		if err != nil {
			fmt.Println("ERROR: key is not a base64-encoded value")
			os.Exit(1)
		}
		pk = val
	}

	hash := ripemd160.New()
	hash.Write([]byte(*set))
	if *isInt {
		hash.Write([]byte{byte(INTEGER)})
		binary.Write(hash, binary.BigEndian, uint64(pk.(int)))
	} else if *isFloat {
		hash.Write([]byte{byte(FLOAT)})
		binary.Write(hash, binary.BigEndian, pk.(float64))
	} else if *isBytes {
		hash.Write([]byte{byte(BLOB)})
		hash.Write(pk.([]byte))
	} else {
		hash.Write([]byte{byte(STRING)})
		hash.Write([]byte(pk.(string)))
	}
	digest := hash.Sum(nil)
	b64 := base64.StdEncoding.EncodeToString(digest)
	fmt.Printf("Digest(hex): %X\n", hash.Sum(nil))
	fmt.Printf("Digest(base64): %s\n", b64)
	partitionId := (uint16(digest[1])<<8 | uint16(digest[0])) << 4 >> 4
	fmt.Printf("PartitionId: %d\n", partitionId)
}
