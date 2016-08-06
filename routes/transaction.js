var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  //res.render('index', { title: 'Express' });
  var t = [1,2,3,4,5];
  res.json(t);
});

router.get('/hash/:txnHash', function(req, res, next) {
	var txn = req.params.txnHash;
	req.db.get('_shf_transaction_hash_'+txn, function(error, result) {
		if(!error && result != undefined) {
			console.log("Got database stored transaction for ", txn);
			res.json(JSON.parse(result));
		}
		else {
			req.web3.eth.getTransaction(txn,function(error, result) {
				if(!error) {
					res.json(result);
					req.db.set('_shf_transaction_hash_'+txn, JSON.stringify(result), function(error, result) {
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
