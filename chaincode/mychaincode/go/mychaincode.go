package main
import (
	"fmt"
	"strings"
	// "strconv"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)
type SmartContract struct {
}
type users struct {
	StudentID 	string 	`json:"Student"`
	Name 		string 	`json:"Name"` 
	Tel       	string 	`json:"Tel"`
	Status      string 	`json:"Status"`
}
type wallet struct {
	WalletID	string	`json:"walletID"`
	Walletname 	string 	`json:"Walletname"`
	Money 		string 	`json:"Money"` 
	Owner		string	`json:"Owner"`
}
func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}



func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "createuser" {
		// Make payment of X units from A to B
		return t.createuser(stub, args)
	} else if function == "createwallet" {
		// Deletes an entity from its state
		return t.createwallet(stub, args)
	} else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	}
	return shim.Error("Invalid invoke function name. Expecting \"createuser\" \"createwallet\" \"query\"")
}
func (t *SmartContract) createuser(stub shim.ChaincodeStubInterface, args[]string) pb.Response {

	justString := strings.Join(args,"")
	args = strings.Split(justString,"|")
	// 0			1			2		3
	// StudentID	Name		Tel		hashID	
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	if len(args[0]) <= 0 {
		return shim.Error("StudentID = null")
	}
	if len(args[1]) <= 0 {
		return shim.Error("Name = null")
	}
	if len(args[2]) <= 0 {
		return shim.Error("Tel = null")
	}
	if len(args[3]) <= 0 {
		return shim.Error("hash = null")
	}


	Users := users{
			StudentID 	:  args[0],
			Name 		:  args[1],
			Tel   		:  args[2],
			Status    	:  "true",	 
		  }	
	userKey := "StudentID|"+args[3]

	Usersbytes,err := json.Marshal(Users)
	if err != nil {
		return shim.Error("Marshal is error"+err.Error())
	}
	err = stub.PutState(userKey,Usersbytes)
	if err != nil {
		return shim.Error("PutState is error"+err.Error())
	}
	return shim.Success(nil)
}
func (t *SmartContract) createwallet(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	justString := strings.Join(args,"")
	args = strings.Split(justString,"|")
	// 0			1			2			3
	// walletname	Money		Owner		hashID	
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	if len(args[0]) <= 0 {
		return shim.Error("walletname = null")
	}
	if len(args[1]) <= 0 {
		return shim.Error("Money = null")
	}
	if len(args[2]) <= 0 {
		return shim.Error("Owner = null")
	}
	if len(args[3]) <= 0 {
		return shim.Error("hash = null")
	}
	Wallet := wallet{
			Walletname 	:  args[0],
			Money  		:  args[1],
			Owner    	:  args[2],	 
		  }	
	walletKey := "Walletname|"+args[3]
	wallebytes,err := json.Marshal(Wallet)
	if err != nil {
		return shim.Error("Marshal is error"+err.Error())
	}
	err = stub.PutState(walletKey,wallebytes)
	if err != nil {
		return shim.Error("PutState is error"+err.Error())
	}
	return shim.Success(nil)
}
func (t *SmartContract) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string // Entities
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

func main() {
	// // Create a new Smart Contract
	err := shim.Start(new (SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s",err)
	}
}
