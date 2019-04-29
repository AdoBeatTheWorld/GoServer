package protocol

type LoginResult struct {
	Uid        string `json:"uid"`
	Score      int64  `json:"score"`
	Account    string `json:"account"`
	GameServer string `json:"game_server"`
}

type HttpLogin struct {
	Session    string `json:"session"`
	GateServer string `json:"gate"`
}
