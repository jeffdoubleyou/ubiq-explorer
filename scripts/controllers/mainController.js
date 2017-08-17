angular.module('Explorer')
    .controller('mainController', function (SearchService, NetworkService, $scope, $rootScope, $interval) {
        $scope.search=function() {
            var searchString = $scope.searchString;
            SearchService.routeSearch(searchString);
        }
	var updateRecentBlocks = function() {
        NetworkService.getRecentBlocks().then(function(res) {
            $scope.recentBlocks = res.data;
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
			for (var i in data) {
				labels.push(data[i].name + ' ' + parseFloat(data[i].percent).toFixed(2) + '%' + ' ('+data[i].count+')');
				values.push(data[i].percent);
			}
			$scope.labels = labels;
			$scope.data = values;
			$scope.options = { legend: { display: true, position: 'bottom', labels: {  } } };
		});
	}

	updateTopMiners();
});

