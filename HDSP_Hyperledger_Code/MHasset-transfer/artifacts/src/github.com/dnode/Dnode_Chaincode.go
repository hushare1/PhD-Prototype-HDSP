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
type DistributorData struct {
}

var DelIndexStr = "_invid" 
var logger = shim.NewLogger("Distributor_ChainCode")


// Attributes of a Invoice
type Distributor_Data struct {
    Jar_ID string `json:"jarId"`
    Distributor_name string `json:"dname"`
    Distributor_date string `json:"ddate"`
	Distributor_address string `json:"daddress"`
	Distributor_contact string `json:"dcontact"`
	Distributor_licence string `json:"dlicence"`
	Cost string `json:"cost"`
    Transaction_hash string `json:"transaction_hash"`
}

// ===================================================================================
// Main
// ===================================================================================
func main() {
    err := shim.Start(new(DistributorData))
    if err != nil {
        fmt.Printf("Error starting DistributorData chaincode: %s", err)
    }
}



// ===========================
// Init initializes chaincode
// ===========================
func (t *DistributorData) Init(stub shim.ChaincodeStubInterface) pb.Response {
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
func (t *DistributorData) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    function, args := stub.GetFunctionAndParameters()
    fmt.Println("invoke is running " + function)

    // Handle different functions
	
	if function == "jarIDcheck" { //create a new Invoice
        return t.jarIDcheck(stub, args)
	}
    if function == "distributorAdd" { //create a new Invoice
        return t.distributorAdd(stub, args)
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
// jarIDcheck - create a new DistributorData, store into chaincode state
// ==================================================================

func (t *DistributorData) jarIDcheck(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error
	var retVal string 
	retVal = "false"
	
    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	
	jarId := args[0]	
	logger.Info("########### Inside dcc chaincode  ###########")
	logger.Info("jarId", jarId)
	
	  // ==== Check if Invoice already exists ====
	  jarIdAsBytes, err := stub.GetState(jarId)
	  
	  logger.Info("jarIdAsBytes", jarIdAsBytes)
	  logger.Info("err :::", err)
	  
	  if err != nil {
		  return shim.Error("Failed to get jarId: " + err.Error())
	  } else if jarIdAsBytes != nil {
		  fmt.Println("This jar ID already exists: " + jarId)
		  retVal = "true"
		  retValJSONasBytes, err := json.Marshal(retVal)
		  logger.Info("retValJSONasBytes: ", retValJSONasBytes)
		  if err == nil {
			return shim.Success(retValJSONasBytes)
		  }
	  }
		return shim.Success(jarIdAsBytes)
}

// ==================================================================
// distributorAdd - create a new DistributorData, store into chaincode state
// ==================================================================

func (t *DistributorData) distributorAdd(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error

    if len(args) != 8 {
        return shim.Error("Incorrect number of arguments. Expecting 8")
	}
	
	jarId := args[0]		   
    dname := args[1]
    ddate := args[2]	
    daddress := args[3]
	dcontact := args[4]
	dlicence := args[5]
	cost := args[6]
	hash := args[7]

	  // ==== Check if Invoice already exists ====
	  jarIdAsBytes, err := stub.GetState(jarId)
	  if err != nil {
		  return shim.Error("Failed to get jarId: " + err.Error())
	  } else if jarIdAsBytes != nil {
		  fmt.Println("This jar ID already exists: " + jarId)
		  return shim.Error("This jar ID already exists: " + jarId)
	  } else if err == nil {
		// calling pcc chaincode to check if JAR ID is already registered in pnode
			logger.Info("########### calling producer chaincode chaincode  ###########")
			//jarId
			
			chainCodeArgs := util.ToChaincodeArgs("jarIDcheck",jarId)
			logger.Info("########### chainCodeArgs passed to pcc chaincode  ###########")
			logger.Info(chainCodeArgs)
			response := stub.InvokeChaincode("pcc", chainCodeArgs, "mychannel")

			if response.Status != shim.OK {
			logger.Info("###########  dcc call had an error ###########")
			   return shim.Error(response.Message)
			}
			logger.Info("########### called producer chaincode successfully ###########")
			logger.Info("Invoke chaincode successful. Got response %s", (response.Status))
			logger.Info("Invoke chaincode successful. Got response %s", (response.Message))
			logger.Info("Invoke chaincode successful. Got payload %s", (response.Payload))
			
			if (response.Payload != nil) {
			    logger.Info("Having value", response.Payload)
			} else if (response.Payload == nil) {
			      return shim.Error("This jar ID doesn't exists in Producer Node, Please first enroll JAR in Producer node: " + jarId)
			}
	  }
	  
    // ====marshal to JSON ====
	distributorData := &Distributor_Data{jarId,dname,ddate,daddress,dcontact,dlicence,cost,hash}
	

	distributorDataJSONasBytes, err := json.Marshal(distributorData)
    if err != nil {
        return shim.Error(err.Error())
    }
   
    // === Save product to state ===
    err = stub.PutState(jarId, distributorDataJSONasBytes)
    if err != nil {
        return shim.Error(err.Error())
    }

    eventMessage := "{ \"jarId\" : \""+jarId+"\", \"message\" : \"jar record created succcessfully\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }
	    // ==== invoice is saved and indexed. Return success  ====
		fmt.Println("- end distributorAdd")
		return shim.Success(distributorDataJSONasBytes)
}