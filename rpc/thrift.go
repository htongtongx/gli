package rpc

import (
	"fmt"
	"log"
	"os"

	"github.com/apache/thrift/lib/go/thrift"
)

func Serve(listenAddr string, processor thrift.TProcessor) {

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

	transport, err := thrift.NewTServerSocket(listenAddr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// processor := vpnserver.NewVpnServerProcessor(NewVPN(d, ad))
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	// log.Println("Starting the simple server... on ", listenAddr)
	err = server.Serve()
	if err != nil {
		log.Println(listenAddr + "监听失败:" + err.Error())
		return
	}
	fmt.Println("Sucessful the simple server... on ", listenAddr)
}

func GetClient(hostPort string) (tc thrift.TClient, ts *thrift.TSocket, err error) {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport, err := thrift.NewTSocket(hostPort)
	// defer transport.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, nil, err
	}
	useTransport, err := transportFactory.GetTransport(transport)
	if err != nil {
		log.Println(err.Error())
		return nil, nil, err
	}
	// client := vpnserver.NewVpnServerClientFactory(useTransport, protocolFactory)
	err = transport.Open()
	if err != nil {
		log.Println(os.Stderr, "Error opening socket to "+hostPort, " ", err)
		return nil, nil, err
	}
	return thrift.NewTStandardClient(protocolFactory.GetProtocol(useTransport), protocolFactory.GetProtocol(useTransport)), transport, nil
}
