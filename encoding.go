package hotstuff

import "encoding/binary"

// 64位编码
func EncodeUint64(u uint64) []byte {
	rst := make([]byte, 8)
	binary.BigEndian.PutUint64(rst, u)
	return rst
}

// 64位解码
func DecodeUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}
