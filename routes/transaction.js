var express = require('express');
var router = express.Router();

router.get('/hash/:txnHash', function(req, res, next) {
	var txn = req.params.txnHash;
	req.db.get('explorer::transaction_hash_'+txn, function(error, result) {
        result = JSON.parse(result);
		if(!error && result != undefined && result.blockNumber > 0) {
			console.log("Got database stored transaction for ", txn);
			res.json(result);
		}
		else {
			req.web3.eth.getTransaction(txn,function(error, result) {
				if(!error) {
					res.json(result);
					req.db.set('explorer::transaction_hash_'+txn, JSON.stringify(result), function(error, result) {
						if(!error) {
							console.log("Stored transaction hash ", txn);
						}
						else {
							console.log("Error storing transaction hash", txn, error);
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
