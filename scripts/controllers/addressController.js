angular.module('Explorer').controller('addressController', function (AddressService, $rootScope, $scope, $location, $routeParams) {

	 $scope.isPoolAccount = false;
	 $scope.action = $routeParams.action;
	 $scope.tabs = [
            {
              slug: 'summary',
              title: "Summary",
              content: "Address Summary",
              template: "template/Address/Summary.html"
            }, 
            {
              slug: 'incoming',
              title: "Incoming Transactions",
              content: "Incoming Transactions",
              template: "template/Address/Incoming.html"
            }, 
            {
              slug: 'outgoing',
              title: "Outgoing Transactions",
              content: "Outgoing Transactions",
              template: "template/Address/Outgoing.html"
            },
            {
              slug: 'mined',
              title: "Mined Blocks",
              content: "Mined Blocks",
              template: "template/Address/Mined.html"
            }
        ];

        $scope.init=function()
        {
            $scope.balanceHistory = [];
            $scope.addressId=$routeParams.addressId;

            if($scope.addressId!==undefined) { 
                $rootScope.title += " for "+$scope.addressId;
                $scope.getBalance($scope.addressId);
            }
            else {
                $location.path("/");
            }

            $scope.onTabSelected = function(route) {
                   console.log("Tab select ", route+'/'+$scope.addressId);
            }
        };

        $scope.setPoolAccount = function() {
            $scope.isPoolAccount = true;
        }

        $scope.getBalance = function(address) {
            AddressService.getBalance(address).then(function(res) {
                if(res.data && res.data.balance)
                    $scope.balance = res.data.balance;
                else
                    $scope.balance = "N/A";
            });
        }

        $scope.setIncomingTransactions = function(cnt) {
            $scope.incomingTransactionCount = cnt;
        }

        $scope.setOutgoingTransactions = function(cnt) {
            $scope.outgoingTransactionCount = cnt;
        }

        $scope.init();
    });

angular.module('Explorer').controller('AddressIncomingController', function(AddressService, $scope) {
	$scope.currentPage = 1;
	$scope.maxSize = 10;
	$scope.displayLimit = 10;

	$scope.pageChanged = function() {
		var page = $scope.currentPage;
		page = page - 1;
		var start = page * $scope.displayLimit;
		var end = start + $scope.displayLimit - 1;

        AddressService.getIncomingTransactions($scope.addressId, start, end).then(function(res) {
            if(res && res.data)
                $scope.incomingTransactions = res.data;
            else
                $scope.incomingTransactions = [];
        });
	}
	
	$scope.transactionsInit = function () {
        AddressService.getIncomingTransactionsTotal($scope.addressId).then(function(res) {
            if(res && res.data) {
                $scope.totalTransactions = res.data.transactions;
                $scope.setIncomingTransactions(res.data.transactions);
                if(res.data.transactions == 0)
                    $scope.pages = 0;
                else
                    $scope.pages = Math.floor(res.data.transactions/100) + 1;
            }
            else {
                $scope.pages = 0;           
            }
        });
	}

	$scope.pageChanged();
});

angular.module('Explorer').controller('AddressOutgoingController', function(AddressService, $scope) {
	$scope.currentPage = 1;
	$scope.maxSize = 10;
	$scope.displayLimit = 10;

	$scope.pageChanged = function() {
		var page = $scope.currentPage;
		page = page - 1;
		var start = page * $scope.displayLimit;
		var end = start + $scope.displayLimit - 1;

        AddressService.getOutgoingTransactions($scope.addressId, start, end).then(function(res) {
            if(res && res.data)
                $scope.outgoingTransactions = res.data;
            else
                $scope.outgoingTransactions = [];
        });
	}
	
	$scope.transactionsInit = function () {
        AddressService.getOutgoingTransactionsTotal($scope.addressId).then(function(res) {
            if(res && res.data) {
                $scope.totalTransactions = res.data.transactions;
                $scope.setOutgoingTransactions(res.data.transactions);
                if(res.data.transactions == 0)
                    $scope.pages = 0;
                else
                    $scope.pages = Math.floor(res.data.transactions/100) + 1;
            }
            else {
                $scope.pages = 0;
            }
        });
	}

	$scope.pageChanged();
});

angular.module('Explorer').controller('AddressMinedController', function(AddressService, $scope) {
	$scope.currentPage = 1;
	$scope.maxSize = 10;
	$scope.displayLimit = 10;

   	$scope.pageChanged = function() {
		var page = $scope.currentPage;
		page = page - 1;
		var start = page * $scope.displayLimit;
		var end = start + $scope.displayLimit - 1;
        AddressService.getMinedBlocks($scope.addressId, start, end).then(function(res) {
            if(res && res.data)
                $scope.minedBlocks = res.data;
            else
                $scope.minedBlocks = [];
        });
    }
	
	$scope.minedBlocksInit = function () {
        AddressService.getMinedBlocksTotal($scope.addressId).then(function(res) {
            if(res && res.data) {
                $scope.totalBlocks = res.data.blocks;
                if(res.data.blocks == 0)
                    $scope.pages = 0;
                else
                    $scope.pages = Math.floor(res.data.blocks/100) + 1;
            }
            else {
                $scope.pages = 0;
            }
        });
	}

	$scope.pageChanged();
});

angular.module('Explorer').controller('AddressBalanceChartController', function(AddressService, $scope) {
    AddressService.getBalanceHistory($scope.addressId).then(function(res) {
		$scope.labels = [];
		$scope.series = ['Balance Change'];
		$scope.data = [];
		for(var i in res.data) {
			$scope.labels.push("");
			$scope.data.push(res.data[i]);
		}
        // Chart options
		$scope.options = { 
			responsive: true, 
			maintainAspectRatio: false,
			elements: {
				point: {
					radius: 0
			       },
				line: {
					borderWidth: 6
				}
   			},
			scales : {
				xAxes : [ {
				    gridLines : {
					display : false
				    }
				} ]
			}
		};
	});
});

angular.module('Explorer').controller('AddressTabsController', function(AddressService, PoolAlertsService, $scope, $interval, $timeout, $injector) {
	var poolAddressData = function() {
        AddressService.getPoolAccount($scope.addressId).then(function(res) {
            $scope.minerData = res.data;
            if($scope.poolRoundShares) {
                $scope.roundShare = parseInt(res.data.roundShares/$scope.poolRoundShares*100);
            }
            else {
                $scope.roundShare = 0;
            }
            if(!$scope.isPoolAccount) {
                var tabIndex = $scope.tabs.push({
                    slug: 'pool',
                    title: "Pool Stats",
                    content: "Pool Stats",
                    template: "template/Address/Pool.html",
                });
                $scope.setPoolAccount();
                // Inject pool stats mostly just for calculating round share
                var poolService=$injector.get('PoolStatsService');
                PoolAlertsService.getAlerts($scope.addressId).success(function(alerts, status) {
                    $scope.minerAlerts = alerts.alerts;
                })
                .error(function(err, status) {
                    console.log("Failed to get alerts", err, status);
                });

                // Select mining tab if we followed a mining link
                if($scope.action == 'pool') {
                    $timeout(function() {
                        $scope.ActiveAddressTab = tabIndex-1;
                    });
                }
	        }
		})
	};

    $scope.removeAlert = function(address, email, idx) {
        console.log("Remove alert", address, email, idx);
        PoolAlertsService.removeAlert(address, email).then(function(res, err) {
            if(res && res.data) {
                console.log("Removed", res.data);
                $scope.minerAlerts.splice(idx, 1);
                $rootScope.$emit("Removed");
            }
            else {
                console.log("Failed to remove alert", err);
                alert("Unable to remove alert, please try again");
            }
        });
    }

    $scope.addAlert = function(address) {
        var email = $('#alertEmail').val();
        console.log("Add alert", address, email);
        PoolAlertsService.addAlert(address, email).then(function(res, err) {
            if(res.data && res.data.status == 1) {
                if(!$scope.minerAlerts)
                    $scope.minerAlerts = [];
                $scope.minerAlerts.unshift(email);
                $('#alertEmail').val("");
            }
            else {
                alert("Failed to add alert. "+data.error);
            }
        });
    }
	poolAddressData();
	$scope.poolAddressInterval = $interval(poolAddressData, 5000);
});
