package crypto

import (
	"github.com/dshulyak/go-hotstuff/types"
	"github.com/kilic/bls12-381/blssig"
)
// 门限签名验证器
func NewBLS12381Verifier(threshold int, pubkeys []blssig.PublicKey) *BLS12381Verifier {
	return &BLS12381Verifier{
		threshold: threshold,
		pubkeys:   pubkeys,
	}
}

type BLS12381Verifier struct {
	threshold int
	pubkeys   []blssig.PublicKey
}


func (v *BLS12381Verifier) Verify(idx uint64, msg, sig []byte) bool {
	if idx >= uint64(len(v.pubkeys)) {
		return false
	}
	// 获取公钥组内其中一个公钥
	key := v.pubkeys[idx]
	// 解压签名
	signature, err := blssig.NewSignatureFromCompresssed(sig)
	if err != nil {
		return false
	}
	m := [32]byte{}
	copy(m[:], msg)
	// 认证
	return blssig.Verify(m, domain, signature, &key)
}

// 整体门限签名认证
func (v *BLS12381Verifier) VerifyAggregated(msg []byte, asig *types.AggregatedSignature) bool {
	// 个数要满足
	if len(asig.Voters) < v.threshold {
		return false
	}
	pubs := make([]*blssig.PublicKey, 0, len(asig.Voters))
	for _, voter := range asig.Voters {
		if voter >= uint64(len(v.pubkeys)) {
			return false
		}
		pubs = append(pubs, &v.pubkeys[voter])
	}
	sig, err := blssig.NewSignatureFromCompresssed(asig.Sig)
	if err != nil {
		return false
	}
	m := [32]byte{}
	copy(m[:], msg)
	return blssig.VerifyAggregateCommon(m, domain, pubs, sig)
}

// 将一个部门签名合并到整体签名里面
func (v *BLS12381Verifier) Merge(asig *types.AggregatedSignature, voter uint64, sig []byte) {
	if voter >= uint64(len(v.pubkeys)) {
		return
	}
	for _, v := range asig.Voters {
		if v == voter {
			return
		}
	}
	if asig.Sig != nil {
		sig1, err := blssig.NewSignatureFromCompresssed(asig.Sig)
		if err != nil {
			return
		}
		sig2, err := blssig.NewSignatureFromCompresssed(sig)
		if err != nil {
			return
		}
		sig1 = blssig.AggregateSignature(sig1, sig2)
		asig.Sig = blssig.SignatureToCompressed(sig1)
	} else {
		asig.Sig = sig
	}
	asig.Voters = append(asig.Voters, voter)
}
