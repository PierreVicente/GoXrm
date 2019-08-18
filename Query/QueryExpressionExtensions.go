package Query

import (
	"fmt"
	"github.com/beevik/etree"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

func (qe *QueryExpression) ToFetchXml() string {

	xdoc := qe.getFetchXmlDocument()
	str, err := xdoc.WriteToString()
	if err == nil {
		return str
	}
	return ""

}

func (this *QueryExpression) getFetchXmlDocument() etree.Document {
	fetchDoc := etree.NewDocument()
	xmlFetch := fetchDoc.CreateElement("fetch")
	xmlFetch.CreateAttr("version", "1.5")
	xmlFetch.CreateAttr("output-format", "xml-platform")
	xmlFetch.CreateAttr("mapping", "logical")
	xmlFetch.CreateAttr("no-lock", strconv.FormatBool(this.NoLock))
	if this.PageInfo.PagingCookie != "" {
		xmlFetch.CreateAttr("paging-cookie", this.PageInfo.PagingCookie)
	}

	if !this.AggregateCount && !this.AggregateAvg && !this.AggregateSum {
		xmlFetch.CreateAttr("page", strconv.FormatInt(this.PageInfo.PageNumber, 10))
		xmlFetch.CreateAttr("count", strconv.FormatInt(this.PageInfo.Count, 10))

		if this.PageInfo.PagingCookie != "" {
			xmlFetch.CreateAttr("paging-cookie", this.PageInfo.PagingCookie)
		}
		if this.PageInfo.ReturnTotalRecordCount {
			xmlFetch.CreateAttr("returntotalrecordcount", "true")
		}
	} else {
		xmlFetch.CreateAttr("aggregate", "true")
	}

	xmlEntity := xmlFetch.CreateElement("entity")
	xmlEntity.CreateAttr("name", strings.ToLower(this.EntityName))
	//processAttributes(this.ColumnSet, xmlEntity, this.AggregateCount, this.AggregateAvg, this.AggregateAvg)

	attributeNodes := processAttributes(this.ColumnSet, this.AggregateCount, this.AggregateAvg, this.AggregateAvg)
	for _, attn := range attributeNodes {
		xmlEntity.AddChild(attn.Copy())
	}

	//filterNode := xmlEntity.CreateElement("filter")
	//_t := "and"
	//switch this.Criteria.FilterOperator {
	//case LogicalOperator_Or:
	//	_t = "or"
	//	break
	//}
	//filterNode.CreateAttr("type", _t)

	if len(this.Criteria.Conditions) > 0 || len(this.Criteria.Filters) > 0 {
		filterNode := processFilter(this.Criteria)
		xmlEntity.AddChild(filterNode.Copy())
	}

	orderNodes := processOrders(this.Orders)
	for _, ordn := range orderNodes {
		xmlEntity.AddChild(ordn.Copy())
	}

	for _, lke := range this.LinkEntities {
		linkEntityNode := processLinkEntity(lke, this.AggregateCount, this.AggregateAvg, this.AggregateAvg)
		xmlEntity.AddChild(linkEntityNode.Copy())

		//linkEntityNode := xmlEntity.CreateElement("link-entity")
		//processLinkEntity(lke, linkEntityNode, this.AggregateCount, this.AggregateAvg, this.AggregateAvg)

	}

	return *fetchDoc
}

func processOrders(orders []OrderExpression) []etree.Element {
	var res []etree.Element

	if orders == nil {
		return res
	}
	if len(orders) == 0 {
		return res
	}

	for _, oe := range orders {
		orderNode := etree.Element{Tag: "order"}
		orderNode.CreateAttr("attribute", oe.AttributeName)
		if oe.OrderType == OrderType_Descending {
			orderNode.CreateAttr("descending", "true")
		}
		res = append(res, orderNode)
	}
	return res
}

func processAttributes(cols ColumnSet, aggregateCount bool, aggregateAvg bool, aggregateSum bool) []etree.Element {
	var res []etree.Element

	if cols.AllColumns {
		res = append(res, etree.Element{Tag: "all-attributes"})
		return res
	}

	for _, selfld := range cols.Colmuns {

		if !aggregateCount && !aggregateAvg && !aggregateSum {
			attrNode := etree.Element{Tag: "attribute"}
			attrNode.CreateAttr("name", selfld)
			res = append(res, attrNode)
		}
		if aggregateCount {
			attrNode := etree.Element{Tag: "attribute"}
			attrNode.CreateAttr("name", selfld)
			attrNode.CreateAttr("aggregate", "count")
			attrNode.CreateAttr("alias", selfld+"_count")
			res = append(res, attrNode)
		}
		if aggregateSum {
			attrNode := etree.Element{Tag: "attribute"}
			attrNode.CreateAttr("name", selfld)
			attrNode.CreateAttr("aggregate", "sum")
			attrNode.CreateAttr("alias", selfld+"_sum")
			res = append(res, attrNode)
		}
		if aggregateAvg {
			attrNode := etree.Element{Tag: "attribute"}
			attrNode.CreateAttr("name", selfld)
			attrNode.CreateAttr("aggregate", "avg")
			attrNode.CreateAttr("alias", selfld+"_avg")
			res = append(res, attrNode)
		}
	}
	return res
}

func processLinkEntity(lkE LinkEntity, aggregateCount bool, aggregateAvg bool, aggregateSum bool) etree.Element {

	lkNode := etree.Element{Tag: "link-entity"}

	lkNode.CreateAttr("name", strings.ToLower(lkE.LinkToEntityName))
	lkNode.CreateAttr("from", strings.ToLower(lkE.LinkToAttributeName))
	lkNode.CreateAttr("to", strings.ToLower(lkE.LinkFromAttributeName))
	alias := ""
	if lkE.EntityAlias == "" {
		g, _ := uuid.NewUUID()
		alias = lkE.LinkToEntityName + strings.ReplaceAll(g.String(), "-", "")
	} else {
		alias = lkE.EntityAlias
	}
	lkNode.CreateAttr("alias", alias)

	attributeNodes := processAttributes(lkE.Columns, aggregateCount, aggregateAvg, aggregateSum)
	for _, attn := range attributeNodes {
		lkNode.AddChild(attn.Copy())
	}

	if len(lkE.LinkCriteria.Conditions) > 0 || len(lkE.LinkCriteria.Filters) > 0 {
		filterNode := processFilter(lkE.LinkCriteria)
		lkNode.AddChild(filterNode.Copy())
	}

	orderNodes := processOrders(lkE.Orders)
	for _, ordn := range orderNodes {
		lkNode.AddChild(ordn.Copy())
	}

	for _, lk2 := range lkE.LinkEntities {
		lkEnode2 := processLinkEntity(lk2, aggregateCount, aggregateAvg, aggregateSum)
		lkNode.AddChild(lkEnode2.Copy())
	}
	return lkNode
}

func processFilter(filter FilterExpression) etree.Element {
	filterNode := etree.Element{Tag: "filter"}

	t := "and"
	if filter.FilterOperator == LogicalOperator_Or {
		t = "or"
	}
	filterNode.CreateAttr("type", t)

	for _, condexpr := range filter.Conditions {

		conditionxml := filterNode.CreateElement("condition")
		conditionxml.CreateAttr("attribute", condexpr.AttributeName)
		strOperator := getOperatorString(condexpr.Operator)
		conditionxml.CreateAttr("operator", strOperator)

		if len(condexpr.Values) > 1 {
			for _, o := range condexpr.Values {
				valueNode := conditionxml.CreateElement("value")
				valueNode.CreateText(fmt.Sprint(o))
			}
		} else if len(condexpr.Values) == 1 {
			if condexpr.Values[0] == "" {
				strOperator = getOperatorString(ConditionOperator_Null)
			} else {
				strvalue := ""
				switch condexpr.Values[0].(type) {
				case time.Time:
					v := condexpr.Values[0].(time.Time)
					//strvalue = v.Format("yyyy-MM-dd HH:mm:ss")
					strvalue = v.Format("2006-01-02 15:04:05")
				default:
					strvalue = fmt.Sprint(condexpr.Values[0])
				}

				conditionxml.CreateAttr("value", strvalue)
			}
		}

		for _, chldFilter := range filter.Filters {

			if len(chldFilter.Conditions) > 0 || len(chldFilter.Filters) > 0 {
				newFilter2 := processFilter(chldFilter)
				filterNode.AddChild(newFilter2.Copy())
			}
		}
	}
	return filterNode
}

func getOperatorString(operatorType int) string {

	switch operatorType {
	case ConditionOperator_Equal:
		return "eq"
	case ConditionOperator_NotEqual:
		return "ne"
	case ConditionOperator_GreaterThan:
		return "gt"
	case ConditionOperator_GreaterEqual:
		return "ge"
	case ConditionOperator_LessEqual:
		return "le"
	case ConditionOperator_LessThan:
		return "lt"
	case ConditionOperator_Like:
		return "like"
	case ConditionOperator_NotLike:
		return "not-like"
	case ConditionOperator_In:
		return "in"
	case ConditionOperator_NotIn:
		return "not-in"
	case ConditionOperator_Between:
		return "between"
	case ConditionOperator_NotBetween:
		return "not-between"
	case ConditionOperator_Null:
		return "null"
	case ConditionOperator_NotNull:
		return "not-null"
	case ConditionOperator_Yesterday:
		return "yesterday"
	case ConditionOperator_Today:
		return "today"
	case ConditionOperator_Tomorrow:
		return "tomorrow"
	case ConditionOperator_Last7Days:
		return "last-seven-days"
	case ConditionOperator_Next7Days:
		return "next-seven-days"
	case ConditionOperator_LastWeek:
		return "last-week"
	case ConditionOperator_NextWeek:
		return "next-week"
	case ConditionOperator_ThisWeek:
		return "this-week"
	case ConditionOperator_LastMonth:
		return "last-month"
	case ConditionOperator_NextMonth:
		return "next-month"
	case ConditionOperator_ThisMonth:
		return "this-month"
	case ConditionOperator_On:
		return "on"
	case ConditionOperator_OnOrBefore:
		return "on-or-before"
	case ConditionOperator_OnOrAfter:
		return "on-or-after"
	case ConditionOperator_LastYear:
		return "last-year"
	case ConditionOperator_NextYear:
		return "next-year"
	case ConditionOperator_ThisYear:
		return "this-year"
	case ConditionOperator_LastXHours:
		return "last-x-hours"
	case ConditionOperator_NextXHours:
		return "next-x-hours"
	case ConditionOperator_LastXDays:
		return "last-x-days"
	case ConditionOperator_NextXDays:
		return "next-x-days"
	case ConditionOperator_LastXWeeks:
		return "last-x-weeks"
	case ConditionOperator_NextXWeeks:
		return "next-x-weeks"
	case ConditionOperator_LastXMonths:
		return "last-x-months"
	case ConditionOperator_NextXMonths:
		return "next-x-months"
	case ConditionOperator_LastXYears:
		return "last-x-years"
	case ConditionOperator_NextXYears:
		return "next-x-years"
	case ConditionOperator_EqualUserId:
		return "eq-userid"
	case ConditionOperator_NotEqualUserId:
		return "ne-userid"
	case ConditionOperator_EqualUserTeams:
		return "eq-userteams"
	case ConditionOperator_EqualUserOrUserTeams:
		return "eq-useroruserteamsteams"
	case ConditionOperator_Under:
		return "under"
	case ConditionOperator_UnderOrEqual:
		return "eq-or-under"
	case ConditionOperator_NotUnder:
		return "not-under"
	case ConditionOperator_Above:
		return "above"
	case ConditionOperator_AboveOrEqual:
		return "eq-or-above"
	case ConditionOperator_EqualUserOrUserHierarchy:
		return "eq-useroruserhierarchy"
	case ConditionOperator_EqualUserOrUserHierarchyAndTeams:
		return "eq-useroruserhierarchyandteams"
	case ConditionOperator_EqualBusinessId:
		return "eq-businessid"
	case ConditionOperator_NotEqualBusinessId:
		return "ne-businessid"
	case ConditionOperator_EqualUserLanguage:
		return "eq-userlanguage"
	case ConditionOperator_ThisFiscalYear:
		return "this-fiscal-year"
	case ConditionOperator_ThisFiscalPeriod:
		return "this-fiscal-period"
	case ConditionOperator_NextFiscalYear:
		return "next-fiscal-year"
	case ConditionOperator_NextFiscalPeriod:
		return "next-fiscal-period"
	case ConditionOperator_LastFiscalYear:
		return "last-fiscal-year"
	case ConditionOperator_LastFiscalPeriod:
		return "last-fiscal-period"
	case ConditionOperator_LastXFiscalYears:
		return "last-x-fiscal-years"
	case ConditionOperator_LastXFiscalPeriods:
		return "last-x-fiscal-periods"
	case ConditionOperator_NextXFiscalYears:
		return "next-x-fiscal-years"
	case ConditionOperator_NextXFiscalPeriods:
		return "next-x-fiscal-periods"
	case ConditionOperator_InFiscalYear:
		return "in-fiscal-year"
	case ConditionOperator_InFiscalPeriod:
		return "in-fiscal-period"
	case ConditionOperator_InFiscalPeriodAndYear:
		return "in-fiscal-period-and-year"
	case ConditionOperator_InOrBeforeFiscalPeriodAndYear:
		return "in-or-before-fiscal-period-and-year"
	case ConditionOperator_InOrAfterFiscalPeriodAndYear:
		return "in-or-after-fiscal-period-and-year"
	case ConditionOperator_BeginsWith:
		return "begins-with"
	case ConditionOperator_DoesNotBeginWith:
		return "not-begin-with"
	case ConditionOperator_EndsWith:
		return "ends-with"
	case ConditionOperator_DoesNotEndWith:
		return "not-end-with"
	case ConditionOperator_OlderThanXMonths:
		return "olderthan-x-months"
	case ConditionOperator_OlderThanXYears:
		return "olderthan-x-years"
	case ConditionOperator_OlderThanXWeeks:
		return "olderthan-x-weeks"
	case ConditionOperator_OlderThanXDays:
		return "olderthan-x-days"
	case ConditionOperator_OlderThanXHours:
		return "olderthan-x-hours"
	case ConditionOperator_OlderThanXMinutes:
		return "olderthan-x-minutes"
	case ConditionOperator_ContainValues:
		return "contain-values"
	case ConditionOperator_DoesNotContainValues:
		return "not-contain-values"
	}
	return "Not supported operator by this library"

}
