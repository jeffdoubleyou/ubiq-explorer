angular.module('Explorer')
    .controller('addressInfosCtrl', function ($rootScope, $scope, $location, $routeParams, $q) {

        
        $scope.init=function()
        {
            $scope.addressId=$routeParams.addressId;

            console.log("ADDRESS ID:", $scope.addressId);

        };

        $scope.init();

    });
        
});
