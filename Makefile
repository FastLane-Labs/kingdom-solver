include .env

run-solver:
	RPC_URL=$(RPC_URL) OPERATIONS_RELAY_URL=$(OPERATIONS_RELAY_URL) KINGDOM_DAPP_CONTROL=$(KINGDOM_DAPP_CONTROL) SOLVER_PK=$(SOLVER_PK) SOLVER_CONTRACT_ADDRESS=$(SOLVER_CONTRACT_ADDRESS) go run ./cmd/start.go solver

trigger-auction:
	RPC_URL=$(RPC_URL) AUCTIONEER_URL=$(AUCTIONEER_URL) KINGDOM_DAPP_CONTROL=$(KINGDOM_DAPP_CONTROL) POOL_ADDRESS=$(POOL_ADDRESS) TOKEN_OUT_ADDRESS=$(TOKEN_OUT_ADDRESS) go run ./cmd/start.go trigger-auction

contracts-bindings:
	abigen --abi ./contract/kingdomdappcontrol/abi.json --pkg kingdomdappcontrol --type KingdomDAppControl --out ./contract/kingdomdappcontrol/kingdomdappcontrol.go
	abigen --abi ./contract/demosolver/abi.json --pkg demosolver --type DemoSolver --out ./contract/demosolver/demosolver.go
	abigen --abi ./contract/ramsesv2pool/abi.json --pkg ramsesv2pool --type RamsesV2Pool --out ./contract/ramsesv2pool/ramsesv2pool.go
