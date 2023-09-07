package leveldb_txn

import (
	"fmt"

	"github.com/magiconair/properties"
	"github.com/pingcap/go-ycsb/pkg/ycsb"
)

const (
	kvType = "tikv.type"
)

type leveldbTxnCreator struct {
}

func (c leveldbTxnCreator) Create(p *properties.Properties) (ycsb.DB, error) {

	tp := p.GetString(kvType, "txn")
	fmt.Println("=====================  leveldb Txn ============================")
	switch tp {
	case "raw":
		panic("not implement yet")
		return nil, nil
		///return createRawDB(p)
	case "txn":
		return createTxnDB(p)
	default:
		return nil, fmt.Errorf("unsupported type %s", tp)
	}
}

func init() {
	ycsb.RegisterDBCreator("leveldb_txn", leveldbTxnCreator{})
}
