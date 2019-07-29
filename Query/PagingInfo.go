package Query

type PagingInfo struct {
	PageNumber             int32
	Count                  int32
	ReturnTotalRecordCount bool
	PagingCookie           string
}

func New(pageNumber int32, pagingcookie string) PagingInfo {
	return new(PagingInfo{PageNumber: pageNumber, PagingCookie: pagingcookie, Count: 5000, ReturnTotalRecordCount: false})
}
