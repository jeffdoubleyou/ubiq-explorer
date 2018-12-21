angular.module('Explorer').service('ExchangeService', function($http) {
    this.getExchangeHistory = function(symbol) {
        return $http.get('/api/v1/exchange/history?symbol='+symbol);
    }
    this.getExchangeList = function() {
	return $http.get('/api/v1/exchange/list');
    }
});


