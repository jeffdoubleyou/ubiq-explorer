angular.module('Explorer').service('RichListService', function($rootScope, $interval, $http) {
    this.getRichList = function() {
        return $http.get('/api/v1/balance/richList');
    }
});
