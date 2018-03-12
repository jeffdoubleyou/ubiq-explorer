angular.module('Explorer')
    .controller('PendingTransactionController', function (NetworkService, $scope, $rootScope, $interval) {

    $scope.updatePendingTransactions = function() {
        NetworkService.getPendingTransactions().then(function(res) {
            if(res && res.data) {
                $scope.Transactions = res.data;
            }
        });
    }
    $scope.updatePendingTransactions();

	
});


