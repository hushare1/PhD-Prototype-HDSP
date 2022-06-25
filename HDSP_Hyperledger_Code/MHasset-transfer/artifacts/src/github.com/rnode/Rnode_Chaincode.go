package main

import (
    "fmt"
    // "strconv"
    "encoding/json"
	"github.com/hyperledger/fabric/common/util"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

// Invoice implements a simple chaincode to manage Invoice
type RetailorData struct {
}

var DelIndexStr = "_invid" 
var logger = shim.NewLogger("Retailor_ChainCode")



// Attributes of a Invoice
type Retailor_Data struct {
    Jar_ID string `json:"jarId"`
    Retailor_name string `json:"rname"`
    Retailor_date string `json:"rdate"`
	Retailor_address string `json:"raddress"`
	Retailor_contact string `json:"rcontact"`
	Cost string `json:"cost"`
    Transaction_hash string `json:"transaction_hash"`
}



// ===================================================================================
// Main
// ===================================================================================
func main() {
    err := shim.Start(new(RetailorData))
    if err != nil {
        fmt.Printf("Error starting RetailorData chaincode: %s", err)
    }
}



// ===========================
// Init initializes chaincode
// ===========================
func (t *RetailorData) Init(stub shim.ChaincodeStubInterface) pb.Response {
    var empty []string
    var err error
    jsonAsBytes, _ := json.Marshal(empty)                               //marshal an emtpy array of strings to clear the index
    err = stub.PutState(DelIndexStr, jsonAsBytes)
    if err != nil {
        return shim.Error(err.Error())
    }
    eventMessage := "{ \"message\" : \"Chaincode is deployed successfully.\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }
    return shim.Success(nil)
}


// ========================================
// Invoke - Our entry point for Invocations
// ========================================
func (t *RetailorData) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    function, args := stub.GetFunctionAndParameters()
    fmt.Println("invoke is running " + function)

    // Handle different functions
    if function == "retailorAdd" { //create a new Invoice
        return t.retailorAdd(stub, args)
	}
	 
    eventMessage := "{ \"message\" : \"Received unknown function invocation\", \"code\" : \"503\"}"
    err := stub.SetEvent("errEvent", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("invoke did not find func: " + function) //error
    return shim.Error("Received unknown function invocation")
}


// ==================================================================
// retailorAdd - create a new RetailorData, store into chaincode state
// ==================================================================

func (t *RetailorData) retailorAdd(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error

    if len(args) != 7 {
        return shim.Error("Incorrect number of arguments. Expecting 7")
	}
	
	jarId := args[0]		   
    rname := args[1]
    rdate := args[2]	
    raddress := args[3]
	rcontact := args[4]
	cost := args[5]
	hash := args[6]

	  // ==== Check if Invoice already exists ====
	  jarIdAsBytes, err := stub.GetState(jarId)
	  if err != nil {
		  return shim.Error("Failed to get jarId: " + err.Error())
	  } else if jarIdAsBytes != nil {
		  fmt.Println("This jar Id order already exists: " + jarId)
		  return shim.Error("TThis jar Id order already exists: " + jarId)
	  } else if err == nil {
		// calling dcc chaincode to check if JAR ID is already registered in dnode
			logger.Info("########### calling Distributor chaincode chaincode  ###########")
			//jarId
			
			chainCodeArgs := util.ToChaincodeArgs("jarIDcheck",jarId)
			logger.Info("########### chainCodeArgs passed to dcc chaincode  ###########")
			logger.Info(chainCodeArgs)
			response := stub.InvokeChaincode("dcc", chainCodeArgs, "mychannel")

			if response.Status != shim.OK {
			logger.Info("###########  rcc call had an error ###########")
			   return shim.Error(response.Message)
			}
			logger.Info("########### called Distributor chaincode successfully ###########")
			logger.Info("Invoke chaincode successful. Got response %s", (response.Status))
			logger.Info("Invoke chaincode successful. Got response %s", (response.Message))
			logger.Info("Invoke chaincode successful. Got payload %s", (response.Payload))
			
			if (response.Payload != nil) {
			    logger.Info("Having value", response.Payload)
			} else if (response.Payload == nil) {
			      return shim.Error("This jar ID doesn't exists in Distributor Node, Please first enroll JAR in Distributor node: " + jarId)
			}
	  }
	  
	
    // ====marshal to JSON ====
	retailorData := &Retailor_Data{jarId,rname,rdate,raddress,rcontact,cost,hash}
	

	retailorDataJSONasBytes, err := json.Marshal(retailorData)
    if err != nil {
        return shim.Error(err.Error())
    }
   
    // === Save product to state ===
    err = stub.PutState(jarId, retailorDataJSONasBytes)
    if err != nil {
        return shim.Error(err.Error())
    }

    eventMessage := "{ \"jarId\" : \""+jarId+"\", \"message\" : \"jar record created succcessfully\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }
	    // ==== invoice is saved and indexed. Return success  ====
		fmt.Println("- end retailorAdd")
		return shim.Success(retailorDataJSONasBytes)
}