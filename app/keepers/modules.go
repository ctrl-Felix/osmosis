package keepers

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	packetforward "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v7/packetforward"
	transfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v7/modules/core"
	ibcclientclient "github.com/cosmos/ibc-go/v7/modules/core/02-client/client"
	tendermint "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"

	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/capability"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	icq "github.com/cosmos/ibc-apps/modules/async-icq/v7"
	ica "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts"

	"github.com/cosmos/cosmos-sdk/x/consensus"

	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

	_ "github.com/osmosis-labs/osmosis/v20/client/docs/statik"
	clclient "github.com/osmosis-labs/osmosis/v20/x/concentrated-liquidity/client"
	concentratedliquidity "github.com/osmosis-labs/osmosis/v20/x/concentrated-liquidity/clmodule"
	cwpoolclient "github.com/osmosis-labs/osmosis/v20/x/cosmwasmpool/client"
	cosmwasmpoolmodule "github.com/osmosis-labs/osmosis/v20/x/cosmwasmpool/module"
	downtimemodule "github.com/osmosis-labs/osmosis/v20/x/downtime-detector/module"
	"github.com/osmosis-labs/osmosis/v20/x/gamm"
	gammclient "github.com/osmosis-labs/osmosis/v20/x/gamm/client"
	"github.com/osmosis-labs/osmosis/v20/x/ibc-rate-limit/ibcratelimitmodule"
	"github.com/osmosis-labs/osmosis/v20/x/incentives"
	incentivesclient "github.com/osmosis-labs/osmosis/v20/x/incentives/client"
	"github.com/osmosis-labs/osmosis/v20/x/lockup"
	"github.com/osmosis-labs/osmosis/v20/x/mint"
	poolincentives "github.com/osmosis-labs/osmosis/v20/x/pool-incentives"
	poolincentivesclient "github.com/osmosis-labs/osmosis/v20/x/pool-incentives/client"
	poolmanagerclient "github.com/osmosis-labs/osmosis/v20/x/poolmanager/client"
	poolmanager "github.com/osmosis-labs/osmosis/v20/x/poolmanager/module"
	"github.com/osmosis-labs/osmosis/v20/x/protorev"
	superfluid "github.com/osmosis-labs/osmosis/v20/x/superfluid"
	superfluidclient "github.com/osmosis-labs/osmosis/v20/x/superfluid/client"
	"github.com/osmosis-labs/osmosis/v20/x/tokenfactory"
	"github.com/osmosis-labs/osmosis/v20/x/twap/twapmodule"
	"github.com/osmosis-labs/osmosis/v20/x/txfees"
	txfeesclient "github.com/osmosis-labs/osmosis/v20/x/txfees/client"
	valsetprefmodule "github.com/osmosis-labs/osmosis/v20/x/valset-pref/valpref-module"
	"github.com/osmosis-labs/osmosis/x/epochs"
	ibc_hooks "github.com/osmosis-labs/osmosis/x/ibc-hooks"

	"github.com/skip-mev/block-sdk/x/auction"
)

// AppModuleBasics returns ModuleBasics for the module BasicManager.
var AppModuleBasics = []module.AppModuleBasic{
	auth.AppModuleBasic{},
	genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
	bank.AppModuleBasic{},
	capability.AppModuleBasic{},
	staking.AppModuleBasic{},
	mint.AppModuleBasic{},
	downtimemodule.AppModuleBasic{},
	distr.AppModuleBasic{},
	gov.NewAppModuleBasic(
		[]govclient.ProposalHandler{
			paramsclient.ProposalHandler,
			upgradeclient.LegacyProposalHandler,
			upgradeclient.LegacyCancelProposalHandler,
			poolincentivesclient.UpdatePoolIncentivesHandler,
			poolincentivesclient.ReplacePoolIncentivesHandler,
			ibcclientclient.UpdateClientProposalHandler,
			ibcclientclient.UpgradeProposalHandler,
			superfluidclient.SetSuperfluidAssetsProposalHandler,
			superfluidclient.RemoveSuperfluidAssetsProposalHandler,
			superfluidclient.UpdateUnpoolWhitelistProposalHandler,
			gammclient.ReplaceMigrationRecordsProposalHandler,
			gammclient.UpdateMigrationRecordsProposalHandler,
			gammclient.CreateCLPoolAndLinkToCFMMProposalHandler,
			gammclient.SetScalingFactorControllerProposalHandler,
			clclient.CreateConcentratedLiquidityPoolProposalHandler,
			clclient.TickSpacingDecreaseProposalHandler,
			cwpoolclient.UploadCodeIdAndWhitelistProposalHandler,
			cwpoolclient.MigratePoolContractsProposalHandler,
			txfeesclient.SubmitUpdateFeeTokenProposalHandler,
			poolmanagerclient.DenomPairTakerFeeProposalHandler,
			incentivesclient.HandleCreateGroupsProposal,
		},
	),
	params.AppModuleBasic{},
	crisis.AppModuleBasic{},
	slashing.AppModuleBasic{},
	authzmodule.AppModuleBasic{},
	consensus.AppModuleBasic{},
	auction.AppModuleBasic{},
	ibc.AppModuleBasic{},
	upgrade.AppModuleBasic{},
	evidence.AppModuleBasic{},
	transfer.AppModuleBasic{},
	vesting.AppModuleBasic{},
	gamm.AppModuleBasic{},
	poolmanager.AppModuleBasic{},
	twapmodule.AppModuleBasic{},
	concentratedliquidity.AppModuleBasic{},
	protorev.AppModuleBasic{},
	txfees.AppModuleBasic{},
	incentives.AppModuleBasic{},
	lockup.AppModuleBasic{},
	poolincentives.AppModuleBasic{},
	epochs.AppModuleBasic{},
	superfluid.AppModuleBasic{},
	tokenfactory.AppModuleBasic{},
	valsetprefmodule.AppModuleBasic{},
	wasm.AppModuleBasic{},
	icq.AppModuleBasic{},
	ica.AppModuleBasic{},
	ibc_hooks.AppModuleBasic{},
	ibcratelimitmodule.AppModuleBasic{},
	packetforward.AppModuleBasic{},
	cosmwasmpoolmodule.AppModuleBasic{},
	tendermint.AppModuleBasic{},
}
