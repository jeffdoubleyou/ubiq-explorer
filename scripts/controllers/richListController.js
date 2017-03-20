angular.module('Explorer').controller('richListController', function (RichListService, $scope) {
    $scope.init = function() {
        RichListService.getRichList().then(function(res) {
            $scope.richList = res.data;
        });
    }
});
