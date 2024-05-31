package job

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shoe-shark/shoe-shark-service/eth"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/abi_repository"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/constants"
	"github.com/shoe-shark/shoe-shark-service/mods/contracts/dao"
	dao2 "github.com/shoe-shark/shoe-shark-service/mods/points/dao"
	"github.com/shoe-shark/shoe-shark-service/repository"
	log "github.com/sirupsen/logrus"
	"math/big"
	"time"
)

type ShoeSharkContractJob struct {
	client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
}

func NewJob(client *ethclient.Client, privateKey *ecdsa.PrivateKey) *ShoeSharkContractJob {
	return &ShoeSharkContractJob{
		client:     client,
		privateKey: privateKey,
	}
}

func (cj *ShoeSharkContractJob) StartJob() {
	// 启动定时任务的 goroutine
	log.Info("[ContractJob] ", "Starting Contract job")
	go cj.scheduleTask()
}

func (cj *ShoeSharkContractJob) scheduleTask() {
	// 计算当前时间到执行时间的间隔
	duration := 5 * time.Minute

	// 创建一个定时器
	timer := time.NewTimer(duration)

	for {
		select {
		case <-timer.C:
			// 执行任务
			log.Info("[ContractJob] ", " Running job")
			cj.startGrantPointsJob()

			// 重置定时器
			timer.Reset(duration)
		}
	}
}

func (cj *ShoeSharkContractJob) startGrantPointsJob() {
	// 获取所需要同步的积分
	rp := repository.GetPGRepository()
	pointsLogs, err := dao2.GetUnSyncedLogs(rp)
	if err != nil {
		log.Error("[同步积分失败][获取积分异常]  ", err)
		return
	}

	if len(pointsLogs) == 0 {
		log.Info("[同步积分]无需要同步积分")
		return
	}

	accountPoints := make(map[string]int64)
	idsMap := make(map[uint]bool)

	for _, log := range pointsLogs {
		accountPoints[log.AccountAddress] += log.Points
		idsMap[log.ID] = true
	}

	// 将 map 中的 keys 转换为 slice
	idsToUpdate := make([]uint, 0, len(idsMap))
	for id := range idsMap {
		idsToUpdate = append(idsToUpdate, id)
	}

	rewardPointContractAddress, err := dao.GetContractAddressByContractName(rp, constants.ShoeSharkRewardPoint)
	if err != nil {
		log.Error("un find reward point contracts entity", "error", err)
	}
	if len(rewardPointContractAddress) == 0 {
		log.Info("[同步积分]未获取到ShoeSharkRewardPoint合约地址")
		return
	}

	nftContractAddress, err := dao.GetContractAddressByContractName(rp, constants.ShoeSharkNft)
	if err != nil {
		log.Error("un find nft contracts entity", "error", err)
	}

	if len(nftContractAddress) == 0 {
		log.Info("[同步积分]未获取到ShoeSharkNft合约地址")
		return
	}

	pointBiz, err := abi_repository.NewShoeSharkRewardPointRepository(cj.client, rewardPointContractAddress, cj.privateKey)
	if err != nil {
		log.Error("init ShoeSharkRewardPointRepository error: ", err)
	}

	nftBiz, err := abi_repository.NewShoeSharkNftRepository(cj.client, nftContractAddress, cj.privateKey)
	if err != nil {
		log.Error("init ShoeSharkNftRepository error: ", err)
	}

	var accounts []common.Address
	var points []*big.Int
	// 遍历 map，将 keys 和 values 分别追加到切片中
	for account, point := range accountPoints {
		accountAddress := common.HexToAddress(account)
		accounts = append(accounts, accountAddress)

		linkAccountPoints := pointBiz.GetPoints(accountAddress)
		points = append(points, big.NewInt(0).Add(linkAccountPoints, big.NewInt(point)))
	}

	err = pointBiz.SetPoints(accounts, points)
	if err != nil {
		log.Error("[同步积分失败] init pointBiz error: ", err)
		return
	}

	// 获取所有账户生成root
	accountsHex, err := dao2.GetAllAccountAddresses(rp)
	if err != nil {
		log.Error("[同步积分失败] GetAllAccountAddresses error: ", err)
		return
	}

	merkleRoot, _, _ := eth.BuildMerkleTree(accountsHex)
	err = nftBiz.SetMerkleRoot(merkleRoot)

	// 打印每个账户的积分总和
	// 更新记录的同步状态
	if len(idsToUpdate) > 0 {
		err = dao2.UpdateIsSyncLink(rp, idsToUpdate, true)
		if err != nil {
			log.Error("[同步积分失败][修改同步状态失败]  ", err)
			return
		}
	}

	//err = dao2.AddPointsBatch(rp, accountsHex, points)
	//if err != nil {
	//	log.Error("[同步积分失败][修改用户积分] ", err)
	//}

	err = dao.InsertAccounts(rp, accountsHex)
	if err != nil {
		log.Error("[同步积分失败][记录白名单账户列表失败] ", err)
	}
}
