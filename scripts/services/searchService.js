angular.module('Explorer').service('SearchService', function($location) {
    this.routeSearch = function(searchString) {
        searchString = searchString.toLowerCase();
        var Transaction = /[0-9a-zA-Z]{64}?/;
        var Address =  /(0x)?[a-zA-Z0-9]{40}/; 
        var Block = /[0-9]{1,7}?/;

        if(Transaction.test(searchString)) {
            $location.path('/transaction/'+searchString);
        }
        else if(Address.test(searchString)) {
            $location.path('/address/'+searchString);
        }
        else if(Block.test(searchString)) {
            $location.path('/block/'+searchString);
        }
    }
});


