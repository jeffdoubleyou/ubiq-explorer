db.getCollectionNames().forEach(function(collection) {
    if(collection.match(/^exchange/)) {
        print(collection)
    }
});
