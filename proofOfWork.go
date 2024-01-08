package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
)

// 难度，表示哈希值前24位为零
const targetBits = 24

// 避免nounce溢出
const maxNounce = math.MaxInt64

// 该区块的哈希值必须要小于target
// target的解释：等于1（左移256-targetBits），这样就可以直接比较哈希值与target的大小，如果哈希值小于target则可以添加区块
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}

// 计算hash要用到的数据，前一个区块的hash，交易数据，时间戳，难度，nounce值
func (pow *ProofOfWork) prepareData(nounce int) []byte {
	transactionsJSON, err := json.Marshal(pow.block.Transactions)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			transactionsJSON,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nounce)),
		},
		[]byte{},
	)
	return data
}

// 工作量证明即计算并找到对应的nounce值
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nounce := 0
	for nounce < maxNounce {
		data := pow.prepareData(nounce)
		hash = sha256.Sum256(data)
		//将哈希值转化成大整数
		hashInt.SetBytes(hash[:])
		//比较哈希值与target，如果小于target返回-1退出循环
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nounce++
		}
	}
	return nounce, hash[:]
}

/*
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}
*/
