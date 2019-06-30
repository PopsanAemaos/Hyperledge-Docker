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
	Walletname 	string 	`json:"Walletname"`
	Money 		string 	`json:"Money"` 
	Owner		string	`json:"Owner"`
}
type docplant struct {
	DocID					string				`json:"DocID"`
	Name					string				`json:"Name"`
	ID						string				`json:"ID"`
	address					string				`json:"Address"`
	Tel						string				`json:"Tel"`
	Endorse					string				`json:"Endorse"`
	Garden					string				`json:"Garden"`
	Classcheck				string				`json:"Classcheck"`
	Datecheck				string				`json:"Datecheck"`
	Timecheck				timecheck			`json:"Timecheck"`
	Infoendorse				infoendorse			`json:"Infoendorse"`
	Docfarm					docfarm				`json:"Docfarm"`
	Information				string				`json:"Infoendorse"`
	Actionstandard			[]actionsandard		`json:"Actionstandard"`
	Login					[]login				`json:"Login"`
	Plantname				string				`json:"Plantname"`
	Knowledge				[]knowledge			`json:"Knowledge"`
	Scope					scope				`json:"Scope"`
	Officer					string				`json:"Officer"`
	Dateofficer				string				`json:"Date"`
	Farmer					string				`json:"Farmer"`
	Datefarmer				string				`json:"Date"`
	Certification			string				`json:"Certification"`
	Time					string				`json:"Time"`
	Plant					string				`json:"Plant"`
	Noplant					string				`json:"Noplant"`
	Nextcheck 				string				`json:"Nextcheck"`
}
type timecheck struct {
	Start					string	`json:"Start"`
	End						string	`json:"End"`
}
type infoendorse struct {
	Roll					string	`json:"Roll"`
	Name					string	`json:"Name"`	
	Relation				string	`json:"Relation"`
}
type docfarm struct {
	Factor					factor	`json:"Factor"`
	Diagram					diagram	`json:"Diagram"`
}
type actionStandard struct {
	Dependency				string	`json:"Dependency"`
	Poit					string	`json:"Poit"`
	Expland					string	`json:"Expland"`
}
type login struct {
	East					string	`json:"East"`
	West					string	`json:"West"`
	North					string	`json:"North"`
	South					string	`json:"South"`
}
type knowledge struct {
	Dependency				string	`json:"Dependency"`
	Poit					string	`json:"Poit"`
	Expland					string	`json:"Expland"`
}
type scope struct {
	Scope1					string	`json:"Scope1"`
    Scope2					string	`json:"Scope2"`
    Scope3  				string	`json:"Scope3"`
    Scope4 					string	`json:"Scope4"`
	Scope5 					string	`json:"Scope5"`
	Expland 				string	`json:"Expland"`
}
type factor struct {
	Factor					string	`json:"Factor"`
	Save					string	`json:"Save"`
}
type diagram struct {
	Diagram					string	`json:"Diagram"`
	Save					string	`json:"Save"`
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
	} else if function == "createstandard" {
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
func (t *SmartContract)createstandard(stub shim.ChaincodeStubInterface, args[]string) pb.Response {

	justString := strings.Join(args,"")
	args = strings.Split(justString,"|")	
	if len(args) != 27 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	if len(args[0]) <= 0 {
		return shim.Error("StudentID = null")
	}
	timecheck, err := gettimecheck(args[9])
	if err != nil {
		return shim.Error("timecheck" + err.Error())
	}
	infoendorse, err := getinfoendorse(args[10])
	if err != nil {
		return shim.Error("infoendorse" + err.Error())
	}
	docfarm, err := getdocfarm(args[11])
	if err != nil {
		return shim.Error("docfarm" + err.Error())
	}
	actionstandard, err := getactionstandard(args[13])
	if err != nil {
		return shim.Error(methodName + err.Error())
	}
	login, err := getlogin(args[14])
	if err != nil {
		return shim.Error(methodName + err.Error())
	}
	knowledge, err := getknowledge(args[16])
	if err != nil {
		return shim.Error(methodName + err.Error())
	}
	scope, err := getinfoendorse(args[17])
	if err != nil {
		return shim.Error(methodName + err.Error())
	}
	

	Docplant := docplant{
			DocID					:  args[0],
			Name					:  args[1],
			ID						:  args[2],
			address					:  args[3],
			Tel						:  args[4],
			Endorse					:  args[5],
			Garden					:  args[6],
			Classcheck				:  args[7],
			Datecheck				:  args[8],
			Timecheck				:timecheck,
			Infoendorse				:infoendorse,			
			Docfarm					:docfarm,				
			Information				:  args[12],
			Actionstandard			:actionstandard,
			Login					:Login,			
			Plantname				:  args[15],
			Knowledge				:knowledge			
			Scope					:scope,				
			Officer					:  args[18],
			Dateofficer				:  args[19],
			Farmer					:  args[20],
			Datefarmer				:  args[21],
			Certification			:  args[22],
			Time					:  args[23],
			Noplant					:  args[24],
			Nextcheck 				:  args[25]
			HashID					:  args[26]
		}	  

	Key := "DocID|"+args[26]

	bytes,err := json.Marshal(Docplant)
	if err != nil {
		return shim.Error("Marshal is error"+err.Error())
	}
	err = stub.PutState(Key,bytes)
	if err != nil {
		return shim.Error("PutState is error"+err.Error())
	}
	return shim.Success(nil)
}
func gettimecheck(timechecks) (timecheck, error) {

	var timecheckAsStruct timecheck
	var jsonData = []byte(timechecks)
	err := json.Unmarshal(jsonData, &timecheckAsStruct)
	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
	}
	return timecheckAsStruct, nil
}
func getinfoendorse(infoendorses) (infoendorse, error) {
	var infoendorseAsStruct infoendorse
	var jsonData = []byte(infoendorses)
	err := json.Unmarshal(jsonData, &infoendorseAsStruct)
	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
	}
	return infoendorseAsStruct, nil
}
func getdocfarm(docfarms) (docfarm, error) {
	var docfarmAsStruct docfarm
		var jsonData = []byte(docfarms)
		err := json.Unmarshal(jsonData, &docfarmAsStruct)
		if err != nil {
			fmt.Printf("There was an error decoding the json. err = %s", err)
		}
		return docfarmAsStruct, nil
}
func getactionstandard(actionStandards) (actionstandard, error) {
	var actionstandardAsStruct actionstandard
	var jsonData = []byte(actionStandards)
	err := json.Unmarshal(jsonData, &actionstandardAsStruct)
	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
	}
	return actionstandardAsStruct, nil
}
func getlogin(logins) (login, error) {
	var loginAsStruct login
	var jsonData = []byte(logins)
	err := json.Unmarshal(jsonData, &loginAsStruct)
	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
	}
	return loginAsStruct, nil
}
func getknowledge(knowledges) (knowledge, error) {
	var knowledgeAsStruct knowledge
		var jsonData = []byte(knowledges)
		err := json.Unmarshal(jsonData, &knowledgeAsStruct)
		if err != nil {
			fmt.Printf("There was an error decoding the json. err = %s", err)
		}
		return knowledgeAsStruct, nil
}
func getscope(scopes) (scope, error) {
	var scopeAsStruct scope
	var jsonData = []byte(scopes)
	err := json.Unmarshal(jsonData, &scopeAsStruct)
	if err != nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
	}
	return scopeAsStruct, nil
}


func main() {
	// // Create a new Smart Contract
	err := shim.Start(new (SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s",err)
	}
}
