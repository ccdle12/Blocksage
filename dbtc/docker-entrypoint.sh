#!/bin/bash
set -e

if [[ "$1" == "bitcoin-cli" || "$1" == "bitcoin-tx" || "$1" == "bitcoind" || "$1" == "test_bitcoin" ]]; then
    mkdir -p "$BITCOIN_DATA"

    if [[ ! -s "$BITCOIN_DATA/bitcoin.conf" ]]; then
	cat <<-EOF > "$BITCOIN_DATA/bitcoin.conf"
maxconnections=12
maxuploadtarget=20
rpcuser=${BITCOIN_RPC_USER:}
rpcpassword=${BITCOIN_RPC_PASSWORD:}
server=1
printtoconsole=1
rpcallowip=::/0
rpcport=8332
port=8333
	EOF
        chown bitcoin:bitcoin "$BITCOIN_DATA/bitcoin.conf"
    fi

    chown -R bitcoin "$BITCOIN_DATA"
    ln -sfn "$BITCOIN_DATA" /home/bitcoin/.bitcoin
    chown -h bitcoin:bitcoin /home/bitcoin/.bitcoin

    exec gosu bitcoin "$@"
fi

exec "$@"
