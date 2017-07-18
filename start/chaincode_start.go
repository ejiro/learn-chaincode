/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
    // Get the args from the transaction proposal
    args := stub.GetStringArgs()
    
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	return shim.Success(nil)
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    // Extract the function and args from the transaction proposal
    function, args := stub.GetFunctionAndParameters()
    
	fmt.Println("invoke is running " + function)
	
	if len(args) > 0{
        fmt.Println("args is running " + args[0])
	}

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub)
	}
	fmt.Println("invoke did not find func: " + function)					//error

	return shim.Error("Received Unknown Smart Contract function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface) peer.Response {
    // Extract the function and args from the transaction proposal
    function, args := stub.GetFunctionAndParameters()
    
	fmt.Println("query is running " + function)
	
    if len(args) > 0{
        fmt.Println("args is running " + args[0])
	}


	// Handle different functions
	if function == "dummy_query" {											//read a variable
		fmt.Println("hi there " + function)						//error
		return shim.Success(nil)
	}
	fmt.Println("query did not find func: " + function)						//error

	return shim.Error("Received unknown function query: " + function)
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
