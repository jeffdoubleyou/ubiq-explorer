var express = require('express');
var path = require('path');
var favicon = require('serve-favicon');
var logger = require('morgan');
var cookieParser = require('cookie-parser');
var bodyParser = require('body-parser');
var Redis = require('redis');
//var redis = Redis.createClient(6379, '127.0.0.1');
var redis = Redis.createClient(6379, 'ubiq.hujoqs.0001.usw2.cache.amazonaws.com');
var web3 = require('web3');
var Web3 = new web3();
Web3.setProvider(new web3.providers.HttpProvider("http://127.0.0.1:8888"));

var routes = require('./routes/index');
var users = require('./routes/users');
var address = require('./routes/address');
var network = require('./routes/network');
var transaction = require('./routes/transaction');
var block = require('./routes/block');
var pool = require('./routes/pool');

var app = express();

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'jade');

// uncomment after placing your favicon in /public
//app.use(favicon(path.join(__dirname, 'public', 'favicon.ico')));
app.use(logger('dev'));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(cookieParser());
app.use('/api', express.static(__dirname + '/public/apidoc'));

app.use(function(req,res,next) {
    req.db = redis;
    req.web3 = Web3;
    next();
});

app.use('/', routes);
app.use('/api/users', users);
app.use('/api/address', address);
app.use('/api/network', network);
app.use('/api/transaction', transaction);
app.use('/api/block', block);
app.use('/api/pool', pool);

// catch 404 and forward to error handler
app.use(function(req, res, next) {
  var err = new Error('Method Not Found');
  err.status = 404;
  next(err);
});

// error handlers

// development error handler
// will print stacktrace
if (app.get('env') === 'development') {
  app.use(function(err, req, res, next) {
    res.status(err.status || 500);
    res.render('error', {
      message: err.message,
      error: {}
    });
  });
}

// production error handler
// no stacktraces leaked to user
app.use(function(err, req, res, next) {
  res.status(err.status || 500);
  res.render('error', {
    message: err.message + " ERRRRR",
    error: {}
  });
});


module.exports = app;
