package Query

import (
	"fmt"
	"strings"
	"time"
)

func (this QueryExpression) ToODataQuery(paginginfo string) string {
	res := ""

	selectstr := processAttributesOData(this.ColumnSet)
	if selectstr != "" {
		res += "$select=" + selectstr
	}

	ord := processOrdersOData(this.Orders)
	if ord != "" {
		if res != "" {
			res += "&"
		}
		res += "$orderby=" + ord
	}

	return res
}

func processOrdersOData(orders []OrderExpression) string {
	if orders == nil {
		return ""
	}
	if len(orders) == 0 {
		return ""
	}
	res := ""

	for _, oe := range orders {
		if res != "" {
			res += ","
		}
		res += oe.AttributeName
		strOrd := "asc"
		if oe.OrderType == OrderType_Descending {
			strOrd = "desc"
		}
		res += " " + strOrd
	}
	return res
}

func processAttributesOData(cols ColumnSet) string {
	if cols.AllColumns {
		return ""
	}

	return strings.Join(cols.Colmuns, ",")
}

func processFilterOData(filter FilterExpression) string {
	res := ""
	op := " and "
	if filter.FilterOperator == LogicalOperator_Or {
		op = " or "
	}

	ncond := 0

	for _, condexpr := range filter.Conditions {
		if ncond > 0 {
			res += op
		}

		res += getOperatorStringOData(condexpr.AttributeName, condexpr.Operator, condexpr.Values)

		for _, chldFilter := range filter.Filters {
			if ncond > 0 {
				res += op
			}
			res += "(" + processFilterOData(chldFilter) + ")"
		}
		ncond++
	}

	if res != "" {
		return "(" + res + ")"
	}

	return res
}

func stringifyIEnumerable(values []interface{}, quotestr string) string {
	var lst []string

	for _, val := range values {
		lst = append(lst, stringifySingleVal(val, quotestr))
	}

	return strings.Join(lst, ",")
}

func stringifySingleVal(value interface{}, quoteStr string) string {

	if value == nil {
		return "null"
	}

	if dt, okDate := value.(time.Time); okDate {
		return "'" + dt.Format(time.RFC3339) + "'"
	}

	if fl, okFloat := value.(float64); okFloat {
		return quoteStr + fmt.Sprint(fl) + quoteStr
	}

	if integ, okInt := value.(int64); okInt {
		return quoteStr + fmt.Sprintf("%v", integ) + quoteStr
	}

	if str, okStr := value.(string); okStr {
		return "'" + str + "'"
	}

	return "unknown_type"
}

func getOperatorStringOData(fieldName string, operatorType int, values []interface{}) string {
	dynprefix := "Microsoft.Dynamics.CRM."

	switch operatorType {
	case ConditionOperator_Equal:
		if len(values) == 1 {
			return fieldName + " eq " + stringifySingleVal(values[0], "")
		} else {
			return getOperatorStringOData(fieldName, ConditionOperator_In, values)
		}
		break
	case ConditionOperator_NotEqual:
		if len(values) == 1 {
			return fieldName + " ne " + stringifySingleVal(values[0], "")
		} else {
			return getOperatorStringOData(fieldName, ConditionOperator_NotIn, values)
		}
	case ConditionOperator_GreaterThan:
		return fieldName + " gt " + stringifySingleVal(values[0], "")
	case ConditionOperator_GreaterEqual:
		return fieldName + " ge " + stringifySingleVal(values[0], "")
	case ConditionOperator_LessEqual:
		return fieldName + " le " + stringifySingleVal(values[0], "")
	case ConditionOperator_LessThan:
		return fieldName + " lt " + stringifySingleVal(values[0], "")

	case ConditionOperator_Like:
		return "contains(" + fieldName + "," + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_NotLike:
		return "not contains(" + fieldName + "," + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_BeginsWith:
		return "startswith(" + fieldName + "," + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_DoesNotBeginWith:
		return "not startswith(" + fieldName + "," + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_EndsWith:
		return "endswith(" + fieldName + "," + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_DoesNotEndWith:
		return "not endswith(" + fieldName + "," + stringifySingleVal(values[0], "") + ")"

	case ConditionOperator_In:
		if len(values) == 1 {
			return getOperatorStringOData(fieldName, ConditionOperator_Equal, values)
		} else {
			return dynprefix + "In" + "(PropertyName=" + fieldName + ",PropertyValues=[" + stringifyIEnumerable(values, "") + "])"
		}
	case ConditionOperator_NotIn:
		if len(values) == 1 {
			return getOperatorStringOData(fieldName, ConditionOperator_NotEqual, values)
		} else {
			return "not " + getOperatorStringOData(fieldName, ConditionOperator_In, values)
		}
	case ConditionOperator_Between:
		return dynprefix + "Between" + "(PropertyName=" + fieldName + ",PropertyValues=[" + stringifyIEnumerable(values, "") + "])"
	case ConditionOperator_NotBetween:
		return "not " + getOperatorStringOData(fieldName, ConditionOperator_Between, values)
	case ConditionOperator_Null:
		return fieldName + " null"
	case ConditionOperator_NotNull:
		return fieldName + " ne null"
	case ConditionOperator_Yesterday:
		return dynprefix + "Yesterday" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_Today:
		return dynprefix + "Today" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_Tomorrow:
		return dynprefix + "Tomorrow" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_Last7Days:
		return dynprefix + "Last7Days" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_Next7Days:
		return dynprefix + "Next7Days" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_LastWeek:
		return dynprefix + "LastWeek" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_NextWeek:
		return dynprefix + "NextWeek" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_ThisWeek:
		return dynprefix + "ThisWeek" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_LastMonth:
		return dynprefix + "LastMonth" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_NextMonth:
		return dynprefix + "NextMonth" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_ThisMonth:
		return dynprefix + "ThisMonth" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_LastYear:
		return dynprefix + "LastYear" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_NextYear:
		return dynprefix + "NextYear" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_ThisYear:
		return dynprefix + "ThisYear" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_EqualUserId:
		return dynprefix + "EqualUserId" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_EqualUserTeams:
		return dynprefix + "EqualUserTeams" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_EqualBusinessId:
		return dynprefix + "EqualBusinessId" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_EqualUserLanguage:
		return dynprefix + "EqualUserLanguage" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_ThisFiscalYear:
		return dynprefix + "ThisFiscalYear" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_ThisFiscalPeriod:
		return dynprefix + "ThisFiscalPeriod" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_NextFiscalYear:
		return dynprefix + "NextFiscalYear" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_NextFiscalPeriod:
		return dynprefix + "NextFiscalPeriod" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_LastFiscalYear:
		return dynprefix + "LastFiscalYear" + "(PropertyName='" + fieldName + "')"
	case ConditionOperator_LastFiscalPeriod:
		return dynprefix + "LastFiscalPeriod" + "(PropertyName='" + fieldName + "')"

	case ConditionOperator_On:
		return dynprefix + "On" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_OnOrBefore:
		return dynprefix + "OnOrBefore" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_OnOrAfter:
		return dynprefix + "OnOrAfter" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_LastXHours:
		return dynprefix + "LastXHours" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_NextXHours:
		return dynprefix + "NextXHours" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_LastXDays:
		return dynprefix + "LastXDays" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_NextXDays:
		return dynprefix + "NextXDays" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_LastXWeeks:
		return dynprefix + "LastXWeeks" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_NextXWeeks:
		return dynprefix + "NextXWeeks" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_LastXMonths:
		return dynprefix + "LastXMonths" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_NextXMonths:
		return dynprefix + "NextXMonths" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_OlderThanXMonths:
		return dynprefix + "OlderThanXMonths" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_LastXYears:
		return dynprefix + "LastXYears" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_NextXYears:
		return dynprefix + "NextXYears" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_LastXFiscalYears:
		return dynprefix + "LastXFiscalYears" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_LastXFiscalPeriods:
		return dynprefix + "LastXFiscalPeriods" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_NextXFiscalYears:
		return dynprefix + "NextXFiscalYears" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_NextXFiscalPeriods:
		return dynprefix + "NextXFiscalPeriods" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_InFiscalYear:
		return dynprefix + "InFiscalYear" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"
	case ConditionOperator_InFiscalPeriod:
		return dynprefix + "InFiscalPeriod" + "(PropertyName='" + fieldName + "',PropertyValue=" + stringifySingleVal(values[0], "") + ")"

	case ConditionOperator_NotEqualUserId:
		return "not " + getOperatorStringOData(fieldName, ConditionOperator_EqualUserId, values)

	case ConditionOperator_NotEqualBusinessId:
		return "not " + getOperatorStringOData(fieldName, ConditionOperator_EqualBusinessId, values)

	case ConditionOperator_InFiscalPeriodAndYear:
		return dynprefix + "InFiscalPeriodAndYear" + "(PropertyName='" + fieldName + "',PropertyValue1=" + stringifySingleVal(values[0], "") + ",PropertyValue2=" + stringifySingleVal(values[1], "") + ")"
	case ConditionOperator_InOrBeforeFiscalPeriodAndYear:
		return dynprefix + "InOrBeforeFiscalPeriodAndYear" + "(PropertyName='" + fieldName + "',PropertyValue1=" + stringifySingleVal(values[0], "") + ",PropertyValue2=" + stringifySingleVal(values[1], "") + ")"
	case ConditionOperator_InOrAfterFiscalPeriodAndYear:
		return dynprefix + "InOrAfterFiscalPeriodAndYear" + "(PropertyName='" + fieldName + "',PropertyValue1=" + stringifySingleVal(values[0], "") + ",PropertyValue2=" + stringifySingleVal(values[1], "") + ")"

	}

	panic("Not supported operator. Cannot create query string")
}
