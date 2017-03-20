angular.module('Explorer')
    .controller('blockController', function (BlockInfoService, TransactionInfoService, $rootScope, $scope, $location, $routeParams,$q) {

        $scope.transactions = [];
        $scope.init=function()
        {
            $scope.blockId=$routeParams.blockId;

            if($scope.blockId!==undefined) {
                $rootScope.title += " Block # "+$scope.blockId;
                BlockInfoService.getBlock($scope.blockId)
                    .then(function(result){
		    	result = result.data;
                    	$scope.result = result;
			$scope.blockNumber = result.number;	
			if(result.hash===undefined)
				result.hash = 'Pending';
			if(result.miner===undefined)
				result.miner = 'Pending';
		
			if($scope.blockNumber!==undefined){
				$scope.conf = $rootScope.blockNum - $scope.blockNumber + " Confirmations";
				if($scope.conf===0 + " Confirmations"){
					$scope.conf='Unconfirmed';
				}
			}

			angular.forEach(result.transactions, function(txn) {
				TransactionInfoService.getTransaction(txn).then(function(result) {
					$scope.transactions.push(result.data);
				});
			});
                });

            }
            else{
                $location.path("/");
            }
        };
        $scope.init();
});
