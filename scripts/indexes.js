db.runCommand({ createIndexes: "addresses", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.addresses"
	}
]
});
db.runCommand({ createIndexes: "balances", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.balances"
	},
	{
		"v" : 2,
		"key" : {
			"address" : 1,
			"_id" : -1
		},
		"name" : "address_1__id_-1",
		"ns" : "ubiq-explorer.balances"
	}
]
});
db.runCommand({ createIndexes: "blockTimeHistory", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.blockTimeHistory"
	}
]
});
db.runCommand({ createIndexes: "currentBalance", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.currentBalance"
	}
]
});
db.runCommand({ createIndexes: "difficultyHistory", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.difficultyHistory"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_APX", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_APX"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_APX_bak", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_APX_bak"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_BTC", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_BTC"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_BTC_bak", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_BTC_bak"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_CEFS", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_CEFS"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_CEFS_bak", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_CEFS_bak"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_DOT", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_DOT"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_DOT_bak", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_DOT_bak"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_GEO", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_GEO"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_GEO_bak", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_GEO_bak"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_MAR", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_MAR"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_QWARK", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_QWARK"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_QWARK_bak", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_QWARK_bak"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_RICKS", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_RICKS"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_RICKS_bak", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_RICKS_bak"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_SPC", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_SPC"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_UBQ", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_UBQ"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_UBQ_bak", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_UBQ_bak"
	}
]
});
db.runCommand({ createIndexes: "exchangeRate_XCT", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeRate_XCT"
	}
]
});
db.runCommand({ createIndexes: "exchangeSource", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchangeSource"
	}
]
});
db.runCommand({ createIndexes: "exchnageRate_UBQ_tmp", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.exchnageRate_UBQ_tmp"
	}
]
});
db.runCommand({ createIndexes: "hashRateHistory", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.hashRateHistory"
	}
]
});
db.runCommand({ createIndexes: "minedBlocks", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.minedBlocks"
	},
	{
		"v" : 2,
		"key" : {
			"miner" : 1,
			"block" : -1
		},
		"name" : "miner_1_block_-1",
		"ns" : "ubiq-explorer.minedBlocks"
	},
	{
		"v" : 2,
		"key" : {
			"block" : -1
		},
		"name" : "block_-1",
		"ns" : "ubiq-explorer.minedBlocks"
	}
]
});
db.runCommand({ createIndexes: "minedUncles", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.minedUncles"
	},
	{
		"v" : 2,
		"key" : {
			"miner" : 1,
			"block" : -1
		},
		"name" : "miner_1_block_-1",
		"ns" : "ubiq-explorer.minedUncles"
	}
]
});
db.runCommand({ createIndexes: "pools", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.pools"
	}
]
});
db.runCommand({ createIndexes: "tokenBalance", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.tokenBalance"
	}
]
});
db.runCommand({ createIndexes: "tokenTransactions", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.tokenTransactions"
	},
	{
		"v" : 2,
		"key" : {
			"from" : 1,
			"_id" : -1
		},
		"name" : "from_1__id_-1",
		"ns" : "ubiq-explorer.tokenTransactions"
	},
	{
		"v" : 2,
		"key" : {
			"to" : 1,
			"_id" : -1
		},
		"name" : "to_1__id_-1",
		"ns" : "ubiq-explorer.tokenTransactions"
	}
]
});
db.runCommand({ createIndexes: "tokens", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.tokens"
	}
]
});
db.runCommand({ createIndexes: "transactions", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.transactions"
	},
	{
		"v" : 2,
		"key" : {
			"to" : 1,
			"_id" : -1
		},
		"name" : "to_1__id_-1",
		"ns" : "ubiq-explorer.transactions"
	},
	{
		"v" : 2,
		"key" : {
			"from" : 1,
			"_id" : -1
		},
		"name" : "from_1__id_-1",
		"ns" : "ubiq-explorer.transactions"
	},
	{
		"v" : 2,
		"key" : {
			"block" : 1,
			"_id" : -1
		},
		"name" : "block_1__id_-1",
		"ns" : "ubiq-explorer.transactions"
	},
	{
		"v" : 2,
		"key" : {
			"number" : 1,
			"_id" : -1
		},
		"name" : "number_1__id_-1",
		"ns" : "ubiq-explorer.transactions"
	},
	{
		"v" : 2,
		"key" : {
			"hash" : 1
		},
		"name" : "hash_1",
		"ns" : "ubiq-explorer.transactions"
	}
]
});
db.runCommand({ createIndexes: "uncleRateHistory", indexes:
[
	{
		"v" : 2,
		"key" : {
			"_id" : 1
		},
		"name" : "_id_",
		"ns" : "ubiq-explorer.uncleRateHistory"
	}
]
});
