const mongoose = require('mongoose');
const appConfig = require('../app-config/appConfig');
const request = require('request')

var log4js = require('log4js');
var logger = log4js.getLogger('Queries');
const Jar = require('../models/jar');
var invoke = require('../app/invoke-transaction.js');
mongoose.set('useNewUrlParser', true);
mongoose.set('useFindAndModify', false);
mongoose.set('useCreateIndex', true);
mongoose.set('useUnifiedTopology', true);

class Queries {

    async createJarRecord(jarId, umf, jarWt, batch, pdate, lab, floral, tutin, pname, paddress, pcontact, plicence, cost, systemDate) {

        return new Promise((resolve, reject) => {
            logger.debug("============ Inside jar record creation ================");
            Jar.findOne({ jarId: jarId }, (error, jar) => {
                if (error) {
                    logger.error("error: " + error);
                    reject("Internal server error while finding jar.");
                } else if (jar) {
                   resolve("jar already exists");
                    //TODO -- Update query should be here
                } else {
                    // var createDate = Date.now();
                    console.log("********* Inserting query");
                    var test = " "
                    Jar.insertMany({
                        _id: jarId,
                        jarId: jarId,
                        umf: umf,
                        jarWt: jarWt,
                        batch: batch,
                        producerdate: pdate,
                        labTest: lab,
                        floralType: floral,
                        tutinLevel: tutin,
                        producername: pname,
                        producerAdd: paddress,
                        producerContact: pcontact,
                        producerLicense:  plicence,
                        producerCost: cost,
                        distributorname:  test,
                        distributorDate: systemDate,
                        distributorAdd:  test,
                        distributorContact:  test,
                        distributorLicense:  test,
                        distributorCost: test,
                        retailername:  test,
                        retailerDate: systemDate,
                        retailerAdd:  test,
                        retailerContact:  test,
                        retailerCost: test
                    }, (error, jar) => {
                        if (error) {
                            reject("Internal server error while creating Jar.");
                        } else if (jar) {
                            logger.debug("Following jar record created: ", jar);
                            resolve(jar);
                        }
                    });
                }
                console.log("Inserted data: ", jar);
            });
        }) // End of Promise
    }

    async updateDnode(jarId, dname, ddate, daddress, dcontact, dlicence, cost) {

        return new Promise((resolve, reject) => {
            logger.debug("============ Inside jar record creation ================")
            console.log("********* Inserting query");
            console.log("jarId", jarId);
            console.log("dname", dname);
                    Jar.findOneAndUpdate({jarId: jarId},
                        { $set: {distributorname: dname, distributorDate: ddate, distributorAdd: daddress, distributorContact: dcontact, 
                        distributorLicense: dlicence, distributorCost: cost}},
                        {new:true})       
                        .then((jar)=>{
                        if(jar) {
                          logger.debug("Following jar record UPDATED: ", jar);
                          resolve(jar);
                        } else {
                          reject({success:false,data:"no such user exist"});
                        }
                     }).catch((err)=>{
                         reject(err);
                     })
                });
}
async updateRnode(jarId, rname, rdate, raddress, rcontact, cost) {

    return new Promise((resolve, reject) => {
        logger.debug("============ Inside jar record creation ================")
        console.log("********* Inserting query");
        console.log("jarId", jarId);
        console.log("dname", rname);
                Jar.findOneAndUpdate({jarId: jarId},
                    { $set: {retailername: rname, retailerDate: rdate, retailerAdd: raddress, retailerContact: rcontact, 
                    retailerCost: cost}},
                    {new:true})       
                    .then((jar)=>{
                    if(jar) {
                      logger.debug("Following jar record UPDATED: ", jar);
                      resolve(jar);
                    } else {
                      reject({success:false,data:"no such user exist"});
                    }
                 }).catch((err)=>{
                     reject(err);
                 })
            });
}

}

module.exports = Queries;
