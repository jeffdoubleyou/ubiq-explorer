angular.module('Explorer').service('NetworkService', function($rootScope, $interval, $http) {

    var networkService = this;

    this.getHashRateHistory = function() {
        return $http.get('/api/network/hashratehistory');
    }

    this.getDifficultyHistory = function() {
        return $http.get('/api/network/difficultyhistory');
    }

    this.getBlockTimeHistory = function() {
        return $http.get('/api/network/blocktimehistory');
    }

    this.getRecentBlocks = function() {
        return $http.get('/api/network/recentblocks');
    }

    this.getRecentTransactions = function() {
        return $http.get('/api/network/recenttxns');
    }

    this.getTopMiners = function() {
        return $http.get('/api/network/topminers');
    }

	this.updateNetworkHashRate = function() {
		$http({ method: 'GET', url: '/api/network/hashrate'})
			.success(function(data, status) {
				$rootScope.networkHashRate = data.hashrate;
		})
		.error(function(data, status) {
			$rootScope.networkHashRate = 'N/A';
		})
	};

	this.updateDifficulty = function() {
	    $http({method : 'GET',url : '/api/network/difficulty'})
		.success(function(data, status) {
			$rootScope.networkDifficulty = data.difficulty;
		})
		.error(function(data, status) {
			$rootScope.networkDifficulty = 'N/A';
		})
	}

	this.updateBlockTime = function() {
	    $http({method : 'GET',url : '/api/network/blocktime'})
		.success(function(data, status) {
			var blocktime = parseFloat(data.blocktime).toFixed(2);
			$rootScope.networkBlockTime = blocktime;
		})
		.error(function(data, status) {
			$rootScope.networkBlockTime = 'N/A';
		})
	}

	this.updateExchangeRate = function() {
	    $http({method : 'GET',url : '/api/network/exchangerate'})
		.success(function(data, status) {
			$rootScope.btc = parseFloat(data.btc);
			$rootScope.usd = parseFloat(data.usd).toFixed(4);
			$rootScope.shf_usd = parseFloat($rootScope.btc*$rootScope.usd).toFixed(8);
		})
		.error(function(data, status) {
			$rootScope.btc = 'N/A';
			$rootScope.usd = 'N/A';
			$rootScope.shf_usd = 'N/A';
		})
	}

	this.updateCurrentBlock = function() {
	    $http({method : 'GET',url : '/api/network/lastblock'})
		.success(function(data, status) {
			$rootScope.blockNum = data.result;
		})
		.error(function(data, status) {
			$rootScope.blockNum = 'N/A';
		})
	}

	this.updateCurrentBlock();
	this.updateNetworkHashRate();
	this.updateExchangeRate();
	this.updateBlockTime();
	this.updateDifficulty();

	$interval(networkService.updateCurrentBlock, 5000);
	$interval(networkService.updateNetworkHashRate, 5000);
	$interval(networkService.updateExchangeRate, 15000);
	$interval(networkService.updateBlockTime, 5000);
	$interval(networkService.updateDifficulty, 5000);
});

angular.module('Explorer').service('BlockInfoService', function($rootScope, $http, $q) {
	this.getBlock = function(num) {
		return $http.get('/api/block/'+num).then(function(data, status) {
			return data;
		});
	}
});

angular.module('Explorer').service('TransactionInfoService', function($rootScope, $http, $q) {
	this.getTransaction = function(txn) {
		return $http.get('/api/transaction/hash/'+txn).then(function(data, status) {
			return data;
		});
	}
});

