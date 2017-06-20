angular.module('Explorer').controller('networkPoolController', function (NetworkPoolService, $rootScope, $scope, $interval) {
    $scope.init = function() {
        NetworkPoolService.getPoolStats().then(function(res) {
            console.log(res);
            $scope.pools = res.data.pools;
        });
    }

    $scope.totalHashRate = 0;
    $scope.totalMiners = 0;
    $scope.addHashRate = function(hr) {
        $scope.totalHashRate = $scope.totalHashRate + hr; 
    }
    $scope.addMiners = function(cnt) {
        $scope.totalMiners = $scope.totalMiners + cnt;
    }
    $scope.init(); 

});
