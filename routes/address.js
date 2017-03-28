var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  //res.render('index', { title: 'Express' });
  var t = [1,2,3,4,5];
  res.json(t);
});

/**
* @api {get} /api/transactions/from:address.:start.:limit Get address information
* @apiName from
* @apiGroup transactions
*/

router.get('/transactions/from/:addressId.:start.:end', function(req, res, next) {
  var start = req.params.start ? req.params.start : 1;
  var end = req.params.end ? req.params.end : -1;
  var addressId = req.params.addressId;
  if(addressId.substring(0,2) != '0x') {
	addressId = '0x'+addressId;
  }
  req.db.lrange('explorer::txn_from_' + addressId, start, end, function(error, result) {
        var resArray = [];
        for(var i in result) {
            resArray.push(JSON.parse(result[i]));
        }
        res.json(resArray);
  });
});

router.get('/transactions/from/total/:addressId', function(req, res, next) {
	var addressId = req.params.addressId;
	if(addressId.substring(0,2) != '0x') {
		adressId = '0x'+addressId;
	}
	req.db.llen('explorer::txn_from_'+addressId, function(error, result) {
		var data = {
			error: error,
			transactions: result
		};
		res.json(data);
	});
});

router.get('/transactions/to/:addressId.:start.:end', function(req, res, next) {
  var start = req.params.start ? req.params.start : 1;
  var end = req.params.end ? req.params.end : -1;
  var addressId = req.params.addressId;
  if(addressId.substring(0,2) != '0x') {
	addressId = '0x'+addressId;
  }
  req.db.lrange('explorer::txn_to_' + addressId, start, end, function(error, result) {
        var resArray = [];
        for(var i in result) {
            resArray.push(JSON.parse(result[i]));
        }
        res.json(resArray);
  });
});

router.get('/transactions/to/total/:addressId', function(req, res, next) {
	var addressId = req.params.addressId;
	if(addressId.substring(0,2) != '0x') {
		adressId = '0x'+addressId;
	}
	req.db.llen('explorer::txn_to_'+addressId, function(error, result) {
		var data = {
			error: error,
			transactions: result
		};
		res.json(data);
	});
});

router.get('/transactions/balance/:addressId', function(req, res, next) {
  var addressId = req.params.addressId;
  if(addressId.substring(0,2) != '0x') {
	addressId = '0x'+addressId;
  }
  req.db.llen('explorer::balance_'+addressId, function(error, result) {
	var end = result;
	var start = result - 500;
	if(start < 0) {
		start = 0;
	}
	req.db.lrange('explorer::balance_'+addressId, start, end, function(error, result) {
		res.json(result);
	});
  });

});

router.get('/mined/:addressId.:start.:end', function(req, res, next) {
  var start = req.params.start ? req.params.start : 1;
  var end = req.params.end ? req.params.end : -1;
  var addressId = req.params.addressId;
  if(addressId.substring(0,2) != '0x') {
	addressId = '0x'+addressId;
  }
    req.db.lrange('explorer::block_miner_'+ addressId, start, end, function(error, result) {
        var resArray = [];
        for(var i in result) {
            resArray.push(JSON.parse(result[i]));
        }
        res.json(resArray);
    });
});

router.get('/mined/blocks/:addressId', function(req, res, next) {
	var addressId = req.params.addressId;
	if(addressId.substring(0,2) != '0x') {
		adressId = '0x'+addressId;
	}
	req.db.llen('explorer::block_miner_'+addressId, function(error, result) {
		var data = {
			error: error,
			blocks: result
		};
		res.json(data);
	});
});

router.get('/minechart/:addressId', function(req, res, next) {
	var desired_days = 30;
	var addressId = req.params.addressId;
	var days = [];
	var have_days = 0

	if(addressId.substring(0,2) != '0x') {
		addressId = '0x'+addressId;
	}

	req.db.lrange('explorer::block_miner_'+addressId,0,1000, function(error, result) {
				
	});
	
});

router.get('/balance/:addressId', function(req, res, next) {
  var addressId = req.params.addressId;
  if(addressId.substring(0,2) != '0x') {
	addressId = '0x'+addressId;
  }
  req.db.llen('explorer::balance_'+addressId, function(error, result) {
	req.db.lindex('explorer::balance_'+addressId, -1, function(error, result) {
		var data = {
			error: error,
			balance: result ? result : 0
		}
		res.json(data);
	});
  });
});

router.get('/balance2/:addressId', function(req, res, next) {
  var addressId = req.params.addressId;
  if(addressId.substring(0,2) != '0x') {
	addressId = '0x'+addressId;
  }
  req.web3.eth.getBalance(addressId,function(error, result) {
	if(!error) {
		data = {
			error: null,
			balance: result ? result : 0
		};
		res.json(data);
	}
	else {
		res.json(error);
	}
  });
});

router.get('/richlist', function(req, res, next) {
    req.db.lrange('explorer::richlist', 0, -1, function(error, result) {
        var resArray = [];
        for(var i in result) {
            var parts = result[i].split(':');
            resArray.push({address: parts[0], balance: parts[1]});
        }
        res.json(resArray.reverse());
    });
});

module.exports = router;
module.exports = router;
