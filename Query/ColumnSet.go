package Query

type ColumnSet struct {
	AllColumns bool
	Colmuns    []string
}

func NewColmunSetAll(allcols bool) ColumnSet {
	return ColumnSet{AllColumns: allcols, Colmuns: []string{}}

}

func NewColumnSetCols(cols []string) ColumnSet {
	return ColumnSet{AllColumns: false, Colmuns: cols}
}
