angular.module('Explorer').controller('HashRateHistoryController', function (NetworkService, $scope) {
	$scope.labels = [];
	$scope.data = [];
	$scope.series = ['Hashrate Evolution'];
    NetworkService.getHashRateHistory().then(function(res) {
        for (var i in res.data.reverse()) {
            $scope.labels.push("");
            $scope.data.push(res.data[i].value/100000000)
        }
    });

	$scope.options = { 
			responsive: true, 
			maintainAspectRatio: true,
			elements: {
				point: {
					radius: 0
			       },
				line: {
					borderWidth: 0,
                    tension: 1
				}
   			},
			scales : {
				xAxes : [ {
				    gridLines : {
					display : false
				    }
				} ]
			}
		};
});
