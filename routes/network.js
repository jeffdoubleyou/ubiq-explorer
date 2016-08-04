var express = require('express');
var request = require('request');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  //res.render('index', { title: 'Express' });
  var t = [1,2,3,4,5];
  res.json(t);
});

router.get('/lastblock', function(req, res, next) {
	req.db.lrange('_shf_recent_blocks', 0, 0, function(error, result) {
		result = JSON.parse(result[0]);
		var data = {
			error: error,
			result: result.block
		};
		res.json(data);
	});
});

router.get('/hashrate', function(req, res, next) {
  //res.render('index', { title: 'Express' });
  req.db.get('_shf_current_hashrate', function(error, result) {
    var data = {
        error: error,
        hashrate: result
    };
    res.json(data);
  });
});

router.get('/blocktime', function(req, res, next) {
  req.db.get('_shf_current_blocktime', function(error, result) {
    var data = {
        error: error,
        blocktime: result
    };
    res.json(data);
  });
});

router.get('/difficulty', function(req, res, next) {
  req.db.get('_shf_current_difficulty', function(error, result) {
    var data = {
        error: error,
        difficulty: result
    };
    res.json(data);
  });
});

router.get('/hashratehistory', function(req, res, next) {
	req.db.lrange('_shf_history_hashrate', 0, 2500, function(error, result) {
		var resArray = [];
		for (var i in result) {
			resArray.push(result[i]);
		}
		res.json(resArray);
	});
});

router.get('/difficultyhistory', function(req, res, next) {
	req.db.lrange('_shf_history_difficulty', 0, 2500, function(error, result) {
		var resArray = [];
		for (var i in result) {
			resArray.push(result[i]);
		}
		res.json(resArray);
	});
});

router.get('/blocktimehistory', function(req, res, next) {
	req.db.lrange('_shf_history_blocktime', 0, 2500, function(error, result) {
		var resArray = [];
		for (var i in result) {
			resArray.push(result[i]);
		}
		res.json(resArray);
	});
});


router.get('/recentblocks', function(req, res, next) {
	req.db.lrange('_shf_recent_blocks', 0, 10, function(error, result) {
		var resArray = [];
		for(var i in result) {
		    resArray.push(JSON.parse(result[i]));
		}
		console.log(resArray);
		res.json(resArray);
	});
});

router.get('/recenttxns', function(req, res, next) {
	req.db.lrange('_shf_recent_transactions', 0, 10, function(error, result) {
		var resArray = [];
		for(var i in result) {
		    resArray.push(JSON.parse(result[i]));
		}
		console.log(resArray);
		res.json(resArray);
	});
});

router.get('/topminers', function(req, res, next) {
	req.db.lrange('_shf_top_miners', 0, 10, function(error, result) {
		var resArray = [];
		for(var i in result) {
		    resArray.push(JSON.parse(result[i]));
		}
		console.log(resArray);
		res.json(resArray);
	});
});

router.get('/exchangerate', function(req, res, next) {
    var now = new Date() / 1000;
    var expires = 0;
    var btc = 0;
    var usd = 0;

    console.log("Get exchange rate");
    req.db.get('_shf_exchange_time', function(error, result) {
        if(!error) {
            expires = result;
        }
    
        if(now > expires) {
            console.log("Exchange rate expired at " + expires + " it is currently " + now);
            request.get('https://bittrex.com/api/v1.1/public/getticker?market=BTC-SHF', function(err, result, body) {
                console.log("BTC REQUEST SETNT");
                if(!err && result.statusCode == 200) {
                    body = JSON.parse(body);
                    btc = body.result.Last;
                    request.get('https://bittrex.com/api/v1.1/public/getticker?market=USDT-BTC', function(err, result, body) {
                        console.log("USD REQUEST SETNT");
                        if(!err && result.statusCode == 200) {
                            body = JSON.parse(body);
                            usd = parseFloat(body.result.Last);
                            var data = {usd: usd, btc: btc};
                            req.db.set('_shf_exchange_rate', JSON.stringify(data), function(error, result) {
                                console.log("Inserted exchange rate data");
                                req.db.set('_shf_exchange_time', now+60, function(error, result) {
                                    console.log("Inserted exchange rate expiration time");
                                    res.json(data);
                                });
                            });
                        }
                        else
                        {
                            console.log("Unable to retrieve exchange data: " + err + " Response code: " + result.statusCode);
                            res.json({error: err});
                        }
                    });
                }
                else
                {
                    console.log("Unable to retrieve exchange data: " + err + " Response code: " + result.statusCode);
                    res.json({error: err});
                }
            });
        }
        else
        {
            console.log("Cache is still good at " + expires + " it is now " + now);
            req.db.get('_shf_exchange_rate', function(error, result) {
                if(!error) {
                    res.json(JSON.parse(result));
                }
                else {
                    res.json({error: err});
                }
            });
        }
    });
});
module.exports = router;
