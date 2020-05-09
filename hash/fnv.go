package hash

import (
	"hash"
	"hash/fnv"
)

func fnv32(text string) (uint32, error) {
	hash := fnv.New32()
	return fnvSum32(hash, text)
}

func fnv32a(text string) (uint32, error) {
	hash := fnv.New32a()
	return fnvSum32(hash, text)
}

func fnv64(text string) (uint64, error) {
	hash := fnv.New64()
	return fnvSum64(hash, text)
}

func fnv64a(text string) (uint64, error) {
	hash := fnv.New64a()
	return fnvSum64(hash, text)
}

func fnvSum32(hash hash.Hash32, text string) (uint32, error) {
	_, err := hash.Write([]byte(text))
	return hash.Sum32(), err
}

func fnvSum64(hash hash.Hash64, text string) (uint64, error) {
	_, err := hash.Write([]byte(text))
	return hash.Sum64(), err
}
