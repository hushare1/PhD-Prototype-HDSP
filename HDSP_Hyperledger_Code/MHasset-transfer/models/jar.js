var mongoose = require('mongoose');

var jarSchema = mongoose.Schema({

        _id: { type: String, required: true, maxlength: 200 },
        jarId: { type: Number, maxlength: 15 },
        umf: { type: Number, maxlength: 3 },
        jarWt: { type: Number, maxlength: 3 },
        batch: { type: String, maxlength: 30 },
        producerdate: { type: Date },
        labTest: { type: String, maxlength: 35 },
        floralType: { type: String, maxlength: 11 },
        tutinLevel: { type: Number, },
        producername: { type: String },
		producerAdd: { type: String },
		producerContact: { type: String },
		producerLicense: { type: String },
		producerCost: { type: Number },
		distributorname: { type: String },
		distributorDate: { type: Date },
		distributorAdd: { type: String },
		distributorContact: { type: String },
		distributorLicense: { type: String },
		distributorCost: { type: Number },
		retailername: { type: String },
		retailerDate: { type: Date },
		retailerAdd: { type: String },
		retailerContact: { type: String },
		retailerCost: { type: Number }
		
});

module.exports = mongoose.model('Jar', jarSchema);
