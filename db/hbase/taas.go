package hbase

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/pingcap/errors"
	"github.com/pingcap/go-ycsb/db/taas"
	"net"
	"sync/atomic"
	"time"
)

import (
	"context"
)

func (db *txnDB) TxnCommit(ctx context.Context, table string, keys []string, values []map[string][]byte) error {
	for taas.InitOk == 0 {
		time.Sleep(50)
	}

	t1 := time.Now().UnixNano()
	atomic.AddUint64(&taas.TotalTransactionCounter, 1)

	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(&thrift.TConfiguration{
		TBinaryStrictRead:  thrift.BoolPtr(true),
		TBinaryStrictWrite: thrift.BoolPtr(true),
	})
	transport := thrift.NewTSocketConf(net.JoinHostPort(taas.HbaseServerIp, PORT), &thrift.TConfiguration{
		ConnectTimeout: time.Second * 5,
		SocketTimeout:  time.Second * 5,
	})
	client := NewTHBaseServiceClientFactory(transport, protocolFactory)
	err := transport.Open()
	if err != nil {
		return err
	}
	defer func(transport *thrift.TSocket) {
		err := transport.Close()
		if err != nil {
			return
		}
	}(transport)

	var readOpNum, writeOpNum uint64 = 0, 0
	time1 := time.Now()
	for i, key := range keys {
		if values[i] == nil { //read
			readOpNum++
			rowKey := db.getRowKey(table, key)
			time2 := time.Now()
			rowData, err := client.Get(ctx, []byte(table), &TGet{Row: rowKey})
			if err != nil {
				return err
			} else if rowData == nil {
				return errors.New("txn read failed")
			}
			timeLen2 := time.Now().Sub(time2)
			atomic.AddUint64(&taas.TikvReadLatency, uint64(timeLen2))
			//fmt.Println("; Read, key : " + string(rowKey) + " Data : " + string(rowData))
		} else {
			writeOpNum++
			rowKey := db.getRowKey(table, key)
			var cvarr []*TColumnValue
			finalData, err1 := db.r.Encode(nil, values[i])
			if err1 != nil {
				return err1
			}
			cvarr = append(cvarr, &TColumnValue{
				Family:    []byte("entire"),
				Qualifier: []byte(""),
				Value:     finalData,
			})
			tempTPut := TPut{Row: rowKey, ColumnValues: cvarr}
			err = client.Put(ctx, []byte(table), &tempTPut)
			if err != nil {
				fmt.Println(err)
			}
		}

	}

	timeLen := time.Now().Sub(time1)
	atomic.AddUint64(&taas.TikvTotalLatency, uint64(timeLen))
	//fmt.Println("; read op : " + strconv.FormatUint(readOpNum, 10) + ", write op : " + strconv.FormatUint(writeOpNum, 10))

	t2 := uint64(time.Now().UnixNano() - t1)
	taas.TotalLatency += t2
	atomic.AddUint64(&taas.TotalReadCounter, readOpNum)
	atomic.AddUint64(&taas.TotalUpdateCounter, writeOpNum)
	atomic.AddUint64(&taas.SuccessReadCounter, readOpNum)
	atomic.AddUint64(&taas.SuccessUpdateCounter, writeOpNum)
	atomic.AddUint64(&taas.SuccessTransactionCounter, 1)

	return nil
}
