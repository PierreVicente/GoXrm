package Query

type PagingInfo struct {
	PageNumber             int64
	Count                  int64
	ReturnTotalRecordCount bool
	PagingCookie           string
}

func New(pageNumber int64, pagingcookie string) PagingInfo {
	var p PagingInfo
	p.PageNumber = pageNumber
	p.PagingCookie = pagingcookie
	p.Count = 5000
	p.ReturnTotalRecordCount = false
	return p
	//return new(PagingInfo{PageNumber: pageNumber, PagingCookie: pagingcookie, Count: 5000, ReturnTotalRecordCount: false})
}
