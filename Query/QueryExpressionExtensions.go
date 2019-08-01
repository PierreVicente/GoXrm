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
	return qe.getFetchXmlString()
}

func (qe *QueryExpression) getFetchXmlString() string {
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
	fetchDoc.CreateAttr("version", "1.5")
	fetchDoc.CreateAttr("output-format", "xml-platform")
	fetchDoc.CreateAttr("mapping", "logical")
	fetchDoc.CreateAttr("no-lock", strconv.FormatBool(this.NoLock))
	if this.PageInfo.PagingCookie != "" {
		xmlFetch.CreateAttr("paging-cookie", this.PageInfo.PagingCookie)
	}

	if !this.AggregateCount && !this.AggregateAvg && !this.AggregateSum {
		xmlFetch.CreateAttr("page", string(this.PageInfo.PageNumber))
		xmlFetch.CreateAttr("count", string(this.PageInfo.Count))

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

	processAttributes(this.ColumnSet, xmlEntity, this.AggregateCount, this.AggregateAvg, this.AggregateAvg)

	filterNode := xmlEntity.CreateElement("filter")

	_t := "and"
	switch this.Criteria.FilterOperator {
	case LogicalOperator_Or:
		_t = "or"
		break
	}

	filterNode.CreateAttr("type", _t)

	processFilter(this.Criteria, filterNode)

	processOrders(this.Orders, xmlEntity)

	for _, lke := range this.LinkEntities {
		linkEntityNode := xmlEntity.CreateElement("link-entity")
		processLinkEntity(lke, linkEntityNode, this.AggregateCount, this.AggregateAvg, this.AggregateAvg)
	}

	return *fetchDoc
}

func processOrders(orders []OrderExpression, entityNode *etree.Element) {
	if orders == nil {
		return
	}
	if len(orders) == 0 {
		return
	}

	for _, oe := range orders {
		orderNode := entityNode.CreateElement("order")
		orderNode.CreateAttr("attribute", oe.AttributeName)
		if oe.OrderType == OrderType_Descending {
			orderNode.CreateAttr("descending", "true")
		}
	}
}

func processAttributes(cols ColumnSet, pNode *etree.Element, aggregateCount bool, aggregateAvg bool, aggregateSum bool) {
	if cols.AllColumns {
		pNode.CreateElement("all-attributes")
		return
	}

	for _, selfld := range cols.Colmuns {

		if !aggregateCount && !aggregateAvg && !aggregateSum {
			attrNode := pNode.CreateElement("attribute")
			attrNode.CreateAttr("name", selfld)
		}
		if aggregateCount {
			attrNode := pNode.CreateElement("attribute")
			attrNode.CreateAttr("name", selfld)
			attrNode.CreateAttr("aggregate", "count")
			attrNode.CreateAttr("alias", selfld+"_count")
			//pNode.AppendChild(attrNode);
		}
		if aggregateSum {
			attrNode := pNode.CreateElement("attribute")
			attrNode.CreateAttr("name", selfld)
			attrNode.CreateAttr("aggregate", "sum")
			attrNode.CreateAttr("alias", selfld+"_sum")
			//pNode.AppendChild(attrNode);
		}
		if aggregateAvg {
			attrNode := pNode.CreateElement("attribute")
			attrNode.CreateAttr("name", selfld)
			attrNode.CreateAttr("aggregate", "avg")
			attrNode.CreateAttr("alias", selfld+"_avg")
			//pNode.AppendChild(attrNode);
		}
	}
}

func processLinkEntity(lkE LinkEntity, lkNode *etree.Element, aggregateCount bool, aggregateAvg bool, aggregateSum bool) {
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

	//if (!string.IsNullOrEmpty(_alias) && !aliasMap.ContainsKey(_alias))
	//{
	//    aliasMap.Add(_alias, lkE.LinkToEntityName.ToLower());
	//}

	processAttributes(lkE.Columns, lkNode, aggregateCount, aggregateAvg, aggregateSum)

	filterNode := lkNode.CreateElement("filter")
	t := "and"
	if lkE.LinkCriteria.FilterOperator == LogicalOperator_Or {
		t = "or"
	}
	filterNode.CreateAttr("type", t)

	processFilter(lkE.LinkCriteria, filterNode)
	//lkNode.AppendChild(filterNode);

	processOrders(lkE.Orders, lkNode)

	for _, lk2 := range lkE.LinkEntities {
		lkEnode2 := lkNode.CreateElement("link-entity")
		processLinkEntity(lk2, lkEnode2, aggregateCount, aggregateAvg, aggregateSum)
		//lkNode.AppendChild(lkEnode2);
	}
}

func processFilter(filter FilterExpression, filterNode *etree.Element) {
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
			newFilter := filterNode.CreateElement("filter")
			v := "and"
			if chldFilter.FilterOperator == LogicalOperator_Or {
				v = "or"
			}
			newFilter.CreateAttr("type", v)

			processFilter(chldFilter, newFilter)
		}
	}
}

func getOperatorString(operatorType int32) string {

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
	case ConditionOperator_OlderThanXMonths:
		return "olderthan-x-months"
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
	}
	return "Not supported operator by this library"

}
