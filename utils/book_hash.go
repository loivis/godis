package utils

import (
	"hash/fnv"
)

// BookHash ...
func BookHash(site, book string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(site + book))
	return hash.Sum32()
}
