angular.module('Explorer')
    .controller('RecentTransactionController', function (NetworkService, $scope, $rootScope, $interval) {


    if(!$scope.Cursor)
        $scope.Cursor = "";
    $scope.Transactions = [];
	$scope.pageChanged = function() {
        if(!$scope.scrolling) {
            $scope.scrolling = true;
            NetworkService.getRecentTransactions(25, $scope.Cursor).then(function(res) {
                if(res && res.data) {
                    if(res.data.End)
                        $scope.Cursor = res.data.End;
                    angular.forEach(res.data.Transactions, function(txn) {
                        $scope.Transactions.push(txn);
                    });
                    $scope.TotalTransactions = res.data.Total;
                }
                $scope.scrolling = false;
            });
        }
	}
	
	$scope.pageChanged();
});

angular.module('Explorer')
    .controller('RecentTokenTransactionController', function (NetworkService, $scope, $rootScope, $interval) {


    if(!$scope.Cursor)
        $scope.Cursor = "";
    $scope.Transactions = [];
	$scope.pageChanged = function() {
        if(!$scope.scrolling) {
            $scope.scrolling = true;
            NetworkService.getRecentTokenTransactions(25, $scope.Cursor).then(function(res) {
                if(res && res.data) {
                    if(res.data.End)
                        $scope.Cursor = res.data.End;
                    angular.forEach(res.data.Transactions, function(txn) {
                        $scope.Transactions.push(txn);
                    });
                    $scope.TotalTransactions = res.data.Total;
                }
                $scope.scrolling = false;
            });
        }
	}
	
	$scope.pageChanged();
});
