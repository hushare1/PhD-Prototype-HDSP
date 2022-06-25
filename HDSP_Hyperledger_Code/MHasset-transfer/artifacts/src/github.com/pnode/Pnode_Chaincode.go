package main

import (
    "fmt"
    // "strconv"
    "encoding/json"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

// Invoice implements a simple chaincode to manage Invoice
type ProducerData struct {
}

var DelIndexStr = "_invid" 

var logger = shim.NewLogger("Producer_ChainCode")

// Attributes of a Invoice
type Producer_Data struct {
    Jar_ID string `json:"jarId"`
    UMF string `json:"umf"`
	Jar_Wt string `json:"jarWt"`
	Batch string `json:"batch"`
	Producer_date string `json:"pdate"`
	Lab_test string `json:"lab"`
    Floral_Type string `json:"floral"`
    Tutin string `json:"tutin"`
    Producer_name string `json:"pname"`
	Producer_address string `json:"paddress"`
	Producer_contact string `json:"pcontact"`
	Producer_licence string `json:"plicence"`
	Cost string `json:"cost"`
    Transaction_date string `json:"date"`
    Transaction_hash string `json:"transaction_hash"`
}



// ===================================================================================
// Main
// ===================================================================================
func main() {
    err := shim.Start(new(ProducerData))
    if err != nil {
        fmt.Printf("Error starting ProducerData chaincode: %s", err)
    }
}



// ===========================
// Init initializes chaincode
// ===========================
func (t *ProducerData) Init(stub shim.ChaincodeStubInterface) pb.Response {
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
func (t *ProducerData) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    function, args := stub.GetFunctionAndParameters()
    fmt.Println("invoke is running " + function)

    // Handle different functions
    if function == "jarIDcheck" { //create a new Invoice
        return t.jarIDcheck(stub, args)
	}
	
	if function == "producerAdd" { //create a new Invoice
        return t.producerAdd(stub, args)
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
// jarIDcheck - create a new ProducerData, store into chaincode state
// ==================================================================

func (t *ProducerData) jarIDcheck(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error
	var retVal string 
	retVal = "false"
	
    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	
	jarId := args[0]	
	logger.Info("########### Inside PCC chaincode  ###########")
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
// producerAdd - create a new ProducerData, store into chaincode state
// ==================================================================

func (t *ProducerData) producerAdd(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error

    if len(args) != 15 {
        return shim.Error("Incorrect number of arguments. Expecting 15")
	}
	
	jarId := args[0]	
	umf := args[1]	
	jarWt := args[2]	
	batch := args[3]	
	pdate := args[4]	
	lab := args[5]	
    floral := args[6]
    tutin := args[7]    
    pname := args[8]
    paddress := args[9]
	pcontact := args[10]
	plicence := args[11]
	cost := args[12]
	systemDate := args[13]
	hash := args[14]

	  // ==== Check if Invoice already exists ====
	  jarIdAsBytes, err := stub.GetState(jarId)
	  if err != nil {
		  return shim.Error("Failed to get jarId: " + err.Error())
	  } else if jarIdAsBytes != nil {
		  fmt.Println("This jar ID already exists: " + jarId)
		  return shim.Error("This jar ID already exists: " + jarId)
	  }
	  
	//   pdate_date, _err := strconv.Atoi(pdate)   //Date 1
    //   if _err != nil {
    //       return shim.Error(_err.Error())
    //   }
	  
	//   transaction_date, _err := strconv.Atoi(systemDate)   //Date 1
    //   if _err != nil {
    //       return shim.Error(_err.Error())
    //   }
	  
    // ====marshal to JSON ====
	producerData := &Producer_Data{jarId,umf,jarWt,batch,pdate,lab,floral,tutin,pname,paddress,pcontact,plicence,cost,systemDate,hash}
	

	producerDataJSONasBytes, err := json.Marshal(producerData)
    if err != nil {
        return shim.Error(err.Error())
    }
   
    // === Save product to state ===
    err = stub.PutState(jarId, producerDataJSONasBytes)
    if err != nil {
        return shim.Error(err.Error())
    }

    eventMessage := "{ \"jarId\" : \""+jarId+"\", \"message\" : \"jar record created succcessfully\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }
	    // ==== invoice is saved and indexed. Return success  ====
		fmt.Println("- end producerAdd")
		return shim.Success(producerDataJSONasBytes)
}