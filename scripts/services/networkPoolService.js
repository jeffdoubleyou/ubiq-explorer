angular.module('Explorer').service('NetworkPoolService', function($rootScope, $interval, $http) {
    var networkPoolService = this;

    this.getPoolStats = function() {
        return $http.get('/api/pool/poolstats');
    }
});
