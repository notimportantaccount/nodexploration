# nodexploration


https://github.com/notimportantaccount/nodexploration

## Running the node

### Consensus
first with checkpoint sync
```
./prysm.sh beacon-chain --checkpoint-sync-url=https://beaconstate.ethstaker.cc  --genesis-beacon-api-url=https://beaconstate.ethstaker.cc --execution-endpoint=http://localhost:8551 --jwt-secret=~/.ethereum/geth/jwtsecret --suggested-fee-recipient=0x5D599f39c1afb84FD20D46981fd868D2cE5CCF8C
```

then
```
./prysm.sh beacon-chain --execution-endpoint=http://localhost:8551 --jwt-secret=~/.ethereum/geth/jwtsecret --suggested-fee-recipient=0x5D599f39c1afb84FD20D46981fd868D2cE5CCF8C
```

Us
garchery.eth
0x5D599f39c1afb84FD20D46981fd868D2cE5CCF8C

### Execution
```
geth --http --authrpc.addr localhost --authrpc.port 8551 --authrpc.vhosts localhost --authrpc.jwtsecret ~/.ethereum/geth/jwtsecret
```

Attach a node terminal

geth attach .ethereum/geth.ipc
Commands
eth.getBlock("6517331").transactions.length

## Go contract binding

https://geth.ethereum.org/docs/developers/dapp-developer/native-bindings

### Run a go script retrieving past deposit aave events

go run aaveDepositEvents.go