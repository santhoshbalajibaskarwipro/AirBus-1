/*/*
-a-Licensed to the Apache Software Foundation (ASF) under one
or more contributor license Forms.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
"errors"
"fmt"
"strconv"
"encoding/json"

"github.com/hyperledger/fabric/core/chaincode/shim"
"github.com/hyperledger/fabric/core/util"
)

// Proposal example simple Chaincode implementation
type manage_proposal struct {
}

var approved_proposal_entry = "approved_proposal_entry"				//name for the key/value that will store a list of all known  Tier3 Form

type proposal struct{
								// Attributes of a Form 
	proposal_id string `json:"proposal_id"`	
	region string `json:"region"`
	country string `json:"country"`
	
}
// ============================================================================================================================
// Main - start the chaincode for Form management
// ============================================================================================================================
func main() {			
	err := shim.Start(new(manage_proposal))
	if err != nil {
		fmt.Printf("Error starting Form management chaincode: %s", err)
	}
}
// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *manage_proposal) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var msg string
	var err error
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	// Initialize the chaincode
	msg = args[0]
	fmt.Println("manage_proposal chaincode is deployed successfully.");
	
	// Write the state to the ledger
	err = stub.PutState("abc", []byte(msg))	//making a test var "abc", I find it handy to read/write to it right away to test the network
	if err != nil {
		return nil, err
	}
	var proposal_form_empty []string
	proposal_form_empty_json_as_bytes, _ := json.Marshal(proposal_form_empty)								//marshal an emtpy array of strings to clear the index
	err = stub.PutState(approved_proposal_entry, proposal_form_empty_json_as_bytes)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
// ============================================================================================================================
// Run - Our entry Formint for Invocations - [LEGACY] obc-peer 4/25/2016
// ============================================================================================================================
func (t *manage_proposal) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("run is running " + function)
	return t.Invoke(stub, function, args)
}
// ============================================================================================================================
// Invoke - Our entry Formint for Invocations
// ============================================================================================================================
func (t *manage_proposal) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "create_proposal_id" {											//create a new Form
		return t.create_proposal_id(stub, args)
	} 
	fmt.Println("invoke did not find func: " + function)	
	jsonResp := "Error : Received unknown function invocation: "+ function 				//error
	return nil, errors.New(jsonResp)
}
// ============================================================================================================================
// create Form - create a new Form for Tier-3, store into chaincode state
// ============================================================================================================================
func (t *manage_proposal) createForm_Tier3(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	if len(args) != 10 {
		return nil, errors.New("Incorrect number of arguments. Expecting 9")
	}
	fmt.Println("Creating a new Form for Tier-3")
	if len(args[0]) <= 0 {
		return nil, errors.New("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return nil, errors.New("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return nil, errors.New("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return nil, errors.New("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return nil, errors.New("5th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return nil, errors.New("6th argument must be a non-empty string")
	}
	if len(args[6]) <= 0 {
		return nil, errors.New("7th argument must be a non-empty string")
	}
	if len(args[7]) <= 0 {
		return nil, errors.New("8th argument must be a non-empty string")
	}
	if len(args[8]) <= 0 {
		return nil, errors.New("9th argument must be a non-empty string")
	}
	if len(args[9]) <= 0 {
		return nil, errors.New("10th argument must be a non-empty string")
	}
	
	FAA_formNumber := args[0] // FAA_formNumber or FAA_formNumberber
	quantity := args[1]
	FAA_formUrl := args[2]
	fileHash:=args[3]
	user := args[4]
	itemType := args[5]
	part_number := args[6]
	total_approvedQty := args[7]
	approvalDate	:= args[8]
	authorization_number := args[9]
	userType := "Tier-3"	
	qty,err := strconv.Atoi(quantity)
	if err != nil {
		return nil, errors.New("Error while converting string 'quantity' to int ")
	}
	approvedQty,err := strconv.Atoi(total_approvedQty)
	if err != nil {
		return nil, errors.New("{\"Error\":\"Error while converting string 'approvedQty' to int \"}")
	}
	if(qty > approvedQty){
		jsonResp := "Error: Quantity should be less than Total Approved Quantity"
		return nil,errors.New(jsonResp)
	}	
		
	//build the Form json string manually
	input := 	`{`+
		`"FAA_formNumber": "` + FAA_formNumber + `" , `+
		`"quantity": "` + quantity + `" , `+ 
		`"FAA_formUrl": "` + FAA_formUrl + `" , `+ 
		`"fileHash": "`+fileHash+ `" , `+ 
		`"user": "` + user + `" , `+
		`"itemType": "` + itemType + `" , `+
		`"part_number": "` + part_number + `" , `+ 
		`"total_approvedQty": "` + total_approvedQty + `" , `+ 
		`"approvalDate": "` + approvalDate + `" , `+ 	
		`"authorization_number": "` + authorization_number + `" , `+ 
		`"userType": "` + userType + `"`+
		`}`
		fmt.Println("input: " + input)
		fmt.Print("input in bytes array: ")
		fmt.Println([]byte(input))
	err = stub.PutState(FAA_formNumber, []byte(input))									//store Form with FAA_formNumber as key
	if err != nil {
		return nil, err
	}
	//get the Form index
	Tier3FormIndexAsBytes, err := stub.GetState(Tier3FormIndexStr)
	if err != nil {
		return nil, errors.New("Failed to get Tier-2 Form index")
	}
	var Tier3FormIndex []string
	fmt.Print("Tier3FormIndexAsBytes: ")
	fmt.Println(Tier3FormIndexAsBytes)
	
	json.Unmarshal(Tier3FormIndexAsBytes, &Tier3FormIndex)							//un stringify it aka JSON.parse()
	fmt.Print("Tier3FormIndex after unmarshal..before append: ")
	fmt.Println(Tier3FormIndex)
	//append
	Tier3FormIndex = append(Tier3FormIndex, FAA_formNumber)									//add Form transID to index list
	fmt.Println("! Tier-3 Form index after appending FAA_formNumber: ", Tier3FormIndex)
	jsonAsBytes, _ := json.Marshal(Tier3FormIndex)
	fmt.Print("jsonAsBytes: ")
	fmt.Println(jsonAsBytes)
	err = stub.PutState(Tier3FormIndexStr, jsonAsBytes)						//store name of Form
	if err != nil {
		return nil, err
	}

	fmt.Println("Tier-3 Form created successfully.")
	return nil, nil
}
