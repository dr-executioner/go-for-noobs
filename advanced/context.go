package advanced

import (
	"context"
	"fmt"
	"time"
)

func Worker(ctx context.Context, id int) {
	select {
	case <-ctx.Done():
		fmt.Printf("Worker %d is running \n", id)
		
		return
	default:
		fmt.Printf("Worker %d is still running \n", id)
		time.Sleep(10 * time.Second)
	}
}
