package main

import (
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    client, err := ethclient.Dial("http://localhost:8545")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("we have a connection")
    fmt.Println("******************************************")
    fmt.Println("Checking Account Balance")
    balance(client)
    fmt.Println("******************************************")
    fmt.Println("Setting Up Wallet")
    wallet()
    fmt.Println("******************************************")
    fmt.Println("Checking if Address is a Smart Contract Type or User Type")
    checkAddress(client)
    fmt.Println("******************************************")
    fmt.Println("Querying Block")
    queryBlock(client)
    fmt.Println("******************************************")
    fmt.Println("Querying Transaction")
    queryTx(client)
    fmt.Println("******************************************")
    fmt.Println("Transferring Ethereum")
    transferETH(client)
    fmt.Println("******************************************")
    fmt.Println("Transferring Token")
    transferToken(client)
    fmt.Println("******************************************")
}