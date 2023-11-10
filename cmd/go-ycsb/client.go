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

// Modifyed by singheart
package main

import (
	"fmt"
	//"github.com/golang/protobuf/proto"
	//zmq "github.com/pebbe/zmq4"
	//"github.com/pingcap/go-ycsb/db/taas_proto"
	//"github.com/pingcap/go-ycsb/db/taas_tikv"
	// "github.com/pingcap/go-ycsb/pkg/workload"
	//"io/ioutil"
	//"log"
	"github.com/pingcap/go-ycsb/db/taas"
	"github.com/pingcap/go-ycsb/pkg/workload"
	"strconv"
	"time"

	_ "github.com/pingcap/go-ycsb/db/taas_tikv"
	"github.com/pingcap/go-ycsb/pkg/client"
	"github.com/pingcap/go-ycsb/pkg/measurement"
	"github.com/pingcap/go-ycsb/pkg/prop"
	"github.com/spf13/cobra"
)

func runClientCommandFunc(cmd *cobra.Command, args []string, doTransactions bool, command string) {
	dbName := args[0]

	initialGlobal(dbName, func() {
		fmt.Println("****************** Connect To Taas ***************************")
		taas.SetConfig(globalProps)
		for i := 0; i < taas.ClientNum; i++ {
			taas.ChanList = append(taas.ChanList, make(chan string, 100000))
		}
		go taas.SendTxnToTaas()
		go taas.ListenFromTaas()
		for i := 0; i < taas.UnPackNum; i++ {
			go taas.UnPack()
		}
		taas.InitOk = 1
		fmt.Println("****************** Connect To Taas Finished ******************")

		doTransFlag := "true"
		if !doTransactions {
			doTransFlag = "false"
		}
		globalProps.Set(prop.DoTransactions, doTransFlag)
		globalProps.Set(prop.Command, command)

		if cmd.Flags().Changed("threads") {
			// We set the threadArg via command line.
			globalProps.Set(prop.ThreadCount, strconv.Itoa(threadsArg))
		}

		if cmd.Flags().Changed("target") {
			globalProps.Set(prop.Target, strconv.Itoa(targetArg))
		}

		if cmd.Flags().Changed("interval") {
			globalProps.Set(prop.LogInterval, strconv.Itoa(reportInterval))
		}
	})

	fmt.Println("***************** properties *****************")
	for key, value := range globalProps.Map() {
		fmt.Printf("\"%s\"=\"%s\"\n", key, value)
	}
	fmt.Println("**********************************************")

	c := client.NewClient(globalProps, globalWorkload, globalDB)
	start := time.Now()
	c.Run(globalContext)
	fmt.Println("**********************************************")
	fmt.Printf("Run finished, takes %s\n", time.Now().Sub(start))
	measurement.Output()
	fmt.Printf("[Read] TotalOpNum %d, SuccessOpNum %d FailedOpNum %d SuccessRate %f\n",
		taas.TotalReadCounter, taas.SuccessReadCounter, taas.FailedReadCounter,
		float64(taas.SuccessReadCounter)/float64(taas.TotalReadCounter))
	fmt.Printf("[Update] TotalOpNum %d, SuccessOpNum %d FailedOpNum %d SuccessRate %f\n",
		taas.TotalUpdateCounter, taas.SuccessUpdateCounter, taas.FailedUpdateounter,
		float64(taas.SuccessUpdateCounter)/float64(taas.TotalUpdateCounter))
	fmt.Printf("[Transaction] TotalOpNum %d, SuccessOpNum %d FailedOpNum %d SuccessRate %f\n",
		taas.TotalTransactionCounter, taas.SuccessTransactionCounter, taas.FailedTransactionCounter,
		float64(taas.SuccessTransactionCounter)/float64(taas.TotalTransactionCounter))
	fmt.Printf("[Op] ReadOpNum %d, UpdateOpNum %d UpdateRate %f\n",
		workload.TotalReadCounter, workload.TotalUpdateCounter,
		float64(workload.TotalUpdateCounter)/float64(workload.TotalReadCounter+workload.TotalUpdateCounter))
	fmt.Printf("[Op] TikvTotalTime %d, TikvReadTime %d\n, ",
		taas.TikvTotalLatency, taas.TikvReadLatency)
	fmt.Printf("[Txn] TikvAvgTime %f\n",
		float64(taas.TikvTotalLatency)/float64(taas.TotalTransactionCounter))
	fmt.Printf("[Txn] TotalFailedLatency %d, TotalSuccessLatency %d, AvgFailedLatency %f, AvgSuccessLatency %f\n",
		taas.TotalFailedLatency, taas.TotalSuccessLatency, float64(taas.TotalFailedLatency)/float64(taas.FailedTransactionCounter), float64(taas.TotalSuccessLatency)/float64(taas.SuccessTransactionCounter))
}

func runLoadCommandFunc(cmd *cobra.Command, args []string) {
	runClientCommandFunc(cmd, args, false, "load")
}

func runTransCommandFunc(cmd *cobra.Command, args []string) {
	runClientCommandFunc(cmd, args, true, "run")
}

var (
	threadsArg     int
	targetArg      int
	reportInterval int
)

func initClientCommand(m *cobra.Command) {
	m.Flags().StringSliceVarP(&propertyFiles, "property_file", "P", nil, "Spefify a property file")
	m.Flags().StringArrayVarP(&propertyValues, "prop", "p", nil, "Specify a property value with name=value")
	m.Flags().StringVar(&tableName, "table", "", "Use the table name instead of the default \""+prop.TableNameDefault+"\"")
	m.Flags().IntVar(&threadsArg, "threads", 1, "Execute using n threads - can also be specified as the \"threadcount\" property")
	m.Flags().IntVar(&targetArg, "target", 0, "Attempt to do n operations per second (default: unlimited) - can also be specified as the \"target\" property")
	m.Flags().IntVar(&reportInterval, "interval", 10, "Interval of outputting measurements in seconds")
}

func newLoadCommand() *cobra.Command {
	m := &cobra.Command{
		Use:   "load db",
		Short: "YCSB load benchmark",
		Args:  cobra.MinimumNArgs(1),
		Run:   runLoadCommandFunc,
	}

	initClientCommand(m)
	return m
}

func newRunCommand() *cobra.Command {
	m := &cobra.Command{
		Use:   "run db",
		Short: "YCSB run benchmark",
		Args:  cobra.MinimumNArgs(1),
		Run:   runTransCommandFunc,
	}

	initClientCommand(m)
	return m
}
