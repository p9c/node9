package hardfork

import (
	"encoding/hex"

	"github.com/p9c/node9/pkg/chain/config/netparams"
	"github.com/p9c/node9/pkg/util"
)

// Payee is an address and amount
type Payee struct {
	Address util.Address
	Amount  util.Amount
}

var (
	// The following prepares hard fork disbursement payout transactions
	tn = &netparams.TestNet3Params
	mn = &netparams.MainNetParams
	// Payees are the list of payments to be made on the hard fork activation on
	// mainnet
	Payees = []Payee{
		{Addr("ag7s5bmcA8XoP1CcS1QPjiD4C5hhMWATik", mn), Amt(4400)},
	}
	// TestnetPayees are the list of payments to be made on the hard fork
	// activation in the testnet
	//
	// these are made using the following seed for testnet
	// f4d2c4c542bb52512ed9e6bbfa2d000e576a0c8b4ebd1acafd7efa37247366bc
	TestnetPayees = []Payee{
		{Addr("8K73LTaMHZmwwqe4vTHu7wm7QtwusvRCwC", tn), Amt(100)},
		// {Addr("8JEEhaMxJf4dZh5rvVCVSA7JKeYBvy8fir", tn), Amt(15500)},
		// {Addr("8bec3m8qpMePrBPHDAyCrkSm7TanGX8yWW", tn), Amt(1223)},
		// {Addr("8MCLEWq8pjXikrpb9rF9M5DpnpaoWPUD2W", tn), Amt(4000)},
		// {Addr("8cYGvT7km339nVukTj3ztfyQDFEHFivBNk", tn), Amt(2440)},
		// {Addr("8YUAAfUeS2mqUnsfiwDwQcEbMfM3tazKr7", tn), Amt(100)},
		// {Addr("8MMam6gxH1ns5LqASfhkHfRV2vsQaoM9VC", tn), Amt(8800)},
		// {Addr("8JABYpdqqyRD5FbACtMJ3XF5HJ38jaytrk", tn), Amt(422)},
		// {Addr("8MUnJMYi5Fo7Bm5Pmpr7JjdL3ZDJ7wqmXJ", tn), Amt(5000)},
		// {Addr("8d2RLbCBE8CiF4DetVuRfFFLEJJaXYjhdH", tn), Amt(30000)},
	}
	// CorePubkeyBytes is the address and public keys for the core dev
	// disbursement
	CorePubkeyBytes = [][]byte{
		//nWo
		Key("021a00c7e054279124e2d3eb8b64a58f1fda515464cd8df3c0823d2ff2931ebf37"),
		// loki
		Key("0387484f75bc5e45092b1334684def6b47f3dba1566b4b87f62d11c73d8f98db3e"),
		// trax0r
		Key("02daf0bda15f83899f4ebb62fd837c2dd2368ec8ed90ed0f050054d75d35935c99"),
	}
	// CoreAmount is the amount paid into the dev pool
	CoreAmount = Amt(30000)
	// TestnetCorePubkeyBytes are the addresses for the 3 of 4 multisig payment
	// for dev costs
	//
	// these are made using the following seed for testnet
	// f4d2c4c542bb52512ed9e6bbfa2d000e576a0c8b4ebd1acafd7efa37247366bc
	TestnetCorePubkeyBytes = [][]byte{
		// "8cL2fDzTSMu9Cd2rFi1dceWitQheAaCgTs",
		Key("03f040c0cff7918415974f05154c8ffe126ad93db7216103fb6f4080dc3bcf4803"),
		// "8YUDhyrcGrQk4PMpxnaNk1XyLRpQLa7N47",
		Key("03f5a5ff1ce0564c7f4565a108220ebac9bd544b44e79ca5a2a805e585d8297cc6"),
		// "8Rpf7CT4ikJQqXRpSp4EnAypKmidhHADN2",
		Key("022976653e490cea689faafa899aa41b6295c32a5fb3e02d0fa201ac698e0c0c24"),
		// "8Yw41PD1A3RviyjFQc38L9VufZasDU1pY8",
		Key("029ed2885ea597fddea070a5c4c9f40900a514f67f9d5f662aa7b556e8bc5a26f8"),
	}
	// TestnetCoreAmount is the amount paid into the dev pool
	TestnetCoreAmount = Amt(30000)
)

func Amt(f float64) (amt util.Amount) {
	amt, err := util.NewAmount(f)
	if err != nil {
		panic(err)
	}
	return
}

func Addr(addr string, defaultNet *netparams.Params) (out util.Address) {
	out, err := util.DecodeAddress(addr, defaultNet)
	if err != nil {
		panic(err)
	}
	return
}

func Key(key string) (out []byte) {
	out, err := hex.DecodeString(key)
	if err != nil {
		panic(err)
	}
	return
}
