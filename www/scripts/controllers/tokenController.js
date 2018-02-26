angular.module('Explorer')
    .controller('tokenController', function (NetworkService, $scope) {
    
    var tokenList = function() {
        NetworkService.getTokens().then(function(res) {
            if(res && res.data)
                $scope.Tokens = res.data.Tokens
            else
                $scope.Tokens = [];
        });
    }

    tokenList();
});
