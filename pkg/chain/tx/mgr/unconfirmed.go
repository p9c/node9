package wtxmgr

import (
	chainhash "github.com/p9c/node9/pkg/chain/hash"
	"github.com/p9c/node9/pkg/chain/wire"
	"github.com/p9c/node9/pkg/log"
	walletdb "github.com/p9c/node9/pkg/wallet/db"
)

// insertMemPoolTx inserts the unmined transaction record.  It also marks
// previous outputs referenced by the inputs as spent.
func (s *Store) insertMemPoolTx(ns walletdb.ReadWriteBucket, rec *TxRecord) error {
	// Check whether the transaction has already been added to the
	// unconfirmed bucket.
	if existsRawUnmined(ns, rec.Hash[:]) != nil {
		// TODO: compare serialized txs to ensure this isn't a hash
		// collision?
		return nil
	}
	// Since transaction records within the store are keyed by their
	// transaction _and_ block confirmation, we'll iterate through the
	// transaction's outputs to determine if we've already seen them to
	// prevent from adding this transaction to the unconfirmed bucket.
	for i := range rec.MsgTx.TxOut {
		k := canonicalOutPoint(&rec.Hash, uint32(i))
		if existsRawUnspent(ns, k) != nil {
			return nil
		}
	}
	log.INFO("inserting unconfirmed transaction", rec.Hash)
	v, err := valueTxRecord(rec)
	if err != nil {
		return err
	}
	err = putRawUnmined(ns, rec.Hash[:], v)
	if err != nil {
		return err
	}
	for _, input := range rec.MsgTx.TxIn {
		prevOut := &input.PreviousOutPoint
		k := canonicalOutPoint(&prevOut.Hash, prevOut.Index)
		err = putRawUnminedInput(ns, k, rec.Hash[:])
		if err != nil {
			return err
		}
	}
	// TODO: increment credit amount for each credit (but those are unknown
	// here currently).
	return nil
}

// removeDoubleSpends checks for any unmined transactions which would introduce
// a double spend if tx was added to the store (either as a confirmed or unmined
// transaction).  Each conflicting transaction and all transactions which spend
// it are recursively removed.
func (s *Store) removeDoubleSpends(ns walletdb.ReadWriteBucket, rec *TxRecord) error {
	for _, input := range rec.MsgTx.TxIn {
		prevOut := &input.PreviousOutPoint
		prevOutKey := canonicalOutPoint(&prevOut.Hash, prevOut.Index)
		doubleSpendHashes := fetchUnminedInputSpendTxHashes(ns, prevOutKey)
		for _, doubleSpendHash := range doubleSpendHashes {
			doubleSpendVal := existsRawUnmined(ns, doubleSpendHash[:])
			// If the spending transaction spends multiple outputs
			// from the same transaction, we'll find duplicate
			// entries within the store, so it's possible we're
			// unable to find it if the conflicts have already been
			// removed in a previous iteration.
			if doubleSpendVal == nil {
				continue
			}
			var doubleSpend TxRecord
			doubleSpend.Hash = doubleSpendHash
			err := readRawTxRecord(
				&doubleSpend.Hash, doubleSpendVal, &doubleSpend,
			)
			if err != nil {
				return err
			}
			log.DEBUG(
				"removing double spending transaction", doubleSpend.Hash)
			if err := RemoveConflict(ns, &doubleSpend); err != nil {
				return err
			}
		}
	}
	return nil
}

func // RemoveConflict removes an unmined transaction record and all spend
// chains deriving from it from the store.
// This is designed to remove transactions that would otherwise result in
// double spend conflicts if left in the store,
// and to remove transactions that spend coinbase transactions on reorgs.
RemoveConflict(ns walletdb.ReadWriteBucket, rec *TxRecord) error {
	// For each potential credit for this record, each spender (if any) must
	// be recursively removed as well.  Once the spenders are removed, the
	// credit is deleted.
	for i := range rec.MsgTx.TxOut {
		k := canonicalOutPoint(&rec.Hash, uint32(i))
		spenderHashes := fetchUnminedInputSpendTxHashes(ns, k)
		for _, spenderHash := range spenderHashes {
			spenderVal := existsRawUnmined(ns, spenderHash[:])
			// If the spending transaction spends multiple outputs
			// from the same transaction, we'll find duplicate
			// entries within the store, so it's possible we're
			// unable to find it if the conflicts have already been
			// removed in a previous iteration.
			if spenderVal == nil {
				continue
			}
			var spender TxRecord
			spender.Hash = spenderHash
			err := readRawTxRecord(&spender.Hash, spenderVal, &spender)
			if err != nil {
				return err
			}
			log.DEBUGF(
				"transaction %v is part of a removed conflict chain -- removing as well %s",
				spender.Hash)
			if err := RemoveConflict(ns, &spender); err != nil {
				return err
			}
		}
		if err := deleteRawUnminedCredit(ns, k); err != nil {
			return err
		}
	}
	// If this tx spends any previous credits (either mined or unmined), set
	// each unspent.  Mined transactions are only marked spent by having the
	// output in the unmined inputs bucket.
	for _, input := range rec.MsgTx.TxIn {
		prevOut := &input.PreviousOutPoint
		k := canonicalOutPoint(&prevOut.Hash, prevOut.Index)
		if err := deleteRawUnminedInput(ns, k); err != nil {
			return err
		}
	}
	return deleteRawUnmined(ns, rec.Hash[:])
}

func // UnminedTxs returns the underlying transactions for all unmined
// transactions which are not known to have been mined in a block.
// Transactions are guaranteed to be sorted by their dependency order.
(s *Store) UnminedTxs(ns walletdb.ReadBucket) ([]*wire.MsgTx, error) {
	recSet, err := s.unminedTxRecords(ns)
	if err != nil {
		return nil, err
	}
	recs := dependencySort(recSet)
	txs := make([]*wire.MsgTx, 0, len(recs))
	for _, rec := range recs {
		txs = append(txs, &rec.MsgTx)
	}
	return txs, nil
}
func
(s *Store) unminedTxRecords(ns walletdb.ReadBucket) (map[chainhash.Hash]*TxRecord, error) {
	unmined := make(map[chainhash.Hash]*TxRecord)
	err := ns.NestedReadBucket(bucketUnmined).ForEach(func(k, v []byte) error {
		var txHash chainhash.Hash
		err := readRawUnminedHash(k, &txHash)
		if err != nil {
			return err
		}
		rec := new(TxRecord)
		err = readRawTxRecord(&txHash, v, rec)
		if err != nil {
			return err
		}
		unmined[rec.Hash] = rec
		return nil
	})
	return unmined, err
}

func // UnminedTxHashes returns the hashes of all transactions not known to
// have been mined in a block.
(s *Store) UnminedTxHashes(ns walletdb.ReadBucket) ([]*chainhash.Hash, error) {
	return s.unminedTxHashes(ns)
}
func (s *Store) unminedTxHashes(ns walletdb.ReadBucket) ([]*chainhash.Hash, error) {
	var hashes []*chainhash.Hash
	err := ns.NestedReadBucket(bucketUnmined).ForEach(func(k, v []byte) error {
		hash := new(chainhash.Hash)
		err := readRawUnminedHash(k, hash)
		if err == nil {
			hashes = append(hashes, hash)
		}
		return err
	})
	return hashes, err
}
