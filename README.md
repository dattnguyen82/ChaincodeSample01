# Sample Chaincode for Predix Blockchain Service
This is a sample smart contract (chaincode) for the [Predix Blockchain Service](https://docs.predix.io/en-US/content/service/security/blockchain_as_a_service/)


## Prerequesites

* [GO Lang](https://golang.org/)
* [HyperLedger](http://hyperledger-fabric.readthedocs.io/en/release-1.0/getting_started.html#install-prerequisites) - Optional

## Build & Deploy

### Clone this repo using Go:

	go get github.com/dattnguyen82/ChaincodeSample01

### Compile source
In your GO path you should have a directory called:  **$GOPATH/src/github.com/dattnguyen82/predix-sample-chaincode**
Run this command:

	go build predix-blockchain-sample01.go

### Deploy
To deploy your smart contract, compress the predix-sample-chaincode directory contents with tar.gz

	cd $GOPATH/src/github.com/dattnguyen82/predix-sample-chaincode
	tar czvf predix-sample-chaincode.tgz *

Deploy to Predix using this [guide](https://docs.predix.io/en-US/content/service/security/blockchain_as_a_service/using-blockchain-as-a-service#task_77c2dbb7-4ba7-4a1a-9628-359c3fd0800e)

    curl -X PUT \
    https://{predix-url}/v1/chaincodes/{smart-contract-name} \
    -H 'authorization: bearer {token}' \
    -H 'content-type: multipart/form-data;' \
    -H 'predix-zone-id: {predix-zone-id}' \
    -F chaincode=@{path} \
    -F 'args=["{predix-zone-id}"]'

## Postman Collection

A Postman collection and environment have been added for convenience.

### Environment parameters

* uaa-url - UAA base url

* uaa-auth - UAA token retrieval authorization

* token - Access token from UAA

* api-url - Base URL for predix blockchain service e.g. *predix-blockchain-api.run.aws-usw02-pr.ice.predix.io*

* tenant - Predix-Zone-Id from VCAP

* user-id - new user Id

* default-user - username from VCAP

* default-password - onetimepassword from VCAP

* chain-code - name of chaincode to deploy

* tx-id - transaction id from chaincode invocation

* deploy-arguments - chaincode deployment argument e.g. *["arg1", "arg2", "arg3" ...]*,
