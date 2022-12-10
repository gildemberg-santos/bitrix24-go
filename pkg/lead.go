package pkg

type Phone struct {
	Value     string `json:"VALUE"`
	ValueType string `json:"VALUE_TYPE"`
}

type Web struct {
	Value     string `json:"VALUE"`
	ValueType string `json:"VALUE_TYPE"`
}

type Lead struct {
	Title        string  `json:"TITLE"`
	Name         string  `json:"NAME"`
	SecondName   string  `json:"SECOND_NAME"`
	LastName     string  `json:"LAST_NAME"`
	StatusId     string  `json:"STATUS_ID"`
	Opened       string  `json:"OPENED"`
	AssignedById string  `json:"ASSIGNED_BY_ID"`
	CompanyId    string  `json:"COMPANY_ID"`
	Opportunity  string  `json:"OPPORTUNITY"`
	Phone        []Phone `json:"PHONE"`
	Web          []Web   `json:"WEB"`
}
