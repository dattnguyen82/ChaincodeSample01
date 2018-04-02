package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type DocumentTracker struct {
}

type fn func() ([]byte, error)

//Init the chaincode asigned the value "0" to the counter in the state.
func (t *DocumentTracker) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	err := stub.PutState("zoneID", []byte(args[0]))
	if err != nil {
		return nil, err
	}
	return nil, err
}

func (t *DocumentTracker) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
 	if function == "store" {
		f := func() ([]byte, error){
			err := stub.PutState(args[0], []byte(args[1]))
			return nil, err
		}
		return t.verifyAndExecute(stub, f)
	}
	return nil, errors.New("Invalid invoke function")
}

func (t *DocumentTracker) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "read" {
		var err error
		// Get the state from the ledger
		dataBytes, err := stub.GetState(args[0])
		
		if err != nil {
			jsonResp := "{\"Error\":\"Invalid data query \"}"
			return nil, errors.New(jsonResp)
		}

		if dataBytes == nil {
			jsonResp := "{\"Error\":\"No data stored\"}"
			return nil, errors.New(jsonResp)
		}

		//jsonResp := "{\"Name\":\"store\",\"Amount\":\"" + string(dataBytes) + "\"}"
		//fmt.Printf("Query Response:%s\n", jsonResp)
		return dataBytes, nil
	}
	return nil, errors.New("Invalid query function")
}

//Attribute reader from Predix chaincode example
func (t *DocumentTracker) getAttributeFromCertificate(stub shim.ChaincodeStubInterface) (string, []byte, error) {
	attributeName := "zone"
	val, err := stub.ReadCertAttribute(attributeName)
	fmt.Printf("%s => %v error %v \n", attributeName, string(val), err)
	if err != nil {
		return "", nil, errors.New("Error reading cert attribute:" + err.Error())
	}
	return attributeName, val, nil
}

//ACS function from Predix chaincode example
func (t *DocumentTracker) verifyAndExecute(stub shim.ChaincodeStubInterface, f fn) ([]byte, error) {
	zoneID, err := stub.GetState("zoneID")
	if err != nil {
		return nil, err
	}
	isOK, _ := stub.VerifyAttribute("zone", zoneID)
	if isOK {
		return f()
	}
	return nil, nil
}

func main() {
	err := shim.Start(new(DocumentTracker))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
