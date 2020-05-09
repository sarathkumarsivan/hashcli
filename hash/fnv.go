package hash

import (
	"hash"
	"hash/fnv"
)

func FNV32(text string) (uint32, error) {
	hash := fnv.New32()
	return fnv32(hash, text)
}

func FNV32a(text string) (uint32, error) {
	hash := fnv.New32a()
	return fnv32(hash, text)
}

func FNV64(text string) (uint64, error) {
	hash := fnv.New64()
	return fnv64(hash, text)
}

func FNV64a(text string) (uint64, error) {
	hash := fnv.New64a()
	return fnv64(hash, text)
}

func fnv32(hash hash.Hash32, text string) (uint32, error) {
	_, err := hash.Write([]byte(text))
	return hash.Sum32(), err
}

func fnv64(hash hash.Hash64, text string) (uint64, error) {
	_, err := hash.Write([]byte(text))
	return hash.Sum64(), err
}
