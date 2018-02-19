angular.module('Explorer').controller('DifficultyHistoryController', function (NetworkService, $scope) {
	$scope.labels = [];
	$scope.data = [];
	$scope.series = ['Hashrate Evolution'];

    NetworkService.getDifficultyHistory().then(function(res) {
        for (var i in res.data.reverse()) {
            $scope.labels.push("");
            $scope.data.push(res.data[i].value)
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
					borderWidth: 1,
                    borderJoinStyle: 'round',
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
