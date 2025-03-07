//SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import {SolverBase} from "@atlas/solver/SolverBase.sol";

interface IERC20 {
    function approve(address spender, uint256 amount) external returns (bool);
}

interface IRamsesV2Pool {
    function currentFee() external view returns (uint24);
}

interface IUniswapV3Router {
    struct ExactInputSingleParams {
        address tokenIn;
        address tokenOut;
        uint24 fee;
        address recipient;
        uint256 deadline;
        uint256 amountIn;
        uint256 amountOutMinimum;
        uint160 sqrtPriceLimitX96;
    }

    function exactInputSingle(ExactInputSingleParams memory params) external payable returns (uint256 amountOut);
}

contract DemoSolver is SolverBase {
    address public immutable router;

    constructor(address weth_, address atlas_, address router_) SolverBase(weth_, atlas_, msg.sender) {
        router = router_;
    }

    function demoSwap(
        address pool,
        address tokenIn,
        address tokenOut,
        uint256 amountIn,
        uint24 originalFee,
        bool checkCurrentFee
    ) external {
        require(msg.sender == _owner || msg.sender == address(this), "Invalid caller");

        if (checkCurrentFee) {
            require(IRamsesV2Pool(pool).currentFee() == 0, "Pool fee not 0");
        }

        IERC20(tokenIn).approve(router, amountIn);

        IUniswapV3Router(router).exactInputSingle(
            IUniswapV3Router.ExactInputSingleParams({
                tokenIn: tokenIn,
                tokenOut: tokenOut,
                fee: originalFee,
                recipient: address(this),
                deadline: block.timestamp,
                amountIn: amountIn,
                amountOutMinimum: 0,
                sqrtPriceLimitX96: 0
            })
        );
    }
}
