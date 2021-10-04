package models

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"liquidity_pool_service/contracts/ERC20"
	"liquidity_pool_service/contracts/factory"
	"liquidity_pool_service/contracts/pair"
	"log"
	"math/big"
	"strconv"
	"strings"
)
var callOpt *bind.CallOpts

const (
	InfuraSecretId                = "wss://mainnet.infura.io/ws/v3/2dd74ccabc4f45a9892b85af20187db6"
	UniswapFactoryContractAddress = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
)

type TokenPair struct {
	PairAddress string
	Token0Symbol string
	Token1Symbol string
	ReservePairs []ReservePair
}

type ReservePair struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}

type CreatedPairEvent struct {
	PairAddress common.Address
	AllPairsLength uint
}

func ConnectToSmartContract(address, clientSecretId string) (factoryMain *factory.Main,client *ethclient.Client, err error) {
	client, err = ethclient.Dial(clientSecretId)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	factoryMain, err = factory.NewMain(common.HexToAddress(address), client)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	return factoryMain,client, nil
}

func GetAllPairs(factoryMain *factory.Main,client *ethclient.Client) (tokenPairList []TokenPair, err error){
	//get all pairs length
	length, err := factoryMain.AllPairsLength(callOpt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	lengthInt, _ := strconv.Atoi(length.String())
	for i := 0; i < lengthInt; i++ {
		tokenPair, err := GetPairByLength(factoryMain, client, int64(i))
		if err != nil {
			return nil, err
		}
		tokenPairList = append(tokenPairList, tokenPair)
	}

	return tokenPairList, nil
}

func GetPairByLength(factoryMain *factory.Main, client *ethclient.Client, allPairsLength int64) (tokenPair TokenPair, err error) {
	var reserverPairs []ReservePair
	pairAddress, _ := factoryMain.AllPairs(callOpt, big.NewInt(allPairsLength))
	tokenPair.PairAddress = pairAddress.String()
	pairMain, err := pair.NewMain(pairAddress, client)
	if err != nil {
		return tokenPair, err
	}

	token0Symbol, token1Symbol, err := GetToken0Token1Symbols(pairMain, client)
	if err != nil {
		log.Println(err)
		return tokenPair, err
	}

	tokenPair.Token0Symbol = token0Symbol
	tokenPair.Token1Symbol = token1Symbol
	reserves, err := GetPairReserves(pairMain)
	if err != nil {
		return tokenPair, err
	}

	reserverPairs = append(reserverPairs, reserves)
	tokenPair.ReservePairs = reserverPairs

	return tokenPair, nil
}

func GetToken0Token1Symbols(pairMain *pair.Main, client *ethclient.Client) (string,string,error) {
	token0Address, _ := pairMain.Token0(callOpt)
	token1Address, _ := pairMain.Token1(callOpt)
	token0Main, _ := ERC20.NewMain(token0Address, client)
	token1Main, _ := ERC20.NewMain(token1Address, client)
	token0symbol, _ := token0Main.Symbol(callOpt)
	token1symbol, _ := token1Main.Symbol(callOpt)
	return token0symbol, token1symbol, nil
}

func GetPairReserves(pairMain *pair.Main) (reserves ReservePair, err error) {
	reserveJson, err := pairMain.GetReserves(callOpt)
	if err != nil {
		log.Println(err)
		return reserveJson, err
	}
	return reserveJson, err
}

func SubcribingToCreatedPairEvent(contractAddress common.Address,client *ethclient.Client) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	var event []interface{}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(factory.MainABI))
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			err := contractAbi.UnpackIntoInterface(&event,"PairCreated", vLog.Data)
			fmt.Println(event)
			if err != nil {
				log.Println("Cannot unmarshal")
			}
		}
	}
}


