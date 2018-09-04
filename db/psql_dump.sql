--
-- Table structure for table "blocks"
--

DROP TABLE IF EXISTS "blocks";

CREATE TABLE "blocks" (
    hash text UNIQUE NOT NULL,
    confirmations bigint NOT NULL,
    strippedsize bigint NOT NULL,
    size bigint NOT NULL,
    weight bigint NOT NULL,
    height bigint NOT NULL,
    version bigint NOT NULL,
    versionHex text NOT NULL,
    merkleroot text UNIQUE NOT NULL,
    tx text[] NOT NULL,
    time bigint NOT NULL,
    mediantime bigint NOT NULL,
    nonce bigint NOT NULL,
    bits text NOT NULL,
    difficulty double precision NOT NULL,
    chainwork text NOT NULL,
    nextblockhash text,
    PRIMARY KEY("height")
);
