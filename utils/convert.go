package utils

import "strconv"

func String2Uint64(s string) uint64  {
	r,_ := strconv.ParseInt(s,10,64)
	return uint64(r)
}
