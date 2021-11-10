package operations

import "context"

func (o *Operations) GetBalance(ctx context.Context, ecosystem, address string) (float64, error) {
	balance, err := o.ecosystems.GetBalance(ctx, ecosystem, address)
	if err != nil {
		return 0, err
	}
	return balance.Value, nil
}
