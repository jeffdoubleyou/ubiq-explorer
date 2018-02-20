function sortObject(items, field, reverse) {
        var filtered = [];
        angular.forEach(items, function(v,k) {
          v['key']=k;
          filtered.push(v);
        });
        filtered.sort(function (a, b) {
          return (a[field] > b[field] ? 1 : -1);
        });
        if(reverse) filtered.reverse();
        return filtered;
};
