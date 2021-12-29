package apiv1

import (
	"context"
	"fmt"
)

// ESI formats a schacPersonalUniqueCode from ladok externtUID
func (c *Client) ESI(ctx context.Context, externtUID string) string {
	_, span := c.tp.Start(ctx, "apiv1.ESI")
	defer span.End()

	schacPersonalUniqueCode := fmt.Sprintf(
		"urn:schac:personalUniqueCode:int:esi:ladok.se:externtstudentuid-%s",
		externtUID,
	)
	return schacPersonalUniqueCode
}
