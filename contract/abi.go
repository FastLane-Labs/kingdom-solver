package contract

import (
	"github.com/FastLane-Labs/kingdom-solver/contract/demosolver"
	"github.com/FastLane-Labs/kingdom-solver/contract/kingdomdappcontrol"
)

var (
	KingdomDAppControlAbi, _ = kingdomdappcontrol.KingdomDAppControlMetaData.GetAbi()
	DemoSolverAbi, _         = demosolver.DemoSolverMetaData.GetAbi()
)
