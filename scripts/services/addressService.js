angular.module('Explorer').service('AddressService', function($http) {
    this.getBalance = function(address) {
        return $http.get('/api/v1/balance/get?address='+address);
    }

    this.getIncomingTransactions = function(address, start, limit) {
        return $http.get('/api/v1/transaction/to?address='+address+'&cursor='+start+'&limit='+limit);
    }

    this.getOutgoingTransactions = function(address, start, limit) {
        return $http.get('/api/v1/transaction/from?address='+address+'&cursor='+start+'&limit='+limit);
    }

    this.getMinedBlocks = function(address, start, limit) {
        return $http.get('/api/v1/block/miner?address='+address+'&cursor='+start+'&limit='+limit)
    }

    this.getMinedUncles = function(address, start, limit) {
        return $http.get('/api/v1/uncle/miner?address='+address+'&cursor='+start+'&limit='+limit)
    }

    this.getBalanceHistory = function(address) {
        return $http.get('/api/v1/balance/history?limit=100&address='+address);
    }

    this.getPoolAccount = function(address) {
        return $http.get('/pool/accounts/'+address);
    }

    this.getTokenBalance = function(address) {
        return $http.get('/api/v1/token/balance?address='+address);
    }

    this.getTokenInfo = function(address) {
        return $http.get('/api/v1/token/address?address='+address);
    }

    this.getIncomingTokenTransactions = function(address, cursor, limit) {
        return $http.get('api/v1/token/to?address='+address+'&cursor='+cursor+'&limit='+limit);
    }

    this.getOutgoingTokenTransactions = function(address, cursor, limit) {
        return $http.get('api/v1/token/from?address='+address+'&cursor='+cursor+'&limit='+limit);
    }
});
