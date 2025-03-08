include .env

run-solver:
	RPC_URL=$(RPC_URL) OPERATIONS_RELAY_URL=$(OPERATIONS_RELAY_URL) KINGDOM_DAPP_CONTROL=$(KINGDOM_DAPP_CONTROL) SOLVER_PK=$(SOLVER_PK) SOLVER_CONTRACT_ADDRESS=$(SOLVER_CONTRACT_ADDRESS) go run ./cmd solver

trigger-auction:
	RPC_URL=$(RPC_URL) AUCTIONEER_URL=$(AUCTIONEER_URL) KINGDOM_DAPP_CONTROL=$(KINGDOM_DAPP_CONTROL) POOL_ADDRESS=$(POOL_ADDRESS) TOKEN_OUT_ADDRESS=$(TOKEN_OUT_ADDRESS) go run ./cmd trigger-auction

print-all:
	$(MAKE) print-bonded-atleth
	$(MAKE) print-solver-contract-usdc-balance
	$(MAKE) print-solver-contract-weth-balance

print-bonded-atleth:
	RPC_URL=$(RPC_URL) SOLVER_PK=$(SOLVER_PK) go run ./cmd print-bonded-atleth

deposit-and-bond-atleth:
	RPC_URL=$(RPC_URL) SOLVER_PK=$(SOLVER_PK) go run ./cmd deposit-and-bond-atleth

print-solver-contract-usdc-balance:
	RPC_URL=$(RPC_URL) SOLVER_CONTRACT_ADDRESS=$(SOLVER_CONTRACT_ADDRESS) go run ./cmd print-solver-contract-erc20-balance 0xaf88d065e77c8cC2239327C5EDb3A432268e5831

print-solver-contract-weth-balance:
	RPC_URL=$(RPC_URL) SOLVER_CONTRACT_ADDRESS=$(SOLVER_CONTRACT_ADDRESS) go run ./cmd print-solver-contract-erc20-balance 0x82aF49447D8a07e3bd95BD0d56f35241523fBab1

swap-usdc-to-solver-contract:
	RPC_URL=$(RPC_URL) WETH_ADDRESS=$(WETH_ADDRESS) SOLVER_PK=$(SOLVER_PK) SOLVER_CONTRACT_ADDRESS=$(SOLVER_CONTRACT_ADDRESS) ROUTER_ADDRESS=$(ROUTER_ADDRESS) go run ./cmd swap-erc20-to-solver-contract 0xaf88d065e77c8cC2239327C5EDb3A432268e5831

swap-weth-to-solver-contract:
	RPC_URL=$(RPC_URL) WETH_ADDRESS=$(WETH_ADDRESS) SOLVER_PK=$(SOLVER_PK) SOLVER_CONTRACT_ADDRESS=$(SOLVER_CONTRACT_ADDRESS) ROUTER_ADDRESS=$(ROUTER_ADDRESS) go run ./cmd swap-erc20-to-solver-contract 0x82aF49447D8a07e3bd95BD0d56f35241523fBab1

contracts-bindings:
	abigen --abi ./contract/kingdomdappcontrol/abi.json --pkg kingdomdappcontrol --type KingdomDAppControl --out ./contract/kingdomdappcontrol/kingdomdappcontrol.go
	abigen --abi ./contract/demosolver/abi.json --pkg demosolver --type DemoSolver --out ./contract/demosolver/demosolver.go
	abigen --abi ./contract/ramsesv2pool/abi.json --pkg ramsesv2pool --type RamsesV2Pool --out ./contract/ramsesv2pool/ramsesv2pool.go
	abigen --abi ./contract/erc20/abi.json --pkg erc20 --type Erc20 --out ./contract/erc20/erc20.go
	abigen --abi ./contract/kingdomrouter/abi.json --pkg kingdomrouter --type KingdomRouter --out ./contract/kingdomrouter/kingdomrouter.go
