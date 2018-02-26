# ubiq-explorer
Ubiq.cc - Block Chain Explorer and Pool

This is the API and front end for the Ubiq block explorer at www.ubiq.cc

The API backend is written in golang and uses the beego framework.

## API Requirements:

* Go 1.7+ ( only 1.8 has been tested )
* Beego 1.9
* Beego Bee tool
* Ubiq go node

## UI Requirements

* node / npm
* nginx

## Installing

You need to first install all dependencies and verify that they are working.  I have only ever installed on Ubuntu and there is an install script for Ubuntu 16.04 that will build and install the explorer for you.

If you are using Ubuntu 16.04, simply run:

```
sh scripts/build.sh
```

Assuming that you have dependencies installed, this should build and install everything for you.

If you are not using Ubuntu 16.04, you can manually install everything:

1. Build the block daemon and pool monitor
```
go build -o bin/blockdaemon daemon/blockdaemon.go
go build -o bin/poolstats daemon/pool.go
```
2. Build the UI
```
cd www
npm install
```
3. Package everything
```
bee pack
```
4. Unpack to wherever you'd like to install
```
tar -C /destination/directory -xf ubiq-explorer.tar.gz
```
5. Add scripts/nginx.conf to your NGINX config
6. Add scritps/cron to /etc/cron.d/

## Configuration

* The configuration file can be found in conf/app.conf

Here are the parameters that you'll likely be interested in:

### General settings

* httpport : The port that the API server listens on.  Your NGINX needs to proxy here.
* runmode : Beego run mode [ dev / production ]
* EnableAdmin : Enable the Beego admin interface
* AdminPort : The port for the Beego admin interface

### Mongodb

* url : Mongo DB URL
* max_connections : Max simultaneous, separate MongoDB connections

### Stats

* smooting : This is the number of blocks used to average out current hash rate, block time, difficulty & uncle rate to make charts smoother
* history_window : The number of blocks to keep in the history graphs.  Changing this number requires a drop of the Mongo DB or manually altering the cap on the collections
* miner_window : The number of blocks to show in the miner graph
* use_target_block_time : Use the target block time when calculating hash rate.  This results in a very smooth graph, but I don't believe it's very accurate w/ the difficulty algo of Ubiq
* target_block_time : The target block time to use when use_target_block_time = 1

## Running

* Nginx needs to be running.

### Ubuntu auto install from script

1. Start your wallet
```
sudo service wallet start
```
2. Start the block daemon
```
sudo service blockdaemon start
```
3. Start API server
```
sudo service ubiq-api start
```

### Manual install

1. Start your wallet
2. Start the block daemon
```
cd /destination/directory
./bin/blockdaemon
```
3. Start API server
```
cd /destination/directory
./ubiq-explorer
```

## Other notes

You can add the default Ubiq pool list by running:
```
sh tests/pools.sh
```

You can add the default address list ( those that display names in the UI instead of just address ):
```
sh tests/addresses.sh
```

There are a lot of things that the blockdaemon stores in the DB, so building the DB from scratch takes a very long time ( a couple of days probably ).

