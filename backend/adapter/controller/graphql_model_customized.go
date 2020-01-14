package controller

// 組織
type Organization struct {
	// UUID
	ID string `json:"id"`
	// 名称
	Name string `json:"Name"`
	// 上位組織 ... 別resolver呼び出し
	// 下位組織群 ... 別resolver呼び出し
}

func (Organization) IsNode() {}

// 作品
type Work struct {
	// UUID
	ID string `json:"id"`
	// 作品名
	Name string `json:"name"`
	// 価格（無料は0円）
	Price int `json:"price"`
	// 作成者群（不明な場合もある） ... 別resolver呼び出し
}

func (Work) IsNode() {}

// 作成者
type WorkHolder struct {
	// UUID
	ID string `json:"id"`
	// 姓
	FirstName string `json:"firstName"`
	// 名
	LastName string `json:"lastName"`
	// 姓名（姓と名から動的に生成）
	Name string `json:"name"`
	// ニックネーム
	Nickname *string `json:"nickname"`
	// 所属組織群 ... 別resolver呼び出し
	// 所持作品群 ... 別resolver呼び出し
}

func (WorkHolder) IsNode() {}
