#!/bin/bash

SYMBOL=$1
echo $SYMBOL
COL="exchangeRate_"$SYMBOL
TMP="exchangeRate_"$SYMBOL"_tmp"
BAK="exchangeRate_"$SYMBOL"_bak"

echo $TMP
cat << EOF
db.createCollection("$TMP", { capped: true, max: 2016, size: 500000 });
db.$COL.find().forEach(function(d) { db.$TMP.insert(d)});
db.$COL.renameCollection("$BAK", true);
db.$TMP.renameCollection("$COL", true);
EOF



