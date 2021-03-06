package file

type Node struct {
	Name          string  `json:"name"`          // 股票名字
	Open          string  `json:"open"`          // 开盘价
	Prevclose     string  `json:"prevclose"`     // 昨日收盘价
	Price         string  `json:"price"`         // 当前价格
	High          string  `json:"high"`          // 最高价
	Low           string  `json:"low"`           // 最低价
	Buy           string  `json:"buy"`           // 当前最高购入价
	Sell          string  `json:"sell"`          // 当前最低出售价
	TotalVolume   string  `json:"totalVolume"`   // 成交量
	TotalAmount   string  `json:"totalAmount"`   // 成交金额
	Buyvol1       string  `json:"buyvol1"`       // 买1数量
	Buyprice1     string  `json:"buyprice1"`     // 买1总价
	Buyvol2       string  `json:"buyvol2"`       // 买2数量
	Buyprice2     string  `json:"buyprice2"`     // 买2总价
	Buyvol3       string  `json:"buyvol3"`       // 买3数量
	Buyprice3     string  `json:"buyprice3"`     // 买3总价
	Buyvol4       string  `json:"buyvol4"`       // 买4数量
	Buyprice4     string  `json:"buyprice4"`     // 买4总价
	Buyvol5       string  `json:"buyvol5"`       // 买5数量
	Buyprice5     string  `json:"buyprice5"`     // 买5总价
	Sellvol1      string  `json:"sellvol1"`      // 卖1数量
	Sellprice1    string  `json:"sellprice1"`    // 卖1总价
	Sellvol2      string  `json:"sellvol2"`      // 卖2数量
	Sellprice2    string  `json:"sellprice2"`    // 卖2总价
	Sellvol3      string  `json:"sellvol3"`      // 卖3数量
	Sellprice3    string  `json:"sellprice3"`    // 卖3总价
	Sellvol4      string  `json:"sellvol4"`      // 卖4数量
	Sellprice4    string  `json:"sellprice4"`    // 卖4总价
	Sellvol5      string  `json:"sellvol5"`      // 卖5数量
	Sellprice5    string  `json:"sellprice5"`    // 卖5总价
	Date          string  `json:"date"`          // 日期
	Time          string  `json:"time"`          // 时间
	State         string  `json:"state"`         //
	Stop          bool    `json:"stop"`          // 停牌
	Symbol        string  `json:"symbol"`        // 代码
	Change        float64 `json:"change"`        // 变动
	Percent       float64 `json:"percent"`       // 变动比例
	Amplitude     float64 `json:"amplitude"`     // 振幅
	RpOut         float64 `json:"rp_out"`        // 主力流出
	RpIn          float64 `json:"rp_in"`         // 主力流入
	RpNet         float64 `json:"rp_net"`        // 主力净流入
	RpNetRate     float64 `json:"rp_net_rate"`   // 主力净流入占比
	RxlNet        float64 `json:"rxl_net"`       // 特大单净流入
	RxlNetRate    float64 `json:"rxl_net_rate"`  // 特大单占比
	RlNet         float64 `json:"rl_net"`        // 大单净流入
	RlNetRate     float64 `json:"rl_net_rate"`   // 大单占比
	RmNet         float64 `json:"rm_net"`        // 中单净流入
	RmNetRate     float64 `json:"rm_net_rate"`   // 中单占比
	RsNet         float64 `json:"rs_net"`        // 小单净流入
	RsNetRate     float64 `json:"rs_net_rate"`   // 小单占比
	TotalAmountMF string  `json:"totalAmountMF"` //
	Ddjl          string  `json:"ddjl"`          // 大单净量
	Ddjl3         string  `json:"ddjl3"`         // 大单3日净流
	Ddjl5         string  `json:"ddjl5"`         // 大单5日净流
	Ddjl10        string  `json:"ddjl10"`        // 大单10日净流
	Ddjl20        string  `json:"ddjl20"`        // 大单20日净流
	UpdateTime    int64  `json:"update_time"`   // 最新更新时间
}
