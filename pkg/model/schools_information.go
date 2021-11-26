package model

// Schools consists of all ladok access schools
var Schools = map[string]SchoolInfo{
	"bth": {
		LongNameSv: "Blekinge tekniska högskola",
		LongNameEn: "Blekinge Institute of Technology",
	},
	"cth": {
		LongNameSv: "Chalmers tekniska högskola",
		LongNameEn: "Chalmers University of Technology",
	},
	"ehs": {
		LongNameSv: "Enskilda Högskolan Stockholm",
		LongNameEn: "University College Stockholm",
	},
	"esh": {
		LongNameSv: "Ersta Sköndal Bräcke högskola",
		LongNameEn: "Ersta Sköndal Bräcke University College",
	},
	"fhs": {
		LongNameSv: "Försvarshögskolan",
		LongNameEn: "Swedish Defence University",
	},
	"ghi": {
		LongNameSv: "Gymnastik- och idrottshögskolan",
		LongNameEn: "Swedish School of Sport and Health Sciences",
	},
	"gu": {
		LongNameSv: "Göteborgs universitet",
		LongNameEn: "University of Gothenburg",
	},
	"hda": {
		LongNameSv: "Högskolan Dalarna",
		LongNameEn: "Dalarna University",
	},
	"hb": {
		LongNameSv: "Högskolan i Borås",
		LongNameEn: "University of Borås",
	},
	"hig": {
		LongNameSv: "Högskolan i Gävle",
		LongNameEn: "University of Gävle",
	},
	"hh": {
		LongNameSv: "Högskolan i Halmstad",
		LongNameEn: "Halmstad University",
	},
	"hs": {
		LongNameSv: "Högskolan i Skövde",
		LongNameEn: "University of Skövde",
	},
	"hkr": {
		LongNameSv: "Högskolan Kristianstad",
		LongNameEn: "Kristianstad University",
	},
	"hv": {
		LongNameSv: "Högskolan Väst",
		LongNameEn: "University West",
	},
	"kau": {
		LongNameSv: "Karlstads universitet",
		LongNameEn: "Karlstad University",
	},
	"ki": {
		LongNameSv: "Karolinska institutet",
		LongNameEn: "Karolinska Institute",
	},
	"kf": {
		LongNameSv: "Konstfack",
		LongNameEn: "University of Arts, Crafts and Design",
	},
	"kkh": {
		LongNameSv: "Kungliga Konsthögskolan",
		LongNameEn: "Royal Institute of Art",
	},
	"kmh": {
		LongNameSv: "Kungliga Musikhögskolan i Stockholm",
		LongNameEn: "Royal College of Music in Stockholm",
	},
	"kth": {
		LongNameSv: "Kungliga tekniska högskolan",
		LongNameEn: "Royal Institute of Technology",
	},
	"liu": {
		LongNameSv: "Linköpings universitet",
		LongNameEn: "Linköping University",
	},
	"lnu": {
		LongNameSv: "Linnéuniversitetet",
		LongNameEn: "Linnaeus University",
	},
	"ltu": {
		LongNameSv: "Luleå tekniska universitet",
		LongNameEn: "Luleå University of Technology",
	},
	"lu": {
		LongNameSv: "Lunds universitet",
		LongNameEn: "Lund University",
	},
	"mau": {
		LongNameSv: "Malmö universitet",
		LongNameEn: "Malmö university",
	},
	"miu": {
		LongNameSv: "Mittuniversitetet",
		LongNameEn: "Mid Sweden University",
	},
	"mdh": {
		LongNameSv: "Mälardalens högskola",
		LongNameEn: "Mälardalen University",
	},
	"ni": {
		LongNameSv: "Newmaninstitutet",
		LongNameEn: "Newman Institute University College",
	},
	"rkh": {
		LongNameSv: "Röda Korsets högskola",
		LongNameEn: "The Swedish Red Cross University College",
	},
	"shh": {
		LongNameSv: "Sophiahemmet Högskola",
		LongNameEn: "Sophiahemmet University",
	},
	"hj": {
		LongNameSv: "Stiftelsen Högskolan i Jönköping",
		LongNameEn: "Jönköping University",
	},
	"skh": {
		LongNameSv: "Stockholms konstnärliga högskola",
		LongNameEn: "Stockholm University of the Arts",
	},
	"su": {
		LongNameSv: "Stockholms universitet",
		LongNameEn: "Stockholm University",
	},
	"slu": {
		LongNameSv: "Sveriges lantbruksuniversitet",
		LongNameEn: "Swedish University of Agricultural Sciences",
	},
	"sh": {
		LongNameSv: "Södertörns högskola",
		LongNameEn: "Södertörn University",
	},
	"umu": {
		LongNameSv: "Umeå universitet",
		LongNameEn: "Umeå University",
	},
	"uu": {
		LongNameSv: "Uppsala universitet",
		LongNameEn: "Uppsala University",
	},
	"oru": {
		LongNameSv: "Örebro universitet",
		LongNameEn: "Örebro University",
	},
}

// SchoolsList a list of abbreviated school names
var SchoolsList = []string{"bth", "cth", "ehs", "esh", "fhs", "ghi", "gu", "hda", "hb", "hig", "hh", "hs", "hkr", "hv", "kau", "ki", "kf", "kkh", "kmh", "kth", "liu", "lnu", "ltu", "lu", "mau", "miu", "mdh", "ni", "rkh", "shh", "hj", "skh", "su", "slu", "sh", "umu", "uu", "oru"}

// SchoolInfo collect information about a school
type SchoolInfo struct {
	LongNameSv string `json:"long_name_sv"`
	LongNameEn string `json:"long_name_en"`
}
