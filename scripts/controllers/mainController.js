angular.module('Explorer')
    .controller('mainController', function (SearchService, NetworkService, $scope, $rootScope, $interval) {
        $scope.search=function() {
            var searchString = $scope.searchString;
            SearchService.routeSearch(searchString);
        }
	var updateRecentBlocks = function() {
        NetworkService.getRecentBlocks().then(function(res) {
            $scope.recentBlocks = res.data;
            if(res.data && res.data.End > $rootScope.blockNum)
                $rootScope.blockNum = res.data.End
        });
	}

	updateRecentBlocks();

	$interval(updateRecentBlocks, 15000);

	var updateRecentTransactions = function() {
        NetworkService.getRecentTransactions().then(function(res) {
			$scope.recentTransactions = res.data;
        });
	}

	updateRecentTransactions();
	$interval(updateRecentTransactions, 15000);

	var updateTopMiners = function() {
		var labels = [];
		var values = [];
       
        NetworkService.getTopMiners().then(function(res) { 
            var data = res.data || [];
            angular.forEach(res.data, function(data) {
				labels.push(data.name + ' ' + parseFloat(data.percent).toFixed(2) + '%' + ' ('+data.count+')');
                values.push(data.percent)
            });

			$scope.labels = labels;
			$scope.data = values;
			$scope.options = { legend: { display: true, position: 'bottom', labels: {  } } };
		});
	}

	updateTopMiners();

    if(!$rootScope.knownAddreses) {
        $rootScope.knownAddresses = [];
        NetworkService.getKnownAddresses().then(function(res) {
            angular.forEach(res.data, function(a) {
                $rootScope.knownAddresses["_"+a.address] = a.name
            });
        });
    }
});

