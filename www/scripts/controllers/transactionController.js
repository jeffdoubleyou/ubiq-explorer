angular.module('Explorer')
    .controller('transactionController', function (TransactionInfoService, BlockInfoService, $rootScope, $scope, $location, $routeParams) {

        $scope.init=function()
        {
            $scope.transactionId=$routeParams.transactionId;
            if($scope.transactionId!==undefined) {
                $rootScope.title += " for "+$scope.transactionId;
                TransactionInfoService.getTransaction($scope.transactionId).then(function(res) {
                    $scope.txn = res.data;
                    $scope.txn.txprice = (res.data.gas*res.data.gasPrice);
                    BlockInfoService.getBlock(res.data.blockNumber).then(function(block) {
                        $scope.blockInfo = block.data;
                    });
                });
            }
            else{
                $location.path("/"); 
            }
        };
        $scope.init();


});
