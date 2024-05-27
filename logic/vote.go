package logic

import "bluebell/models"

// 使用簡易版的投票分數
/* 投票的幾種情況
direction=1 時，有兩種情況
	1. 之前沒有投過票，現在投贊成票
	2. 之前投反對票，現在改投贊成票
direction=0 時，有兩種情況
	1. 之前投贊成票，現在要取消投票
	2. 之前投反對票，現在要取消投票
direction=-1 時，有兩種情況
	1. 之前沒有投過票，現在投反對票
	2. 之前投贊成票，現在改投反對票
投票的限制:
每個帖子自發佈之日起一星期內允許用戶投票，超過一星期就不允許再投票
    1. 到期之後將 Redis 中的投票數據存入 mysql 表中
	2. 到期之後刪除 Redis 中的 KeyPostVotedZsetPrefix
*/
// 每個投票的分數固定為 432 分 86400/200 -> 200 張贊成票可以把帖子推到一天的熱門列表
func VoteForPost(userID int64, p *models.ParamVoteData) {
}
