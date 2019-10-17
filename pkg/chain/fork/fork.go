// Package fork handles tracking the hard fork status and is used to determine which consensus rules apply on a block
package fork

import (
	"encoding/hex"
	"math/big"
	"math/rand"
	"time"

	"github.com/p9c/pod/pkg/log"
)

// AlgoParams are the identifying block version number and their minimum target bits
type AlgoParams struct {
	Version int32
	MinBits uint32
	AlgoID  uint32
	NSperOp int64
}

// HardForks is the details related to a hard fork, number, name and activation height
type HardForks struct {
	Number             uint32
	ActivationHeight   int32
	Name               string
	Algos              map[string]AlgoParams
	AlgoVers           map[int32]string
	WorkBase           int64
	TargetTimePerBlock int32
	AveragingInterval  int64
	TestnetStart       int32
}

var (
	// AlgoVers is the lookup for pre hardfork
	//
	AlgoVers = map[int32]string{
		2:   "sha256d",
		514: "scrypt",
	}
	// Algos are the specifications identifying the algorithm used in the
	// block proof
	Algos = map[string]AlgoParams{
		AlgoVers[2]: {
			Version: 2,
			MinBits: MainPowLimitBits,
			NSperOp: 824,
		}, // 824 ns/op
		AlgoVers[514]: {
			Version: 514,
			MinBits: MainPowLimitBits,
			AlgoID:  1,
			NSperOp: 740839,
		}, // 740839 ns/op
	}
	// FirstPowLimit is
	FirstPowLimit = func() big.Int {
		mplb, _ := hex.DecodeString(
			"0fffff0000000000000000000000000000000000000000000000000000000000")
		return *big.NewInt(0).SetBytes(mplb)
	}()
	// FirstPowLimitBits is
	FirstPowLimitBits = BigToCompact(&FirstPowLimit)
	// IsTestnet is set at startup here to be accessible to all other libraries
	IsTestnet bool
	// List is the list of existing hard forks and when they activate
	List = []HardForks{
		{
			Number:           0,
			Name:             "Halcyon days",
			ActivationHeight: 0,
			Algos:            Algos,
			AlgoVers:         AlgoVers,
			WorkBase: func() (out int64) {
				for i := range Algos {
					out += Algos[i].NSperOp
				}
				out /= int64(len(Algos))
				return
			}(),
			TargetTimePerBlock: 300,
			AveragingInterval:  10, // 50 minutes
			TestnetStart:       0,
		},
		{
			Number:           1,
			Name:             "Plan 9 from Crypto Space",
			ActivationHeight: 250000,
			Algos:            P9Algos,
			AlgoVers:         P9AlgoVers,
			WorkBase: func() (out int64) {
				for i := range P9Algos {
					out += P9Algos[i].NSperOp
				}
				out /= int64(len(P9Algos))
				return
			}(),
			TargetTimePerBlock: 36,
			AveragingInterval:  3600,
			TestnetStart:       1,
		},
	}
	// P9AlgoVers is the lookup for after 1st hardfork
	P9AlgoVers = map[int32]string{
		5:  "blake2b",
		6:  "argon2i",
		7:  "cryptonight7v2",
		8:  "keccak",
		9:  "scrypt",
		10: "sha256d",
		11: "skein",
		12: "stribog",
		13: "lyra2rev2",
	}
	// P9Algos is the algorithm specifications after the hard fork
	// given ns/op values are approximate and kopach bench writes them to
	// a file. These should vary on different architectures due to limitations
	// of division bit width and the cache behavior of hash functions, and
	// refers to one thread of execution, how this relates to number of cores
	// will vary also
	P9Algos = map[string]AlgoParams{
		P9AlgoVers[5]:  {5, FirstPowLimitBits, 0, 60425},     // 69495444},
		P9AlgoVers[6]:  {6, FirstPowLimitBits, 1, 66177},     // 79734306},
		P9AlgoVers[7]:  {7, FirstPowLimitBits, 2, 102982451}, // 69968425},
		P9AlgoVers[8]:  {8, FirstPowLimitBits, 3, 64589},     // 71988313},
		P9AlgoVers[9]:  {9, FirstPowLimitBits, 4, 1416611},   // 68395274},
		P9AlgoVers[10]: {10, FirstPowLimitBits, 5, 67040},    // 67460443},
		P9AlgoVers[11]: {11, FirstPowLimitBits, 7, 66803},    // 64433603},
		P9AlgoVers[12]: {12, FirstPowLimitBits, 6, 1858295},  // 69987634},
		P9AlgoVers[13]: {13, FirstPowLimitBits, 8, 134025},   // 64936544},
	}
	// SecondPowLimit is
	SecondPowLimit = func() big.Int {
		mplb, _ := hex.DecodeString(
			"00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
		return *big.NewInt(0).SetBytes(mplb)
	}()
	SecondPowLimitBits = BigToCompact(&SecondPowLimit)
	MainPowLimit       = func() big.Int {
		mplb, _ := hex.DecodeString(
			"00000fffff000000000000000000000000000000000000000000000000000000")
		return *big.NewInt(0).SetBytes(mplb)
	}()
	MainPowLimitBits = BigToCompact(&MainPowLimit)
)

// GetAlgoID returns the 'algo_id' which in pre-hardfork is not the same as the
// block version number, but is afterwards
func GetAlgoID(algoname string, height int32) uint32 {
	if GetCurrent(height) > 1 {
		return P9Algos[algoname].AlgoID
	}
	return Algos[algoname].AlgoID
}

// GetAlgoName returns the string identifier of an algorithm depending on
// hard fork activation status
func GetAlgoName(algoVer int32, height int32) (name string) {
	hf := GetCurrent(height)
	var ok bool
	name, ok = List[hf].AlgoVers[algoVer]
	if hf < 1 && !ok {
		name = "sha256d"
	}
	// INFO("GetAlgoName", algoVer, height, name}
	return
}

// GetAlgoVer returns the version number for a given algorithm (by string name)
// at a given height. If "random" is given, a random number is taken from the
// system secure random source (for randomised cpu mining)
func GetAlgoVer(name string, height int32) (version int32) {
	n := "sha256d"
	hf := GetCurrent(height)
	// INFO("GetAlgoVer", name, height, hf}
	if name == "random" {
		rng := rand.New(rand.NewSource(time.Now().Unix()))
		rn := rng.Intn(len(List[hf].AlgoVers)) + 5
		log.TRACE("random!", rn)
		randomalgover := int32(rn)
		switch hf {
		case 0:
			// INFO("rng", randomalgover, randomalgover }
			switch randomalgover & 1 {
			case 0:
				version = 2
			case 1:
				version = 514
			}
			return
		case 1:
			// INFO("rng", randomalgover, randomalgover }
			actualver := randomalgover
			// INFO("actualver", actualver}
			rndalgo := List[1].AlgoVers[actualver]
			// INFO("algo", rndalgo}
			algo := List[1].Algos[rndalgo].Version
			// INFO("actualalgo", algo}
			return algo
		}
	} else {
		n = name
	}
	version = List[hf].Algos[n].Version
	return
}

// GetAveragingInterval returns the active block interval target based on
// hard fork status
func GetAveragingInterval(height int32) (r int64) {
	r = List[GetCurrent(height)].AveragingInterval
	return
}

// GetCurrent returns the hardfork number code
func GetCurrent(height int32) (curr int) {
	if IsTestnet {
		for i := range List {
			if height >= List[i].TestnetStart {
				curr = i
			}
		}
	} else {
		for i := range List {
			if height >= List[i].ActivationHeight {
				curr = i
			}
		}
	}
	return
}

// GetMinBits returns the minimum diff bits based on height and testnet
func GetMinBits(algoname string, height int32) (mb uint32) {
	log.TRACE("GetMinBits", algoname)
	curr := GetCurrent(height)
	mb = List[curr].Algos[algoname].MinBits
	log.TRACE("minbits", mb)
	return
}

// GetMinDiff returns the minimum difficulty in uint256 form
func GetMinDiff(algoname string, height int32) (md *big.Int) {
	log.TRACE("GetMinDiff", algoname)
	return CompactToBig(GetMinBits(algoname, height))
}

// GetTargetTimePerBlock returns the active block interval target based on
// hard fork status
func GetTargetTimePerBlock(height int32) (r int64) {
	r = int64(List[GetCurrent(height)].TargetTimePerBlock)
	return
}
