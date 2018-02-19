angular.module('Explorer')
    .controller('blockController', function (BlockInfoService, TransactionInfoService, $rootScope, $scope, $location, $routeParams,$q) {

        $scope.transactions = [];
        $scope.init=function()
        {
            $scope.blockId=$routeParams.blockId;
            $scope.uncles = [];
            if($scope.blockId!==undefined) {
                $rootScope.title += " Block # "+$scope.blockId;
                BlockInfoService.getBlock($scope.blockId).then(function(result){
		    var transactions = [];
                    result = result.data;
                    $scope.result = result;
                    $scope.blockNumber = result.number;	
                    if(result.hash===undefined)
                        result.hash = 'Pending';
                    if(result.miner===undefined)
                        result.miner = 'Pending';
                
                    if($scope.blockNumber!==undefined){
                        $scope.confirms = $rootScope.blockNum - $scope.blockNumber + " Confirmations";
                        if($scope.confirms===0 + " Confirmations"){
                            $scope.confirms='Unconfirmed';
                        }
                        BlockInfoService.getUncles($scope.blockNumber).then(function(result) {
                            if(result && result.data && result.data.Uncles) {
                                $scope.uncles = result.data.Uncles;
                            }
                        });
                        BlockInfoService.getTransactions($scope.blockNumber).then(function(result) {
                            if(result && result.data && result.data.Transactions) {
                                $scope.transactions = result.data.Transactions;
                            }
                        });
                    }
                    else {
                        $scope.confirms = 'Pending';
                    }
                });
            }
            else{
                $location.path("/");
            }
        };
        $scope.init();
});
