angular.module('Explorer').service('NetworkService', function($rootScope, $interval, $http) {

    var networkService = this;

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

    this.getRecentBlocks = function() {
        return $http.get('/api/v1/block/list?limit=13');
    }

    this.getRecentTransactions = function() {
        return $http.get('/api/v1/transaction/list?limit=25');
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
	    $http({method : 'GET',url : 'https://api.coinmarketcap.com/v1/ticker/UBIQ/?convert=USD'})
		.success(function(data, status) {
            if(data && data[0]) {
                $rootScope.btc = parseFloat(data[0].price_btc);
                $rootScope.shf_usd = parseFloat(data[0].price_usd).toFixed(8);
            }
		})
		.error(function(data, status) {
			$rootScope.btc = 'N/A';
			$rootScope.usd = 'N/A';
			$rootScope.shf_usd = 'N/A';
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

