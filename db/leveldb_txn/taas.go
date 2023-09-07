package leveldb_txn

//#include ""
import (
	"context"
	"errors"
	"github.com/pingcap/go-ycsb/db/taas"
	"sync/atomic"
	"time"
)

func (db *txnDB) TxnCommit(ctx context.Context, table string, keys []string, values []map[string][]byte) error {
	for taas.InitOk == 0 {
		time.Sleep(50)
	}

	t1 := time.Now().UnixNano()
	atomic.AddUint64(&taas.TotalTransactionCounter, 1)

	var readOpNum, writeOpNum uint64 = 0, 0
	time1 := time.Now()
	for i, key := range keys {
		if values[i] == nil { // 如果values[i]为nil，则表示读取操作
			readOpNum++
			rowKey := db.getRowKey(table, key)
			time2 := time.Now()
			rowData, err := db.client.Get(rowKey)
			//rowData, err := db.db.Get(rowKey, nil)
			timeLen2 := time.Now().Sub(time2)
			atomic.AddUint64(&taas.TikvReadLatency, uint64(timeLen2))
			if err != nil {
				return err
			} else if rowData == nil {
				return errors.New("txn read failed")
			}
		} else {
			writeOpNum++
			rowKey := db.getRowKey(table, key)
			buf := db.bufPool.Get()
			rowData, err := db.r.Encode(buf, values[i])
			if err != nil {
				return err
			}
			db.client.Put(rowKey, rowData)
			if err != nil {
				return err
			}
		}

	}

	timeLen := time.Now().Sub(time1)
	atomic.AddUint64(&taas.TikvTotalLatency, uint64(timeLen))
	//fmt.Println("; read op : " + strconv.FormatUint(readOpNum, 10) + ", write op : " + strconv.FormatUint(writeOpNum, 10))

	t2 := uint64(time.Now().UnixNano() - t1)
	taas.TotalLatency += t2
	//append(latency, t2)
	//result, ok := "Abort", true
	atomic.AddUint64(&taas.TotalReadCounter, uint64(readOpNum))
	atomic.AddUint64(&taas.TotalUpdateCounter, uint64(writeOpNum))

	atomic.AddUint64(&taas.SuccessReadCounter, uint64(readOpNum))
	atomic.AddUint64(&taas.SuccessUpdateCounter, uint64(writeOpNum))
	atomic.AddUint64(&taas.SuccessTransactionCounter, 1)

	return nil
}
