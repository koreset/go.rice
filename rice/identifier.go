package main

import "strconv"

var identifierCount incremental.Uint64

func nextIdentifier() string {
	num := identifierCount.Next()
	return strconv.FormatUint(num, 36) // 0123456789abcdefghijklmnopqrstuvwxyz
}
