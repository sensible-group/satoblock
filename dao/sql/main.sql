-- 基础区块数据按每2100个区块分区保存到clickhouse，此数据保证和最长链一致

-- block
-- ================================================================
-- 区块头，分区内按区块高度height排序、索引。按blk height查询时可确定分区 (快)
DROP TABLE blk_height;
CREATE TABLE IF NOT EXISTS blk_height (
	height       UInt32,
	blkid        FixedString(32),
	previd       FixedString(32),
	merkle       FixedString(32),
	ntx          UInt64,
	blocktime    UInt32,
	bits         UInt32,
	blocksize    UInt32
) engine=MergeTree()
ORDER BY height
PARTITION BY intDiv(height, 2100)
SETTINGS storage_policy = 'prefer_nvme_policy';
-- load
-- cat /data/674936/blk.ch | clickhouse-client -h 192.168.31.236 --database="bsv" --query="INSERT INTO blk_height FORMAT RowBinary"


-- 区块头，分区内按区块blkid排序、索引。按blkid查询时将遍历所有分区 (慢)
DROP TABLE blk;
CREATE TABLE IF NOT EXISTS blk (
	height       UInt32,
	blkid        FixedString(32),
	previd       FixedString(32),
	merkle       FixedString(32),
	ntx          UInt64,
	blocktime    UInt32,
	bits         UInt32,
	blocksize    UInt32
) engine=MergeTree()
ORDER BY blkid
PARTITION BY intDiv(height, 2100)
SETTINGS storage_policy = 'prefer_nvme_policy';
-- insert
-- INSERT INTO blk SELECT * FROM blk_height;


-- tx list
-- ================================================================
-- 区块包含的交易列表，分区内按区块高度height排序、索引。按blk height查询时可确定分区 (快)
DROP TABLE blktx_height;
CREATE TABLE IF NOT EXISTS blktx_height (
	txid         FixedString(32),
	nin          UInt32,
	nout         UInt32,
	txsize       UInt32,
	locktime     UInt32,
	height       UInt32,
	blkid        FixedString(32),
	txidx        UInt64
) engine=MergeTree()
ORDER BY height
PARTITION BY intDiv(height, 2100)
SETTINGS storage_policy = 'prefer_nvme_policy';
-- load
-- cat /data/674936/tx.ch | clickhouse-client -h 192.168.31.236 --database="bsv" --query="INSERT INTO blktx_height FORMAT RowBinary"


-- 区块包含的交易列表，分区内按交易txid排序、索引。仅按txid查询时将遍历所有分区 (慢)
-- 查询需附带height。可配合tx_height表查询
DROP TABLE tx;
CREATE TABLE IF NOT EXISTS tx (
	txid         FixedString(32),
	nin          UInt32,
	nout         UInt32,
	txsize       UInt32,
	locktime     UInt32,
	height       UInt32,
	blkid        FixedString(32),
	txidx        UInt64
) engine=MergeTree()
ORDER BY txid
PARTITION BY intDiv(height, 2100)
SETTINGS storage_policy = 'prefer_nvme_policy';
-- insert
-- INSERT INTO tx SELECT * FROM blktx_height;

-- txout
-- ================================================================
-- 交易输出列表，分区内按交易txid+idx排序、索引，单条记录包括输出的各种细节。仅按txid查询时将遍历所有分区（慢）
-- 查询需附带height，可配合tx_height表查询
DROP TABLE txout;
CREATE TABLE IF NOT EXISTS txout (
	utxid        FixedString(32),
	vout         UInt32,
	address      String,
	genesis      String,
	satoshi      UInt64,
	script_type  String,
	script_pk    String,
	height       UInt32,
	txidx        UInt64
) engine=MergeTree()
ORDER BY (utxid, vout)
PARTITION BY intDiv(height, 2100)
SETTINGS storage_policy = 'prefer_nvme_policy';
-- load
-- cat /data/674936/tx-out.ch | clickhouse-client -h 192.168.31.236 --database="bsv" --query="INSERT INTO txout FORMAT RowBinary"


-- txin
-- ================================================================
-- 交易输入列表，分区内按交易txid+idx排序、索引，单条记录包括输入的各种细节。仅按txid查询时将遍历所有分区（慢）
-- 查询需附带height。可配合tx_height表查询
DROP TABLE txin_full;
CREATE TABLE IF NOT EXISTS txin_full (
	height       UInt32,         --txo 花费的区块高度
	txidx        UInt64,
	txid         FixedString(32),
	idx          UInt32,
	script_sig   String,
	nsequence    UInt32,

	height_txo   UInt32,         --txo 产生的区块高度
	utxidx       UInt64,
	utxid        FixedString(32),
	vout         UInt32,
	address      String,
	genesis      String,
	satoshi      UInt64,
	script_type  String,
	script_pk    String
) engine=MergeTree()
ORDER BY (txid, idx)
PARTITION BY intDiv(height, 2100)
SETTINGS storage_policy = 'prefer_nvme_policy';
-- load
-- cat /data256/674936/tx-in.ch | clickhouse-client -h 192.168.31.236 --database="bsv" --query="INSERT INTO txin_full FORMAT RowBinary"


-- 交易输入的outpoint列表，分区内按outpoint txid+idx排序、索引。用于查询某txo被哪个tx花费，需遍历所有分区（慢）
-- 查询需附带height，需配合txout_spent_height表查询
DROP TABLE txin_spent;
CREATE TABLE IF NOT EXISTS txin_spent (
	height       UInt32,
	txid         FixedString(32),
	idx          UInt32,
	utxid        FixedString(32),
	vout         UInt32
) engine=MergeTree()
ORDER BY (utxid, vout)
PARTITION BY intDiv(height, 2100)
SETTINGS storage_policy = 'prefer_nvme_policy';
-- 创建数据
-- INSERT INTO txin_spent SELECT height, txid, idx, utxid, vout FROM txin_full;
