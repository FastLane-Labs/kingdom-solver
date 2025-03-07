//SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {Test} from "forge-std/Test.sol";
import {IAtlas} from "@atlas/interfaces/IAtlas.sol";
import {DemoSolver} from "../src/DemoSolver.sol";

interface IERC20 {
    function balanceOf(address account) external view returns (uint256);
}

address constant ATLAS = address(0xf0E388C7DFfE14a61280a4E5b84d77be3d2875e3);
address constant ROUTER = address(0xAA23611badAFB62D37E7295A682D21960ac85A90);
address constant WETH_USDC_POOL = address(0x30AFBcF9458c3131A6d051C621E307E6278E4110);
uint24 constant WETH_USDC_ORIGINAL_FEE = 500;
address constant WETH = address(0x82aF49447D8a07e3bd95BD0d56f35241523fBab1);
address constant USDC = address(0xaf88d065e77c8cC2239327C5EDb3A432268e5831);

contract DemoSolverTest is Test {
    address solver = makeAddr("solver");
    DemoSolver demoSolver;

    function setUp() public {
        vm.createSelectFork(vm.envString("ARBITRUM_RPC_URL"), 306_921_677); // Feb-17-2025 06:57:17 AM +UTC

        vm.deal(solver, 100 ether);

        vm.startPrank(solver);
        IAtlas(ATLAS).depositAndBond{value: 100 ether}(100 ether);
        demoSolver = new DemoSolver(WETH, ATLAS, ROUTER);
        vm.stopPrank();
    }

    function test_demoSwap() public {
        uint256 dealAmount = 100 ether;

        deal(WETH, address(demoSolver), dealAmount);
        deal(USDC, address(demoSolver), dealAmount);

        uint256 balanceWethBefore = IERC20(WETH).balanceOf(address(demoSolver));
        uint256 balanceUsdcBefore = IERC20(USDC).balanceOf(address(demoSolver));

        assertEq(balanceWethBefore, dealAmount);
        assertEq(balanceUsdcBefore, dealAmount);

        vm.prank(address(demoSolver));
        demoSolver.demoSwap(WETH_USDC_POOL, WETH, USDC, 100 ether, WETH_USDC_ORIGINAL_FEE, false);

        uint256 balanceWethAfter = IERC20(WETH).balanceOf(address(demoSolver));
        uint256 balanceUsdcAfter = IERC20(USDC).balanceOf(address(demoSolver));

        assertEq(balanceWethAfter, 0);
        assertGt(balanceUsdcAfter, balanceUsdcBefore);
    }
}
