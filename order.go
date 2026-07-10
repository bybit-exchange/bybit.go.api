package bybit_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

type Order struct {
	c                     *Client
	category              string
	symbol                string
	isLeverage            *int
	side                  string
	orderType             string
	qty                   string
	price                 *string
	triggerDirection      *int
	orderFilter           *string
	triggerPrice          *string
	triggerBy             *string
	orderIv               *string
	timeInForce           *string
	positionIdx           *int
	orderLinkId           *string
	takeProfit            *string
	stopLoss              *string
	tpTriggerBy           *string
	slTriggerBy           *string
	reduceOnly            *bool
	closeOnTrigger        *bool
	smpType               *string
	mmp                   *bool
	tpslMode              *string
	tpLimitPrice          *string
	slLimitPrice          *string
	tpOrderType           *string
	slOrderType           *string
	marketUnit            *string
	slippageToleranceType *string
	slippageTolerance     *string
}

func (order *Order) SlippageToleranceType(slippageToleranceType string) *Order {
	order.slippageToleranceType = &slippageToleranceType
	return order
}

func (order *Order) SlippageTolerance(slippageTolerance string) *Order {
	order.slippageTolerance = &slippageTolerance
	return order
}

func (order *Order) TimeInForce(tif string) *Order {
	order.timeInForce = &tif
	return order
}

func (order *Order) IsLeverage(isLeverage int) *Order {
	order.isLeverage = &isLeverage
	return order
}

func (order *Order) TriggerPrice(triggerPrice string) *Order {
	order.triggerPrice = &triggerPrice
	return order
}

func (order *Order) OrderLinkId(orderLinkId string) *Order {
	order.orderLinkId = &orderLinkId
	return order
}

func (order *Order) Price(price string) *Order {
	order.price = &price
	return order
}

func (order *Order) TriggerDirection(direction int) *Order {
	order.triggerDirection = &direction
	return order
}

func (order *Order) OrderFilter(filter string) *Order {
	order.orderFilter = &filter
	return order
}

func (order *Order) TriggerBy(triggerBy string) *Order {
	order.triggerBy = &triggerBy
	return order
}

func (order *Order) OrderIv(iv string) *Order {
	order.orderIv = &iv
	return order
}

func (order *Order) PositionIdx(idx int) *Order {
	order.positionIdx = &idx
	return order
}

func (order *Order) TakeProfit(profit string) *Order {
	order.takeProfit = &profit
	return order
}

func (order *Order) StopLoss(loss string) *Order {
	order.stopLoss = &loss
	return order
}

func (order *Order) TpTriggerBy(triggerBy string) *Order {
	order.tpTriggerBy = &triggerBy
	return order
}

func (order *Order) SlTriggerBy(triggerBy string) *Order {
	order.slTriggerBy = &triggerBy
	return order
}

func (order *Order) ReduceOnly(reduce bool) *Order {
	order.reduceOnly = &reduce
	return order
}

func (order *Order) CloseOnTrigger(close bool) *Order {
	order.closeOnTrigger = &close
	return order
}

func (order *Order) SmpType(smp string) *Order {
	order.smpType = &smp
	return order
}

func (order *Order) Mmp(mmp bool) *Order {
	order.mmp = &mmp
	return order
}

func (order *Order) TpslMode(mode string) *Order {
	order.tpslMode = &mode
	return order
}

func (order *Order) TpLimitPrice(price string) *Order {
	order.tpLimitPrice = &price
	return order
}

func (order *Order) SlLimitPrice(price string) *Order {
	order.slLimitPrice = &price
	return order
}

func (order *Order) TpOrderType(orderType string) *Order {
	order.tpOrderType = &orderType
	return order
}

func (order *Order) SlOrderType(orderType string) *Order {
	order.slOrderType = &orderType
	return order
}

func (order *Order) MarketUnit(marketUnit string) *Order {
	order.marketUnit = &marketUnit
	return order
}

func (order *Order) Do(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/v5/order/create",
		secType:  secTypeSigned,
	}
	m := params{
		"category":  order.category,
		"symbol":    order.symbol,
		"side":      order.side,
		"orderType": order.orderType,
		"qty":       order.qty,
	}
	if order.price != nil {
		m["price"] = *order.price
	}
	if order.isLeverage != nil {
		m["isLeverage"] = *order.isLeverage
	}
	if order.triggerDirection != nil {
		m["triggerDirection"] = *order.triggerDirection
	}
	if order.orderFilter != nil {
		m["orderFilter"] = *order.orderFilter
	}
	if order.triggerPrice != nil {
		m["triggerPrice"] = *order.triggerPrice
	}
	if order.triggerBy != nil {
		m["triggerBy"] = *order.triggerBy
	}
	if order.orderIv != nil {
		m["orderIv"] = *order.orderIv
	}
	if order.timeInForce != nil {
		m["timeInForce"] = *order.timeInForce
	}
	if order.positionIdx != nil {
		m["positionIdx"] = *order.positionIdx
	}
	if order.orderLinkId != nil {
		m["orderLinkId"] = *order.orderLinkId
	}
	if order.takeProfit != nil {
		m["takeProfit"] = *order.takeProfit
	}
	if order.stopLoss != nil {
		m["stopLoss"] = *order.stopLoss
	}
	if order.tpTriggerBy != nil {
		m["tpTriggerBy"] = *order.tpTriggerBy
	}
	if order.slTriggerBy != nil {
		m["slTriggerBy"] = *order.slTriggerBy
	}
	if order.reduceOnly != nil {
		m["reduceOnly"] = *order.reduceOnly
	}
	if order.closeOnTrigger != nil {
		m["closeOnTrigger"] = *order.closeOnTrigger
	}
	if order.smpType != nil {
		m["smpType"] = *order.smpType
	}
	if order.mmp != nil {
		m["mmp"] = *order.mmp
	}
	if order.tpslMode != nil {
		m["tpslMode"] = *order.tpslMode
	}
	if order.tpLimitPrice != nil {
		m["tpLimitPrice"] = *order.tpLimitPrice
	}
	if order.slLimitPrice != nil {
		m["slLimitPrice"] = *order.slLimitPrice
	}
	if order.tpOrderType != nil {
		m["tpOrderType"] = *order.tpOrderType
	}
	if order.slOrderType != nil {
		m["slOrderType"] = *order.slOrderType
	}
	if order.marketUnit != nil {
		m["marketUnit"] = *order.marketUnit
	}
	if order.slippageTolerance != nil {
		m["slippageTolerance"] = *order.slippageTolerance
	}
	if order.slippageToleranceType != nil {
		m["slippageToleranceType"] = *order.slippageToleranceType
	}
	r.setParams(m)
	data, err := order.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ServerResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
