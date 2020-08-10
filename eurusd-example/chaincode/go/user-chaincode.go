package main

import (
	"fmt"

	oraclizeapi "https://github.com/provable-things/fabric-api"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract defines the Smart Contract structure
type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) error {
	return nil
}


func (s *SmartContract) FetchEURUSDviaOraclize(ctx contractapi.TransactionContextInterface) error {
	fmt.Println("============= START : Calling the oraclize chaincode =============")
	var datasource = "URL"                                                                  // Setting the Oraclize datasource
	var query = "json(https://min-api.cryptocompare.com/data/price?fsym=EUR&tsyms=USD).USD" // Setting the query
	result, proof := oraclizeapi.OraclizeQuery_sync(ctx.GetStub(), datasource, query, oraclizeapi.TLSNOTARY)
	fmt.Printf("proof: %s", proof)
	fmt.Printf("\nresult: %s\n", result)
	fmt.Println("Do something with the result...")
	fmt.Println("============= END : Calling the oraclize chaincode =============")
	
	return nil
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error creating chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting chaincode: %s", err.Error())
	}
}