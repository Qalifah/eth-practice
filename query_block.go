package main

import (
    "context"
    "fmt"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/ethclient"
)

func queryBlock(client *ethclient.Client) {
    header, err := client.HeaderByNumber(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(header.Number.String())

    blockNumber := big.NewInt(0)
    block, err := client.BlockByNumber(context.Background(), blockNumber)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(block.Number().Uint64())     
    fmt.Println(block.Time())       
    fmt.Println(block.Difficulty().Uint64()) 
    fmt.Println(block.Hash().Hex())        
    fmt.Println(len(block.Transactions())) 

    count, err := client.TransactionCount(context.Background(), block.Hash())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(count) 
}
