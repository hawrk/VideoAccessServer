package utils

import (
	"VideoAccessServer/internal/config"
	Murhash "github.com/rryqszq4/go-murmurhash"
)

func GetMurhash2(str string) uint32 {
	var key = []byte(str)
	return Murhash.MurmurHash2(key,config.MUR_HASH_SEED)
}

func GetMurhas(str string) uint32 {
	var key = []byte(str)
	return Murhash.MurmurHash1(key, config.MUR_HASH_SEED)
}
