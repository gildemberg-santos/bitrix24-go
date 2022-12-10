package lead

type Lead struct {
	Address            string  `json:"ADDRESS"`
	Address2           string  `json:"ADDRESS_2"`
	AddressCity        string  `json:"ADDRESS_CITY"`
	AddressCountry     string  `json:"ADDRESS_COUNTRY"`
	AddressCountryCode string  `json:"ADDRESS_COUNTRY_CODE"`
	AddressPostalCode  string  `json:"ADDRESS_POSTAL_CODE"`
	AddressProvice     string  `json:"ADDRESS_PROVICE"`
	AddressRegion      string  `json:"ADDRESS_REGION"`
	AssignedById       string  `json:"ASSIGNED_BY_ID"`
	Birthday           string  `json:"BIRTHDAY"`
	Comments           string  `json:"COMMENTS"`
	CompanyId          string  `json:"COMPANY_ID"`
	CompanyTitle       string  `json:"COMPANY_TITLE"`
	ContactId          string  `json:"CONTACT_ID"`
	CreatedById        string  `json:"CREATED_BY_ID"`
	DateCreate         string  `json:"DATE_CREATE"`
	DateModify         string  `json:"DATE_MODIFY"`
	Email              []Email `json:"EMAIL"`
	HasEmail           string  `json:"HAS_EMAIL"`
	HasPhone           string  `json:"HAS_PHONE"`
	Honorific          string  `json:"HONORIFIC"`
	Id                 string  `json:"ID"`
	Im                 []Im    `json:"IM"`
	IsReturnCastomer   string  `json:"IS_RETURN_CASTOMER"`
	ModifyById         string  `json:"MODIFY_BY_ID"`
	MovedTime          string  `json:"MOVED_TIME"`
	Name               string  `json:"NAME"`
	Opened             string  `json:"OPENED"`
	OriginatorId       string  `json:"ORIGINATOR_ID"`
	OriginId           string  `json:"ORIGIN_ID"`
	OriginVersion      string  `json:"ORIGIN_VERSION"`
	Phone              []Phone `json:"PHONE"`
	Photo              string  `json:"PHOTO"`
	Post               string  `json:"POST"`
	SecondName         string  `json:"SECOND_NAME"`
	SourceDescription  string  `json:"SOURCE_DESCRIPTION"`
	SourceId           string  `json:"SOURCE_ID"`
	StatusDescription  string  `json:"STATUS_DESCRIPTION"`
	StatusId           string  `json:"STATUS_ID"`
	StatusSemantics    string  `json:"STATUS_SEMANTICS"`
	Title              string  `json:"TITLE"`
	UtmCampaign        string  `json:"UTM_CAMPAIGN"`
	UtmContent         string  `json:"UTM_CONTENT"`
	UtmMedium          string  `json:"UTM_MEDIUM"`
	UtmSource          string  `json:"UTM_SOURCE"`
	UtmTerm            string  `json:"UTM_TERM"`
	Web                []Web   `json:"WEB"`
	LastName           string  `json:"LAST_NAME"`
	Opportunity        string  `json:"OPPORTUNITY"`
}

type Phone struct {
	Value     string `json:"VALUE"`
	ValueType string `json:"VALUE_TYPE"`
}

type Web struct {
	Value     string `json:"VALUE"`
	ValueType string `json:"VALUE_TYPE"`
}

type Email struct {
	Value     string `json:"VALUE"`
	ValueType string `json:"VALUE_TYPE"`
}

type Im struct {
	Value     string `json:"VALUE"`
	ValueType string `json:"VALUE_TYPE"`
}
