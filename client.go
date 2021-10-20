package main

import (
    "context"
    "fmt"
    "log"
    "math"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    client, err := ethclient.Dial("http://localhost:8545")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("we have a connection")

    account := common.HexToAddress("0xf59867553F6cDD694158205eA047ffe61f517B97")
    balance, err := client.BalanceAt(context.Background(), account, nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(balance)

    fbalance := new(big.Float)
    fbalance.SetString(balance.String())
    ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
    fmt.Println(ethValue) 

    pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(pendingBalance)
}