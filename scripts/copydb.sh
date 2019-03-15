#!/bin/bash
. "scripts/shini.sh"

SRC=$1

echo "Copying block data from $SRC"

declare -A config

__shini_parsed() {
    if [ "$1" != "" ]; then
        config["$1 $2"]="$3"
    else
        config["$2"]="$3"
    fi
}

shini_parse "conf/app.conf"

checkMongo() {
    db=$(echo ${config['mongodb url']} | cut -d/ -f4)
    declare -r cmd="mongo --quiet $SRC/$db --eval='db'"
    declare -i ret
    declare output
    echo -n "Checking to see if we can connect to $SRC/$db : "
    output=$((eval $cmd) 2>&1)
    ret=$?
    if [ ${ret} != 0 ]; then
        echo "FAIL - $output"
        exit $ret
    else
        echo "OK"
    fi
}

exportExchangeData() {
    host=$(echo ${config['mongodb url']} | cut -d/ -f3)
    db=$(echo ${config['mongodb url']} | cut -d/ -f4)
    cmd="mongo --quiet ${config['mongodb url']} scripts/exchangeList.js"
    for collection in $(eval $cmd); do
        echo -n "Export $collection : ";
        dump_cmd="mongodump --quiet --host=$host --db=$db --collection=$collection --out=scripts/exchangeData"
        output=$((eval $dump_cmd) 2>&1)
        ret=$?
        if [ ${ret} != 0 ]; then
            echo "FAIL - $output"
            exit $ret
        else
            echo "OK"
        fi
    done
}

importExchangeData() {
    host=$(echo ${config['mongodb url']} | cut -d/ -f3)
    db=$(echo ${config['mongodb url']} | cut -d/ -f4)
    import_cmd="mongorestore --quiet --host=$host --db=$db scripts/exchangeData/$db"
    output=$((eval $import_cmd) 2>&1)
    ret=$?
    if [ {$ret} != 0 ]; then
        echo "FAIL - $output"
        exit $ret
    else
        echo "OK"
    fi
}

exportVerifiedTokens() {
    declare -r cmd="go run tools/exportVerifiedTokens.go"
    echo -n "Exporting verified tokens : "
    declare -i ret
    declare output
    output=$((eval $cmd) 2>&1)
    ret=$?
    if [ ${ret} != 0 ]; then
        echo "FAIL - $output"
        exit $ret
    else
        echo "OK"
    fi    
}

importVerifiedTokens() {
    declare -r cmd="sh scripts/verified_tokens.sh"
    declare -i ret
    declare output
    output=$((eval $cmd) 2>&1)
    ret=$?
    if [ ${ret} != 0 ]; then
        echo "FAIL - $output"
        exit $ret
    else
        echo "OK"
    fi
}

exportAddresses() {
    declare -r cmd="go run tools/exportAddresses.go"
    echo -n "Exporting addresses : "
    declare -i ret
    declare output
    output=$((eval $cmd) 2>&1)
    ret=$?
    if [ ${ret} != 0 ]; then
        echo "FAIL - $output"
        exit $ret
    else
        echo "OK"
    fi    
}

importAddresses() {
    declare -r cmd="sh scripts/addresses.sh"
    declare -i ret
    declare output
    output=$((eval $cmd) 2>&1)
    ret=$?
    if [ ${ret} != 0 ]; then
        echo "FAIL - $output"
        exit $ret
    else
        echo "OK"
    fi
}

exportPools() {
    declare -r cmd="go run tools/exportPools.go"
    echo -n "Exporting pools : "
    declare -i ret
    declare output
    output=$((eval $cmd) 2>&1)
    ret=$?
    if [ ${ret} != 0 ]; then
        echo "FAIL - $output"
        exit $ret
    else
        echo "OK"
    fi    
}

importPools() {
    declare -r cmd="sh scripts/pools.sh"
    declare -i ret
    declare output
    output=$((eval $cmd) 2>&1)
    ret=$?
    if [ ${ret} != 0 ]; then
        echo "FAIL - $output"
        exit $ret
    else
        echo "OK"
    fi
}

exportIndexes() {
    declare -r cmd="mongo --quiet ${config['mongodb url']} scripts/exportIndexes.js > scripts/indexes.js"
    echo -n "Export indexes from ${config['mongodb url']} : "
    declare -i ret
    declare output
    output=$((eval $cmd) 2>&1)
    ret=$?
    if [ ${ret} != 0 ]; then
        echo "FAIL - $output"
        exit $ret
    else
        echo "OK"
    fi
}

importIndexes() {
    declare -r cmd="mongo --quiet ${config['mongodb url']} scripts/indexes.js"
    echo -n "Import indexes to ${config['mongodb url']} : "
    declare -i ret
    declare output
    output=$((eval $cmd) 2>&1)
    ret=$?
    if [ ${ret} != 0 ]; then
        echo "FAIL - $output"
        exit $ret
    else
        echo "OK"
    fi
}

dropDatabase() {
    declare -r cmd="echo mongo --quiet ${config['mongodb url']} --eval=\"db.dropDatabase()\""
    echo -n "Dropping database from ${config['mongodb url']} : "
    declare -i ret
    declare output
    output=$((eval $cmd) 2>&1)
    ret=$?
    if [ ${ret} != 0 ]; then
        echo "FAIL - $output"
        exit $ret
    else
        echo "OK"
    fi
}

importBlockData() {
    host=$(echo ${config['mongodb url']} | cut -d/ -f3)
    db=$(echo ${config['mongodb url']} | cut -d/ -f4)
    declare -r cmd="mongodump --quiet --db=$db --host=$SRC --archive | mongorestore --quiet --db=$db --host=$host --archive"
    declare -i ret
    decoare output
    output=$((eval $cmd) 2>&1)
    ret=$?
    if [ ${ret} != 0 ]; then
        echo "FAIL - $output"
        exit $ret
    else
        echo "OK"
    fi
}

commitBackups() {
    declare -r cmd="git commit scripts/addresses.sh scripts/pools.sh scripts/verifiedTokens.sh -m 'Update backups during import of new block DB'"
    echo -n "Committing changes to backup scripts: "
    declare -i ret
    declare output
    output=$((eval $cmd) 2>&1)
    ret=$?
    if [ ${ret} != 0 ]; then
        echo "FAIL - $output"
        exit $ret
    else
        echo "OK"
    fi
}

checkMongo
exportExchangeData
exportIndexes
exportAddresses
exportPools
exportVerifiedTokens
commitBackups
exit

dropDatabase
importBlockData
importExchangeData
importAddresses
importPools
importVerifiedTokens
importIndexes

