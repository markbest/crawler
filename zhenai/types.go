package zhenai

type Profile struct {
	Id            int    `json:"id"`
	Url           string `json:"url"`
	Name          string `json:"name"`
	Age           int    `json:"age"`
	Height        int    `json:"height"`
	Income        string `json:"income"`
	Marriage      string `json:"marriage"`
	Education     string `json:"education"`
	WorkCity      string `json:"work_city"`
	Profession    string `json:"profession"`
	Constellation string `json:"constellation"`
	NativePlace   string `json:"native_place"`
}
