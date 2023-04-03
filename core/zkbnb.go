package core

import (
	"math/big"

	"github.com/bnb-chain/zkbnb-eth-rpc/rpc"
	"github.com/ethereum/go-ethereum/common"
)

type ZkBNBClient struct {
	Instance    *ZkBNB
	Provider    *rpc.ProviderClient
	Constructor TransactOptsConstructor
}

func NewZkBNBClient(provider *rpc.ProviderClient, address string, constructor TransactOptsConstructor) (*ZkBNBClient, error) {
	instance, err := NewZkBNB(common.HexToAddress(address), *provider)
	if err != nil {
		return nil, err
	}
	client := &ZkBNBClient{
		Instance:    instance,
		Provider:    provider,
		Constructor: constructor,
	}
	return client, nil
}

/*
	CommitBlocks: commit blocks
*/
func (c *ZkBNBClient) CommitBlocks(lastBlock StorageStoredBlockInfo, commitBlocksInfo []ZkBNBCommitBlockInfo,
	gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := c.Constructor.ConstructTransactOpts(c.Provider, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := c.Instance.CommitBlocks(transactOpts, lastBlock, commitBlocksInfo)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	CommitBlocks: commit blocks
*/
func (c *ZkBNBClient) CommitBlocksWithNonce(lastBlock StorageStoredBlockInfo, commitBlocksInfo []ZkBNBCommitBlockInfo,
	gasPrice *big.Int, gasLimit uint64, nonce uint64) (txHash string, err error) {
	transactOpts, err := c.Constructor.ConstructTransactOptsWithNonce(gasPrice, gasLimit, nonce)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := c.Instance.CommitBlocks(transactOpts, lastBlock, commitBlocksInfo)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	Estimate Gas for commit blocks operation
*/
func (c *ZkBNBClient) EstimateCommitGasWithNonce(lastBlock StorageStoredBlockInfo, commitBlocksInfo []ZkBNBCommitBlockInfo,
	gasPrice *big.Int, gasLimit uint64, nonce uint64,
) (gas uint64, err error) {
	transactOpts, err := c.Constructor.ConstructTransactOptsWithNonce(gasPrice, gasLimit, nonce)
	if err != nil {
		return 0, err
	}
	// Only for gas estimation and set NoSend = true
	transactOpts.NoSend = true
	// call initialize
	tx, err := c.Instance.CommitBlocks(transactOpts, lastBlock, commitBlocksInfo)
	if err != nil {
		return 0, err
	}
	return tx.Gas(), nil
}

/*
	VerifyAndExecuteBlocks: verify and execute blocks
*/
func (c *ZkBNBClient) VerifyAndExecuteBlocks(verifyAndExecuteBlocksInfo []ZkBNBVerifyAndExecuteBlockInfo, proofs []*big.Int,
	gasPrice *big.Int, gasLimit uint64,
) (txHash string, err error) {
	transactOpts, err := c.Constructor.ConstructTransactOpts(c.Provider, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := c.Instance.VerifyAndExecuteBlocks(transactOpts, verifyAndExecuteBlocksInfo, proofs)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	VerifyAndExecuteBlocks: verify and execute blocks
*/
func (c *ZkBNBClient) VerifyAndExecuteBlocksWithNonce(verifyAndExecuteBlocksInfo []ZkBNBVerifyAndExecuteBlockInfo, proofs []*big.Int,
	gasPrice *big.Int, gasLimit uint64, nonce uint64) (txHash string, err error) {
	transactOpts, err := c.Constructor.ConstructTransactOptsWithNonce(gasPrice, gasLimit, nonce)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := c.Instance.VerifyAndExecuteBlocks(transactOpts, verifyAndExecuteBlocksInfo, proofs)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	Estimate Gas for verifying and executing blocks with kms signature facility
*/
func (c *ZkBNBClient) EstimateVerifyAndExecuteWithNonce(verifyAndExecuteBlocksInfo []ZkBNBVerifyAndExecuteBlockInfo, proofs []*big.Int,
	gasPrice *big.Int, gasLimit uint64, nonce uint64) (gas uint64, err error) {
	transactOpts, err := c.Constructor.ConstructTransactOptsWithNonce(gasPrice, gasLimit, nonce)
	if err != nil {
		return 0, err
	}
	// Only for gas estimation and set NoSend = true
	transactOpts.NoSend = true
	// call initialize
	tx, err := c.Instance.VerifyAndExecuteBlocks(transactOpts, verifyAndExecuteBlocksInfo, proofs)
	if err != nil {
		return 0, err
	}
	return tx.Gas(), nil
}

/*
	RevertBlocks: revert blocks
*/
func (c *ZkBNBClient) RevertBlocks(revertBlocks []StorageStoredBlockInfo, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := c.Constructor.ConstructTransactOpts(c.Provider, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	tx, err := c.Instance.RevertBlocks(transactOpts, revertBlocks)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	PerformDesert: perform desert
*/
func (c *ZkBNBClient) PerformDesert(storedBlockInfo StorageStoredBlockInfo, nftRoot *big.Int, assetExitData DesertVerifierAssetExitData,
	accountExitData DesertVerifierAccountExitData, assetMerkleProof [16]*big.Int, accountMerkleProof [32]*big.Int,
	gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := c.Constructor.ConstructTransactOpts(c.Provider, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := c.Instance.PerformDesert(transactOpts, storedBlockInfo, nftRoot, assetExitData, accountExitData, assetMerkleProof, accountMerkleProof)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	PerformDesertNft: perform desert nft
*/
func (c *ZkBNBClient) PerformDesertNft(storedBlockInfo StorageStoredBlockInfo, assetRoot *big.Int, accountExitData DesertVerifierAccountExitData,
	exitNfts []DesertVerifierNftExitData, accountMerkleProof [32]*big.Int, nftMerkleProofs [][40]*big.Int,
	gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := c.Constructor.ConstructTransactOpts(c.Provider, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := c.Instance.PerformDesertNft(transactOpts, storedBlockInfo, assetRoot, accountExitData, exitNfts, accountMerkleProof, nftMerkleProofs)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	WithdrawPendingBalance: withdraw pending balance
*/
func (c *ZkBNBClient) WithdrawPendingBalance(owner common.Address, token common.Address, amount *big.Int,
	gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := c.Constructor.ConstructTransactOpts(c.Provider, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := c.Instance.WithdrawPendingBalance(transactOpts, owner, token, amount)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	WithdrawPendingNFTBalance: withdraw pending nft balance
*/
func (c *ZkBNBClient) WithdrawPendingNFTBalance(nftIndex *big.Int, gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := c.Constructor.ConstructTransactOpts(c.Provider, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := c.Instance.WithdrawPendingNFTBalance(transactOpts, nftIndex)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	CancelOutstandingDepositsForExodusMode: cancel outstanding deposit
*/
func (c *ZkBNBClient) CancelOutstandingDepositsForExodusMode(priorityRequestId uint64, depositsPubData [][]byte,
	gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := c.Constructor.ConstructTransactOpts(c.Provider, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := c.Instance.CancelOutstandingDepositsForDesertMode(transactOpts, priorityRequestId, depositsPubData)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}

/*
	ActivateDesertMode: activate desert mode
*/
func (c *ZkBNBClient) ActivateDesertMode(gasPrice *big.Int, gasLimit uint64) (txHash string, err error) {
	transactOpts, err := c.Constructor.ConstructTransactOpts(c.Provider, gasPrice, gasLimit)
	if err != nil {
		return "", err
	}
	// call initialize
	tx, err := c.Instance.ActivateDesertMode(transactOpts)
	if err != nil {
		return "", err
	}
	return tx.Hash().String(), nil
}
