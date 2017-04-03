angular.module('Explorer').controller('networkPoolController', function (NetworkPoolService, $rootScope, $scope, $interval) {
    $scope.init = function() {
        NetworkPoolService.getPoolStats().then(function(res) {
            console.log(res);
            $scope.pools = res.data.pools;
        });
    }

    $scope.init(); 

});
