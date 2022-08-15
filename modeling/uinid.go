package modelling

import (
	"hash/fnv"
)

type UinId struct {
	Identifier string `json:"identifier"`
	HashCode   int    `json:"hashCode"`
}

func NewUinId(identifier string) UinId {
	uinId := UinId{identifier, 0}
	uinId.HashCode = hash(identifier)
	return uinId
}

func hash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}
