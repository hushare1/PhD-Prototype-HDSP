var express = require("express")
var app = express()
var body = require("body-parser")
var cors = require("cors")
//var User = require('./jar.model')
var mongoose = require("mongoose");
var db = 'mongodb://localhost:27017/honey';
mongoose.set('useNewUrlParser', true);
mongoose.set('useFindAndModify', false);
mongoose.set('useCreateIndex', true);
mongoose.set('useUnifiedTopology', true);

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
    db1.collection('jars').find({'jarId': tid}).toArray(function(err, result) {
        if (err) throw err;
        console.log(result);
    res.json(result)
    })
    ///
})
 
app.listen(3000)