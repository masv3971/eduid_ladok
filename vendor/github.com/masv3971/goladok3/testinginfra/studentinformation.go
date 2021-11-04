package testinginfra

// StudentUID uid for testing student
var StudentUID = "339A47C0-426D-4012-B83A-6427E9587352"

// JSONStudentinformationStudent ladok reply
var JSONStudentinformationStudent = []byte(`{
			"Avliden": false,
			"Efternamn": "TestEfternamn",
			"ExterntUID": "11111111-2222-0000-0000-000000000000",
			"FelVidEtableringExternt": false,
			"Fodelsedata": "1996-11-05",
			"FolkbokforingsbevakningTillOchMed": "2020-03-13",
			"Fornamn": "TestFornamn",
			"KonID": 1,
			"LarosateID": 96,
			"Personnummer": "199611052383",
			"SenastAndradAv": "testEppn@ladok3.ladok.se",
			"SenastSparad": "2012-01-11T12:45:45",
			"Uid": "11111111-2222-0000-0000-000000000000",
			"UnikaIdentifierare": {
				"LarosateID": 96,
				"UnikIdentifierare": [{
					"LarosateID": 96,
					"SenastAndradAv": "testEppn@example.com",
					"SenastSparad": "2012-01-11T12:45:45",
					"Typ": "studentinformation.domain.unikidentifieraretyp.passnummer",
					"Uid": "11111111-2222-0000-0000-000000000000",
					"Varde": "1234",
					"link": [{
						"method": "POST",
						"uri": "https://api.mit.ladok.se:443/test",
						"mediaType": "application/vnd.ladok+xml",
						"rel": "http://schemas.ladok.se"
					}]
				}],
				"link": [{
					"method": "POST",
					"uri": "https://api.mit.ladok.se:443/test",
					"mediaType": "application/vnd.ladok+xml",
					"rel": "http://schemas.ladok.se"
				}]
			},
			"link": [{
				"method": "POST",
				"uri": "https://api.mit.ladok.se:443/test",
				"mediaType": "application/vnd.ladok+xml",
				"rel": "http://schemas.ladok.se"
			}]
		}`)
