var express = require('express');
var request = require('request');
var router = express.Router();

router.get('/hashratehistory', function(req, res, next) {
	req.db.lrange('explorer::pool_hashrate', 0, 1000, function(error, result) {
		var resArray = [];
		for (var i in result) {
			resArray.push(result[i]);
		}
		res.json(resArray);
	});
});

router.get('/minerhistory/:address', function(req, res, next) {
    var address = req.params.address;
    req.db.zrange('ubq:hashrate:'+address, 0, -1, function(error, result) {
        var miner = {};
        for (var i in result) {
            var info = result[i].split(':');
            if(!miner.hasOwnProperty(info[1]))
                miner[info[1]] = [];
            miner[info[1]].push(info[2]/1000/1000);
        }
        res.json(miner);
    });
});

router.get('/addalert/:address/:email', function(req, res, next) {
    var address = req.params.address;
    var email = req.params.email;
    var ip = req.headers['x-forwarded-for'] || 
         req.connection.remoteAddress || 
         req.socket.remoteAddress ||
         req.connection.socket.remoteAddress;
    console.log("Adding alert for ", address, email, ip);
    if (/^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/.test(email)) {
        req.db.hset('alert:'+address, email, ip, function(error, result) {
            if(error) {
                res.json({ status: 0, error: "Unable to add alert, please try again later" });
            }
            else {
                res.json({ status: 1 });
            }
        });
    }
    else
    {
        res.json({ status: 0, error: "Invalid email address: "+email });
    }
});

router.get('/getalert/:address', function(req, res, next) {
    var address = req.params.address;
    var ip = req.headers['x-forwarded-for'] || 
         req.connection.remoteAddress || 
         req.socket.remoteAddress ||
         req.connection.socket.remoteAddress;
    console.log("Get alerts for ", address, " from ", ip);
    req.db.hgetall('alert:'+address, function(error, result) {
        if(error || !result) {
            res.json({ status: 0 }); 
        }
        else {
            var data = [];
            for(var email in result) {
                var alertIP = result[email];
                console.log("There is an alert found for email ", email, " at IP ", alertIP);
                if(ip == alertIP)
                    data.push(email);
            }
            res.json({ status: 1, alerts: data });
        }
    });
});

router.get('/removealert/:address/:email', function(req, res, next) {
    var address = req.params.address;
    var alertEmail = req.params.email;

    var ip = req.headers['x-forwarded-for'] || 
         req.connection.remoteAddress || 
         req.socket.remoteAddress ||
         req.connection.socket.remoteAddress;
    console.log("Delete alerts for ", alertEmail, " on ", address, " from ", ip);
    req.db.hgetall('alert:'+address, function(error, result) {
        if(error || !result) {
            res.json({ status: 0 }); 
        }
        else {
            var data = [];
            for(var email in result) {
                var alertIP = result[email];
                console.log("There is an alert found for email ", email, " at IP ", alertIP);
                if(ip == alertIP && email == alertEmail) {
                    console.log("Going to remove ", email, " for ", alertIP);
                    data.push(email);
                    req.db.hdel('alert:'+address, email, function(error, result) {
                        console.log("Removing ", email);
                    });
                }
            }
            res.json({ status: 1, emails: data });
        }
    });
});


module.exports = router;
