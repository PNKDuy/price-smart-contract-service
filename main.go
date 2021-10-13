package main

import (
	"liquidity_pool_service/connector"
	"log"
)

func main() {
	//connector.ConnectToClickHouse()
	//_, client, err := model.ConnectToSmartContract(model.UniswapFactoryContractAddress,models.InfuraSecretId)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//contractAddress := common.HexToAddress(model.UniswapFactoryContractAddress)

	//model.SubscribingToCreatedPairEvent(contractAddress, client)
	//subscribing event log
	//query := ethereum.FilterQuery{
	//	Addresses: []common.Address{contractAddress},
	//}
	//
	//logs := make(chan types.Log)
	//sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//contractAbi, err := abi.JSON(strings.NewReader(factory.MainABI))
	//
	//for {
	//	select {
	//	case err := <-sub.Err():
	//		log.Fatal(err)
	//	case vLog := <-logs:
	//		evt, err := contractAbi.Unpack("PairCreated", vLog.Data)
	//		fmt.Println(evt[1])
	//		//err := json.Unmarshal(vLog.Data, &event)
	//		if err != nil {
	//			log.Println("Cannot unmarshal")
	//			return
	//		}
	//	}
	//}
	////subscribing event log



	//tokenPairList, err := models.GetAllPairs(factoryMain, client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for i, _ := range tokenPairList {
	//	fmt.Println(tokenPairList[i].Token0Symbol + tokenPairList[i].Token0Symbol)
	//}
	db, err := connector.ConnectToClickHouse()
	if err != nil {
		log.Println(err)
	}
	println(db.Name())

}


