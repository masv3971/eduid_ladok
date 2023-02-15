package goladok3

import (
	"context"

	"github.com/masv3971/goladok3/ladoktypes"
)

// CheckPermission compare ladok permissions with ps
func (c *Client) CheckPermission(ctx context.Context, myPermissions Permissions) error {
	permissionErrors := ladoktypes.PermissionErrors{}

	egna, _, err := c.Kataloginformation.GetAnvandarbehorighetEgna(ctx)
	if err != nil {
		return err
	}

	if len(egna.Anvandarbehorighet) < 1 {
		return ladoktypes.ErrNotSufficientPermissions
	}

	ladokProfile, _, err := c.Kataloginformation.GetBehorighetsprofil(ctx, &GetBehorighetsprofilerReq{UID: egna.Anvandarbehorighet[0].BehorighetsprofilRef.UID})
	if err != nil {
		return err
	}
	permissions, err := c.permissionUnify(ctx, *ladokProfile, myPermissions)
	if err != nil {
		return err
	}

	for permissionID, data := range permissions {
		myPermission, ok := data["my"]
		if !ok {
			// continue if ladok has permission not specified in "my"
			continue
		}

		ladokPermission, ok := data["ladok"]
		if !ok {
			// Ladok does not have the required permission
			permissionErrors = append(permissionErrors, ladoktypes.PermissionError{
				Msg:                 "Missing ladok permission",
				MissingPermissionID: permissionID,
				PermissionLevel:     c.translatePermission(ctx, data["my"]),
			})
			continue
		}

		if !c.comparePermission(ctx, ladokPermission, myPermission) {
			// ladokPermission does not reach myPermission
			myPermission := data["my"]
			permissionErrors = append(permissionErrors, ladoktypes.PermissionError{
				Msg:                 "Not sufficient permission",
				MissingPermissionID: permissionID,
				PermissionLevel:     c.translatePermission(ctx, myPermission),
			})
		}
	}

	if len(permissionErrors) > 0 {
		return permissionErrors
	}
	return nil
}

// comparePermission compare l with m permission.
func (c *Client) comparePermission(ctx context.Context, l, m int64) bool {
	if l == m {
		return true
	}

	switch l {
	case 4:
		if m == 6 {
			return true
		}
	case 6:
		if m == 4 {
			return false
		}
	}

	return false
}

// permissionUnify convert ladok permission structure to something that's easier to compare.
func (c *Client) permissionUnify(ctx context.Context, l ladoktypes.KataloginformationBehorighetsprofil, p Permissions) (permissions map[int64]map[string]int64, err error) {
	if len(l.Systemaktiviteter) == 0 {
		return nil, ladoktypes.ErrNotSufficientPermissions
	}
	if len(p) == 0 {
		return nil, ladoktypes.ErrNoPermissionProvided
	}

	permissions = make(map[int64]map[string]int64)

	parse := func(permission, className string, classMap map[string]int64, key int64, store map[int64]map[string]int64) {
		switch permission {
		case "rattighetsniva.las":
			classMap[className] = 4
		case "rattighetsniva.lokal":
			classMap[className] = 6
		default:
			classMap[className] = 0
		}
	}

	for key, permission := range p {
		classMyMap := make(map[string]int64)
		parse(permission, "my", classMyMap, key, permissions)
		permissions[key] = classMyMap
	}

	for _, sys := range l.Systemaktiviteter {
		classMap := make(map[string]int64)
		key := sys.ID
		permission := sys.Rattighetsniva
		parse(permission, "ladok", classMap, key, permissions)
		val, ok := permissions[key]
		if ok {
			val["ladok"] = classMap["ladok"]
			permissions[key] = val
		} else {
			permissions[key] = classMap
		}
	}

	return permissions, nil
}

func (c *Client) translatePermission(ctx context.Context, p int64) string {
	switch p {
	case 0:
		return "NoPermission"
	case 4:
		return "rattighetsniva.las"
	case 6:
		return "rattighetsniva.lokal"
	default:
		return "Undefined"
	}
}

// PermissionID is a map of permissionID and permission
var PermissionID = map[int64]string{
	1008:   "examen.allt.las",
	11004:  "kataloginformation.las",
	21008:  "kataloginformationbehorighet.behorigheter.allt.las",
	21010:  "kataloginformationbehorighet.anvandare.las",
	31001:  "resultat.las",
	51001:  "studiedeltagande.las",
	51012:  "studiedeltagande.lasa_studieavgiftsbetalning_anmalningsavgiftsbetalning",
	51014:  "studiedeltagande.lasa_aktivitet_och_finansiering",
	51051:  "studiedeltagande.las.aktorer",
	51053:  "studiedeltagande.las.anknytning",
	51056:  "studiedeltagande.las.ej_hanterad_antagning",
	51058:  "studiedeltagande.las.dokumenteradebeslut",
	51060:  "studiedeltagande.las.grupper",
	51063:  "studiedeltagande.las.anteckning",
	51066:  "studiedeltagande.las.studiebehorigheter",
	51067:  "studiedeltagande.las.tillfallesantagningar",
	51070:  "studiedeltagande.kontrollera.mot.kurs",
	61001:  "studentinformation.lasa",
	61009:  "studentinformation.student.lasa",
	71001:  "utbildningsinformation.allman.las",
	90019:  "uppfoljning.feeds",
	91001:  "uppfoljning.allman.las",
	91011:  "uppfoljning.population.student.sok",
	91012:  "uppfoljning.population.student.addresslista",
	101030: "arendestod.las",
	860131: "extintegration.lasa",
}

func (c *Client) translateID(p int64) string {
	d, ok := PermissionID[p]
	if !ok {
		return "undefined"
	}
	return d
}

func (c *Client) environment(ctx context.Context) (string, error) {
	switch c.certificate.Subject.OrganizationalUnit[1] {
	case ladoktypes.EnvIntTestAPI:
		return ladoktypes.EnvIntTestAPI, nil
	case ladoktypes.EnvProdAPI:
		return ladoktypes.EnvProdAPI, nil
	case ladoktypes.EnvTestAPI:
		return ladoktypes.EnvTestAPI, nil
	default:
		return "", ladoktypes.ErrNoEnvFound
	}
}

// StudentDegree is a student degree.
type StudentDegree struct {
	Name string `json:"name"`
}

// MyStudentDegrees array of student degrees.
type MyStudentDegrees []StudentDegree

// MarshalPDF marshal MyStudentDegrees to PDF.
func (degrees *MyStudentDegrees) MarshalPDF() {}

// GetMyStudentDegrees get student data.
func (c *Client) GetMyStudentDegrees(ctx context.Context) (MyStudentDegrees, error) {
	myStudentDegrees := []StudentDegree{}

	return myStudentDegrees, nil
}

// IsStudentReq is a request to check if a user is a student.
type IsStudentReq struct {
	UID          string `validate:"required_without_all=Personnummer ExterntUID"`
	ExterntUID   string `validate:"required_without_all=Personnummer UID"`
	Personnummer string `validate:"required_without_all=UID ExterntUID"`
}

// IsStudent check if requested user is a student.
func (c *Client) IsStudent(ctx context.Context, req *IsStudentReq) (bool, error) {
	getStudentReq := &GetStudentReq{
		UID:          req.UID,
		ExterntUID:   req.ExterntUID,
		Personnummer: req.Personnummer,
	}
	_, _, err := c.Studentinformation.GetStudent(ctx, getStudentReq)
	if err != nil {
		return false, err
	}
	return false, nil
}
