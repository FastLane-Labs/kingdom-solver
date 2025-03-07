// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ramsesv2pool

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RamsesV2PoolMetaData contains all meta data concerning the RamsesV2Pool contract.
var RamsesV2PoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"_advancePeriod\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"boostInfos\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"totalBoostAmount\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"totalVeRamAmount\",\"type\":\"int128\",\"internalType\":\"int128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"boostInfos\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"boostAmount\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"veRamAmount\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"secondsDebtX96\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"boostedSecondsDebtX96\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"boostedLiquidity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"amount\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"amount\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"amount\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"veRamTokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"collect\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"amount0Requested\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"amount1Requested\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"amount1\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"collect\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"amount0Requested\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"amount1Requested\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"amount1\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"collectProtocol\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount0Requested\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"amount1Requested\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"amount1\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"currentFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeGrowthGlobal0X128\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeGrowthGlobal1X128\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"flash\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseObservationCardinalityNext\",\"inputs\":[{\"name\":\"observationCardinalityNext\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_factory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nfpManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_veRam\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_voter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token0\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token1\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"_tickSpacing\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"sqrtPriceX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"liquidity\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxLiquidityPerTick\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"amount\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"amount\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"veRamTokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nfpManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"observations\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"blockTimestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tickCumulative\",\"type\":\"int56\",\"internalType\":\"int56\"},{\"name\":\"secondsPerLiquidityCumulativeX128\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"initialized\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"secondsPerBoostedLiquidityPeriodX128\",\"type\":\"uint160\",\"internalType\":\"uint160\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"observe\",\"inputs\":[{\"name\":\"secondsAgos\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[{\"name\":\"tickCumulatives\",\"type\":\"int56[]\",\"internalType\":\"int56[]\"},{\"name\":\"secondsPerLiquidityCumulativeX128s\",\"type\":\"uint160[]\",\"internalType\":\"uint160[]\"},{\"name\":\"secondsPerBoostedLiquidityPeriodX128s\",\"type\":\"uint160[]\",\"internalType\":\"uint160[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"periodCumulativesInside\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[{\"name\":\"secondsPerLiquidityInsideX128\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"secondsPerBoostedLiquidityInsideX128\",\"type\":\"uint160\",\"internalType\":\"uint160\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"periods\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"previousPeriod\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"startTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"lastTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"endSecondsPerLiquidityCumulativeX128\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"endSecondsPerBoostedLiquidityCumulativeX128\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"boostedInRange\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"positionPeriodDebt\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[{\"name\":\"secondsDebtX96\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"boostedSecondsDebtX96\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"positionPeriodSecondsInRange\",\"inputs\":[{\"name\":\"period\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[{\"name\":\"periodSecondsInsideX96\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"periodBoostedSecondsInsideX96\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"positions\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"_liquidity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"feeGrowthInside0LastX128\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeGrowthInside1LastX128\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensOwed0\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"tokensOwed1\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"attachedVeRamId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolFees\",\"inputs\":[],\"outputs\":[{\"name\":\"token0\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"token1\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readStorage\",\"inputs\":[{\"name\":\"slots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"returnData\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setFee\",\"inputs\":[{\"name\":\"_fee\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slot0\",\"inputs\":[],\"outputs\":[{\"name\":\"sqrtPriceX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"tick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"observationIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"observationCardinality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"observationCardinalityNext\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"feeProtocol\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"unlocked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"snapshotCumulativesInside\",\"inputs\":[{\"name\":\"tickLower\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[{\"name\":\"tickCumulativeInside\",\"type\":\"int56\",\"internalType\":\"int56\"},{\"name\":\"secondsPerLiquidityInsideX128\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"secondsPerBoostedLiquidityInsideX128\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"secondsInside\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"zeroForOne\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"amountSpecified\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"amount0\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"amount1\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tickBitmap\",\"inputs\":[{\"name\":\"wordPosition\",\"type\":\"int16\",\"internalType\":\"int16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tickSpacing\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ticks\",\"inputs\":[{\"name\":\"tick\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[{\"name\":\"liquidityGross\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"liquidityNet\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"boostedLiquidityGross\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"boostedLiquidityNet\",\"type\":\"int128\",\"internalType\":\"int128\"},{\"name\":\"feeGrowthOutside0X128\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeGrowthOutside1X128\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tickCumulativeOutside\",\"type\":\"int56\",\"internalType\":\"int56\"},{\"name\":\"secondsPerLiquidityOutsideX128\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"secondsOutside\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"initialized\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token0\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token1\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"veRam\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"voter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Burn\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tickLower\",\"type\":\"int24\",\"indexed\":true,\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"indexed\":true,\"internalType\":\"int24\"},{\"name\":\"amount\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"amount0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Collect\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tickLower\",\"type\":\"int24\",\"indexed\":true,\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"indexed\":true,\"internalType\":\"int24\"},{\"name\":\"amount0\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"amount1\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollectProtocol\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount0\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"amount1\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Flash\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"paid0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"paid1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreaseObservationCardinalityNext\",\"inputs\":[{\"name\":\"observationCardinalityNextOld\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"observationCardinalityNextNew\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialize\",\"inputs\":[{\"name\":\"sqrtPriceX96\",\"type\":\"uint160\",\"indexed\":false,\"internalType\":\"uint160\"},{\"name\":\"tick\",\"type\":\"int24\",\"indexed\":false,\"internalType\":\"int24\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Mint\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tickLower\",\"type\":\"int24\",\"indexed\":true,\"internalType\":\"int24\"},{\"name\":\"tickUpper\",\"type\":\"int24\",\"indexed\":true,\"internalType\":\"int24\"},{\"name\":\"amount\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"amount0\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount1\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetFeeProtocol\",\"inputs\":[{\"name\":\"feeProtocol0Old\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"feeProtocol1Old\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"feeProtocol0New\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"feeProtocol1New\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Swap\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount0\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"amount1\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"sqrtPriceX96\",\"type\":\"uint160\",\"indexed\":false,\"internalType\":\"uint160\"},{\"name\":\"liquidity\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"},{\"name\":\"tick\",\"type\":\"int24\",\"indexed\":false,\"internalType\":\"int24\"}],\"anonymous\":false}]",
}

// RamsesV2PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use RamsesV2PoolMetaData.ABI instead.
var RamsesV2PoolABI = RamsesV2PoolMetaData.ABI

// RamsesV2Pool is an auto generated Go binding around an Ethereum contract.
type RamsesV2Pool struct {
	RamsesV2PoolCaller     // Read-only binding to the contract
	RamsesV2PoolTransactor // Write-only binding to the contract
	RamsesV2PoolFilterer   // Log filterer for contract events
}

// RamsesV2PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type RamsesV2PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RamsesV2PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RamsesV2PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RamsesV2PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RamsesV2PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RamsesV2PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RamsesV2PoolSession struct {
	Contract     *RamsesV2Pool     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RamsesV2PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RamsesV2PoolCallerSession struct {
	Contract *RamsesV2PoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RamsesV2PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RamsesV2PoolTransactorSession struct {
	Contract     *RamsesV2PoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RamsesV2PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type RamsesV2PoolRaw struct {
	Contract *RamsesV2Pool // Generic contract binding to access the raw methods on
}

// RamsesV2PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RamsesV2PoolCallerRaw struct {
	Contract *RamsesV2PoolCaller // Generic read-only contract binding to access the raw methods on
}

// RamsesV2PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RamsesV2PoolTransactorRaw struct {
	Contract *RamsesV2PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRamsesV2Pool creates a new instance of RamsesV2Pool, bound to a specific deployed contract.
func NewRamsesV2Pool(address common.Address, backend bind.ContractBackend) (*RamsesV2Pool, error) {
	contract, err := bindRamsesV2Pool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RamsesV2Pool{RamsesV2PoolCaller: RamsesV2PoolCaller{contract: contract}, RamsesV2PoolTransactor: RamsesV2PoolTransactor{contract: contract}, RamsesV2PoolFilterer: RamsesV2PoolFilterer{contract: contract}}, nil
}

// NewRamsesV2PoolCaller creates a new read-only instance of RamsesV2Pool, bound to a specific deployed contract.
func NewRamsesV2PoolCaller(address common.Address, caller bind.ContractCaller) (*RamsesV2PoolCaller, error) {
	contract, err := bindRamsesV2Pool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RamsesV2PoolCaller{contract: contract}, nil
}

// NewRamsesV2PoolTransactor creates a new write-only instance of RamsesV2Pool, bound to a specific deployed contract.
func NewRamsesV2PoolTransactor(address common.Address, transactor bind.ContractTransactor) (*RamsesV2PoolTransactor, error) {
	contract, err := bindRamsesV2Pool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RamsesV2PoolTransactor{contract: contract}, nil
}

// NewRamsesV2PoolFilterer creates a new log filterer instance of RamsesV2Pool, bound to a specific deployed contract.
func NewRamsesV2PoolFilterer(address common.Address, filterer bind.ContractFilterer) (*RamsesV2PoolFilterer, error) {
	contract, err := bindRamsesV2Pool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RamsesV2PoolFilterer{contract: contract}, nil
}

// bindRamsesV2Pool binds a generic wrapper to an already deployed contract.
func bindRamsesV2Pool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RamsesV2PoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RamsesV2Pool *RamsesV2PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RamsesV2Pool.Contract.RamsesV2PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RamsesV2Pool *RamsesV2PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.RamsesV2PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RamsesV2Pool *RamsesV2PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.RamsesV2PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RamsesV2Pool *RamsesV2PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RamsesV2Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RamsesV2Pool *RamsesV2PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RamsesV2Pool *RamsesV2PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.contract.Transact(opts, method, params...)
}

// BoostInfos is a free data retrieval call binding the contract method 0x0d63237f.
//
// Solidity: function boostInfos(uint256 period) view returns(uint128 totalBoostAmount, int128 totalVeRamAmount)
func (_RamsesV2Pool *RamsesV2PoolCaller) BoostInfos(opts *bind.CallOpts, period *big.Int) (struct {
	TotalBoostAmount *big.Int
	TotalVeRamAmount *big.Int
}, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "boostInfos", period)

	outstruct := new(struct {
		TotalBoostAmount *big.Int
		TotalVeRamAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalBoostAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalVeRamAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BoostInfos is a free data retrieval call binding the contract method 0x0d63237f.
//
// Solidity: function boostInfos(uint256 period) view returns(uint128 totalBoostAmount, int128 totalVeRamAmount)
func (_RamsesV2Pool *RamsesV2PoolSession) BoostInfos(period *big.Int) (struct {
	TotalBoostAmount *big.Int
	TotalVeRamAmount *big.Int
}, error) {
	return _RamsesV2Pool.Contract.BoostInfos(&_RamsesV2Pool.CallOpts, period)
}

// BoostInfos is a free data retrieval call binding the contract method 0x0d63237f.
//
// Solidity: function boostInfos(uint256 period) view returns(uint128 totalBoostAmount, int128 totalVeRamAmount)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) BoostInfos(period *big.Int) (struct {
	TotalBoostAmount *big.Int
	TotalVeRamAmount *big.Int
}, error) {
	return _RamsesV2Pool.Contract.BoostInfos(&_RamsesV2Pool.CallOpts, period)
}

// BoostInfos0 is a free data retrieval call binding the contract method 0x4860e05f.
//
// Solidity: function boostInfos(uint256 period, bytes32 key) view returns(uint128 boostAmount, int128 veRamAmount, int256 secondsDebtX96, int256 boostedSecondsDebtX96)
func (_RamsesV2Pool *RamsesV2PoolCaller) BoostInfos0(opts *bind.CallOpts, period *big.Int, key [32]byte) (struct {
	BoostAmount           *big.Int
	VeRamAmount           *big.Int
	SecondsDebtX96        *big.Int
	BoostedSecondsDebtX96 *big.Int
}, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "boostInfos0", period, key)

	outstruct := new(struct {
		BoostAmount           *big.Int
		VeRamAmount           *big.Int
		SecondsDebtX96        *big.Int
		BoostedSecondsDebtX96 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BoostAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.VeRamAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SecondsDebtX96 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BoostedSecondsDebtX96 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BoostInfos0 is a free data retrieval call binding the contract method 0x4860e05f.
//
// Solidity: function boostInfos(uint256 period, bytes32 key) view returns(uint128 boostAmount, int128 veRamAmount, int256 secondsDebtX96, int256 boostedSecondsDebtX96)
func (_RamsesV2Pool *RamsesV2PoolSession) BoostInfos0(period *big.Int, key [32]byte) (struct {
	BoostAmount           *big.Int
	VeRamAmount           *big.Int
	SecondsDebtX96        *big.Int
	BoostedSecondsDebtX96 *big.Int
}, error) {
	return _RamsesV2Pool.Contract.BoostInfos0(&_RamsesV2Pool.CallOpts, period, key)
}

// BoostInfos0 is a free data retrieval call binding the contract method 0x4860e05f.
//
// Solidity: function boostInfos(uint256 period, bytes32 key) view returns(uint128 boostAmount, int128 veRamAmount, int256 secondsDebtX96, int256 boostedSecondsDebtX96)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) BoostInfos0(period *big.Int, key [32]byte) (struct {
	BoostAmount           *big.Int
	VeRamAmount           *big.Int
	SecondsDebtX96        *big.Int
	BoostedSecondsDebtX96 *big.Int
}, error) {
	return _RamsesV2Pool.Contract.BoostInfos0(&_RamsesV2Pool.CallOpts, period, key)
}

// BoostedLiquidity is a free data retrieval call binding the contract method 0x9fdb2616.
//
// Solidity: function boostedLiquidity() view returns(uint128)
func (_RamsesV2Pool *RamsesV2PoolCaller) BoostedLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "boostedLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BoostedLiquidity is a free data retrieval call binding the contract method 0x9fdb2616.
//
// Solidity: function boostedLiquidity() view returns(uint128)
func (_RamsesV2Pool *RamsesV2PoolSession) BoostedLiquidity() (*big.Int, error) {
	return _RamsesV2Pool.Contract.BoostedLiquidity(&_RamsesV2Pool.CallOpts)
}

// BoostedLiquidity is a free data retrieval call binding the contract method 0x9fdb2616.
//
// Solidity: function boostedLiquidity() view returns(uint128)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) BoostedLiquidity() (*big.Int, error) {
	return _RamsesV2Pool.Contract.BoostedLiquidity(&_RamsesV2Pool.CallOpts)
}

// CurrentFee is a free data retrieval call binding the contract method 0xda3c300d.
//
// Solidity: function currentFee() view returns(uint24)
func (_RamsesV2Pool *RamsesV2PoolCaller) CurrentFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "currentFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentFee is a free data retrieval call binding the contract method 0xda3c300d.
//
// Solidity: function currentFee() view returns(uint24)
func (_RamsesV2Pool *RamsesV2PoolSession) CurrentFee() (*big.Int, error) {
	return _RamsesV2Pool.Contract.CurrentFee(&_RamsesV2Pool.CallOpts)
}

// CurrentFee is a free data retrieval call binding the contract method 0xda3c300d.
//
// Solidity: function currentFee() view returns(uint24)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) CurrentFee() (*big.Int, error) {
	return _RamsesV2Pool.Contract.CurrentFee(&_RamsesV2Pool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_RamsesV2Pool *RamsesV2PoolCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_RamsesV2Pool *RamsesV2PoolSession) Factory() (common.Address, error) {
	return _RamsesV2Pool.Contract.Factory(&_RamsesV2Pool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) Factory() (common.Address, error) {
	return _RamsesV2Pool.Contract.Factory(&_RamsesV2Pool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_RamsesV2Pool *RamsesV2PoolCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_RamsesV2Pool *RamsesV2PoolSession) Fee() (*big.Int, error) {
	return _RamsesV2Pool.Contract.Fee(&_RamsesV2Pool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) Fee() (*big.Int, error) {
	return _RamsesV2Pool.Contract.Fee(&_RamsesV2Pool.CallOpts)
}

// FeeGrowthGlobal0X128 is a free data retrieval call binding the contract method 0xf3058399.
//
// Solidity: function feeGrowthGlobal0X128() view returns(uint256)
func (_RamsesV2Pool *RamsesV2PoolCaller) FeeGrowthGlobal0X128(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "feeGrowthGlobal0X128")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGrowthGlobal0X128 is a free data retrieval call binding the contract method 0xf3058399.
//
// Solidity: function feeGrowthGlobal0X128() view returns(uint256)
func (_RamsesV2Pool *RamsesV2PoolSession) FeeGrowthGlobal0X128() (*big.Int, error) {
	return _RamsesV2Pool.Contract.FeeGrowthGlobal0X128(&_RamsesV2Pool.CallOpts)
}

// FeeGrowthGlobal0X128 is a free data retrieval call binding the contract method 0xf3058399.
//
// Solidity: function feeGrowthGlobal0X128() view returns(uint256)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) FeeGrowthGlobal0X128() (*big.Int, error) {
	return _RamsesV2Pool.Contract.FeeGrowthGlobal0X128(&_RamsesV2Pool.CallOpts)
}

// FeeGrowthGlobal1X128 is a free data retrieval call binding the contract method 0x46141319.
//
// Solidity: function feeGrowthGlobal1X128() view returns(uint256)
func (_RamsesV2Pool *RamsesV2PoolCaller) FeeGrowthGlobal1X128(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "feeGrowthGlobal1X128")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeGrowthGlobal1X128 is a free data retrieval call binding the contract method 0x46141319.
//
// Solidity: function feeGrowthGlobal1X128() view returns(uint256)
func (_RamsesV2Pool *RamsesV2PoolSession) FeeGrowthGlobal1X128() (*big.Int, error) {
	return _RamsesV2Pool.Contract.FeeGrowthGlobal1X128(&_RamsesV2Pool.CallOpts)
}

// FeeGrowthGlobal1X128 is a free data retrieval call binding the contract method 0x46141319.
//
// Solidity: function feeGrowthGlobal1X128() view returns(uint256)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) FeeGrowthGlobal1X128() (*big.Int, error) {
	return _RamsesV2Pool.Contract.FeeGrowthGlobal1X128(&_RamsesV2Pool.CallOpts)
}

// LastPeriod is a free data retrieval call binding the contract method 0xd340ef8a.
//
// Solidity: function lastPeriod() view returns(uint256)
func (_RamsesV2Pool *RamsesV2PoolCaller) LastPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "lastPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastPeriod is a free data retrieval call binding the contract method 0xd340ef8a.
//
// Solidity: function lastPeriod() view returns(uint256)
func (_RamsesV2Pool *RamsesV2PoolSession) LastPeriod() (*big.Int, error) {
	return _RamsesV2Pool.Contract.LastPeriod(&_RamsesV2Pool.CallOpts)
}

// LastPeriod is a free data retrieval call binding the contract method 0xd340ef8a.
//
// Solidity: function lastPeriod() view returns(uint256)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) LastPeriod() (*big.Int, error) {
	return _RamsesV2Pool.Contract.LastPeriod(&_RamsesV2Pool.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_RamsesV2Pool *RamsesV2PoolCaller) Liquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "liquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_RamsesV2Pool *RamsesV2PoolSession) Liquidity() (*big.Int, error) {
	return _RamsesV2Pool.Contract.Liquidity(&_RamsesV2Pool.CallOpts)
}

// Liquidity is a free data retrieval call binding the contract method 0x1a686502.
//
// Solidity: function liquidity() view returns(uint128)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) Liquidity() (*big.Int, error) {
	return _RamsesV2Pool.Contract.Liquidity(&_RamsesV2Pool.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_RamsesV2Pool *RamsesV2PoolCaller) MaxLiquidityPerTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "maxLiquidityPerTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_RamsesV2Pool *RamsesV2PoolSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _RamsesV2Pool.Contract.MaxLiquidityPerTick(&_RamsesV2Pool.CallOpts)
}

// MaxLiquidityPerTick is a free data retrieval call binding the contract method 0x70cf754a.
//
// Solidity: function maxLiquidityPerTick() view returns(uint128)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) MaxLiquidityPerTick() (*big.Int, error) {
	return _RamsesV2Pool.Contract.MaxLiquidityPerTick(&_RamsesV2Pool.CallOpts)
}

// NfpManager is a free data retrieval call binding the contract method 0x98bbc3c7.
//
// Solidity: function nfpManager() view returns(address)
func (_RamsesV2Pool *RamsesV2PoolCaller) NfpManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "nfpManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NfpManager is a free data retrieval call binding the contract method 0x98bbc3c7.
//
// Solidity: function nfpManager() view returns(address)
func (_RamsesV2Pool *RamsesV2PoolSession) NfpManager() (common.Address, error) {
	return _RamsesV2Pool.Contract.NfpManager(&_RamsesV2Pool.CallOpts)
}

// NfpManager is a free data retrieval call binding the contract method 0x98bbc3c7.
//
// Solidity: function nfpManager() view returns(address)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) NfpManager() (common.Address, error) {
	return _RamsesV2Pool.Contract.NfpManager(&_RamsesV2Pool.CallOpts)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 index) view returns(uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized, uint160 secondsPerBoostedLiquidityPeriodX128)
func (_RamsesV2Pool *RamsesV2PoolCaller) Observations(opts *bind.CallOpts, index *big.Int) (struct {
	BlockTimestamp                       uint32
	TickCumulative                       *big.Int
	SecondsPerLiquidityCumulativeX128    *big.Int
	Initialized                          bool
	SecondsPerBoostedLiquidityPeriodX128 *big.Int
}, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "observations", index)

	outstruct := new(struct {
		BlockTimestamp                       uint32
		TickCumulative                       *big.Int
		SecondsPerLiquidityCumulativeX128    *big.Int
		Initialized                          bool
		SecondsPerBoostedLiquidityPeriodX128 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockTimestamp = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.TickCumulative = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityCumulativeX128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Initialized = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.SecondsPerBoostedLiquidityPeriodX128 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 index) view returns(uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized, uint160 secondsPerBoostedLiquidityPeriodX128)
func (_RamsesV2Pool *RamsesV2PoolSession) Observations(index *big.Int) (struct {
	BlockTimestamp                       uint32
	TickCumulative                       *big.Int
	SecondsPerLiquidityCumulativeX128    *big.Int
	Initialized                          bool
	SecondsPerBoostedLiquidityPeriodX128 *big.Int
}, error) {
	return _RamsesV2Pool.Contract.Observations(&_RamsesV2Pool.CallOpts, index)
}

// Observations is a free data retrieval call binding the contract method 0x252c09d7.
//
// Solidity: function observations(uint256 index) view returns(uint32 blockTimestamp, int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128, bool initialized, uint160 secondsPerBoostedLiquidityPeriodX128)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) Observations(index *big.Int) (struct {
	BlockTimestamp                       uint32
	TickCumulative                       *big.Int
	SecondsPerLiquidityCumulativeX128    *big.Int
	Initialized                          bool
	SecondsPerBoostedLiquidityPeriodX128 *big.Int
}, error) {
	return _RamsesV2Pool.Contract.Observations(&_RamsesV2Pool.CallOpts, index)
}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulativeX128s, uint160[] secondsPerBoostedLiquidityPeriodX128s)
func (_RamsesV2Pool *RamsesV2PoolCaller) Observe(opts *bind.CallOpts, secondsAgos []uint32) (struct {
	TickCumulatives                       []*big.Int
	SecondsPerLiquidityCumulativeX128s    []*big.Int
	SecondsPerBoostedLiquidityPeriodX128s []*big.Int
}, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "observe", secondsAgos)

	outstruct := new(struct {
		TickCumulatives                       []*big.Int
		SecondsPerLiquidityCumulativeX128s    []*big.Int
		SecondsPerBoostedLiquidityPeriodX128s []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TickCumulatives = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.SecondsPerLiquidityCumulativeX128s = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.SecondsPerBoostedLiquidityPeriodX128s = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulativeX128s, uint160[] secondsPerBoostedLiquidityPeriodX128s)
func (_RamsesV2Pool *RamsesV2PoolSession) Observe(secondsAgos []uint32) (struct {
	TickCumulatives                       []*big.Int
	SecondsPerLiquidityCumulativeX128s    []*big.Int
	SecondsPerBoostedLiquidityPeriodX128s []*big.Int
}, error) {
	return _RamsesV2Pool.Contract.Observe(&_RamsesV2Pool.CallOpts, secondsAgos)
}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] secondsAgos) view returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulativeX128s, uint160[] secondsPerBoostedLiquidityPeriodX128s)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) Observe(secondsAgos []uint32) (struct {
	TickCumulatives                       []*big.Int
	SecondsPerLiquidityCumulativeX128s    []*big.Int
	SecondsPerBoostedLiquidityPeriodX128s []*big.Int
}, error) {
	return _RamsesV2Pool.Contract.Observe(&_RamsesV2Pool.CallOpts, secondsAgos)
}

// PeriodCumulativesInside is a free data retrieval call binding the contract method 0xadd5887e.
//
// Solidity: function periodCumulativesInside(uint32 period, int24 tickLower, int24 tickUpper) view returns(uint160 secondsPerLiquidityInsideX128, uint160 secondsPerBoostedLiquidityInsideX128)
func (_RamsesV2Pool *RamsesV2PoolCaller) PeriodCumulativesInside(opts *bind.CallOpts, period uint32, tickLower *big.Int, tickUpper *big.Int) (struct {
	SecondsPerLiquidityInsideX128        *big.Int
	SecondsPerBoostedLiquidityInsideX128 *big.Int
}, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "periodCumulativesInside", period, tickLower, tickUpper)

	outstruct := new(struct {
		SecondsPerLiquidityInsideX128        *big.Int
		SecondsPerBoostedLiquidityInsideX128 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SecondsPerLiquidityInsideX128 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerBoostedLiquidityInsideX128 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PeriodCumulativesInside is a free data retrieval call binding the contract method 0xadd5887e.
//
// Solidity: function periodCumulativesInside(uint32 period, int24 tickLower, int24 tickUpper) view returns(uint160 secondsPerLiquidityInsideX128, uint160 secondsPerBoostedLiquidityInsideX128)
func (_RamsesV2Pool *RamsesV2PoolSession) PeriodCumulativesInside(period uint32, tickLower *big.Int, tickUpper *big.Int) (struct {
	SecondsPerLiquidityInsideX128        *big.Int
	SecondsPerBoostedLiquidityInsideX128 *big.Int
}, error) {
	return _RamsesV2Pool.Contract.PeriodCumulativesInside(&_RamsesV2Pool.CallOpts, period, tickLower, tickUpper)
}

// PeriodCumulativesInside is a free data retrieval call binding the contract method 0xadd5887e.
//
// Solidity: function periodCumulativesInside(uint32 period, int24 tickLower, int24 tickUpper) view returns(uint160 secondsPerLiquidityInsideX128, uint160 secondsPerBoostedLiquidityInsideX128)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) PeriodCumulativesInside(period uint32, tickLower *big.Int, tickUpper *big.Int) (struct {
	SecondsPerLiquidityInsideX128        *big.Int
	SecondsPerBoostedLiquidityInsideX128 *big.Int
}, error) {
	return _RamsesV2Pool.Contract.PeriodCumulativesInside(&_RamsesV2Pool.CallOpts, period, tickLower, tickUpper)
}

// Periods is a free data retrieval call binding the contract method 0xea4a1104.
//
// Solidity: function periods(uint256 period) view returns(uint32 previousPeriod, int24 startTick, int24 lastTick, uint160 endSecondsPerLiquidityCumulativeX128, uint160 endSecondsPerBoostedLiquidityCumulativeX128, uint32 boostedInRange)
func (_RamsesV2Pool *RamsesV2PoolCaller) Periods(opts *bind.CallOpts, period *big.Int) (struct {
	PreviousPeriod                              uint32
	StartTick                                   *big.Int
	LastTick                                    *big.Int
	EndSecondsPerLiquidityCumulativeX128        *big.Int
	EndSecondsPerBoostedLiquidityCumulativeX128 *big.Int
	BoostedInRange                              uint32
}, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "periods", period)

	outstruct := new(struct {
		PreviousPeriod                              uint32
		StartTick                                   *big.Int
		LastTick                                    *big.Int
		EndSecondsPerLiquidityCumulativeX128        *big.Int
		EndSecondsPerBoostedLiquidityCumulativeX128 *big.Int
		BoostedInRange                              uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PreviousPeriod = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.StartTick = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastTick = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndSecondsPerLiquidityCumulativeX128 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EndSecondsPerBoostedLiquidityCumulativeX128 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.BoostedInRange = *abi.ConvertType(out[5], new(uint32)).(*uint32)

	return *outstruct, err

}

// Periods is a free data retrieval call binding the contract method 0xea4a1104.
//
// Solidity: function periods(uint256 period) view returns(uint32 previousPeriod, int24 startTick, int24 lastTick, uint160 endSecondsPerLiquidityCumulativeX128, uint160 endSecondsPerBoostedLiquidityCumulativeX128, uint32 boostedInRange)
func (_RamsesV2Pool *RamsesV2PoolSession) Periods(period *big.Int) (struct {
	PreviousPeriod                              uint32
	StartTick                                   *big.Int
	LastTick                                    *big.Int
	EndSecondsPerLiquidityCumulativeX128        *big.Int
	EndSecondsPerBoostedLiquidityCumulativeX128 *big.Int
	BoostedInRange                              uint32
}, error) {
	return _RamsesV2Pool.Contract.Periods(&_RamsesV2Pool.CallOpts, period)
}

// Periods is a free data retrieval call binding the contract method 0xea4a1104.
//
// Solidity: function periods(uint256 period) view returns(uint32 previousPeriod, int24 startTick, int24 lastTick, uint160 endSecondsPerLiquidityCumulativeX128, uint160 endSecondsPerBoostedLiquidityCumulativeX128, uint32 boostedInRange)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) Periods(period *big.Int) (struct {
	PreviousPeriod                              uint32
	StartTick                                   *big.Int
	LastTick                                    *big.Int
	EndSecondsPerLiquidityCumulativeX128        *big.Int
	EndSecondsPerBoostedLiquidityCumulativeX128 *big.Int
	BoostedInRange                              uint32
}, error) {
	return _RamsesV2Pool.Contract.Periods(&_RamsesV2Pool.CallOpts, period)
}

// PositionPeriodDebt is a free data retrieval call binding the contract method 0xdfc8b615.
//
// Solidity: function positionPeriodDebt(uint256 period, address recipient, uint256 index, int24 tickLower, int24 tickUpper) view returns(int256 secondsDebtX96, int256 boostedSecondsDebtX96)
func (_RamsesV2Pool *RamsesV2PoolCaller) PositionPeriodDebt(opts *bind.CallOpts, period *big.Int, recipient common.Address, index *big.Int, tickLower *big.Int, tickUpper *big.Int) (struct {
	SecondsDebtX96        *big.Int
	BoostedSecondsDebtX96 *big.Int
}, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "positionPeriodDebt", period, recipient, index, tickLower, tickUpper)

	outstruct := new(struct {
		SecondsDebtX96        *big.Int
		BoostedSecondsDebtX96 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SecondsDebtX96 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BoostedSecondsDebtX96 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PositionPeriodDebt is a free data retrieval call binding the contract method 0xdfc8b615.
//
// Solidity: function positionPeriodDebt(uint256 period, address recipient, uint256 index, int24 tickLower, int24 tickUpper) view returns(int256 secondsDebtX96, int256 boostedSecondsDebtX96)
func (_RamsesV2Pool *RamsesV2PoolSession) PositionPeriodDebt(period *big.Int, recipient common.Address, index *big.Int, tickLower *big.Int, tickUpper *big.Int) (struct {
	SecondsDebtX96        *big.Int
	BoostedSecondsDebtX96 *big.Int
}, error) {
	return _RamsesV2Pool.Contract.PositionPeriodDebt(&_RamsesV2Pool.CallOpts, period, recipient, index, tickLower, tickUpper)
}

// PositionPeriodDebt is a free data retrieval call binding the contract method 0xdfc8b615.
//
// Solidity: function positionPeriodDebt(uint256 period, address recipient, uint256 index, int24 tickLower, int24 tickUpper) view returns(int256 secondsDebtX96, int256 boostedSecondsDebtX96)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) PositionPeriodDebt(period *big.Int, recipient common.Address, index *big.Int, tickLower *big.Int, tickUpper *big.Int) (struct {
	SecondsDebtX96        *big.Int
	BoostedSecondsDebtX96 *big.Int
}, error) {
	return _RamsesV2Pool.Contract.PositionPeriodDebt(&_RamsesV2Pool.CallOpts, period, recipient, index, tickLower, tickUpper)
}

// PositionPeriodSecondsInRange is a free data retrieval call binding the contract method 0x9918fbb6.
//
// Solidity: function positionPeriodSecondsInRange(uint256 period, address owner, uint256 index, int24 tickLower, int24 tickUpper) view returns(uint256 periodSecondsInsideX96, uint256 periodBoostedSecondsInsideX96)
func (_RamsesV2Pool *RamsesV2PoolCaller) PositionPeriodSecondsInRange(opts *bind.CallOpts, period *big.Int, owner common.Address, index *big.Int, tickLower *big.Int, tickUpper *big.Int) (struct {
	PeriodSecondsInsideX96        *big.Int
	PeriodBoostedSecondsInsideX96 *big.Int
}, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "positionPeriodSecondsInRange", period, owner, index, tickLower, tickUpper)

	outstruct := new(struct {
		PeriodSecondsInsideX96        *big.Int
		PeriodBoostedSecondsInsideX96 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PeriodSecondsInsideX96 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PeriodBoostedSecondsInsideX96 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PositionPeriodSecondsInRange is a free data retrieval call binding the contract method 0x9918fbb6.
//
// Solidity: function positionPeriodSecondsInRange(uint256 period, address owner, uint256 index, int24 tickLower, int24 tickUpper) view returns(uint256 periodSecondsInsideX96, uint256 periodBoostedSecondsInsideX96)
func (_RamsesV2Pool *RamsesV2PoolSession) PositionPeriodSecondsInRange(period *big.Int, owner common.Address, index *big.Int, tickLower *big.Int, tickUpper *big.Int) (struct {
	PeriodSecondsInsideX96        *big.Int
	PeriodBoostedSecondsInsideX96 *big.Int
}, error) {
	return _RamsesV2Pool.Contract.PositionPeriodSecondsInRange(&_RamsesV2Pool.CallOpts, period, owner, index, tickLower, tickUpper)
}

// PositionPeriodSecondsInRange is a free data retrieval call binding the contract method 0x9918fbb6.
//
// Solidity: function positionPeriodSecondsInRange(uint256 period, address owner, uint256 index, int24 tickLower, int24 tickUpper) view returns(uint256 periodSecondsInsideX96, uint256 periodBoostedSecondsInsideX96)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) PositionPeriodSecondsInRange(period *big.Int, owner common.Address, index *big.Int, tickLower *big.Int, tickUpper *big.Int) (struct {
	PeriodSecondsInsideX96        *big.Int
	PeriodBoostedSecondsInsideX96 *big.Int
}, error) {
	return _RamsesV2Pool.Contract.PositionPeriodSecondsInRange(&_RamsesV2Pool.CallOpts, period, owner, index, tickLower, tickUpper)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 key) view returns(uint128 _liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1, uint256 attachedVeRamId)
func (_RamsesV2Pool *RamsesV2PoolCaller) Positions(opts *bind.CallOpts, key [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
	AttachedVeRamId          *big.Int
}, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "positions", key)

	outstruct := new(struct {
		Liquidity                *big.Int
		FeeGrowthInside0LastX128 *big.Int
		FeeGrowthInside1LastX128 *big.Int
		TokensOwed0              *big.Int
		TokensOwed1              *big.Int
		AttachedVeRamId          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside0LastX128 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthInside1LastX128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed0 = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TokensOwed1 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AttachedVeRamId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 key) view returns(uint128 _liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1, uint256 attachedVeRamId)
func (_RamsesV2Pool *RamsesV2PoolSession) Positions(key [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
	AttachedVeRamId          *big.Int
}, error) {
	return _RamsesV2Pool.Contract.Positions(&_RamsesV2Pool.CallOpts, key)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 key) view returns(uint128 _liquidity, uint256 feeGrowthInside0LastX128, uint256 feeGrowthInside1LastX128, uint128 tokensOwed0, uint128 tokensOwed1, uint256 attachedVeRamId)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) Positions(key [32]byte) (struct {
	Liquidity                *big.Int
	FeeGrowthInside0LastX128 *big.Int
	FeeGrowthInside1LastX128 *big.Int
	TokensOwed0              *big.Int
	TokensOwed1              *big.Int
	AttachedVeRamId          *big.Int
}, error) {
	return _RamsesV2Pool.Contract.Positions(&_RamsesV2Pool.CallOpts, key)
}

// ProtocolFees is a free data retrieval call binding the contract method 0x1ad8b03b.
//
// Solidity: function protocolFees() view returns(uint128 token0, uint128 token1)
func (_RamsesV2Pool *RamsesV2PoolCaller) ProtocolFees(opts *bind.CallOpts) (struct {
	Token0 *big.Int
	Token1 *big.Int
}, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "protocolFees")

	outstruct := new(struct {
		Token0 *big.Int
		Token1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Token1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProtocolFees is a free data retrieval call binding the contract method 0x1ad8b03b.
//
// Solidity: function protocolFees() view returns(uint128 token0, uint128 token1)
func (_RamsesV2Pool *RamsesV2PoolSession) ProtocolFees() (struct {
	Token0 *big.Int
	Token1 *big.Int
}, error) {
	return _RamsesV2Pool.Contract.ProtocolFees(&_RamsesV2Pool.CallOpts)
}

// ProtocolFees is a free data retrieval call binding the contract method 0x1ad8b03b.
//
// Solidity: function protocolFees() view returns(uint128 token0, uint128 token1)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) ProtocolFees() (struct {
	Token0 *big.Int
	Token1 *big.Int
}, error) {
	return _RamsesV2Pool.Contract.ProtocolFees(&_RamsesV2Pool.CallOpts)
}

// ReadStorage is a free data retrieval call binding the contract method 0xe57c0ca9.
//
// Solidity: function readStorage(bytes32[] slots) view returns(bytes32[] returnData)
func (_RamsesV2Pool *RamsesV2PoolCaller) ReadStorage(opts *bind.CallOpts, slots [][32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "readStorage", slots)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ReadStorage is a free data retrieval call binding the contract method 0xe57c0ca9.
//
// Solidity: function readStorage(bytes32[] slots) view returns(bytes32[] returnData)
func (_RamsesV2Pool *RamsesV2PoolSession) ReadStorage(slots [][32]byte) ([][32]byte, error) {
	return _RamsesV2Pool.Contract.ReadStorage(&_RamsesV2Pool.CallOpts, slots)
}

// ReadStorage is a free data retrieval call binding the contract method 0xe57c0ca9.
//
// Solidity: function readStorage(bytes32[] slots) view returns(bytes32[] returnData)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) ReadStorage(slots [][32]byte) ([][32]byte, error) {
	return _RamsesV2Pool.Contract.ReadStorage(&_RamsesV2Pool.CallOpts, slots)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_RamsesV2Pool *RamsesV2PoolCaller) Slot0(opts *bind.CallOpts) (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "slot0")

	outstruct := new(struct {
		SqrtPriceX96               *big.Int
		Tick                       *big.Int
		ObservationIndex           uint16
		ObservationCardinality     uint16
		ObservationCardinalityNext uint16
		FeeProtocol                uint8
		Unlocked                   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SqrtPriceX96 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tick = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ObservationIndex = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.ObservationCardinality = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.ObservationCardinalityNext = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.FeeProtocol = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Unlocked = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_RamsesV2Pool *RamsesV2PoolSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	return _RamsesV2Pool.Contract.Slot0(&_RamsesV2Pool.CallOpts)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	return _RamsesV2Pool.Contract.Slot0(&_RamsesV2Pool.CallOpts)
}

// SnapshotCumulativesInside is a free data retrieval call binding the contract method 0xa38807f2.
//
// Solidity: function snapshotCumulativesInside(int24 tickLower, int24 tickUpper) view returns(int56 tickCumulativeInside, uint160 secondsPerLiquidityInsideX128, uint160 secondsPerBoostedLiquidityInsideX128, uint32 secondsInside)
func (_RamsesV2Pool *RamsesV2PoolCaller) SnapshotCumulativesInside(opts *bind.CallOpts, tickLower *big.Int, tickUpper *big.Int) (struct {
	TickCumulativeInside                 *big.Int
	SecondsPerLiquidityInsideX128        *big.Int
	SecondsPerBoostedLiquidityInsideX128 *big.Int
	SecondsInside                        uint32
}, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "snapshotCumulativesInside", tickLower, tickUpper)

	outstruct := new(struct {
		TickCumulativeInside                 *big.Int
		SecondsPerLiquidityInsideX128        *big.Int
		SecondsPerBoostedLiquidityInsideX128 *big.Int
		SecondsInside                        uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TickCumulativeInside = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityInsideX128 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerBoostedLiquidityInsideX128 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SecondsInside = *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return *outstruct, err

}

// SnapshotCumulativesInside is a free data retrieval call binding the contract method 0xa38807f2.
//
// Solidity: function snapshotCumulativesInside(int24 tickLower, int24 tickUpper) view returns(int56 tickCumulativeInside, uint160 secondsPerLiquidityInsideX128, uint160 secondsPerBoostedLiquidityInsideX128, uint32 secondsInside)
func (_RamsesV2Pool *RamsesV2PoolSession) SnapshotCumulativesInside(tickLower *big.Int, tickUpper *big.Int) (struct {
	TickCumulativeInside                 *big.Int
	SecondsPerLiquidityInsideX128        *big.Int
	SecondsPerBoostedLiquidityInsideX128 *big.Int
	SecondsInside                        uint32
}, error) {
	return _RamsesV2Pool.Contract.SnapshotCumulativesInside(&_RamsesV2Pool.CallOpts, tickLower, tickUpper)
}

// SnapshotCumulativesInside is a free data retrieval call binding the contract method 0xa38807f2.
//
// Solidity: function snapshotCumulativesInside(int24 tickLower, int24 tickUpper) view returns(int56 tickCumulativeInside, uint160 secondsPerLiquidityInsideX128, uint160 secondsPerBoostedLiquidityInsideX128, uint32 secondsInside)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) SnapshotCumulativesInside(tickLower *big.Int, tickUpper *big.Int) (struct {
	TickCumulativeInside                 *big.Int
	SecondsPerLiquidityInsideX128        *big.Int
	SecondsPerBoostedLiquidityInsideX128 *big.Int
	SecondsInside                        uint32
}, error) {
	return _RamsesV2Pool.Contract.SnapshotCumulativesInside(&_RamsesV2Pool.CallOpts, tickLower, tickUpper)
}

// TickBitmap is a free data retrieval call binding the contract method 0x5339c296.
//
// Solidity: function tickBitmap(int16 wordPosition) view returns(uint256)
func (_RamsesV2Pool *RamsesV2PoolCaller) TickBitmap(opts *bind.CallOpts, wordPosition int16) (*big.Int, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "tickBitmap", wordPosition)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickBitmap is a free data retrieval call binding the contract method 0x5339c296.
//
// Solidity: function tickBitmap(int16 wordPosition) view returns(uint256)
func (_RamsesV2Pool *RamsesV2PoolSession) TickBitmap(wordPosition int16) (*big.Int, error) {
	return _RamsesV2Pool.Contract.TickBitmap(&_RamsesV2Pool.CallOpts, wordPosition)
}

// TickBitmap is a free data retrieval call binding the contract method 0x5339c296.
//
// Solidity: function tickBitmap(int16 wordPosition) view returns(uint256)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) TickBitmap(wordPosition int16) (*big.Int, error) {
	return _RamsesV2Pool.Contract.TickBitmap(&_RamsesV2Pool.CallOpts, wordPosition)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_RamsesV2Pool *RamsesV2PoolCaller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_RamsesV2Pool *RamsesV2PoolSession) TickSpacing() (*big.Int, error) {
	return _RamsesV2Pool.Contract.TickSpacing(&_RamsesV2Pool.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) TickSpacing() (*big.Int, error) {
	return _RamsesV2Pool.Contract.TickSpacing(&_RamsesV2Pool.CallOpts)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 tick) view returns(uint128 liquidityGross, int128 liquidityNet, uint128 boostedLiquidityGross, int128 boostedLiquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, int56 tickCumulativeOutside, uint160 secondsPerLiquidityOutsideX128, uint32 secondsOutside, bool initialized)
func (_RamsesV2Pool *RamsesV2PoolCaller) Ticks(opts *bind.CallOpts, tick *big.Int) (struct {
	LiquidityGross                 *big.Int
	LiquidityNet                   *big.Int
	BoostedLiquidityGross          *big.Int
	BoostedLiquidityNet            *big.Int
	FeeGrowthOutside0X128          *big.Int
	FeeGrowthOutside1X128          *big.Int
	TickCumulativeOutside          *big.Int
	SecondsPerLiquidityOutsideX128 *big.Int
	SecondsOutside                 uint32
	Initialized                    bool
}, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "ticks", tick)

	outstruct := new(struct {
		LiquidityGross                 *big.Int
		LiquidityNet                   *big.Int
		BoostedLiquidityGross          *big.Int
		BoostedLiquidityNet            *big.Int
		FeeGrowthOutside0X128          *big.Int
		FeeGrowthOutside1X128          *big.Int
		TickCumulativeOutside          *big.Int
		SecondsPerLiquidityOutsideX128 *big.Int
		SecondsOutside                 uint32
		Initialized                    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidityGross = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidityNet = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BoostedLiquidityGross = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BoostedLiquidityNet = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthOutside0X128 = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.FeeGrowthOutside1X128 = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TickCumulativeOutside = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.SecondsPerLiquidityOutsideX128 = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.SecondsOutside = *abi.ConvertType(out[8], new(uint32)).(*uint32)
	outstruct.Initialized = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 tick) view returns(uint128 liquidityGross, int128 liquidityNet, uint128 boostedLiquidityGross, int128 boostedLiquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, int56 tickCumulativeOutside, uint160 secondsPerLiquidityOutsideX128, uint32 secondsOutside, bool initialized)
func (_RamsesV2Pool *RamsesV2PoolSession) Ticks(tick *big.Int) (struct {
	LiquidityGross                 *big.Int
	LiquidityNet                   *big.Int
	BoostedLiquidityGross          *big.Int
	BoostedLiquidityNet            *big.Int
	FeeGrowthOutside0X128          *big.Int
	FeeGrowthOutside1X128          *big.Int
	TickCumulativeOutside          *big.Int
	SecondsPerLiquidityOutsideX128 *big.Int
	SecondsOutside                 uint32
	Initialized                    bool
}, error) {
	return _RamsesV2Pool.Contract.Ticks(&_RamsesV2Pool.CallOpts, tick)
}

// Ticks is a free data retrieval call binding the contract method 0xf30dba93.
//
// Solidity: function ticks(int24 tick) view returns(uint128 liquidityGross, int128 liquidityNet, uint128 boostedLiquidityGross, int128 boostedLiquidityNet, uint256 feeGrowthOutside0X128, uint256 feeGrowthOutside1X128, int56 tickCumulativeOutside, uint160 secondsPerLiquidityOutsideX128, uint32 secondsOutside, bool initialized)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) Ticks(tick *big.Int) (struct {
	LiquidityGross                 *big.Int
	LiquidityNet                   *big.Int
	BoostedLiquidityGross          *big.Int
	BoostedLiquidityNet            *big.Int
	FeeGrowthOutside0X128          *big.Int
	FeeGrowthOutside1X128          *big.Int
	TickCumulativeOutside          *big.Int
	SecondsPerLiquidityOutsideX128 *big.Int
	SecondsOutside                 uint32
	Initialized                    bool
}, error) {
	return _RamsesV2Pool.Contract.Ticks(&_RamsesV2Pool.CallOpts, tick)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_RamsesV2Pool *RamsesV2PoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_RamsesV2Pool *RamsesV2PoolSession) Token0() (common.Address, error) {
	return _RamsesV2Pool.Contract.Token0(&_RamsesV2Pool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) Token0() (common.Address, error) {
	return _RamsesV2Pool.Contract.Token0(&_RamsesV2Pool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_RamsesV2Pool *RamsesV2PoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_RamsesV2Pool *RamsesV2PoolSession) Token1() (common.Address, error) {
	return _RamsesV2Pool.Contract.Token1(&_RamsesV2Pool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) Token1() (common.Address, error) {
	return _RamsesV2Pool.Contract.Token1(&_RamsesV2Pool.CallOpts)
}

// VeRam is a free data retrieval call binding the contract method 0x97e9dc31.
//
// Solidity: function veRam() view returns(address)
func (_RamsesV2Pool *RamsesV2PoolCaller) VeRam(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "veRam")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VeRam is a free data retrieval call binding the contract method 0x97e9dc31.
//
// Solidity: function veRam() view returns(address)
func (_RamsesV2Pool *RamsesV2PoolSession) VeRam() (common.Address, error) {
	return _RamsesV2Pool.Contract.VeRam(&_RamsesV2Pool.CallOpts)
}

// VeRam is a free data retrieval call binding the contract method 0x97e9dc31.
//
// Solidity: function veRam() view returns(address)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) VeRam() (common.Address, error) {
	return _RamsesV2Pool.Contract.VeRam(&_RamsesV2Pool.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_RamsesV2Pool *RamsesV2PoolCaller) Voter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RamsesV2Pool.contract.Call(opts, &out, "voter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_RamsesV2Pool *RamsesV2PoolSession) Voter() (common.Address, error) {
	return _RamsesV2Pool.Contract.Voter(&_RamsesV2Pool.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x46c96aac.
//
// Solidity: function voter() view returns(address)
func (_RamsesV2Pool *RamsesV2PoolCallerSession) Voter() (common.Address, error) {
	return _RamsesV2Pool.Contract.Voter(&_RamsesV2Pool.CallOpts)
}

// AdvancePeriod is a paid mutator transaction binding the contract method 0xc2e0f9b2.
//
// Solidity: function _advancePeriod() returns()
func (_RamsesV2Pool *RamsesV2PoolTransactor) AdvancePeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RamsesV2Pool.contract.Transact(opts, "_advancePeriod")
}

// AdvancePeriod is a paid mutator transaction binding the contract method 0xc2e0f9b2.
//
// Solidity: function _advancePeriod() returns()
func (_RamsesV2Pool *RamsesV2PoolSession) AdvancePeriod() (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.AdvancePeriod(&_RamsesV2Pool.TransactOpts)
}

// AdvancePeriod is a paid mutator transaction binding the contract method 0xc2e0f9b2.
//
// Solidity: function _advancePeriod() returns()
func (_RamsesV2Pool *RamsesV2PoolTransactorSession) AdvancePeriod() (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.AdvancePeriod(&_RamsesV2Pool.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x6847456a.
//
// Solidity: function burn(uint256 index, int24 tickLower, int24 tickUpper, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolTransactor) Burn(opts *bind.TransactOpts, index *big.Int, tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.contract.Transact(opts, "burn", index, tickLower, tickUpper, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x6847456a.
//
// Solidity: function burn(uint256 index, int24 tickLower, int24 tickUpper, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolSession) Burn(index *big.Int, tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Burn(&_RamsesV2Pool.TransactOpts, index, tickLower, tickUpper, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x6847456a.
//
// Solidity: function burn(uint256 index, int24 tickLower, int24 tickUpper, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolTransactorSession) Burn(index *big.Int, tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Burn(&_RamsesV2Pool.TransactOpts, index, tickLower, tickUpper, amount)
}

// Burn0 is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 tickLower, int24 tickUpper, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolTransactor) Burn0(opts *bind.TransactOpts, tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.contract.Transact(opts, "burn0", tickLower, tickUpper, amount)
}

// Burn0 is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 tickLower, int24 tickUpper, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolSession) Burn0(tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Burn0(&_RamsesV2Pool.TransactOpts, tickLower, tickUpper, amount)
}

// Burn0 is a paid mutator transaction binding the contract method 0xa34123a7.
//
// Solidity: function burn(int24 tickLower, int24 tickUpper, uint128 amount) returns(uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolTransactorSession) Burn0(tickLower *big.Int, tickUpper *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Burn0(&_RamsesV2Pool.TransactOpts, tickLower, tickUpper, amount)
}

// Burn1 is a paid mutator transaction binding the contract method 0xe3f9b398.
//
// Solidity: function burn(uint256 index, int24 tickLower, int24 tickUpper, uint128 amount, uint256 veRamTokenId) returns(uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolTransactor) Burn1(opts *bind.TransactOpts, index *big.Int, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, veRamTokenId *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.contract.Transact(opts, "burn1", index, tickLower, tickUpper, amount, veRamTokenId)
}

// Burn1 is a paid mutator transaction binding the contract method 0xe3f9b398.
//
// Solidity: function burn(uint256 index, int24 tickLower, int24 tickUpper, uint128 amount, uint256 veRamTokenId) returns(uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolSession) Burn1(index *big.Int, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, veRamTokenId *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Burn1(&_RamsesV2Pool.TransactOpts, index, tickLower, tickUpper, amount, veRamTokenId)
}

// Burn1 is a paid mutator transaction binding the contract method 0xe3f9b398.
//
// Solidity: function burn(uint256 index, int24 tickLower, int24 tickUpper, uint128 amount, uint256 veRamTokenId) returns(uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolTransactorSession) Burn1(index *big.Int, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, veRamTokenId *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Burn1(&_RamsesV2Pool.TransactOpts, index, tickLower, tickUpper, amount, veRamTokenId)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 tickLower, int24 tickUpper, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_RamsesV2Pool *RamsesV2PoolTransactor) Collect(opts *bind.TransactOpts, recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.contract.Transact(opts, "collect", recipient, tickLower, tickUpper, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 tickLower, int24 tickUpper, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_RamsesV2Pool *RamsesV2PoolSession) Collect(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Collect(&_RamsesV2Pool.TransactOpts, recipient, tickLower, tickUpper, amount0Requested, amount1Requested)
}

// Collect is a paid mutator transaction binding the contract method 0x4f1eb3d8.
//
// Solidity: function collect(address recipient, int24 tickLower, int24 tickUpper, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_RamsesV2Pool *RamsesV2PoolTransactorSession) Collect(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Collect(&_RamsesV2Pool.TransactOpts, recipient, tickLower, tickUpper, amount0Requested, amount1Requested)
}

// Collect0 is a paid mutator transaction binding the contract method 0xa02f1069.
//
// Solidity: function collect(address recipient, uint256 index, int24 tickLower, int24 tickUpper, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_RamsesV2Pool *RamsesV2PoolTransactor) Collect0(opts *bind.TransactOpts, recipient common.Address, index *big.Int, tickLower *big.Int, tickUpper *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.contract.Transact(opts, "collect0", recipient, index, tickLower, tickUpper, amount0Requested, amount1Requested)
}

// Collect0 is a paid mutator transaction binding the contract method 0xa02f1069.
//
// Solidity: function collect(address recipient, uint256 index, int24 tickLower, int24 tickUpper, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_RamsesV2Pool *RamsesV2PoolSession) Collect0(recipient common.Address, index *big.Int, tickLower *big.Int, tickUpper *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Collect0(&_RamsesV2Pool.TransactOpts, recipient, index, tickLower, tickUpper, amount0Requested, amount1Requested)
}

// Collect0 is a paid mutator transaction binding the contract method 0xa02f1069.
//
// Solidity: function collect(address recipient, uint256 index, int24 tickLower, int24 tickUpper, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_RamsesV2Pool *RamsesV2PoolTransactorSession) Collect0(recipient common.Address, index *big.Int, tickLower *big.Int, tickUpper *big.Int, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Collect0(&_RamsesV2Pool.TransactOpts, recipient, index, tickLower, tickUpper, amount0Requested, amount1Requested)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x85b66729.
//
// Solidity: function collectProtocol(address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_RamsesV2Pool *RamsesV2PoolTransactor) CollectProtocol(opts *bind.TransactOpts, recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.contract.Transact(opts, "collectProtocol", recipient, amount0Requested, amount1Requested)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x85b66729.
//
// Solidity: function collectProtocol(address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_RamsesV2Pool *RamsesV2PoolSession) CollectProtocol(recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.CollectProtocol(&_RamsesV2Pool.TransactOpts, recipient, amount0Requested, amount1Requested)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x85b66729.
//
// Solidity: function collectProtocol(address recipient, uint128 amount0Requested, uint128 amount1Requested) returns(uint128 amount0, uint128 amount1)
func (_RamsesV2Pool *RamsesV2PoolTransactorSession) CollectProtocol(recipient common.Address, amount0Requested *big.Int, amount1Requested *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.CollectProtocol(&_RamsesV2Pool.TransactOpts, recipient, amount0Requested, amount1Requested)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_RamsesV2Pool *RamsesV2PoolTransactor) Flash(opts *bind.TransactOpts, recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _RamsesV2Pool.contract.Transact(opts, "flash", recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_RamsesV2Pool *RamsesV2PoolSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Flash(&_RamsesV2Pool.TransactOpts, recipient, amount0, amount1, data)
}

// Flash is a paid mutator transaction binding the contract method 0x490e6cbc.
//
// Solidity: function flash(address recipient, uint256 amount0, uint256 amount1, bytes data) returns()
func (_RamsesV2Pool *RamsesV2PoolTransactorSession) Flash(recipient common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Flash(&_RamsesV2Pool.TransactOpts, recipient, amount0, amount1, data)
}

// IncreaseObservationCardinalityNext is a paid mutator transaction binding the contract method 0x32148f67.
//
// Solidity: function increaseObservationCardinalityNext(uint16 observationCardinalityNext) returns()
func (_RamsesV2Pool *RamsesV2PoolTransactor) IncreaseObservationCardinalityNext(opts *bind.TransactOpts, observationCardinalityNext uint16) (*types.Transaction, error) {
	return _RamsesV2Pool.contract.Transact(opts, "increaseObservationCardinalityNext", observationCardinalityNext)
}

// IncreaseObservationCardinalityNext is a paid mutator transaction binding the contract method 0x32148f67.
//
// Solidity: function increaseObservationCardinalityNext(uint16 observationCardinalityNext) returns()
func (_RamsesV2Pool *RamsesV2PoolSession) IncreaseObservationCardinalityNext(observationCardinalityNext uint16) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.IncreaseObservationCardinalityNext(&_RamsesV2Pool.TransactOpts, observationCardinalityNext)
}

// IncreaseObservationCardinalityNext is a paid mutator transaction binding the contract method 0x32148f67.
//
// Solidity: function increaseObservationCardinalityNext(uint16 observationCardinalityNext) returns()
func (_RamsesV2Pool *RamsesV2PoolTransactorSession) IncreaseObservationCardinalityNext(observationCardinalityNext uint16) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.IncreaseObservationCardinalityNext(&_RamsesV2Pool.TransactOpts, observationCardinalityNext)
}

// Initialize is a paid mutator transaction binding the contract method 0x5fa9c925.
//
// Solidity: function initialize(address _factory, address _nfpManager, address _veRam, address _voter, address _token0, address _token1, uint24 _fee, int24 _tickSpacing) returns()
func (_RamsesV2Pool *RamsesV2PoolTransactor) Initialize(opts *bind.TransactOpts, _factory common.Address, _nfpManager common.Address, _veRam common.Address, _voter common.Address, _token0 common.Address, _token1 common.Address, _fee *big.Int, _tickSpacing *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.contract.Transact(opts, "initialize", _factory, _nfpManager, _veRam, _voter, _token0, _token1, _fee, _tickSpacing)
}

// Initialize is a paid mutator transaction binding the contract method 0x5fa9c925.
//
// Solidity: function initialize(address _factory, address _nfpManager, address _veRam, address _voter, address _token0, address _token1, uint24 _fee, int24 _tickSpacing) returns()
func (_RamsesV2Pool *RamsesV2PoolSession) Initialize(_factory common.Address, _nfpManager common.Address, _veRam common.Address, _voter common.Address, _token0 common.Address, _token1 common.Address, _fee *big.Int, _tickSpacing *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Initialize(&_RamsesV2Pool.TransactOpts, _factory, _nfpManager, _veRam, _voter, _token0, _token1, _fee, _tickSpacing)
}

// Initialize is a paid mutator transaction binding the contract method 0x5fa9c925.
//
// Solidity: function initialize(address _factory, address _nfpManager, address _veRam, address _voter, address _token0, address _token1, uint24 _fee, int24 _tickSpacing) returns()
func (_RamsesV2Pool *RamsesV2PoolTransactorSession) Initialize(_factory common.Address, _nfpManager common.Address, _veRam common.Address, _voter common.Address, _token0 common.Address, _token1 common.Address, _fee *big.Int, _tickSpacing *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Initialize(&_RamsesV2Pool.TransactOpts, _factory, _nfpManager, _veRam, _voter, _token0, _token1, _fee, _tickSpacing)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 sqrtPriceX96) returns()
func (_RamsesV2Pool *RamsesV2PoolTransactor) Initialize0(opts *bind.TransactOpts, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.contract.Transact(opts, "initialize0", sqrtPriceX96)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 sqrtPriceX96) returns()
func (_RamsesV2Pool *RamsesV2PoolSession) Initialize0(sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Initialize0(&_RamsesV2Pool.TransactOpts, sqrtPriceX96)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xf637731d.
//
// Solidity: function initialize(uint160 sqrtPriceX96) returns()
func (_RamsesV2Pool *RamsesV2PoolTransactorSession) Initialize0(sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Initialize0(&_RamsesV2Pool.TransactOpts, sqrtPriceX96)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 tickLower, int24 tickUpper, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _RamsesV2Pool.contract.Transact(opts, "mint", recipient, tickLower, tickUpper, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 tickLower, int24 tickUpper, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolSession) Mint(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Mint(&_RamsesV2Pool.TransactOpts, recipient, tickLower, tickUpper, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x3c8a7d8d.
//
// Solidity: function mint(address recipient, int24 tickLower, int24 tickUpper, uint128 amount, bytes data) returns(uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolTransactorSession) Mint(recipient common.Address, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Mint(&_RamsesV2Pool.TransactOpts, recipient, tickLower, tickUpper, amount, data)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa418e9e0.
//
// Solidity: function mint(address recipient, uint256 index, int24 tickLower, int24 tickUpper, uint128 amount, uint256 veRamTokenId, bytes data) returns(uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolTransactor) Mint0(opts *bind.TransactOpts, recipient common.Address, index *big.Int, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, veRamTokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _RamsesV2Pool.contract.Transact(opts, "mint0", recipient, index, tickLower, tickUpper, amount, veRamTokenId, data)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa418e9e0.
//
// Solidity: function mint(address recipient, uint256 index, int24 tickLower, int24 tickUpper, uint128 amount, uint256 veRamTokenId, bytes data) returns(uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolSession) Mint0(recipient common.Address, index *big.Int, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, veRamTokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Mint0(&_RamsesV2Pool.TransactOpts, recipient, index, tickLower, tickUpper, amount, veRamTokenId, data)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa418e9e0.
//
// Solidity: function mint(address recipient, uint256 index, int24 tickLower, int24 tickUpper, uint128 amount, uint256 veRamTokenId, bytes data) returns(uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolTransactorSession) Mint0(recipient common.Address, index *big.Int, tickLower *big.Int, tickUpper *big.Int, amount *big.Int, veRamTokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Mint0(&_RamsesV2Pool.TransactOpts, recipient, index, tickLower, tickUpper, amount, veRamTokenId, data)
}

// SetFee is a paid mutator transaction binding the contract method 0xeabb5622.
//
// Solidity: function setFee(uint24 _fee) returns()
func (_RamsesV2Pool *RamsesV2PoolTransactor) SetFee(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.contract.Transact(opts, "setFee", _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0xeabb5622.
//
// Solidity: function setFee(uint24 _fee) returns()
func (_RamsesV2Pool *RamsesV2PoolSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.SetFee(&_RamsesV2Pool.TransactOpts, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0xeabb5622.
//
// Solidity: function setFee(uint24 _fee) returns()
func (_RamsesV2Pool *RamsesV2PoolTransactorSession) SetFee(_fee *big.Int) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.SetFee(&_RamsesV2Pool.TransactOpts, _fee)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x7b7d549d.
//
// Solidity: function setFeeProtocol() returns()
func (_RamsesV2Pool *RamsesV2PoolTransactor) SetFeeProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RamsesV2Pool.contract.Transact(opts, "setFeeProtocol")
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x7b7d549d.
//
// Solidity: function setFeeProtocol() returns()
func (_RamsesV2Pool *RamsesV2PoolSession) SetFeeProtocol() (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.SetFeeProtocol(&_RamsesV2Pool.TransactOpts)
}

// SetFeeProtocol is a paid mutator transaction binding the contract method 0x7b7d549d.
//
// Solidity: function setFeeProtocol() returns()
func (_RamsesV2Pool *RamsesV2PoolTransactorSession) SetFeeProtocol() (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.SetFeeProtocol(&_RamsesV2Pool.TransactOpts)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_RamsesV2Pool *RamsesV2PoolTransactor) Swap(opts *bind.TransactOpts, recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _RamsesV2Pool.contract.Transact(opts, "swap", recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_RamsesV2Pool *RamsesV2PoolSession) Swap(recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Swap(&_RamsesV2Pool.TransactOpts, recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// Swap is a paid mutator transaction binding the contract method 0x128acb08.
//
// Solidity: function swap(address recipient, bool zeroForOne, int256 amountSpecified, uint160 sqrtPriceLimitX96, bytes data) returns(int256 amount0, int256 amount1)
func (_RamsesV2Pool *RamsesV2PoolTransactorSession) Swap(recipient common.Address, zeroForOne bool, amountSpecified *big.Int, sqrtPriceLimitX96 *big.Int, data []byte) (*types.Transaction, error) {
	return _RamsesV2Pool.Contract.Swap(&_RamsesV2Pool.TransactOpts, recipient, zeroForOne, amountSpecified, sqrtPriceLimitX96, data)
}

// RamsesV2PoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the RamsesV2Pool contract.
type RamsesV2PoolBurnIterator struct {
	Event *RamsesV2PoolBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RamsesV2PoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RamsesV2PoolBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RamsesV2PoolBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RamsesV2PoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RamsesV2PoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RamsesV2PoolBurn represents a Burn event raised by the RamsesV2Pool contract.
type RamsesV2PoolBurn struct {
	Owner     common.Address
	TickLower *big.Int
	TickUpper *big.Int
	Amount    *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolFilterer) FilterBurn(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*RamsesV2PoolBurnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _RamsesV2Pool.contract.FilterLogs(opts, "Burn", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &RamsesV2PoolBurnIterator{contract: _RamsesV2Pool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *RamsesV2PoolBurn, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _RamsesV2Pool.contract.WatchLogs(opts, "Burn", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RamsesV2PoolBurn)
				if err := _RamsesV2Pool.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurn is a log parse operation binding the contract event 0x0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c.
//
// Solidity: event Burn(address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolFilterer) ParseBurn(log types.Log) (*RamsesV2PoolBurn, error) {
	event := new(RamsesV2PoolBurn)
	if err := _RamsesV2Pool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RamsesV2PoolCollectIterator is returned from FilterCollect and is used to iterate over the raw logs and unpacked data for Collect events raised by the RamsesV2Pool contract.
type RamsesV2PoolCollectIterator struct {
	Event *RamsesV2PoolCollect // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RamsesV2PoolCollectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RamsesV2PoolCollect)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RamsesV2PoolCollect)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RamsesV2PoolCollectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RamsesV2PoolCollectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RamsesV2PoolCollect represents a Collect event raised by the RamsesV2Pool contract.
type RamsesV2PoolCollect struct {
	Owner     common.Address
	Recipient common.Address
	TickLower *big.Int
	TickUpper *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollect is a free log retrieval operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_RamsesV2Pool *RamsesV2PoolFilterer) FilterCollect(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*RamsesV2PoolCollectIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _RamsesV2Pool.contract.FilterLogs(opts, "Collect", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &RamsesV2PoolCollectIterator{contract: _RamsesV2Pool.contract, event: "Collect", logs: logs, sub: sub}, nil
}

// WatchCollect is a free log subscription operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_RamsesV2Pool *RamsesV2PoolFilterer) WatchCollect(opts *bind.WatchOpts, sink chan<- *RamsesV2PoolCollect, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _RamsesV2Pool.contract.WatchLogs(opts, "Collect", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RamsesV2PoolCollect)
				if err := _RamsesV2Pool.contract.UnpackLog(event, "Collect", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollect is a log parse operation binding the contract event 0x70935338e69775456a85ddef226c395fb668b63fa0115f5f20610b388e6ca9c0.
//
// Solidity: event Collect(address indexed owner, address recipient, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount0, uint128 amount1)
func (_RamsesV2Pool *RamsesV2PoolFilterer) ParseCollect(log types.Log) (*RamsesV2PoolCollect, error) {
	event := new(RamsesV2PoolCollect)
	if err := _RamsesV2Pool.contract.UnpackLog(event, "Collect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RamsesV2PoolCollectProtocolIterator is returned from FilterCollectProtocol and is used to iterate over the raw logs and unpacked data for CollectProtocol events raised by the RamsesV2Pool contract.
type RamsesV2PoolCollectProtocolIterator struct {
	Event *RamsesV2PoolCollectProtocol // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RamsesV2PoolCollectProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RamsesV2PoolCollectProtocol)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RamsesV2PoolCollectProtocol)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RamsesV2PoolCollectProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RamsesV2PoolCollectProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RamsesV2PoolCollectProtocol represents a CollectProtocol event raised by the RamsesV2Pool contract.
type RamsesV2PoolCollectProtocol struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectProtocol is a free log retrieval operation binding the contract event 0x596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151.
//
// Solidity: event CollectProtocol(address indexed sender, address indexed recipient, uint128 amount0, uint128 amount1)
func (_RamsesV2Pool *RamsesV2PoolFilterer) FilterCollectProtocol(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*RamsesV2PoolCollectProtocolIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RamsesV2Pool.contract.FilterLogs(opts, "CollectProtocol", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &RamsesV2PoolCollectProtocolIterator{contract: _RamsesV2Pool.contract, event: "CollectProtocol", logs: logs, sub: sub}, nil
}

// WatchCollectProtocol is a free log subscription operation binding the contract event 0x596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151.
//
// Solidity: event CollectProtocol(address indexed sender, address indexed recipient, uint128 amount0, uint128 amount1)
func (_RamsesV2Pool *RamsesV2PoolFilterer) WatchCollectProtocol(opts *bind.WatchOpts, sink chan<- *RamsesV2PoolCollectProtocol, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RamsesV2Pool.contract.WatchLogs(opts, "CollectProtocol", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RamsesV2PoolCollectProtocol)
				if err := _RamsesV2Pool.contract.UnpackLog(event, "CollectProtocol", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollectProtocol is a log parse operation binding the contract event 0x596b573906218d3411850b26a6b437d6c4522fdb43d2d2386263f86d50b8b151.
//
// Solidity: event CollectProtocol(address indexed sender, address indexed recipient, uint128 amount0, uint128 amount1)
func (_RamsesV2Pool *RamsesV2PoolFilterer) ParseCollectProtocol(log types.Log) (*RamsesV2PoolCollectProtocol, error) {
	event := new(RamsesV2PoolCollectProtocol)
	if err := _RamsesV2Pool.contract.UnpackLog(event, "CollectProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RamsesV2PoolFlashIterator is returned from FilterFlash and is used to iterate over the raw logs and unpacked data for Flash events raised by the RamsesV2Pool contract.
type RamsesV2PoolFlashIterator struct {
	Event *RamsesV2PoolFlash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RamsesV2PoolFlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RamsesV2PoolFlash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RamsesV2PoolFlash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RamsesV2PoolFlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RamsesV2PoolFlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RamsesV2PoolFlash represents a Flash event raised by the RamsesV2Pool contract.
type RamsesV2PoolFlash struct {
	Sender    common.Address
	Recipient common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Paid0     *big.Int
	Paid1     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlash is a free log retrieval operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_RamsesV2Pool *RamsesV2PoolFilterer) FilterFlash(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*RamsesV2PoolFlashIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RamsesV2Pool.contract.FilterLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &RamsesV2PoolFlashIterator{contract: _RamsesV2Pool.contract, event: "Flash", logs: logs, sub: sub}, nil
}

// WatchFlash is a free log subscription operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_RamsesV2Pool *RamsesV2PoolFilterer) WatchFlash(opts *bind.WatchOpts, sink chan<- *RamsesV2PoolFlash, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RamsesV2Pool.contract.WatchLogs(opts, "Flash", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RamsesV2PoolFlash)
				if err := _RamsesV2Pool.contract.UnpackLog(event, "Flash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFlash is a log parse operation binding the contract event 0xbdbdb71d7860376ba52b25a5028beea23581364a40522f6bcfb86bb1f2dca633.
//
// Solidity: event Flash(address indexed sender, address indexed recipient, uint256 amount0, uint256 amount1, uint256 paid0, uint256 paid1)
func (_RamsesV2Pool *RamsesV2PoolFilterer) ParseFlash(log types.Log) (*RamsesV2PoolFlash, error) {
	event := new(RamsesV2PoolFlash)
	if err := _RamsesV2Pool.contract.UnpackLog(event, "Flash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RamsesV2PoolIncreaseObservationCardinalityNextIterator is returned from FilterIncreaseObservationCardinalityNext and is used to iterate over the raw logs and unpacked data for IncreaseObservationCardinalityNext events raised by the RamsesV2Pool contract.
type RamsesV2PoolIncreaseObservationCardinalityNextIterator struct {
	Event *RamsesV2PoolIncreaseObservationCardinalityNext // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RamsesV2PoolIncreaseObservationCardinalityNextIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RamsesV2PoolIncreaseObservationCardinalityNext)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RamsesV2PoolIncreaseObservationCardinalityNext)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RamsesV2PoolIncreaseObservationCardinalityNextIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RamsesV2PoolIncreaseObservationCardinalityNextIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RamsesV2PoolIncreaseObservationCardinalityNext represents a IncreaseObservationCardinalityNext event raised by the RamsesV2Pool contract.
type RamsesV2PoolIncreaseObservationCardinalityNext struct {
	ObservationCardinalityNextOld uint16
	ObservationCardinalityNextNew uint16
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterIncreaseObservationCardinalityNext is a free log retrieval operation binding the contract event 0xac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a.
//
// Solidity: event IncreaseObservationCardinalityNext(uint16 observationCardinalityNextOld, uint16 observationCardinalityNextNew)
func (_RamsesV2Pool *RamsesV2PoolFilterer) FilterIncreaseObservationCardinalityNext(opts *bind.FilterOpts) (*RamsesV2PoolIncreaseObservationCardinalityNextIterator, error) {

	logs, sub, err := _RamsesV2Pool.contract.FilterLogs(opts, "IncreaseObservationCardinalityNext")
	if err != nil {
		return nil, err
	}
	return &RamsesV2PoolIncreaseObservationCardinalityNextIterator{contract: _RamsesV2Pool.contract, event: "IncreaseObservationCardinalityNext", logs: logs, sub: sub}, nil
}

// WatchIncreaseObservationCardinalityNext is a free log subscription operation binding the contract event 0xac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a.
//
// Solidity: event IncreaseObservationCardinalityNext(uint16 observationCardinalityNextOld, uint16 observationCardinalityNextNew)
func (_RamsesV2Pool *RamsesV2PoolFilterer) WatchIncreaseObservationCardinalityNext(opts *bind.WatchOpts, sink chan<- *RamsesV2PoolIncreaseObservationCardinalityNext) (event.Subscription, error) {

	logs, sub, err := _RamsesV2Pool.contract.WatchLogs(opts, "IncreaseObservationCardinalityNext")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RamsesV2PoolIncreaseObservationCardinalityNext)
				if err := _RamsesV2Pool.contract.UnpackLog(event, "IncreaseObservationCardinalityNext", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncreaseObservationCardinalityNext is a log parse operation binding the contract event 0xac49e518f90a358f652e4400164f05a5d8f7e35e7747279bc3a93dbf584e125a.
//
// Solidity: event IncreaseObservationCardinalityNext(uint16 observationCardinalityNextOld, uint16 observationCardinalityNextNew)
func (_RamsesV2Pool *RamsesV2PoolFilterer) ParseIncreaseObservationCardinalityNext(log types.Log) (*RamsesV2PoolIncreaseObservationCardinalityNext, error) {
	event := new(RamsesV2PoolIncreaseObservationCardinalityNext)
	if err := _RamsesV2Pool.contract.UnpackLog(event, "IncreaseObservationCardinalityNext", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RamsesV2PoolInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the RamsesV2Pool contract.
type RamsesV2PoolInitializeIterator struct {
	Event *RamsesV2PoolInitialize // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RamsesV2PoolInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RamsesV2PoolInitialize)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RamsesV2PoolInitialize)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RamsesV2PoolInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RamsesV2PoolInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RamsesV2PoolInitialize represents a Initialize event raised by the RamsesV2Pool contract.
type RamsesV2PoolInitialize struct {
	SqrtPriceX96 *big.Int
	Tick         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 sqrtPriceX96, int24 tick)
func (_RamsesV2Pool *RamsesV2PoolFilterer) FilterInitialize(opts *bind.FilterOpts) (*RamsesV2PoolInitializeIterator, error) {

	logs, sub, err := _RamsesV2Pool.contract.FilterLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return &RamsesV2PoolInitializeIterator{contract: _RamsesV2Pool.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 sqrtPriceX96, int24 tick)
func (_RamsesV2Pool *RamsesV2PoolFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *RamsesV2PoolInitialize) (event.Subscription, error) {

	logs, sub, err := _RamsesV2Pool.contract.WatchLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RamsesV2PoolInitialize)
				if err := _RamsesV2Pool.contract.UnpackLog(event, "Initialize", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialize is a log parse operation binding the contract event 0x98636036cb66a9c19a37435efc1e90142190214e8abeb821bdba3f2990dd4c95.
//
// Solidity: event Initialize(uint160 sqrtPriceX96, int24 tick)
func (_RamsesV2Pool *RamsesV2PoolFilterer) ParseInitialize(log types.Log) (*RamsesV2PoolInitialize, error) {
	event := new(RamsesV2PoolInitialize)
	if err := _RamsesV2Pool.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RamsesV2PoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the RamsesV2Pool contract.
type RamsesV2PoolMintIterator struct {
	Event *RamsesV2PoolMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RamsesV2PoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RamsesV2PoolMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RamsesV2PoolMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RamsesV2PoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RamsesV2PoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RamsesV2PoolMint represents a Mint event raised by the RamsesV2Pool contract.
type RamsesV2PoolMint struct {
	Sender    common.Address
	Owner     common.Address
	TickLower *big.Int
	TickUpper *big.Int
	Amount    *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolFilterer) FilterMint(opts *bind.FilterOpts, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (*RamsesV2PoolMintIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _RamsesV2Pool.contract.FilterLogs(opts, "Mint", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return &RamsesV2PoolMintIterator{contract: _RamsesV2Pool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *RamsesV2PoolMint, owner []common.Address, tickLower []*big.Int, tickUpper []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tickLowerRule []interface{}
	for _, tickLowerItem := range tickLower {
		tickLowerRule = append(tickLowerRule, tickLowerItem)
	}
	var tickUpperRule []interface{}
	for _, tickUpperItem := range tickUpper {
		tickUpperRule = append(tickUpperRule, tickUpperItem)
	}

	logs, sub, err := _RamsesV2Pool.contract.WatchLogs(opts, "Mint", ownerRule, tickLowerRule, tickUpperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RamsesV2PoolMint)
				if err := _RamsesV2Pool.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMint is a log parse operation binding the contract event 0x7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde.
//
// Solidity: event Mint(address sender, address indexed owner, int24 indexed tickLower, int24 indexed tickUpper, uint128 amount, uint256 amount0, uint256 amount1)
func (_RamsesV2Pool *RamsesV2PoolFilterer) ParseMint(log types.Log) (*RamsesV2PoolMint, error) {
	event := new(RamsesV2PoolMint)
	if err := _RamsesV2Pool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RamsesV2PoolSetFeeProtocolIterator is returned from FilterSetFeeProtocol and is used to iterate over the raw logs and unpacked data for SetFeeProtocol events raised by the RamsesV2Pool contract.
type RamsesV2PoolSetFeeProtocolIterator struct {
	Event *RamsesV2PoolSetFeeProtocol // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RamsesV2PoolSetFeeProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RamsesV2PoolSetFeeProtocol)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RamsesV2PoolSetFeeProtocol)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RamsesV2PoolSetFeeProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RamsesV2PoolSetFeeProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RamsesV2PoolSetFeeProtocol represents a SetFeeProtocol event raised by the RamsesV2Pool contract.
type RamsesV2PoolSetFeeProtocol struct {
	FeeProtocol0Old uint8
	FeeProtocol1Old uint8
	FeeProtocol0New uint8
	FeeProtocol1New uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetFeeProtocol is a free log retrieval operation binding the contract event 0x973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133.
//
// Solidity: event SetFeeProtocol(uint8 feeProtocol0Old, uint8 feeProtocol1Old, uint8 feeProtocol0New, uint8 feeProtocol1New)
func (_RamsesV2Pool *RamsesV2PoolFilterer) FilterSetFeeProtocol(opts *bind.FilterOpts) (*RamsesV2PoolSetFeeProtocolIterator, error) {

	logs, sub, err := _RamsesV2Pool.contract.FilterLogs(opts, "SetFeeProtocol")
	if err != nil {
		return nil, err
	}
	return &RamsesV2PoolSetFeeProtocolIterator{contract: _RamsesV2Pool.contract, event: "SetFeeProtocol", logs: logs, sub: sub}, nil
}

// WatchSetFeeProtocol is a free log subscription operation binding the contract event 0x973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133.
//
// Solidity: event SetFeeProtocol(uint8 feeProtocol0Old, uint8 feeProtocol1Old, uint8 feeProtocol0New, uint8 feeProtocol1New)
func (_RamsesV2Pool *RamsesV2PoolFilterer) WatchSetFeeProtocol(opts *bind.WatchOpts, sink chan<- *RamsesV2PoolSetFeeProtocol) (event.Subscription, error) {

	logs, sub, err := _RamsesV2Pool.contract.WatchLogs(opts, "SetFeeProtocol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RamsesV2PoolSetFeeProtocol)
				if err := _RamsesV2Pool.contract.UnpackLog(event, "SetFeeProtocol", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFeeProtocol is a log parse operation binding the contract event 0x973d8d92bb299f4af6ce49b52a8adb85ae46b9f214c4c4fc06ac77401237b133.
//
// Solidity: event SetFeeProtocol(uint8 feeProtocol0Old, uint8 feeProtocol1Old, uint8 feeProtocol0New, uint8 feeProtocol1New)
func (_RamsesV2Pool *RamsesV2PoolFilterer) ParseSetFeeProtocol(log types.Log) (*RamsesV2PoolSetFeeProtocol, error) {
	event := new(RamsesV2PoolSetFeeProtocol)
	if err := _RamsesV2Pool.contract.UnpackLog(event, "SetFeeProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RamsesV2PoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the RamsesV2Pool contract.
type RamsesV2PoolSwapIterator struct {
	Event *RamsesV2PoolSwap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RamsesV2PoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RamsesV2PoolSwap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RamsesV2PoolSwap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RamsesV2PoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RamsesV2PoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RamsesV2PoolSwap represents a Swap event raised by the RamsesV2Pool contract.
type RamsesV2PoolSwap struct {
	Sender       common.Address
	Recipient    common.Address
	Amount0      *big.Int
	Amount1      *big.Int
	SqrtPriceX96 *big.Int
	Liquidity    *big.Int
	Tick         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_RamsesV2Pool *RamsesV2PoolFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*RamsesV2PoolSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RamsesV2Pool.contract.FilterLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &RamsesV2PoolSwapIterator{contract: _RamsesV2Pool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_RamsesV2Pool *RamsesV2PoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *RamsesV2PoolSwap, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RamsesV2Pool.contract.WatchLogs(opts, "Swap", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RamsesV2PoolSwap)
				if err := _RamsesV2Pool.contract.UnpackLog(event, "Swap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwap is a log parse operation binding the contract event 0xc42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67.
//
// Solidity: event Swap(address indexed sender, address indexed recipient, int256 amount0, int256 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick)
func (_RamsesV2Pool *RamsesV2PoolFilterer) ParseSwap(log types.Log) (*RamsesV2PoolSwap, error) {
	event := new(RamsesV2PoolSwap)
	if err := _RamsesV2Pool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
