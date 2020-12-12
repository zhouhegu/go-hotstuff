package crypto

import (
	"crypto/rand"
	"io"

	"github.com/kilic/bls12-381/blssig"
)

type (
	PublicKey = blssig.PublicKey
	PrivateKey = blssig.SecretKey
)

// 生成一对公私钥
func GenerateKey(rng io.Reader) (PrivateKey, *PublicKey, error) {
	if rng == nil {
		rng = rand.Reader
	}
	// 随机生成私钥
	priv, err := blssig.RandSecretKey(rng)
	if err != nil {
		return nil, nil, err
	}
	// 估计私钥生成公钥，通过PublicKeyFromSecretKey生成公钥并返回
	return priv, blssig.PublicKeyFromSecretKey(priv), nil
}

// 生成n对公私钥
func GenerateKeys(rng io.Reader, n int) ([]PublicKey, []PrivateKey, error) {
	pubs := make([]blssig.PublicKey, n)
	privs := make([]blssig.SecretKey, n)
	for i := range privs {
		priv, pub, err := GenerateKey(rng)
		if err != nil {
			return nil, nil, err
		}
		privs[i] = priv
		pubs[i] = *pub
	}
	return pubs, privs, nil
}
