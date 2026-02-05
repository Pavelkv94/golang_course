package payments

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentMethod PaymentMethod
	paymentsInfo  map[int]PaymentInfo
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentMethod: paymentMethod,
		paymentsInfo: make(map[int]PaymentInfo, 0),
	}
}

func (p *PaymentModule) Pay(description string, usd int) int {
	id := p.paymentMethod.Pay(usd)

	info := PaymentInfo{
		Description: description,
		Usd:         usd,
		Canceled:    false,
	}

	p.paymentsInfo[id] = info

	return id
}

func (p *PaymentModule) Cancel(id int) {

	info, ok := p.paymentsInfo[id] // получаем копию
	if !ok {
		return
	}

	p.paymentMethod.Cancel(id)

	info.Canceled = true

	p.paymentsInfo[id] = info

}

func (p *PaymentModule) Info(id int) PaymentInfo {
	info, ok := p.paymentsInfo[id] // получаем копию
	if !ok {
		return PaymentInfo{}
	}

	return info
}

func (p *PaymentModule) AllInfo() map[int]PaymentInfo {
	tempMap := make(map[int]PaymentInfo, len(p.paymentsInfo))

	for k, v := range p.paymentsInfo {
		tempMap[k] = v
	}

	return tempMap
}
