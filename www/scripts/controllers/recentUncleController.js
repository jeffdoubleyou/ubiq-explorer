angular.module('Explorer')
    .controller('RecentUncleController', function (NetworkService, $scope, $interval) {

    if(!$scope.Cursor)
        $scope.Cursor = "";
    $scope.Uncles = [];
	$scope.pageChanged = function() {
        if(!$scope.scrolling) {
            $scope.scrolling = true;
            NetworkService.getRecentUncles(25, $scope.Cursor).then(function(res) {
                if(res && res.data) {
                    if(res.data.End)
                        $scope.Cursor = res.data.End;
                    angular.forEach(res.data.Uncles, function(uncle) {
                        $scope.Uncles.push(uncle);
                    });
                    $scope.TotalUncles = res.data.Total;
                }
                $scope.scrolling = false;
            });
        }
	}
	
	$scope.pageChanged();
});


