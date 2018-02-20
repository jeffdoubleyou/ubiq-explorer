angular.module('Explorer').controller('BlockTimeHistoryController', function (NetworkService, $scope) {
	$scope.labels = [];
	$scope.data = [];
	$scope.series = ['Block Time Evolution'];
    NetworkService.getBlockTimeHistory().then(function(res) {
        //for (var i in res.data.reverse()) {
        for (var i in res.data) {
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
					borderWidth: 0,
                    tension: .5
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
