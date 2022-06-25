## Manuka Honey asset transfer to three peers (honey producer, distributor and retailer) on Hyperledger Fabric 1.4 Network. 
A sample Node.js app to demonstrate **__fabric-client__** & **__fabric-ca-client__** Node.js SDK APIs

### Steps and prerequisites to run the code.
* [Docker](https://www.docker.com/products/overview) - v1.12 or higher
* [Docker Compose](https://docs.docker.com/compose/overview/) - v1.8 or higher
* [Git client](https://git-scm.com/downloads) - needed for clone commands
* **Node.js** v8.4.0 or higher
* [Download Docker images](http://hyperledger-fabric.readthedocs.io/en/latest/samples.html#binaries)

```
cd fabric-samples/MHasset-transfer/
```

Once you have completed the above setup, you will have provisioned a local network with the following docker container configuration:

* 2 CAs
* A SOLO orderer
* 4 peers (2 peers per Org)

#### Artifacts
* Crypto material has been generated using the **cryptogen** tool from Hyperledger Fabric and mounted to all peers, the orderering node and CA containers. More details regarding the cryptogen tool are available [here](http://hyperledger-fabric.readthedocs.io/en/latest/build_network.html#crypto-generator).
* An Orderer genesis block (genesis.block) and channel configuration transaction (mychannel.tx) has been pre generated using the **configtxgen** tool from Hyperledger Fabric and placed within the artifacts folder. More details regarding the configtxgen tool are available [here](http://hyperledger-fabric.readthedocs.io/en/latest/build_network.html#configuration-transaction-generator).


**##Blockchain Hyperledger Setup**

1.	Give command sudo su to go to root. No pwd required. 
2.	Login to Linux server (t2-medium instance) and Go to below folder / directory
/home/ubuntu/MH-Project/MHasset-transfer

	2a.  run command -         rm -r fabric-client-kv-*

	2b. change user name (e.g. Jim3 to Jim4 in app.js file) in testAPIs.sh script.....


3.	Run below command
	./runApp.sh				
	•	This shell script delete existing docker containers and it will again create new docker containers.
	•	This will make node server up. Below output will be generated
	•	Command output: 
	[2020-02-17 08:51:44.097] [INFO] SampleWebApp - ****************** SERVER STARTED ************************
	[2020-02-17 08:51:44.099] [INFO] SampleWebApp - ***************  http://localhost:4000  ******************
	Database connected successfully			 MongoDB is connected successfully
4.	Open new window / duplicate session and go to same folder /home/ubuntu/MH-Project/MHasset-transfer)
	Run below command
	testAPIs.sh
	•	This shell script will enroll users, create channel, add users to channel and deploy the smart contracts
	•	This command will generate unique JWT token. These tokens should be inputted in POSTMAN for testing later 
	•	Command output: 

ORG1 token is 
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODE5NjI3MTQsInVzZXJuYW1lIjoiSmltMSIsIm9yZ05hbWUiOiJQbm9kZSIsImlhdCI6MTU4MTkyNjcxNH0.9aNRfcEbi17JmJjDrJ4rX_01r1Dq0AacPMowdYTjYR4
ORG2 token is 
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODE5NjI3MTUsInVzZXJuYW1lIjoiQmFycnkxIiwib3JnTmFtZSI6IkRub2RlIiwiaWF0IjoxNTgxOTI2NzE1fQ.Rpvr0wRwoo8NqDKKmaPXAxQpoR4FSKP57mEZwtkUH7Y
ORG3 token is 
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODE5NjI3MTUsInVzZXJuYW1lIjoiVG9ueTEiLCJvcmdOYW1lIjoiUm5vZGUiLCJpYXQiOjE1ODE5MjY3MTV9.dxmDbiNXBtVAD6njqCAm5c2iXDZXCCJsjvOSWef77vQ

5.	If all above steps are successful then Blockchain Hyperledger setup is ready for testing



**##Blockchain Hyperledger Testing (Via Postman)**
1.	Open POSTMAN application
2.	Create new POST API request
	•	e.g. http://13.235.244.125:4000/pnode/add
	•	go to Authorization tab. In type select Bearer Token. Enter JWT token which you received in step 3 (for first org)
	•	go to Body tab. Select raw radio button and input data which needs to be passed to Blockchain Hyperledger. Ensure JSON data is selected
	•	once this setup is done. Click on Send button. This will trigger POST API to Blockchain Hyper ledger and you will get POST response back in POSTMAN application
3.	Follow the above step 2 for rest POST API’s request creation

4.	POSTMAN API collections
Please import below POSTMAN API collection for testing
  

## Running the sample program

Run with chaincode written in golang or in node.js.

* Launch the network using docker-compose

```
docker-compose -f artifacts/docker-compose.yaml up
```

npm install
```

* Start the node app on PORT 4000

```
PORT=4000 node app
```

##### Terminal Window 

* Execute the REST APIs from the section using the Postman. 


cd fabric-samples/MHasset-transfer

./runApp.sh

```

* This launches the required network on your local machine
* Installs the fabric-client and fabric-ca-client node modules
* And, starts the node app on PORT 4000

##### Terminal Window 


In order for the following shell script to properly parse the JSON, you must install ``jq``:

instructions [https://stedolan.github.io/jq/](https://stedolan.github.io/jq/)

With the application started in terminal 1, next, test the APIs by executing the script - **testAPIs.sh**:
```
cd fabric-samples/MHasset-transfer

## To use golang chaincode execute the following command

./testAPIs.sh -l golang


```


### Create Channel request

```
curl -s -X POST \
  http://localhost:4000/channels \
  -H "authorization: Bearer <put JSON Web Token here>" \
  -H "content-type: application/json" \
  -d '{
	"channelName":"mychannel",
	"channelConfigPath":"../artifacts/channel/mychannel.tx"
}'
```

Please note that the Header **authorization** must contain the JWT returned from the `POST /users` call

### Join Channel request

```
curl -s -X POST \
  http://localhost:4000/channels/mychannel/peers \
  -H "authorization: Bearer <put JSON Web Token here>" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.example.com","peer1.org1.example.com"]
}'
```
### Install chaincode

```
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer <put JSON Web Token here>" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.example.com","peer1.org1.example.com"],
	"chaincodeName":"mycc",
	"chaincodePath":"github.com/example_cc/go",
	"chaincodeType": "golang",
	"chaincodeVersion":"v0"
}'
```
**NOTE:** *chaincodeType* must be set to **node** when node.js chaincode is used and *chaincodePath* must be set to the location of the node.js chaincode. Also put in the $PWD
```
ex:
curl -s -X POST \
  http://localhost:4000/chaincodes \
  -H "authorization: Bearer <put JSON Web Token here>" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.example.com","peer1.org1.example.com"],
	"chaincodeName":"mycc",
	"chaincodePath":"$PWD/artifacts/src/github.com/example_cc/node",
	"chaincodeType": "node",
	"chaincodeVersion":"v0"
}'
```

### Instantiate chaincode

This is the endorsement policy defined during instantiation.
This policy can be fulfilled when members from both orgs sign the transaction proposal.

```
{
	identities: [{
			role: {
				name: 'member',
				mspId: 'Org1MSP'
			}
		},
		{
			role: {
				name: 'member',
				mspId: 'Org2MSP'
			}
		}
		{
			role: {
				name: 'member',
				mspId: 'Org3MSP'
			}
		}
	],
	policy: {
		'1-of': [{
			'signed-by': 0
		}, {
			'signed-by': 1
		}, {
			'signed-by': 2
		}]
	}
}
```

```
curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes \
  -H "authorization: Bearer <put JSON Web Token here>" \
  -H "content-type: application/json" \
  -d '{
	"chaincodeName":"mycc",
	"chaincodeVersion":"v0",
	"chaincodeType": "golang",
	"args":["a","100","b","200"]
}'
```
**NOTE:** *chaincodeType* must be set to **node** when node.js chaincode is used

### Invoke request

This invoke request is signed by peers from both orgs, *org1* & *org2*.
```
curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/mycc \
  -H "authorization: Bearer <put JSON Web Token here>" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.example.com","peer0.org2.example.com","peer0.org3.example.com"],
	"fcn":"move",
	"args":["a","b","10"]
}'
```
**NOTE:** Ensure that you save the Transaction ID from the response in order to pass this string in the subsequent query transactions.

### Chaincode Query

```
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/mycc?peer=peer0.org1.example.com&fcn=query&args=%5B%22a%22%5D" \
  -H "authorization: Bearer <put JSON Web Token here>" \
  -H "content-type: application/json"
```

### Query Block by BlockNumber

```
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/blocks/1?peer=peer0.org1.example.com" \
  -H "authorization: Bearer <put JSON Web Token here>" \
  -H "content-type: application/json"
```

### Query Transaction by TransactionID

```
curl -s -X GET http://localhost:4000/channels/mychannel/transactions/<put transaction id here>?peer=peer0.org1.example.com \
  -H "authorization: Bearer <put JSON Web Token here>" \
  -H "content-type: application/json"
```
**NOTE**: The transaction id can be from any previous invoke transaction, see results of the invoke request, will look something like `8a95b1794cb17e7772164c3f1292f8410fcfdc1943955a35c9764a21fcd1d1b3`.


### Query ChainInfo

```
curl -s -X GET \
  "http://localhost:4000/channels/mychannel?peer=peer0.org1.example.com" \
  -H "authorization: Bearer <put JSON Web Token here>" \
  -H "content-type: application/json"
```

### Query Installed chaincodes

```
curl -s -X GET \
  "http://localhost:4000/chaincodes?peer=peer0.org1.example.com&type=installed" \
  -H "authorization: Bearer <put JSON Web Token here>" \
  -H "content-type: application/json"
```

### Query Instantiated chaincodes

```
curl -s -X GET \
  "http://localhost:4000/chaincodes?peer=peer0.org1.example.com&type=instantiated" \
  -H "authorization: Bearer <put JSON Web Token here>" \
  -H "content-type: application/json"
```

### Query Channels

```
curl -s -X GET \
  "http://localhost:4000/channels?peer=peer0.org1.example.com" \
  -H "authorization: Bearer <put JSON Web Token here>" \
  -H "content-type: application/json"
```

### Clean the network

The network will still be running at this point. Before starting the network manually again, here are the commands which cleans the containers and artifacts.

```
docker rm -f $(docker ps -aq)
docker rmi -f $(docker images | grep dev | awk '{print $3}')
rm -rf fabric-client-kv-org[1-2]
```

### Network configuration considerations

You have the ability to change configuration parameters by either directly editing the network-config.yaml file or provide an additional file for an alternative target network. The app uses an optional environment variable "TARGET_NETWORK" to control the configuration files to use. For example, if you deployed the target network on Amazon Web Services EC2, you can add a file "fabric-ca-server-config.yaml", and set the "TARGET_NETWORK" environment to 'aws'. The app will pick up the settings inside the "network-config-aws.yaml" file.

#### IP Address** and PORT information

If you choose to customize your docker-compose yaml file by hardcoding IP Addresses and PORT information for your peers and orderer, then you MUST also add the identical values into the fabric-ca-server-config.yaml file. The url and eventUrl settings will need to be adjusted to match your docker-compose yaml file.

```
peer1.org1.example.com:
  url: grpcs://x.x.x.x:7056
  eventUrl: grpcs://x.x.x.x:7058

```

#### Discover IP Address

To retrieve the IP Address for one of your network entities, issue the following command:

```
# this will return the IP Address for peer0
docker inspect peer0 | grep IPAddress
```
