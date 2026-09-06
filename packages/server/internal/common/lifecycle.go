package common

import "context"

func GlobalContext() (context.Context, context.CancelFunc) {

	return context.WithCancel(context.Background())

}
