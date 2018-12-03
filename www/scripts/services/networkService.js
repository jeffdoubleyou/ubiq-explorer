angular.module('Explorer').service('NetworkService', function($rootScope, $interval, $http) {

    var networkService = this;
    $rootScope.exchangeRates = {};

    this.getHashRateHistory = function() {
        return $http.get('/api/v1/stats/hashRateHistory');
    }

    this.getDifficultyHistory = function() {
        return $http.get('/api/v1/stats/difficultyHistory');
    }

    this.getBlockTimeHistory = function() {
        return $http.get('/api/v1/stats/blockTimeHistory');
    }

    this.getUncleRateHistory = function() {
        return $http.get('/api/v1/stats/uncleRateHistory');
    }

    this.getRecentBlocks = function(limit, cursor) {
        return $http.get('/api/v1/block/list?limit='+limit+'&cursor='+cursor);
    }

    this.getRecentUncles = function(limit, cursor) {
        return $http.get('/api/v1/uncle/list?limit='+limit+'&cursor='+cursor);
    }

    this.getRecentTransactions = function(limit, cursor) {
        return $http.get('/api/v1/transaction/list?limit='+limit+'&cursor='+cursor);
    }

    this.getPendingTransactions = function(limit, cursor) {
        return $http.get('/api/v1/transaction/pending');
    }

    this.getRecentTokenTransactions = function(limit, cursor) {
        return $http.get('/api/v1/token/listTransactions?limit='+limit+'&cursor='+cursor);
    }
    this.getTopMiners = function() {
        return $http.get('/api/v1/stats/miners');
    }

    this.getKnownAddresses = function() {
        return $http.get("/api/v1/address/list");
    }

    this.getTokens = function() {
        return $http.get("/api/v1/token/listTokens");
    }

    this.updateStats = function() {
        $http.get('/api/v1/stats/get?blocks=10').then(function(res) {
            $rootScope.blockNum = res.data.lastBlock;
            $rootScope.networkHashRate = res.data.hashRate;
            $rootScope.networkDifficulty = res.data.difficulty;
            $rootScope.networkBlockTime = parseFloat(res.data.blockTime).toFixed(2);
            $rootScope.networkUncleRate = parseFloat(res.data.uncleRate).toFixed(2);
        });
    };

	this.updateExchangeRate = function() {
	    $http({method : 'GET',url : '/api/v1/exchange/list'})
		.success(function(data, status) {
            if(data) {
                angular.forEach(data, function(exchange) {
                    $rootScope.exchangeRates[exchange.symbol] = exchange;
                });
                $rootScope.btc = parseFloat($rootScope.exchangeRates["UBQ"].btc);
                $rootScope.usd = parseFloat($rootScope.exchangeRates["UBQ"].btc*$rootScope.exchangeRates["BTC"].usd).toFixed(6);
            }
		})
		.error(function(data, status) {
			$rootScope.btc = 'N/A';
			$rootScope.usd = 'N/A';
		})
	}

    this.updateStats();
    this.updateExchangeRate();
    $interval(networkService.updateStats, 15000);
    $interval(networkService.updateExchangeRate, 120000);
    
});

angular.module('Explorer').service('BlockInfoService', function($rootScope, $http, $q) {
	this.getBlock = function(num) {
		return $http.get('/api/v1/block/get?block='+num).then(function(data, status) {
			return data;
		});
	}
    this.getUncles = function(num) {
        return $http.get('/api/v1/uncle/block?block='+num).then(function(data, status) {
            return data;
        });
    }
    this.getTransactions = function(num) {
	return $http.get('/api/v1/transaction/block?block='+num).then(function(data, status) {
            return data;
        });
    }
});

angular.module('Explorer').service('TransactionInfoService', function($rootScope, $http, $q) {
	this.getTransaction = function(txn) {
		return $http.get('/api/v1/transaction/get?hash='+txn).then(function(data, status) {
			return data;
		});
	}
});

