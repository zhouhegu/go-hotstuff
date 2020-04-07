package types

import "golang.org/x/crypto/blake2s"

// 消息哈希函数
func (h *Header) Hash() []byte {
	// TODO ideally i want to memoize hash, but it requires wrapping proto-generated structs
	// 新建哈希对象
	digest, err := blake2s.New256(nil)
	if err != nil {
		panic(err.Error())
	}
	bytes, err := h.Marshal()
	if err != nil {
		panic(err.Error())
	}
	// 写入数据
	digest.Write(bytes)
	// TODO 添加到字节组中
	return digest.Sum(make([]byte, 0, digest.Size()))
}
