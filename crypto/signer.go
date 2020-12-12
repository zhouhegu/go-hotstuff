package crypto

import (
	"github.com/kilic/bls12-381/blssig"
)

/*
TODO this signature doesn't have any protection against rogue key attack
*/

var (
	// TODO separate domains for timeout and regular certificate
	domain = [8]byte{1, 1, 1}
)

// 门限签名生成器
func NewBLS12381Signer(priv blssig.SecretKey) *BLS12381Signer {
	return &BLS12381Signer{
		priv: priv,
	}
}

// 门限签名类，包含私钥
type BLS12381Signer struct {
	// 私钥类型为SecretKey
	priv blssig.SecretKey
}

// BLS12381Signer类型的扩展方法，可在BLS12381Signer的实例中调用该方法
func (s *BLS12381Signer) Sign(dst, msg []byte) []byte {
	// 要求消息为32个byte长度
	if len(msg) != 32 {
		panic("message must be exactly 32 bytes")
	}
	m := [32]byte{}
	copy(m[:], msg)
	// 用私钥对消息签名，暂时不懂domain的含义
	sig := blssig.Sign(m, domain, s.priv)
	// this conversion is quit inneficient, and should be done on the caller side when signature will be sent over
	// the wire
	// 添加到dst后面
	return append(dst, blssig.SignatureToCompressed(sig)...)
}
