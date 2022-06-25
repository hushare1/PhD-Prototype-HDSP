const jwt = require('jsonwebtoken');
const express = require('express');
const router = express.Router();
var log4js = require('log4js');
var logger = log4js.getLogger('AuthController');
var hfc = require('fabric-client');
var invoke = require('../app/invoke-transaction.js');
var helper = require('../app/helper.js');
app = express();
app.set('secret', 'thisismysecret');

//endpoint ==> /dnode/add
router.post('/add', async function (req, res) {
		logger.debug('==================== Inside End point : /add ==================');
		var peers = "peer0.dnode.example.com";
		var chaincodeName = "dcc";
		var channelName = "mychannel";
		var fcn = "distributorAdd";
		var username = req.username;
		var orgname = req.orgname
		var args = [];
	
			var today = new Date();
			var dd = String(today.getDate()).padStart(2, '0');
			var mm = String(today.getMonth() + 1).padStart(2, '0'); //January is 0 !
			var yyyy = today.getFullYear();

			//today = dd + '-' + mm + '-' + yyyy;
			let todayString = yyyy + '-' + mm + '-' + dd;
			console.log("TODAY ===== ", todayString)
		//todayString = '2019-05-05'
		console.log("__________@@@@_------")
		console.log("today", today)
		console.log("todayString", todayString)
		var jarId = req.body.jarId;
		var dname = req.body.dname;
		var ddate = todayString;
		var daddress = req.body.daddress;
		var dcontact = req.body.dcontact;
		var dlicence = req.body.dlicence;
		var cost = req.body.cost;
		
		args.push(jarId)		
		args.push(dname)
		args.push(ddate)
		args.push(daddress)
		args.push(dcontact)
		args.push(dlicence)
		args.push(cost)

		console.log("peers", peers)
		console.log("channelName", channelName)
		console.log("chaincodeName", chaincodeName)
		console.log("fcn", fcn)
		console.log("args", args)
		console.log("username", username)
		console.log("orgname", orgname)
		
	
		let message = await invoke.invokeChaincode(peers, channelName, chaincodeName, fcn, args, username, orgname);
		logger.debug("Blockchain data received : parse" + JSON.stringify(message));
		res.send(message);

});


module.exports = router;
