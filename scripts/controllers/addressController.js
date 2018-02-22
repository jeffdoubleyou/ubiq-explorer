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
              slug: 'transactions',
              title: "Transactions",
              content: "Transactions",
              template: "template/Address/Transactions.html"
            },
            {
              slug: 'mined',
              title: "Mined Blocks",
              content: "Mined Blocks",
              template: "template/Address/Mined.html"
            },
            {
                slug: "uncles",
                title: "Mined Uncles",
                content: "Mined Uncles",
                template: "template/Address/Uncles.html"
            }
        ];

        $scope.init=function()
        {
            $scope.balanceHistory = [];
            $scope.addressId=$routeParams.addressId;

            if($scope.addressId!==undefined) { 
                $rootScope.title += " for "+$scope.addressId;
                $scope.getBalance($scope.addressId);
                $scope.getTokenInfo($scope.addressId);
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

        $scope.getTokenInfo = function(address) {
            AddressService.getTokenInfo(address).then(function(res) {
                if(res.data)
                    $scope.tokenInfo = res.data;
            });
        };

        $scope.getBalance = function(address) {
            AddressService.getBalance(address).then(function(res) {
                if(res.data && res.data.Balances)
                    $scope.balance = res.data.Balances[0].balance;
                else
                    $scope.balance = "0";
            });
            AddressService.getTokenBalance(address).then(function(res) {
                if(res.data)
                    $scope.tokenBalance = res.data
                else
                    $scope.tokenBalance = []
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
    if(!$scope.Cursor)
        $scope.Cursor = "";
    $scope.incomingTransactions = [];
	$scope.pageChanged = function() {
        if(!$scope.scrolling) {
            $scope.scrolling = true;
            AddressService.getIncomingTransactions($scope.addressId, $scope.Cursor, 10).then(function(res) {
                if(res && res.data) {
                    if(res.data.End)
                        $scope.Cursor = res.data.End;
                    angular.forEach(res.data.Transactions, function(txn) {
                        $scope.incomingTransactions.push(txn);
                    });
                    $scope.setIncomingTransactions(res.data.Total);
                }
                $scope.scrolling = false;
            });
        }
	}
	
	$scope.pageChanged();
});

angular.module('Explorer').controller('AddressIncomingTokensController', function(AddressService, $scope) {
    if(!$scope.Cursor)
        $scope.Cursor = "";
    $scope.incomingTransactions = [];
	$scope.pageChanged = function() {
        if(!$scope.scrolling) {
            $scope.scrolling = true;
            AddressService.getIncomingTokenTransactions($scope.addressId, $scope.Cursor, 10).then(function(res) {
                if(res && res.data) {
                    if(res.data.End) {
                        $scope.Cursor = res.data.End;
                    }
                    angular.forEach(res.data.Transactions, function(txn) {
                        $scope.incomingTransactions.push(txn);
                    });
                    $scope.incomingTokenTransactionCount = res.data.Total;
                }
                $scope.scrolling = false;
            });
        }
	}
	
	$scope.pageChanged();
});


angular.module('Explorer').controller('AddressOutgoingController', function(AddressService, $scope) {
    if(!$scope.Cursor)
        $scope.Cursor = "";
    $scope.outgoingTransactions = [];
	$scope.pageChanged = function() {
        if(!$scope.scrolling) {
            $scope.scrolling = true;
            AddressService.getOutgoingTransactions($scope.addressId, $scope.Cursor, 10).then(function(res) {
                if(res && res.data) {
                    if(res.data.End)
                        $scope.Cursor = res.data.End;
                    angular.forEach(res.data.Transactions, function(txn) {
                        $scope.outgoingTransactions.push(txn);
                    });
                    $scope.setOutgoingTransactions(res.data.Total);
                }
                $scope.scrolling = false;
            });
        }
	}
	$scope.pageChanged();
});

angular.module('Explorer').controller('AddressOutgoingTokensController', function(AddressService, $scope) {
    if(!$scope.Cursor)
        $scope.Cursor = "";
    $scope.outgoingTransactions = [];
	$scope.pageChanged = function() {
        if(!$scope.scrolling) {
            $scope.scrolling = true;
            AddressService.getOutgoingTokenTransactions($scope.addressId, $scope.Cursor, 10).then(function(res) {
                if(res && res.data) {
                    if(res.data.End)
                        $scope.Cursor = res.data.End;
                    angular.forEach(res.data.Transactions, function(txn) {
                        $scope.outgoingTransactions.push(txn);
                    });
                    $scope.outgoingTokenTransactionCount = res.data.Total;
                }
                $scope.scrolling = false;
            });
        }
	}
	$scope.pageChanged();
});


angular.module('Explorer').controller('AddressMinedController', function(AddressService, $scope) {
    if(!$scope.Cursor)
        $scope.Cursor = "";
    $scope.minedBlocks = [];

   	$scope.pageChanged = function() {
        if(!$scope.scrolling) {
            $scope.scrolling = true;

            AddressService.getMinedBlocks($scope.addressId, $scope.Cursor, 10).then(function(res) {
                if(res && res.data) {
                    if(res.data.End)
                        $scope.Cursor = res.data.End;
                    angular.forEach(res.data.Blocks, function(block) {
                        $scope.minedBlocks.push(block);
                    });
                    $scope.minedBlocksCount = res.data.Total;
                }
                $scope.scrolling = false;
            });
        }
    }
    $scope.minedBlocksInit = function() {
	    $scope.pageChanged();
    }
});

angular.module('Explorer').controller('AddressUnclesController', function(AddressService, $scope) {
    if(!$scope.Cursor)
        $scope.Cursor = "";
    $scope.minedUncles = [];

   	$scope.pageChanged = function() {
        if(!$scope.scrolling) {
            $scope.scrolling = true;

            AddressService.getMinedUncles($scope.addressId, $scope.Cursor, 10).then(function(res) {
                if(res && res.data) {
                    if(res.data.End)
                        $scope.Cursor = res.data.End;
                    angular.forEach(res.data.Uncles, function(uncle) {
                        $scope.minedUncles.push(uncle);
                    });
                    $scope.minedUnclesCount = res.data.Total;
                }
                $scope.scrolling = false;
            });
        }
    }
    $scope.minedUnclesInit = function() {
	    $scope.pageChanged();
    }
});

angular.module('Explorer').controller('AddressBalanceChartController', function(AddressService, $scope) {
    AddressService.getBalanceHistory($scope.addressId).then(function(res) {
		$scope.labels = [];
		$scope.series = ['Balance Change'];
		$scope.data = [];
		for(var i in res.data.Balances.reverse()) {
			$scope.labels.push(res.data.Balances[i].changedBy);
			$scope.data.push(res.data.Balances[i].balance/1000000000000000000);
		}
        // Chart options
		$scope.options = { 
			responsive: true, 
			maintainAspectRatio: false,
			elements: {
				point: {
					radius: 3,
                    backgroundColor: '#837727'
			       },
				line: {
					borderWidth: 3,
                    //backgroundColor: '#528373',
                    borderColor: '#779988'
				}
   			},
			scales : {
				xAxes : [ {
				    gridLines : {
					    display : false
				    },
                    ticks: {
                        display: false,
                        minor: {
                            display: false
                        }
                    }
				} ]
			},
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
