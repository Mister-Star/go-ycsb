package hbase

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/magiconair/properties"
	"github.com/pingcap/go-ycsb/pkg/util"
	"github.com/pingcap/go-ycsb/pkg/ycsb"
	"github.com/tikv/client-go/v2/txnkv/transaction"
	"net"
	"reflect"
	"time"
)

import (
	"context"
)

//const (
//	tikvAsyncCommit = "tikv.async_commit"
//	tikvOnePC       = "tikv.one_pc"
//)

const (
	HOST = "127.0.0.1"
	PORT = "9090"
)

type txnConfig struct {
	asyncCommit bool
	onePC       bool
}

type txnDB struct {
	db      *THBaseServiceClient
	r       *util.RowCodec
	bufPool *util.BufPool
	//needed by HBase
	transport       *thrift.TSocket
	protocolFactory *thrift.TProtocolFactory
}

func createTxnDB(p *properties.Properties) (ycsb.DB, error) {
	bufPool := util.NewBufPool()

	return &txnDB{
		r:       util.NewRowCodec(p),
		bufPool: bufPool,
	}, nil
}

func (db *txnDB) Close() error {
	return db.transport.Close()
}

func (db *txnDB) InitThread(ctx context.Context, _ int, _ int) context.Context {
	return ctx
}

func (db *txnDB) CleanupThread(ctx context.Context) {
}

func (db *txnDB) getRowKey(table string, key string) []byte {
	return util.Slice(fmt.Sprintf("%s:%s", table, key))
}

// seems that HBase don't need this function
func (db *txnDB) beginTxn() (*transaction.KVTxn, error) {
	return nil, nil
}

func (db *txnDB) Read(ctx context.Context, table string, key string, fields []string) (map[string][]byte, error) {
	tx := db.db
	row, err := tx.Get(ctx, []byte(table), &TGet{Row: []byte(key)})

	if err != nil {
		return nil, err
	}

	columnValues := row.ColumnValues
	res := make(map[string][]byte)
	for _, column := range columnValues {
		c := reflect.ValueOf(column).Elem()
		family := c.Field(0)
		value := c.Field(2)
		res[string(family.Interface().([]uint8))] = value.Interface().([]byte)
	}

	return res, nil
}

func (db *txnDB) BatchRead(ctx context.Context, table string, keys []string, fields []string) ([]map[string][]byte, error) {
	tx := db.db

	var tempTGet []*TGet

	rowValues := make([]map[string][]byte, len(keys))
	keyLoc := make(map[string]int)
	for i, key := range keys {
		tempTGet = append(tempTGet, &TGet{
			Row: []byte(key),
		})
		keyLoc[key] = i
	}

	res, err := tx.GetMultiple(ctx, []byte(table), tempTGet)

	if err != nil {
		return nil, err
	}

	for _, columnValues := range res {
		i := keyLoc[string(columnValues.Row)]
		column := columnValues.ColumnValues
		c := reflect.ValueOf(column).Elem()
		family := c.Field(0)
		value := c.Field(2)
		rowValues[i][string(family.Interface().([]uint8))] = value.Interface().([]byte)
	}
	return rowValues, nil
}

// TODO I don't know how to use scanner in HBase
func (db *txnDB) Scan(ctx context.Context, table string, startKey string, count int, fields []string) ([]map[string][]byte, error) {
	panic("calling error, should use func TxnCommit")

	//return nil, nil
}

// in HBase, 'update' is equal to 'put', so only implement Insert() is ok
func (db *txnDB) Update(ctx context.Context, table string, key string, values map[string][]byte) error {
	panic("calling error, should use func TxnCommit")
}

func (db *txnDB) BatchUpdate(ctx context.Context, table string, keys []string, values []map[string][]byte) error {
	panic("calling error, should use func TxnCommit")
	return nil
}

func (db *txnDB) Insert(ctx context.Context, table string, key string, values map[string][]byte) error {

	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(&thrift.TConfiguration{
		TBinaryStrictRead:  thrift.BoolPtr(true),
		TBinaryStrictWrite: thrift.BoolPtr(true),
	})
	transport := thrift.NewTSocketConf(net.JoinHostPort(HOST, PORT), &thrift.TConfiguration{
		ConnectTimeout: time.Second * 5,
		SocketTimeout:  time.Second * 5,
	})
	client := NewTHBaseServiceClientFactory(transport, protocolFactory)
	err := transport.Open()
	if err != nil {
		return err
	}
	defer transport.Close()

	//txnId := atomic.AddUint64(&taas.CSNCounter, 1)
	var cvarr []*TColumnValue
	finalData, err1 := db.r.Encode(nil, values)
	if err1 != nil {
		return err1
	}
	cvarr = append(cvarr, &TColumnValue{
		Family:    []byte("entire"),
		Qualifier: []byte(""),
		Value:     []byte(string(finalData)),
	})

	rowKey := db.getRowKey(table, key)
	tempTPut := TPut{Row: []byte(rowKey), ColumnValues: cvarr}
	err = client.Put(ctx, []byte(table), &tempTPut)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func (db *txnDB) BatchInsert(ctx context.Context, table string, keys []string, values []map[string][]byte) error {
	panic("not finished yet")
}

func (db *txnDB) Delete(ctx context.Context, table string, key string) error {

	client := db.db

	tdelete := TDelete{Row: []byte(key)}
	err := client.DeleteSingle(ctx, []byte(table), &tdelete)
	return err
}

func (db *txnDB) BatchDelete(ctx context.Context, table string, keys []string) error {

	client := db.db

	var tDeletes []*TDelete

	for _, key := range keys {
		tDeletes = append(tDeletes, &TDelete{
			Row: []byte(key),
		})
	}

	_, err := client.DeleteMultiple(ctx, []byte(table), tDeletes)
	return err
}
