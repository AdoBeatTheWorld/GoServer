package proto

type LoginResult struct {
	Uid        string `json:"uid"`
	Score      int64  `json:"score"`
	Account    string `json:"account"`
	GameServer string `json:"game_server"`
}
