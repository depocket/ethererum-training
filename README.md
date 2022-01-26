### Get ABI

1. Go to https://bscscan.com
2. Enter contract address to search input
3. Navigate to 'Contract' tab
4. Copy ABI in 'Contract ABI' section
5. E.g BEP20: https://bscscan.com/address/0xe9e7cea3dedca5984780bafc599bd69add087d56#code

### Generate ABI

1. `abigen --abi=[some ABI filename].abi --pkg=[package name] --out=[output.go]`
2. E.g: `abigen --abi=api/contracts/bep20/bep20.abi --pkg=bep20 --out=api/contracts/bep20/bep20.go`

### Find the RPC Endpoint

1. Search `[Smart chain name] smart chain RPC endpoint`
2. E.g: `Binance smart chain RPC endpoint` => `https://bsc-dataseed1.ninicoin.io/`

### Api example

1. Call contract information: `http://127.0.0.1:3000/token/0xe9e7cea3dedca5984780bafc599bd69add087d56`
2. Get user balance: `http://127.0.0.1:3000/token/0xe9e7cea3dedca5984780bafc599bd69add087d56/balance/0xa0d6d7ee26215d00291d51510040becb9ebb9c4b`
