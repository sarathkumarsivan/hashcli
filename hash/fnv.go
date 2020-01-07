package hash

import (
	"hash"
	"hash/fnv"
)

// Fnv32 returns the 32-bit FNV-1 hash of the given text.
func Fnv32(text string) (uint32, error) {
	hash := fnv.New32()
	return fnv32(hash, text)
}

// Fnv32a returns the 32-bit FNV-1a hash of the given text.
func Fnv32a(text string) (uint32, error) {
	hash := fnv.New32a()
	return fnv32(hash, text)
}

// Fnv64 returns the 64-bit FNV-1 hash of the given text.
func Fnv64(text string) (uint64, error) {
	hash := fnv.New64()
	return fnv64(hash, text)
}

// Fnv64a returns the 64-bit FNV-1a hash of the given text.
func Fnv64a(text string) (uint64, error) {
	hash := fnv.New64a()
	return fnv64(hash, text)
}

// fnv32 returns the 32-bit FNV-1 hash of the given text.
func fnv32(hash hash.Hash32, text string) (uint32, error) {
	_, err := hash.Write([]byte(text))
	return hash.Sum32(), err
}

// fnv64 returns the 64-bit FNV-1 hash of the given text.
func fnv64(hash hash.Hash64, text string) (uint64, error) {
	_, err := hash.Write([]byte(text))
	return hash.Sum64(), err
}
