angular.module('Explorer')
    .controller('RecentBlocksController', function (NetworkService, $scope, $interval) {

    if(!$scope.Cursor)
        $scope.Cursor = "";
    $scope.Blocks = [];
	$scope.pageChanged = function() {
        if(!$scope.scrolling) {
            $scope.scrolling = true;
            NetworkService.getRecentBlocks(25, $scope.Cursor).then(function(res) {
                if(res && res.data) {
                    if(res.data.End)
                        $scope.Cursor = res.data.End;
                    angular.forEach(res.data.Blocks, function(block) {
                        $scope.Blocks.push(block);
                    });
                    $scope.TotalBlocks = res.data.Total;
                }
                $scope.scrolling = false;
            });
        }
	}
	
	$scope.pageChanged();
});


