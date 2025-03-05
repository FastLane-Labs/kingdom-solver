include .env

run-solver:
	RPC_URL=$(RPC_URL) OPERATIONS_RELAY_URL=$(OPERATIONS_RELAY_URL) KINGDOM_DAPP_CONTROL=$(KINGDOM_DAPP_CONTROL) go run ./cmd/start.go solver

trigger-auction:
	RPC_URL=$(RPC_URL) AUCTIONEER_URL=$(AUCTIONEER_URL) KINGDOM_DAPP_CONTROL=$(KINGDOM_DAPP_CONTROL) POOL_ADDRESS=$(POOL_ADDRESS) TOKEN_OUT_ADDRESS=$(TOKEN_OUT_ADDRESS) go run ./cmd/start.go trigger-auction

contracts-bindings:
	abigen --abi ./contract/kingdomdappcontrol/abi.json --pkg kingdomdappcontrol --type KingdomDAppControl --out ./contract/kingdomdappcontrol/kingdomdappcontrol.go
