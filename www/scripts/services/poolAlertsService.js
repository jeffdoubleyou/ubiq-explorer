angular.module('Explorer').service('PoolAlertsService', function($rootScope, $interval, $http) {
    this.getAlerts = function(address) {
        return $http.get('/api/pool/getalert/'+address);
    }

    this.addAlert = function(address, email) {
        return $http.get('/api/pool/addalert/'+address+'/'+email);
    }

    this.removeAlert = function(address, email) {
        return $http.get('/api/pool/removealert/'+address+'/'+email);
    }
});
