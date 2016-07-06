var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  //res.render('index', { title: 'Express' });
  var t = [1,2,3,4,5];
  res.json(t);
});

router.get('/transactions/from/:addressId.:start.:end', function(req, res, next) {
  var start = req.params.start ? req.params.start : 1;
  var end = req.params.end ? req.params.end : -1;
  req.db.lrange('_shf_txn_from_' + req.params.addressId, start, end, function(error, result) {
        var resArray = [];
        for(var i in result) {
            resArray.push(JSON.parse(result[i]));
        }
        console.log(resArray);
        res.json(resArray);
  });
});

router.get('/transactions/to/:addressId.:start.:end', function(req, res, next) {
  var start = req.params.start ? req.params.start : 1;
  var end = req.params.end ? req.params.end : -1;
  req.db.lrange('_shf_txn_to_' + req.params.addressId, start, end, function(error, result) {
        var resArray = [];
        for(var i in result) {
            resArray.push(JSON.parse(result[i]));
        }
        console.log(resArray);
        res.json(resArray);
  });
});

router.get('/mined/:addressId.:start.:end', function(req, res, next) {
  var start = req.params.start ? req.params.start : 1;
  var end = req.params.end ? req.params.end : -1;
    req.db.lrange('_shf_block_miner_'+req.params.addressId, start, end, function(error, result) {
        var resArray = [];
        for(var i in result) {
            resArray.push(JSON.parse(result[i]));
        }
        console.log(resArray);
        res.json(resArray);
    });
});

module.exports = router;
module.exports = router;
