package value_object

type MoneyPlus struct {
	fieldName string
	total     float64
}

func NewMoneyPlus(fieldName string, total float64) MoneyPlus {
	return MoneyPlus{
		fieldName,
		total,
	}
}

func (money *MoneyPlus) GetFieldName() string {
	return money.fieldName
}

func (money *MoneyPlus) GetTotal() float64 {
	return money.total
}
