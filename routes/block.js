var express = require('express');
var router = express.Router();

router.get('/:blockNum', function(req, res, next) {
	var num = req.params.blockNum;
	req.db.get('explorer::block_num_'+num, function(error, result) {
		if(!error && result != undefined) {
			console.log("Got database stored block for ", num);
			res.json(JSON.parse(result));
		}
		else {
			req.web3.eth.getBlock(num,function(error, result) {
				if(!error && result && result.hash) {
					res.json(result);
					req.db.set('explorer::block_num_'+num, JSON.stringify(result), function(error, result) {
						if(!error) {
							console.log("Stored block info for ", num);
						}
						else {
							console.log("Error storing block info for ", num, error);
						}
					});
				}
				else {
					res.json(error);
				}
			});
		}
	});
});

router.get('/uncle/:blockNum', function(req, res, next) {
	var num = req.params.blockNum;
	req.db.lrange('explorer::uncle_block_'+num, 0, -1, function(error, result) {
		if(!error && result != undefined) {
			console.log("Got database stored uncles for ", num, " : ", result);
			req.db.lrange('explorer::uncle_list_'+num, 0, -1, function(error, uncleList) {
		        uncles = [];
				if(!error && uncleList != undefined && uncleList.length > 0) {
					console.log("Got cached uncle list", uncleList);
                    for(var i in uncleList) { 
                        uncles.push(JSON.parse(uncleList[i]));
                    }
					res.json({ uncles: uncles });
				}
				else {
					var i;
					for(i=0; i<result.length; i++) {
						req.web3.eth.getUncle(num, i, function(err, uncle) {
							uncles.push(uncle);
							req.db.lpush('explorer::uncle_list_'+num, JSON.stringify(uncle));
							if(i==result.length) {
								res.json({ uncles: uncles });
							}
						});
					}
				}
			});
		}
		else {
			console.log("No uncles found for block ", num);
			res.json({ uncles: [] });
		}
	});
});

module.exports = router;
