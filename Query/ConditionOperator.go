package Query

const (
	ConditionOperator_Equal = iota
	ConditionOperator_NotEqual
	//
	// Summary:
	//     The value is greater than the compared value. Value = 2.
	ConditionOperator_GreaterThan
	//
	// Summary:
	//     The value is less than the compared value. Value = 3.
	ConditionOperator_LessThan
	//
	// Summary:
	//     The value is greater than or equal to the compared value. Value = 4.
	ConditionOperator_GreaterEqual
	//
	// Summary:
	//     The value is less than or equal to the compared value. Value = 5.
	ConditionOperator_LessEqual
	//
	// Summary:
	//     The character string is matched to the specified pattern. Value = 6.
	ConditionOperator_Like
	//
	// Summary:
	//     The character string does not match the specified pattern. Value = 7.
	ConditionOperator_NotLike
	//
	// Summary:
	//     TheThe value exists in a list of values. Value = 8.
	ConditionOperator_In
	//
	// Summary:
	//     The given value is not matched to a value in a subquery or a list. Value = 9.
	ConditionOperator_NotIn
	//
	// Summary:
	//     The value is between two values. Value = 10.
	ConditionOperator_Between
	//
	// Summary:
	//     The value is not between two values. Value = 11.
	ConditionOperator_NotBetween
	//
	// Summary:
	//     The value is null. Value = 12.
	ConditionOperator_Null
	//
	// Summary:
	//     The value is not null. Value = 13.
	ConditionOperator_NotNull
	//
	// Summary:
	//     The value equals yesterday’s date. Value = 14.
	ConditionOperator_Yesterday
	//
	// Summary:
	//     The value equals today’s date. Value = 15.
	ConditionOperator_Today
	//
	// Summary:
	//     The value equals tomorrow’s date. Value = 16.
	ConditionOperator_Tomorrow
	//
	// Summary:
	//     The value is within the last seven days including today. Value = 17.
	ConditionOperator_Last7Days
	//
	// Summary:
	//     The value is within the next seven days. Value = 18.
	ConditionOperator_Next7Days
	//
	// Summary:
	//     The value is within the previous week including Sunday through Saturday. Value
	//     = 19.
	ConditionOperator_LastWeek
	//
	// Summary:
	//     The value is within the current week. Value = 20.
	ConditionOperator_ThisWeek
	//
	// Summary:
	//     The value is within the next week. Value = 21.
	ConditionOperator_NextWeek
	//
	// Summary:
	//     The value is within the last month including first day of the last month and
	//     last day of the last month. Value = 22.
	ConditionOperator_LastMonth
	//
	// Summary:
	//     The value is within the current month. Value = 23.
	ConditionOperator_ThisMonth
	//
	// Summary:
	//     The value is within the next month. Value = 24.
	ConditionOperator_NextMonth
	//
	// Summary:
	//     The value is on a specified date. Value = 25.
	ConditionOperator_On
	//
	// Summary:
	//     The value is on or before a specified date. Value = 26.
	ConditionOperator_OnOrBefore
	//
	// Summary:
	//     The value is on or after a specified date. Value = 27.
	ConditionOperator_OnOrAfter
	//
	// Summary:
	//     The value is within the previous year. Value = 28.
	ConditionOperator_LastYear
	//
	// Summary:
	//     The value is within the current year. Value = 29.
	ConditionOperator_ThisYear
	//
	// Summary:
	//     The value is within the next year. Value = 30.
	ConditionOperator_NextYear
	//
	// Summary:
	//     The value is within the last X hours. Value =31.
	ConditionOperator_LastXHours
	//
	// Summary:
	//     The value is within the next X (specified value) hours. Value = 32.
	ConditionOperator_NextXHours
	//
	// Summary:
	//     The value is within last X days. Value = 33.
	ConditionOperator_LastXDays
	//
	// Summary:
	//     The value is within the next X (specified value) days. Value = 34.
	ConditionOperator_NextXDays
	//
	// Summary:
	//     The value is within the last X (specified value) weeks. Value = 35.
	ConditionOperator_LastXWeeks
	//
	// Summary:
	//     The value is within the next X weeks. Value = 36.
	ConditionOperator_NextXWeeks
	//
	// Summary:
	//     The value is within the last X (specified value) months. Value = 37.
	ConditionOperator_LastXMonths
	//
	// Summary:
	//     The value is within the next X (specified value) months. Value = 38.
	ConditionOperator_NextXMonths
	//
	// Summary:
	//     The value is within the last X years. Value = 39.
	ConditionOperator_LastXYears
	//
	// Summary:
	//     The value is within the next X years. Value = 40.
	ConditionOperator_NextXYears
	//
	// Summary:
	//     The value is equal to the specified user ID. Value = 41.
	ConditionOperator_EqualUserId
	//
	// Summary:
	//     The value is not equal to the specified user ID. Value = 42.
	ConditionOperator_NotEqualUserId
	//
	// Summary:
	//     The value is equal to the specified business ID. Value = 43.
	ConditionOperator_EqualBusinessId
	//
	// Summary:
	//     The value is not equal to the specified business ID. Value = 44.
	ConditionOperator_NotEqualBusinessId
	//
	// Summary:
	//     No token name is specified &lt;?Comment AL: Bug fix 5/30/12 2012-05-30T11:03:00Z
	//     Id=&#39;2?&gt;internal&lt;?CommentEnd Id=&#39;2&#39; ?&gt;.
	ConditionOperator_ChildOf
	//
	// Summary:
	//     The value is found in the specified bit-mask value. Value = 46.
	ConditionOperator_vMask
	//
	// Summary:
	//     The value is not found in the specified bit-mask value. Value = 47.
	ConditionOperator_NotMask
	//
	// Summary:
	//     For internal use only. Value = 48.
	ConditionOperator_MasksSelect
	//
	// Summary:
	//     The string contains another string. Value = 49. You must use the Contains operator
	//     for only those attributes that are enabled for full-text indexing. Otherwise,
	//     you will receive a generic SQL error message while retrieving data. In a Microsoft
	//     Dynamics 365 default installation, only the attributes of the KBArticle (article)
	//     entity are enabled for full-text indexing.
	ConditionOperator_Contains
	//
	// Summary:
	//     The string does not contain another string. Value = 50.
	ConditionOperator_DoesNotContain
	//
	// Summary:
	//     The value is equal to the language for the user. Value = 51.
	ConditionOperator_EqualUserLanguage
	//
	// Summary:
	//     For internal use only.
	ConditionOperator_NotOn
	//
	// Summary:
	//     The value is older than the specified number of months. Value = 53.
	ConditionOperator_OlderThanXMonths
	//
	// Summary:
	//     The string occurs at the beginning of another string. Value = 54.
	ConditionOperator_BeginsWith
	//
	// Summary:
	//     The string does not begin with another string. Value = 55.
	ConditionOperator_DoesNotBeginWith
	//
	// Summary:
	//     The string ends with another string. Value = 56.
	ConditionOperator_EndsWith
	//
	// Summary:
	//     The string does not end with another string. Value = 57.
	ConditionOperator_DoesNotEndWith
	//
	// Summary:
	//     The value is within the current fiscal year . Value = 58.
	ConditionOperator_ThisFiscalYear
	//
	// Summary:
	//     The value is within the current fiscal period. Value = 59.
	ConditionOperator_ThisFiscalPeriod
	//
	// Summary:
	//     The value is within the next fiscal year. Value = 60.
	ConditionOperator_NextFiscalYear
	//
	// Summary:
	//     The value is within the next fiscal period. Value = 61.
	ConditionOperator_NextFiscalPeriod
	//
	// Summary:
	//     The value is within the last fiscal year. Value = 62.
	ConditionOperator_LastFiscalYear
	//
	// Summary:
	//     The value is within the last fiscal period. Value = 63.
	ConditionOperator_LastFiscalPeriod
	//
	// Summary:
	//     The value is within the last X (specified value) fiscal periods. Value = 0x40.
	ConditionOperator_LastXFiscalYears
	//
	// Summary:
	//     The value is within the last X (specified value) fiscal periods. Value = 65.
	ConditionOperator_LastXFiscalPeriods
	//
	// Summary:
	//     The value is within the next X (specified value) fiscal years. Value = 66.
	ConditionOperator_NextXFiscalYears
	//
	// Summary:
	//     The value is within the next X (specified value) fiscal period. Value = 67.
	ConditionOperator_NextXFiscalPeriods
	//
	// Summary:
	//     The value is within the specified year. Value = 68.
	ConditionOperator_InFiscalYear
	//
	// Summary:
	//     The value is within the specified fiscal period. Value = 69.
	ConditionOperator_InFiscalPeriod
	//
	// Summary:
	//     The value is within the specified fiscal period and year. Value = 70.
	ConditionOperator_InFiscalPeriodAndYear
	//
	// Summary:
	//     The value is within or before the specified fiscal period and year. Value = 71.
	ConditionOperator_InOrBeforeFiscalPeriodAndYear
	//
	// Summary:
	//     The value is within or after the specified fiscal period and year. Value = 72.
	ConditionOperator_InOrAfterFiscalPeriodAndYear
	//
	// Summary:
	//     The record is owned by teams that the user is a member of. Value = 73.
	ConditionOperator_EqualUserTeams
	//
	// Summary:
	//     The record is owned by a user or teams that the user is a member of. Value =
	//     74.
	ConditionOperator_EqualUserOrUserTeams
	//
	// Summary:
	//     Returns all child records below the referenced record in the hierarchy. Value
	//     = 76.
	ConditionOperator_Under
	//
	// Summary:
	//     Returns all records not below the referenced record in the hierarchy. Value =
	//     77.
	ConditionOperator_NotUnder
	//
	// Summary:
	//     Returns the referenced record and all child records below it in the hierarchy.
	//     Value = 79.
	ConditionOperator_UnderOrEqual
	//
	// Summary:
	//     Returns all records in referenced record&#39;s hierarchical ancestry line. Value
	//     = 75.
	ConditionOperator_Above
	//
	// Summary:
	//     Returns the referenced record and all records above it in the hierarchy. Value
	//     = 78.
	ConditionOperator_AboveOrEqual
	//
	// Summary:
	//     When hierarchical security models are used, Equals current user or their reporting
	//     hierarchy. Value = 80.
	ConditionOperator_EqualUserOrUserHierarchy
	//
	// Summary:
	//     When hierarchical security models are used, Equals current user and his teams
	//     or their reporting hierarchy and their teams. Value = 81.
	ConditionOperator_EqualUserOrUserHierarchyAndTeams
	//
	ConditionOperator_OlderThanXYears
	//
	ConditionOperator_OlderThanXWeeks
	//
	ConditionOperator_OlderThanXDays
	//
	ConditionOperator_OlderThanXHours
	//
	ConditionOperator_OlderThanXMinutes
	//
	ConditionOperator_ContainValues
	//
	ConditionOperator_DoesNotContainValues
)
