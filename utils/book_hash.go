package utils

import (
	"hash/fnv"
)

// Hash ...
func Hash(args ...string) uint32 {
	var s string
	hash := fnv.New32a()
	for _, arg := range args {
		s += arg
	}
	hash.Write([]byte(s))
	return hash.Sum32()
}
