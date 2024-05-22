package redis

const (
	KeyPrefix              = "bluebell:"
	KeyPostTimeZset        = "post:time"   // zset;帖子以發佈時間為分數
	KeyPostScoreZset       = "post:score"  // zset;帖子以投票分數為分數
	KeyPostVotedZsetPrefix = "post:voted:" // zset;記錄用戶及投票類型;參數是 post id
)
