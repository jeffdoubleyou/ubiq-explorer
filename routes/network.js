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
	req.db.lrange('explorer::recent_blocks', 0, 0, function(error, result) {
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
  req.db.get('explorer::current_hashrate', function(error, result) {
    var data = {
        error: error,
        hashrate: result*1000*1000*1000
    };
    res.json(data);
  });
});

router.get('/blocktime', function(req, res, next) {
  req.db.get('explorer::current_blocktime', function(error, result) {
    var data = {
        error: error,
        blocktime: result
    };
    res.json(data);
  });
});

router.get('/difficulty', function(req, res, next) {
  req.db.get('explorer::current_difficulty', function(error, result) {
    var data = {
        error: error,
        difficulty: result*1000*1000*1000
    };
    res.json(data);
  });
});

router.get('/hashratehistory', function(req, res, next) {
	req.db.lrange('explorer::history_hashrate', 0, 2500, function(error, result) {
		var resArray = [];
		for (var i in result) {
			resArray.push(result[i]);
		}
		res.json(resArray);
	});
});

router.get('/difficultyhistory', function(req, res, next) {
	req.db.lrange('explorer::history_difficulty', 0, 2500, function(error, result) {
		var resArray = [];
		for (var i in result) {
			resArray.push(result[i]);
		}
		res.json(resArray);
	});
});

router.get('/uncleratehistory', function(req, res, next) {
	req.db.lrange('explorer::history_unclerate', 0, 2500, function(error, result) {
		var resArray = [];
		for (var i in result) {
			resArray.push(result[i]);
		}
		res.json(resArray);
	});
});


router.get('/blocktimehistory', function(req, res, next) {
	req.db.lrange('explorer::history_blocktime', 0, 2500, function(error, result) {
		var resArray = [];
		for (var i in result) {
			resArray.push(result[i]);
		}
		res.json(resArray);
	});
});


router.get('/recentblocks', function(req, res, next) {
	req.db.lrange('explorer::recent_blocks', 0, 10, function(error, result) {
		var resArray = [];
		for(var i in result) {
		    resArray.push(JSON.parse(result[i]));
		}
		res.json(resArray);
	});
});

router.get('/recenttxns', function(req, res, next) {
	req.db.lrange('explorer::recent_transactions', 0, 25, function(error, result) {
		var resArray = [];
		for(var i in result) {
		    resArray.push(JSON.parse(result[i]));
		}
		res.json(resArray);
	});
});

router.get('/topminers', function(req, res, next) {
	req.db.lrange('explorer::top_miners', 0, -1, function(error, result) {
		var resArray = [];
		for(var i in result) {
		    resArray.push(JSON.parse(result[i]));
		}
		res.json(resArray);
	});
});

router.get('/unclerate', function(req, res, next) {
    req.db.get('explorer::current_unclerate', function(error, result) {
        var data = {
            unclerate: result,
            error: error
        }
        res.json(data);
    });
});

router.get('/exchangerate', function(req, res, next) {
    var now = new Date() / 1000;
    var expires = 0;
    var btc = 0;
    var usd = 0;

    console.log("Get exchange rate");
    req.db.get('explorer::exchange_time', function(error, result) {
        if(!error) {
            expires = result;
        }
    
        if(now > expires) {
            console.log("Exchange rate expired at " + expires + " it is currently " + now);
            request.get('https://bittrex.com/api/v1.1/public/getticker?market=BTC-UBQ', function(err, result, body) {
                console.log("BTC REQUEST SETNT");
                if(!err && result.statusCode == 200) {
                    body = JSON.parse(body);
		    if(body.result)
                    	btc = body.result.Last;
		    else
			btc = 0;
                    request.get('https://bittrex.com/api/v1.1/public/getticker?market=USDT-BTC', function(err, result, body) {
                        console.log("USD REQUEST SETNT");
                        if(!err && result.statusCode == 200) {
                            body = JSON.parse(body);
                            usd = parseFloat(body.result.Last);
                            var data = {usd: usd, btc: btc};
                            req.db.set('explorer::exchange_rate', JSON.stringify(data), function(error, result) {
                                console.log("Inserted exchange rate data");
                                req.db.set('explorer::exchange_time', now+60, function(error, result) {
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
            req.db.get('explorer::exchange_rate', function(error, result) {
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
