angular.module('Explorer').controller('UncleRateHistoryController', function (NetworkService, $scope) {
	$scope.labels = [];
	$scope.data = [];
	$scope.series = ['Unclerate Evolution'];
    NetworkService.getUncleRateHistory().then(function(res) {
        for (var i in res.data.reverse()) {
            $scope.labels.push("");
            $scope.data.push(res.data[i].value/10);
        }
    });
    $scope.colors = [ '#0ca579', '#00ea90', '#333333' ];
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
