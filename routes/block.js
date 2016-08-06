var express = require('express');
var router = express.Router();

router.get('/:blockNum', function(req, res, next) {
	var num = req.params.blockNum;
	req.db.get('_shf_block_num_'+num, function(error, result) {
		if(!error && result != undefined) {
			console.log("Got database stored block for ", num);
			res.json(JSON.parse(result));
		}
		else {
			req.web3.eth.getBlock(num,function(error, result) {
				if(!error) {
					res.json(result);
					req.db.set('_shf_block_num_'+num, JSON.stringify(result), function(error, result) {
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

module.exports = router;
