'use strict';
var knownAddresses = {};
angular.module('Explorer', ['ngMaterial', 'md.data.table', 'ngRoute','chart.js', 'ui.bootstrap', 'infinite-scroll'])
  .config(['$routeProvider', '$httpProvider', '$locationProvider',
    function($routeProvider, $httpProvider, $locationProvider) {
        $routeProvider.
            when('/', {
                templateUrl: 'views/main.html',
                controller: 'mainController',
                title: 'Network Overview'
            }).
            when('/block/:blockId', {
                templateUrl: 'views/block.html',
                controller: 'blockController',
                title: 'Ubiq Block Information'
            }).
            when('/transaction/:transactionId', {
                templateUrl: 'views/transaction.html',
                controller: 'transactionController',
                title: 'Ubiq Transaction Information'
            }).
            when('/address/:addressId/:action?', {
                templateUrl: 'views/address.html',
                controller: 'addressController',
                title: 'Ubiq Address Information'
            }).
            when('/hashrate', {
                templateUrl: 'views/hashRateHistory.html',
                controller: 'HashRateHistoryController',
                title: 'Ubiq Hashrate Evolution'
            }).
            when('/difficulty', {
                templateUrl: 'views/difficultyHistory.html',
                controller: 'DifficultyHistoryController',
                title: 'Ubiq Difficulty Evolution'
            }).
            when('/blocktime', {
                templateUrl: 'views/blockTimeHistory.html',
                controller: 'BlockTimeHistoryController',
                title: 'Ubiq Block Time Evolution'
            }).
            when('/unclerate', {
                templateUrl: 'views/uncleRateHistory.html',
                controller: 'UncleRateHistoryController',
                title: 'Ubiq Uncle Rate Evolution'
            }).
            when('/addressinfo/summary/:addressId', {
                templateUrl: 'views/addressInfoSummary.html'
                //controller: 'addressInfoSummary'
            }).
            when('/addressinfo/incoming/:addressId', {
                templateUrl: 'views/addressInfoIncoming.html',
                controller: 'addressInfoIncoming'
            }).
            when('/addressinfo/outgoing/:addressId', {
                templateUrl: 'views/addressInfoOutgoing.html',
                controller: 'addressInfoOutgoing'
            }).
            when('/addressinfo/mining/:addressId', {
                templateUrl: 'views/addressInfoMining.html',
                controller: 'addressInfoMining'
            }).
	        when('/pool', {
		        templateUrl: 'views/pool.html',
		        controller: 'poolController'
	        }).
            when('/minerpool', {
                templateUrl: 'views/pool.html',
                controller: 'poolController',
                title: 'Ubiq Mining Pool'
            }).
            when('/richlist', {
                templateUrl: 'views/richList.html',
                controller: 'richListController',
                title: 'Ubiq Rich List'
            }).
            when('/networkpools', {
                templateUrl: 'views/networkPools.html',
                controller: 'networkPoolController',
                title: 'Ubiq Network Pools Overview'
            }).
            when('/tokens', {
                templateUrl: 'views/tokens.html',
                controller: 'tokenController',
                title: 'Ubiq Tokens'
            }).
            when('/recent_transactions', {
                templateUrl: 'views/recentTransactions.html',
                controller: 'RecentTransactionController',
                title: 'Recent Ubiq Transactions'
            }).
            when('/recent_token_transactions', {
                templateUrl: 'views/recentTokenTransactions.html',
                controller: 'RecentTokenTransactionController',
                title: 'Recent Ubiq Token Transactions'
            }).
            when('/recent_blocks', {
                templateUrl: 'views/recentBlocks.html',
                controller: 'RecentBlocksController',
                title: 'Recent Ubiq Blocks'
            }).
            when('/recent_uncles', {
                templateUrl: 'views/recentUncles.html',
                controller: 'RecentUncleController',
                title: 'Recent Ubiq Uncles'
            }).
            when('/pending_transactions', {
                templateUrl: 'views/pendingTransactions.html',
                controller: 'PendingTransactionController',
                title: 'Currently Pending Ubiq Transactions'
            }).
            when('/watched_addresses', {
                templateUrl: 'views/watchedAddresses.html',
                controller: 'mainController',
                title: 'Watched Ubiq Addresses'
            }).
            otherwise({
                redirectTo: '/'
            });
            $locationProvider.html5Mode({ enabled: true, requireBase: true }).hashPrefix('!');
        }])
        .run(function($rootScope, $location, $route, $http) {
            $rootScope.$on("$routeChangeSuccess", function(currentRoute, previousRoute){
            $rootScope.title = 'UBIQ.cc - Block Chain Explorer : '+$route.current.title;
            angular.module('infinite-scroll').value('THROTTLE_MILLISECONDS', 1000)

        }); 
});
