package level_up

type ConfigItem struct {
	ID          int32 `json:"id"`
	Exp         int64 `json:"exp"`
	AddExp      int64 `json:"addexp"`
	AddGold     int32 `json:"addgold"`
	CostYuanBao int32 `json:"costyuanbao"`
}
