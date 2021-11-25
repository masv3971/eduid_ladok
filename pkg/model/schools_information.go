package model

// Schools consists of all ladok access schools
var Schools = map[string]SchoolInfo{
	"bth": {
		LongNameSv: "blekinge tekniska högskola",
		LongNameEn: "",
	},
	"cth": {
		LongNameSv: "Chalmers tekniska högskola",
		LongNameEn: "",
	},
	"ehs": {
		LongNameSv: "Enskilda Högskolan Stockholm",
		LongNameEn: "",
	},
	"esh": {
		LongNameSv: "Ersta Sköndal Bräcke högskola",
		LongNameEn: "",
	},
	"fhs": {
		LongNameSv: "Försvarshögskolan",
		LongNameEn: "",
	},
	"ghi": {
		LongNameSv: "Gymnastik- och idrottshögskolan",
		LongNameEn: "",
	},
	"gu": {
		LongNameSv: "Göteborgs universitet",
		LongNameEn: "",
	},
	"hda": {
		LongNameSv: "Högskolan Dalarna",
		LongNameEn: "",
	},
	"hb": {
		LongNameSv: "Högskolan i Borås",
		LongNameEn: "",
	},
	"hig": {
		LongNameSv: "Högskolan i Gävle",
		LongNameEn: "",
	},
	"hh": {
		LongNameSv: "Högskolan i Halmstad",
		LongNameEn: "",
	},
	"hs": {
		LongNameSv: "Högskolan i Skövde",
		LongNameEn: "",
	},
	"hkr": {
		LongNameSv: "Högskolan Kristianstad",
		LongNameEn: "",
	},
	"hv": {
		LongNameSv: "Högskolan Väst",
		LongNameEn: "",
	},
	"kau": {
		LongNameSv: "Karlstads universitet",
		LongNameEn: "",
	},
	"ki": {
		LongNameSv: "Karolinska institutet",
		LongNameEn: "",
	},
	"kf": {
		LongNameSv: "Konstfack",
		LongNameEn: "",
	},
	"kkh": {
		LongNameSv: "Kungl. Konsthögskolan",
		LongNameEn: "",
	},
	"kmh": {
		LongNameSv: "Kungl. Musikhögskolan i Stockholm",
		LongNameEn: "",
	},
	"kth": {
		LongNameSv: "kungliga tekniska högskolan",
		LongNameEn: "",
	},
	"liu": {
		LongNameSv: "Linköpings universitet",
		LongNameEn: "",
	},
	"lnu": {
		LongNameSv: "linne universitetet",
		LongNameEn: "",
	},
	"ltu": {
		LongNameSv: "Luleå tekniska universitet",
		LongNameEn: "",
	},
	"lu": {
		LongNameSv: "Lunds universitet",
		LongNameEn: "",
	},
	"mau": {
		LongNameSv: "Malmö universitet",
		LongNameEn: "",
	},
	"miu": {
		LongNameSv: "Mittuniversitetet",
		LongNameEn: "",
	},
	"mdh": {
		LongNameSv: "Mälardalens högskola",
		LongNameEn: "",
	},
	"ni": {
		LongNameSv: "Newmaninstitutet",
		LongNameEn: "",
	},
	"rkh": {
		LongNameSv: "Röda Korsets högskola",
		LongNameEn: "",
	},
	"shh": {
		LongNameSv: "Sophiahemmet Högskola",
		LongNameEn: "",
	},
	"hj": {
		LongNameSv: "Stiftelsen Högskolan i Jönköping",
		LongNameEn: "",
	},
	"skh": {
		LongNameSv: "Stockholms konstnärliga högskola",
		LongNameEn: "",
	},
	"su": {
		LongNameSv: "Stockholms universitet",
		LongNameEn: "",
	},
	"slu": {
		LongNameSv: "Sveriges lantbruksuniversitet",
		LongNameEn: "",
	},
	"sh": {
		LongNameSv: "Södertörns högskola",
		LongNameEn: "",
	},
	"umu": {
		LongNameSv: "Umeå universitet",
		LongNameEn: "",
	},
	"uu": {
		LongNameSv: "Uppsala universitet",
		LongNameEn: "",
	},
	"oru": {
		LongNameSv: "Örebro universitet",
		LongNameEn: "",
	},
}

// SchoolsList a list of abbreviated school names
var SchoolsList = []string{"bth", "cth", "ehs", "esh", "fhs", "ghi", "gu", "hda", "hb", "hig", "hh", "hs", "hkr", "hv", "kau", "ki", "kf", "kkh", "kmh", "kth", "liu", "lnu", "ltu", "lu", "mau", "miu", "mdh", "ni", "rkh", "shh", "hj", "skh", "su", "slu", "sh", "umu", "uu", "oru"}

// SchoolInfo collect information about a school
type SchoolInfo struct {
	LongNameSv string `json:"long_name_sv"`
	LongNameEn string `json:"long_name_en"`
}
