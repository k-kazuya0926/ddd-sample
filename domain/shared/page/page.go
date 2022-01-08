package page

type Page struct {
	// 取得したエンティティ
	Items []interface{}

	// ページング情報
	Paging Paging
}

type Paging struct {
	// 指定した条件に該当する全件数
	TotalCount uint64

	// 1ページあたりの件数
	PageSize uint64

	// 取得結果のページ番号
	PageNumber uint64
}

type PagingCondition struct {
	PageSize   uint64
	PageNumber uint64
}
