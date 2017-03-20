angular.module('Explorer').service('PoolStatsService', function($rootScope, $interval, $http) {
	this.getPoolStats = function() {
        console.log("GET Pool stats");
	    $http({method : 'GET',url : 'http://www.ubiq.cc/pool/stats'})
		.success(function(data, status) {
			var blocks = parseInt(data.immatureTotal)+parseInt(data.candidatesTotal);
			$rootScope.$emit('poolBlocks', {blocks: blocks, immature: data.immatureTotal, candidates: data.candidatesTotal, total: data.maturedTotal });
			$rootScope.$emit('activeMiners', data.minersTotal);
			$rootScope.poolMiners = data.minersTotal;
			$rootScope.poolHashrate = parseFloat(data.hashrate/1000/1000).toFixed(2) + ' MH';
			$rootScope.poolDifficulty = parseFloat(data.nodes[0].difficulty/1000/1000/1000).toFixed(2) + ' G';
			$rootScope.poolLastBlock = data.stats.lastBlockFound;
			$rootScope.poolRoundShares = data.stats.roundShares;
			$rootScope.poolRoundDifficulty = parseFloat(data.nodes[0].difficulty);
			$rootScope.poolRoundVariance = $rootScope.poolRoundShares ? parseInt($rootScope.poolRoundShares/$rootScope.poolRoundDifficulty*100) : 0;
			var epoch = (30000 - ($rootScope.blockNum % 30000)) * 1000 * $rootScope.networkBlockTime;
			$rootScope.poolEpochSwitch = (Date.now() / 1000)+(epoch/1000);
		})
		.error(function(data, status) {
			$rootScope.$emit('Notification', 'Pool API is temporarily down - Mining is not affected by this');
		})
	}

    this.getHashRateHistory = function() {
        $http({method : 'GET', url : 'http://www.ubiq.cc/api/pool/hashrateHistory'})
            .success(function(data, status) {
                $rootScope.$emit('poolHashRateHistory', data.reverse());
            })
            .error(function(data, status) {
                console.log("Error getting pool hashrate history", status);
            })
    }

	this.DataLoader = function(section) {
		var dataLoaderURL = '/pool/'+section;
		$http({method: 'GET', url: dataLoaderURL})
			.success(function(data, status) {
				var serviceName = 'pool' + section[0].toUpperCase() + section.substring(1) + 'Result';
				console.log("Service name ", serviceName);
				$rootScope.$emit(serviceName, data);
			})
			.error(function(data, status) {
				$rootScope.$emit('Notification', 'Pool API is temporarily down - Mining is not affected by this');
			})
	}

	this.getPoolStats();
    this.getHashRateHistory();

	var poolStatsService = this;
	$interval(poolStatsService.getPoolStats, 5000);
    $interval(poolStatsService.getHashRateHistory, 15000);

});

