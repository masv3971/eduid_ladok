package apiv1

import "fmt"

// ESI formats a schacPersonalUniqueCode from ladok externtUID
func ESI(externtUID string) string {
	schacPersonalUniqueCode := fmt.Sprintf(
		"urn:schac:personalUniqueCode:int:esi:ladok.se:externtstudentuid-%s",
		externtUID,
	)
	return schacPersonalUniqueCode
}
