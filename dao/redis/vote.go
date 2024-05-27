package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

// 使用簡易版的投票分數
/* 投票的幾種情況
direction=1 時，有兩種情況
	1. 之前沒有投過票，現在投贊成票 --> 更新分數和投票記錄，差值的絕對值是 1 +432
	2. 之前投反對票，現在改投贊成票 --> 更新分數和投票記錄，差值的絕對值是 2 +432*2
direction=0 時，有兩種情況
	1. 之前投反對票，現在要取消投票 --> 更新分數和投票記錄，差值的絕對值是 1 +432
	2. 之前投贊成票，現在要取消投票 --> 更新分數和投票記錄，差值的絕對值是 1 -432
direction=-1 時，有兩種情況
	1. 之前沒有投過票，現在投反對票 --> 更新分數和投票記錄，差值的絕對值是 1 -432
	2. 之前投贊成票，現在改投反對票  --> 更新分數和投票記錄，差值的絕對值是 2 -432*2
投票的限制:
每個帖子自發佈之日起一星期內允許用戶投票，超過一星期就不允許再投票
	1. 到期之後將 Redis 中的投票數據存入 mysql 表中
	2. 到期之後刪除 Redis 中的 KeyPostVotedZsetPrefix
*/

const (
	oneweekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 // 每個投票的分數
)

var (
	ErrVoteTimeExpire = errors.New("投票時間已過")
)

func VoteForPost(userID, postID string, value float64) error {
	// 1. 判斷投票限制
	// 從 Redis 中取出帖子的發佈時間
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZset), postID).Val()
	if float64(time.Now().Unix())-postTime > oneweekInSeconds {
		return ErrVoteTimeExpire
	}
	// 2. 更新帖子的分數
	// 先查詢當前用戶給當前該帖子的投票紀錄
	ov := rdb.ZScore(getRedisKey(KeyPostVotedZsetPrefix+postID), userID).Val()
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)
	_, err := rdb.ZIncrBy(getRedisKey(KeyPostScoreZset), op*diff*scorePerVote, postID).Result()
	if ErrVoteTimeExpire != nil {
		return err
	}
	// 3. 紀錄用戶為該帖子投票的記錄
	if value == 0 {
		_, err = rdb.ZRem(getRedisKey(KeyPostVotedZsetPrefix+postID), postID).Result()
	} else {
		_, err = rdb.ZAdd(getRedisKey(KeyPostVotedZsetPrefix+postID), redis.Z{
			Score:  value, // 贊成票或反對票
			Member: userID,
		}).Result()
	}
	return err
}
