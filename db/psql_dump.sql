--
-- Table structure for table "blocks"
--

DROP TABLE IF EXISTS "blocks";

CREATE TABLE "blocks" (
    hash text UNIQUE NOT NULL,
    strippedsize bigint NOT NULL,
    size bigint NOT NULL,
    weight bigint NOT NULL,
    height bigint NOT NULL,
    version bigint NOT NULL,
    versionHex text NOT NULL,
    merkleroot text UNIQUE NOT NULL,
    time bigint NOT NULL,
    mediantime bigint NOT NULL,
    nonce bigint NOT NULL,
    bits text NOT NULL,
    difficulty double precision NOT NULL,
    chainwork text NOT NULL,
    nextblockhash text,
    PRIMARY KEY("hash")
);

DROP TABLE IF EXISTS "transactions";

CREATE TABLE "transactions" (
    blockhash text NOT NULL, 
    txid     text NOT NULL,              
	hash     text NOT NULL,         
	version  int NOT NULL,          
	size     int NOT NULL,                
	vsize    int NOT NULL,       
	locktime bigint NOT NULL,
    PRIMARY KEY("hash"),
    FOREIGN KEY ("blockhash") REFERENCES blocks ("hash")         
);

DROP TABLE IF EXISTS "inputs";

CREATE TABLE "inputs" (
    txhash text NOT NULL,
    inputtxid text NOT NULL,
	vout bigint NOT NULL,
	asm text NOT NULL,
	hex text NOT NULL,
    sequence bigint NOT NULL,
    FOREIGN KEY("txhash") REFERENCES testtransactions ("hash")
);

DROP TABLE IF EXISTS "outputs";

CREATE TABLE "outputs" (
    txhash text NOT NULL,
    value double precision,
	n integer,
    asm text NOT NULL,
	hex text NOT NULL,
    reqsigs integer,
    type text,
    addresses text[],
    FOREIGN KEY("txhash") REFERENCES testtransactions ("hash")
);


DROP TABLE IF EXISTS "testblocks";

CREATE TABLE "testblocks" (
    hash text UNIQUE NOT NULL,
    strippedsize bigint NOT NULL,
    size bigint NOT NULL,
    weight bigint NOT NULL,
    height bigint NOT NULL,
    version bigint NOT NULL,
    versionHex text NOT NULL,
    merkleroot text UNIQUE NOT NULL,
    time bigint NOT NULL,
    mediantime bigint NOT NULL,
    nonce bigint NOT NULL,
    bits text NOT NULL,
    difficulty double precision NOT NULL,
    chainwork text NOT NULL,
    nextblockhash text,
    PRIMARY KEY("hash")
);

DROP TABLE IF EXISTS "testtransactions";

CREATE TABLE "testtransactions" (
    blockhash text NOT NULL, 
    txid     text UNIQUE NOT NULL,
	hash     text NOT NULL,
	version  int NOT NULL,          
	size     int NOT NULL,                
	vsize    int NOT NULL,       
	locktime bigint NOT NULL,
    PRIMARY KEY("hash"),
    FOREIGN KEY ("blockhash") REFERENCES testblocks ("hash")            
);

DROP TABLE IF EXISTS "testinputs";

CREATE TABLE "testinputs" (
    txhash text NOT NULL,
    inputtxid text NOT NULL,
	vout bigint NOT NULL,
	asm text NOT NULL,
	hex text NOT NULL,
    sequence bigint NOT NULL,
    FOREIGN KEY("txhash") REFERENCES testtransactions ("hash")
);

DROP TABLE IF EXISTS "testoutputs";

CREATE TABLE "testoutputs" (
    txhash text NOT NULL,
    value double precision,
	n integer,
    asm text NOT NULL,
	hex text NOT NULL,
    reqsigs integer,
    type text,
    addresses text[],
    FOREIGN KEY("txhash") REFERENCES testtransactions ("hash")
);
