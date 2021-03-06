angular.module('Explorer').controller('ExchangeController', function (ExchangeService, $scope, $routeParams, $rootScope) {
	$scope.symbol = $routeParams.symbol;
	$scope.base = $routeParams.base;
	$scope.labels = [];
	$scope.data = [];
	$scope.series = ['Price History for '+$scope.symbol];
	$scope.baseCurrency = 'BTC';
        $scope.colors = [ '#0ca579', '#00ea90', '#333333' ];
	if($scope.symbol == 'BTC') {
		$scope.baseCurrency = 'USD';
	}
	$scope.basePrice = ($scope.base == 'USD') ? $rootScope.exchangeRates["BTC"].usd : $rootScope.exchangeRates[$scope.base].btc;

	ExchangeService.getExchangeHistory($scope.symbol).then(function(res) {
		for(var i in res.data.reverse()) {
			var d = new Date(res.data[i].timestamp*1000);
			$scope.labels.push(d.toLocaleString());
			var rate = 0;
			if(res.data[i] && res.data[i].btc > 0) {
				rate = res.data[i].btc;
			} else {
				if($scope.symbol == 'BTC')
					rate = res.data[i].usd;
				else
					rate = res.data[i].usd*$rootScope.exchangeRates["BTC"].btc;
			}
			if($scope.base == $scope.baseCurrency) {
				$scope.data.data.push(rate);
			} else {
				if($scope.base == "USD") {
					$scope.data.push(rate*$scope.basePrice);
				} else {
					$scope.data.push(rate/$scope.basePrice);
				}
			}
		}
		$scope.source = res.data[0].source;
		$scope.price = res.data[res.data.length-1][$scope.baseCurrency.toLowerCase()];
		if($scope.base != $scope.baseCurrency) {
			if($scope.base == "USD") {
				$scope.price = $scope.price*$scope.basePrice;
			} else {
				$scope.price = $scope.price/$scope.basePrice;
			}
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
