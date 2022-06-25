var mongoose = require('mongoose');
var Schema = mongoose.Schema;

var userSchema = new Schema({
    id: Number,
    producerName: String,
    pDate: Date,
    phash: String,
    distributorName: String,
    dDate: Date,
    dhash: String,
    retailerName: String,
    rDate: Date,
    rhash: String   
});

module.exports = mongoose.model('User',userSchema);