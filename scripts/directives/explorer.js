
function truncate(input, max) {
        input = input || "";
        max = max || 0;
        var out;
        if(input.length > max && max > 0)
        {
            out = input.substring(0, max/2-2) + '....' + input.substring(input.length-(max/2-2), input.length);
        }
        else {
            out = input;
        }
        return out;
}

angular.module('Explorer').directive('formatAddress', function(AddressService) {
    return {
        template: "<span>{{name}}</span>",
        scope: {
            addressId: "=",
            maxLength: "="
        },
        link: function(scope) {
            if(knownAddresses["_miner_"+scope.addressId]) {
                scope.name = truncate(knownAddresses["_miner_"+scope.addressId], scope.maxLength);
            }
            else { 
                console.log("Get name for ", scope.addressId);
                AddressService.getAddressName(scope.addressId).then(function(res) {
                    if(res && res.data && res.data.name) {
                        knownAddresses["_miner_"+scope.addressId] = res.data.name;
                        scope.name = truncate(res.data.name, scope.maxLength);
                    }
                    else {
                        scope.name = truncate(scope.addressId, scope.maxLength);
                        knownAddresses["_miner_"+scope.addressId] = scope.addressId;
                    }
                });
            }
        }
    };
});


