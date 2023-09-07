// Copyright 2018 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package taas_hbase

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/magiconair/properties"
	"github.com/pingcap/go-ycsb/db/taas"
	"github.com/pingcap/go-ycsb/pkg/ycsb"
	"net"
	"time"
)

//const (
//	tikvPD = "tikv.pd"
//	// raw, txn, or coprocessor
//	tikvType       = "tikv.type"
//	tikvConnCount  = "tikv.conncount"
//	tikvBatchSize  = "tikv.batchsize"
//	tikvAPIVersion = "tikv.apiversion"
//)

type taasHbaseCreator struct {
}

func (c taasHbaseCreator) Create(p *properties.Properties) (ycsb.DB, error) {
	fmt.Println("=====================  Taas - HBase  ============================")

	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(&thrift.TConfiguration{
		TBinaryStrictRead:  thrift.BoolPtr(true),
		TBinaryStrictWrite: thrift.BoolPtr(true),
	})
	transport := thrift.NewTSocketConf(net.JoinHostPort(taas.HbaseServerIp, PORT), &thrift.TConfiguration{
		ConnectTimeout: time.Second * 5,
		SocketTimeout:  time.Second * 5,
	})
	NewTHBaseServiceClientFactory(transport, protocolFactory)
	err := transport.Open()
	if err != nil {
		panic("connect to hbase error")
	}
	defer transport.Close()

	return createTxnDB(p)
}

func init() {
	ycsb.RegisterDBCreator("taas_hbase", taasHbaseCreator{})
}
