package taas

import (
	"fmt"
	"github.com/magiconair/properties"
	zmq "github.com/pebbe/zmq4"
	"github.com/spf13/viper"
	"log"
)

//#include ""
import (
	"github.com/golang/protobuf/proto"

	"github.com/pingcap/go-ycsb/db/taas_proto"
)

var TaaSServerNum = 1
var TaasServerIps []string

var LocalServerIp = "127.0.0.1"
var TaasServerIp = "127.0.0.1"
var StorageServerIp = "127.0.0.1"
var HbaseServerIp = "127.0.0.1"
var LevelDBServerIP = "127.0.0.1"

var OpNum = 10
var ClientNum = 64
var UnPackNum = 16

type TaasTxn struct {
	TaaSID            uint64
	GzipedTransaction []byte
}

var TaasTxnCH = make(chan TaasTxn, 100000)
var UnPackCH = make(chan string, 100000)

var ChanList []chan string
var InitOk uint64 = 0
var CSNCounter uint64 = 0
var SuccessTransactionCounter, FailedTransactionCounter, TotalTransactionCounter uint64 = 0, 0, 0
var SuccessReadCounter, FailedReadCounter, TotalReadCounter uint64 = 0, 0, 0
var SuccessUpdateCounter, FailedUpdateounter, TotalUpdateCounter uint64 = 0, 0, 0
var TotalLatency, TikvReadLatency, TikvTotalLatency, TotalSuccessLatency, TotalFailedLatency uint64 = 0, 0, 0, 0, 0
var latency []uint64

func SetConfig(globalProps *properties.Properties) {
	TaasServerIp = globalProps.GetString("taasServerIp", "127.0.0.1")
	LocalServerIp = globalProps.GetString("localServerIp", "127.0.0.1")
	StorageServerIp = globalProps.GetString("storageServerIp", "127.0.0.1")
	HbaseServerIp = globalProps.GetString("hbaseServerIp", "127.0.0.1")
	LevelDBServerIP = globalProps.GetString("levelDBServerIp", "127.0.0.1")

	OpNum = globalProps.GetInt("opNum", 1)
	ClientNum = globalProps.GetInt("threadcount", 64)
	UnPackNum = globalProps.GetInt("unpackNum", 16)

	TaaSConfig := viper.New()
	//TaaSConfig.SetConfigFile("workloads/TaaS_config.yml")
	TaaSConfig.AddConfigPath(".")
	TaaSConfig.AddConfigPath("workloads/")
	TaaSConfig.SetConfigName("TaaS_config")
	TaaSConfig.SetConfigType("yaml")

	if err := TaaSConfig.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(err, "找不到配置文件..")
		} else {
			fmt.Println(err, "配置文件出错..")
		}
	}

	TaaSServerNum = TaaSConfig.GetInt("taas.servernum")
	TaasServerIps = TaaSConfig.GetStringSlice("taas.serverips")
	fmt.Println(TaasServerIps)
	//for k, v := range TaasServerIps {
	//	TaasServerIps[k] = v
	//}
	LocalServerIp = TaaSConfig.GetString("client.localserverip")
	StorageServerIp = TaaSConfig.GetString("client.storageserverip")
	HbaseServerIp = TaaSConfig.GetString("client.hbaseserverip")
	LevelDBServerIP = TaaSConfig.GetString("client.leveldbserverip")

	OpNum = TaaSConfig.GetInt("Client.OpNum")
	//ClientNum = TaaSConfig.GetInt("Client.ClientNum")
	UnPackNum = TaaSConfig.GetInt("Client.UnPackNum")

	allSettings := TaaSConfig.AllSettings()
	fmt.Println(allSettings)

	fmt.Println("localServerIp : " + LocalServerIp + ", hbaseServerIp " + HbaseServerIp + ", levelDBServerIp " + LevelDBServerIP)

}

func SendTxnToTaas() {
	var sockets []*zmq.Socket
	for i := 0; i < TaaSServerNum; i++ {
		socket, _ := zmq.NewSocket(zmq.PUSH)
		err := socket.SetSndbuf(1000000000)
		if err != nil {
			return
		}
		err = socket.SetRcvbuf(1000000000)
		if err != nil {
			return
		}
		err = socket.SetSndhwm(1000000000)
		if err != nil {
			return
		}
		err = socket.SetRcvhwm(1000000000)
		if err != nil {
			return
		}
		err = socket.Connect("tcp://" + TaasServerIps[i] + ":5551")
		if err != nil {
			fmt.Println("taas.go 97")
			log.Fatal(err)
		}
		fmt.Println("Taas Send " + TaasServerIps[i])
		sockets = append(sockets, socket)
	}

	for {
		value, ok := <-TaasTxnCH
		if ok {
			_, err := sockets[value.TaaSID%uint64(TaaSServerNum)].Send(string(value.GzipedTransaction), 0)
			//fmt.Println("taas send thread")
			if err != nil {
				return
			}
		} else {
			fmt.Println("taas.go 109")
			log.Fatal(ok)
		}
	}
}

func ListenFromTaas() {
	socket, err := zmq.NewSocket(zmq.PULL)
	err = socket.SetSndbuf(1000000000)
	if err != nil {
		return
	}
	err = socket.SetRcvbuf(1000000000)
	if err != nil {
		return
	}
	err = socket.SetSndhwm(1000000000)
	if err != nil {
		return
	}
	err = socket.SetRcvhwm(1000000000)
	if err != nil {
		return
	}
	err = socket.Bind("tcp://*:5552")
	fmt.Println("连接Taas Listen")
	if err != nil {
		log.Fatal(err)
	}
	for {
		taasReply, err := socket.Recv(0)
		if err != nil {
			fmt.Println("taas.go 115")
			log.Fatal(err)
		}
		UnPackCH <- taasReply
	}
}

func UGZipBytes(in []byte) ([]byte, error) {
	//reader, err := gzip.NewReader(bytes.NewReader(in))
	//if err != nil {
	//	return nil, err
	//}
	//defer reader.Close()
	//out, _ := ioutil.ReadAll(reader)
	//return out, nil
	return in, nil
}

func GzipBytes(in []byte) ([]byte, error) {
	//var bufferBeforeGzip bytes.Buffer
	//bufferBeforeGzip.Reset()
	//gw := gzip.NewWriter(&bufferBeforeGzip)
	//_, err := gw.Write(in)
	//if err != nil {
	//	return nil, err
	//}
	//err = gw.Close()
	//if err != nil {
	//	return nil, err
	//}
	//return bufferBeforeGzip.Bytes(), nil
	return in, nil
}

func UnPack() {
	for {
		taasReply, ok := <-UnPackCH
		if ok {
			UnGZipedReply, err := UGZipBytes([]byte(taasReply))
			if err != nil {
				panic("UnGzip error")
			}
			testMessage := &taas_proto.Message{}
			err = proto.Unmarshal(UnGZipedReply, testMessage)
			if err != nil {
				fmt.Println("taas.go 142")
				log.Fatal(err)
			}
			replyMessage := testMessage.GetReplyTxnResultToClient()
			ChanList[replyMessage.ClientTxnId%uint64(ClientNum)] <- replyMessage.GetTxnState().String()
		} else {
			fmt.Println("taas.go 148")
			log.Fatal(ok)
		}
	}
}
