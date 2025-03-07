// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

import "forge-std/Script.sol";
import "forge-std/Test.sol";

import {DemoSolver} from "../src/DemoSolver.sol";

contract DeployDemoSolverScript is Test {
    function run() external {
        console.log("\n=== DEPLOYING DemoSolver ===\n");

        uint256 deployerPrivateKey = vm.envUint("DEPLOYER_PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);

        console.log("Deployer address: \t\t\t\t", deployer);

        address atlasAddress = vm.envAddress("ATLAS_ADDRESS");
        address wethAddress = vm.envAddress("WETH_ADDRESS");
        address routerAddress = vm.envAddress("ROUTER_ADDRESS");

        require(atlasAddress != address(0), "ATLAS_ADDRESS is not set");
        require(wethAddress != address(0), "WETH_ADDRESS is not set");
        require(routerAddress != address(0), "ROUTER_ADDRESS is not set");

        console.log("Atlas address: \t\t\t\t", atlasAddress);
        console.log("Weth address: \t\t\t\t", wethAddress);
        console.log("Router address: \t\t\t\t", routerAddress);

        console.log("Deploying from deployer Account...");

        vm.startBroadcast(deployerPrivateKey);

        DemoSolver demoSolver = new DemoSolver(wethAddress, atlasAddress, routerAddress);

        vm.stopBroadcast();

        console.log("Contracts deployed by deployer:");
        console.log("DemoSolver: \t\t\t\t", address(demoSolver));
        console.log("\n");
    }
}
