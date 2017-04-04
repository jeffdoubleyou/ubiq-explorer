angular.module('Explorer').controller('poolController', function (PoolStatsService, $rootScope, $scope, $interval) {
    $scope.options = {
            maintainAspectRatio: false,
            elements: {
                point: {
                    radius: 0
                   },
                line: {
                    borderWidth: 6,
                    tension: 1
                },
            },
            scales : {
                xAxes : [ {
                    gridLines : {
                        display : false
                    }
                }],
                yAxes : [ {
                    gridLines: {
                        display: false
                    }
                }]
            },
    };

    $rootScope.$on('poolHashRateHistory', function(event, hashrate) {
        if(!$scope.hashRateHistoryLabels)
        {
            $scope.hashRateHistoryLabels = [];
            for (var i in hashrate) {
                $scope.hashRateHistoryLabels.push("");
            }
        }
		$scope.hashRateHistoryData = hashrate;
    });

	$rootScope.$on('poolBlocks', function(event, blocks) {
		$scope.poolBlocks = blocks.blocks;
		$scope.candidates = blocks.candidates;
		$scope.immature = blocks.immature;
		$scope.totalBlocks = blocks.total;
	});
	$rootScope.$on('activeMiners', function(event, miners) {
		$scope.activeMiners = miners;
	});
	$scope.dataLoaderInterval;
	$scope.dataLoaderURL;

	$scope.PoolTabSelect = function(tab) {
		if($scope.dataLoaderInterval) {
			$interval.cancel($scope.dataLoaderInterval);
		}
		if(tab == 'blocks' || tab == 'payments' || tab == 'miners') {
			$scope.dataLoaderInterval = $scope.startDataLoader(tab);
		}
	}

	$scope.startDataLoader = function(section) {
		PoolStatsService.DataLoader(section);
		return $interval(PoolStatsService.DataLoader, 5000, 0, true, section);
	}
});

angular.module('Explorer').controller('poolIndexController', function(PoolStatsService) {

});

angular.module('Explorer').controller('poolBlocksController', function($rootScope, $scope) {
	$rootScope.$on('poolBlocksResult', function(event, result) {
        $scope.luck = result.luck;
		$scope.matured = result.matured;
		$scope.maturedTotal = result.maturedTotal;
		$scope.immature = result.immature;
		$scope.immatureTotal = result.immatureTotal;
		$scope.candidates = result.candidates;
		$scope.candidatesTotal = result.candidatesTotal;
	});
});


angular.module('Explorer').controller('poolMinersController', function($rootScope, $scope) {
	$rootScope.$on('poolMinersResult', function(event, result) {
		$scope.minersHashRate = result.hashrate;
		$scope.minersList = result.miners;
		$scope.minersTotal = result.minersTotal;
	});
});

angular.module('Explorer').controller('poolPaymentsController', function($rootScope, $scope) {
	$rootScope.$on('poolPaymentsResult', function(event, result) {
		$scope.paymentsTotal = result.paymentsTotal;
		$scope.payments = result.payments;
	});
});

