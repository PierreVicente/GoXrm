package Query

import(
	"github.com/beevik/etree"
	"strconv"
	"strings"
)



func(qe *QueryExpression) ToFetchXml() string{
	return qe.getFetchXmlString()
}

func (qe *QueryExpression) getFetchXmlString() string {

}

func (this *QueryExpression) getFetchXmlDocument() etree.Document {
	fetchDoc := etree.NewDocument()
	xmlFetch := fetchDoc.CreateElement("fetch")
	fetchDoc.CreateAttr("version", "1.5")
	fetchDoc.CreateAttr("output-format", "xml-platform")
	fetchDoc.CreateAttr("mapping", "logical")
	fetchDoc.CreateAttr("no-lock", strconv.FormatBool(this.NoLock))
	if (this.PageInfo.PagingCookie != "")	{
		xmlFetch.CreateAttr("paging-cookie", this.PageInfo.PagingCookie)
	}

	if (!this.AggregateCount && !this.AggregateAvg && !this.AggregateSum){
		xmlFetch.CreateAttr("page", string(this.PageInfo.PageNumber))
		xmlFetch.CreateAttr("count", string(this.PageInfo.Count))

		if (this.PageInfo.PagingCookie != ""){
			xmlFetch.CreateAttr("paging-cookie", this.PageInfo.PagingCookie)
		}
		if (this.PageInfo.ReturnTotalRecordCount){
			xmlFetch.CreateAttr("returntotalrecordcount", "true");
		}
	} else {
		xmlFetch.CreateAttr("aggregate", "true");
	}

	xmlEntity := xmlFetch.CreateElement("entity");
	xmlEntity.CreateAttr("name", strings.ToLower(this.EntityName))

	ProcessAttributes(this.ColumnSet, xmlEntity);

	filterNode := xmlEntity.CreateElement("filter");

	_t := "and"
	switch this.Criteria.FilterOperator {
	case LogicalOperator_Or:
		_t = "or"
		break
	}

	filterNode.CreateAttr("type", _t)

	ProcessFilter(this.Criteria, filterNode);

	processOrders(this.Orders, xmlEntity);

	for _, lke := range  this.LinkEntities {
		linkEntityNode := xmlEntity.CreateElement("link-entity")
		ProcessLinkEntity(lke, linkEntityNode)
	}

	return *fetchDoc;
}

func processOrders(orders []OrderExpression, entityNode etree.Element) {
	if (orders == nil) { return }
	if (len(orders) == 0) { return }

	for _, oe := range orders{
		orderNode := entityNode.CreateElement("order")
		orderNode.CreateAttr( "attribute", oe.AttributeName);
		if (oe.OrderType == OrderType_Descending)
		{
			orderNode.CreateAttr("descending", "true")
		}
	}
}