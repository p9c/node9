package netsync

import (
	"fmt"
	"sync"
	"time"

	"github.com/p9c/pod/pkg/log"
	"github.com/p9c/pod/pkg/util"
)

type // blockProgressLogger provides periodic logging for other services in
	// order to show users progress of certain "actions" involving some or all
	// current blocks. Ex: syncing to best chain, indexing all blocks, etc.
	blockProgressLogger struct {
		receivedLogBlocks int64
		receivedLogTx     int64
		lastBlockLogTime  time.Time
		subsystemLogger   *log.Logger
		progressAction    string
		sync.Mutex
	}

func // newBlockProgressLogger returns a new block progress logger.
// The progress message is templated as follows:  {progressAction }
// {numProcessed} {blocks|block} in the last {timePeriod}
// ({numTxs}, height {lastBlockHeight}, {lastBlockTimeStamp})
newBlockProgressLogger(progressMessage string) *blockProgressLogger {
	return &blockProgressLogger{
		lastBlockLogTime: time.Now(),
		progressAction:   progressMessage,
	}
}

func // LogBlockHeight logs a new block height as an information message to
// show progress to the user. In order to prevent spam,
// it limits logging to one message every 10 seconds with duration and totals
// included.
(b *blockProgressLogger) LogBlockHeight(block *util.Block) {
	b.Lock()
	defer b.Unlock()
	b.receivedLogBlocks++
	b.receivedLogTx += int64(len(block.MsgBlock().Transactions))
	now := time.Now()
	duration := now.Sub(b.lastBlockLogTime)
	if duration < time.Second*2 {
		return
	}
	// Truncate the duration to 10s of milliseconds.
	durationMillis := int64(duration / time.Millisecond)
	tDuration := 10 * time.Millisecond * time.Duration(durationMillis/10)
	// Log information about new block height.
	blockStr := "blocks"
	if b.receivedLogBlocks == 1 {
		blockStr = "block "
	}
	txStr := "transactions"
	if b.receivedLogTx == 1 {
		txStr = "transaction "
	}
	log.INFOF(
		"%s %6d %s in the last %s (%6d %s, height %8d, %s)",
		b.progressAction,
		b.receivedLogBlocks,
		blockStr,
		fmt.Sprintf("%0.1fs", tDuration.Seconds()),
		b.receivedLogTx,
		txStr, block.Height(),
		block.MsgBlock().Header.Timestamp,
	)
	b.receivedLogBlocks = 0
	b.receivedLogTx = 0
	b.lastBlockLogTime = now
}
func
(b *blockProgressLogger) SetLastLogTime(time time.Time) {
	b.lastBlockLogTime = time
}
