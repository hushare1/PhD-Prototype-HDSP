var express = require("express")
var app = express()
var body = require("body-parser")
var cors = require("cors")
var User = require('./jar.model')
var mongoose = require("mongoose");
var db = 'mongodb://localhost:27017/test';

app.use(body.urlencoded({extended:true}))
app.use(body.json());
app.use(cors())

mongoose.connect(db);

mongoose.connection.once('connected', function() {
	console.log("Database connected successfully")
});

app.use(function (req, res, next) {
      // Website you wish to allow to connect
    res.setHeader('Access-Control-Allow-Origin', '*');
      // Request methods you wish to allow
    res.setHeader('Access-Control-Allow-Methods', 'GET, POST, OPTIONS, PUT, PATCH, DELETE');
      next();
  });


app.get('/jardata/:id', function(req, res){
	console.log(req.params)
	console.log(typeof req.params.id)
	console.log(req.query)
	var tid = Number(req.params.id)
	console.log("tid: ", req.params.id)
	var db1 = mongoose.connection 

// to fetch single record from mongoDB
	db1.collection('avi').find({'id': tid}).toArray(function(err, result) {
              if (err) throw err;
              console.log(result);
   	      res.json(result)
       })
///
})

app.post('/addjar', function(req, res){
    console.log(req.body)
    const user1 = new User ({
     id: req.body.id,
     producerName : req.body.pname,
     pDate : new Date(), 
     phash : req.body.phash,
     distributorName : req.body.dname,
     dDate : new Date(), 
     dhash : req.body.dhash,
     retailerName : req.body.rname,
     rDate : new Date(), 
     rhash : req.body.rhash
     })	
 
 console.log(user1)
 var db1 = mongoose.connection
           db1.collection('avi').insertOne(user1);	
     res.json("Data inserted")
 })
 
app.listen(3000)
