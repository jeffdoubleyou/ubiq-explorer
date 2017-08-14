angular.module('Explorer').service('AddressService', function($http) {
    this.getBalance = function(address) {
        return $http.get('/api/address/balance2/'+address);
    }

    this.getIncomingTransactions = function(address, start, end) {
        return $http.get('/api/address/transactions/to/'+address+'.'+start+'.'+end);
    }

    this.getIncomingTransactionsTotal = function(address) {
        return $http.get('/api/address/transactions/to/total/'+address);
    }

    this.getOutgoingTransactions = function(address, start, end) {
        return $http.get('/api/address/transactions/from/'+address+'.'+start+'.'+end);
    }

    this.getOutgoingTransactionsTotal = function(address) {
        return $http.get('/api/address/transactions/from/total/'+address);
    }

    this.getMinedBlocks = function(address, start, end) {
        return $http.get('/api/address/mined/'+address+'.'+start+'.'+end);
    }

    this.getMinedBlocksTotal = function(address) {
        return $http.get('/api/address/mined/blocks/'+address);
    }

    this.getBalanceHistory = function(address) {
        return $http.get('/api/address/transactions/balance/'+address);
    }

    this.getPoolAccount = function(address) {
        return $http.get('/pool/accounts/'+address);
    }

    this.getAddressName = function(address) {
        return $http.get('/api/address/name/'+address);
    }
});
