package handlers

// getAlpha3Code converts ISO 3166-1 alpha-2 to alpha-3
func getAlpha3Code(alpha2 string) string {
	mapping := map[string]string{
		// North America
		"US": "USA", "CA": "CAN", "MX": "MEX",
		// South America
		"AR": "ARG", "BR": "BRA", "CL": "CHL", "CO": "COL", "PE": "PER",
		"VE": "VEN", "EC": "ECU", "BO": "BOL", "PY": "PRY", "UY": "URY",
		// Western Europe
		"GB": "GBR", "DE": "DEU", "FR": "FRA", "ES": "ESP", "IT": "ITA",
		"NL": "NLD", "BE": "BEL", "CH": "CHE", "AT": "AUT", "PT": "PRT",
		"IE": "IRL", "LU": "LUX",
		// Northern Europe
		"SE": "SWE", "NO": "NOR", "DK": "DNK", "FI": "FIN", "IS": "ISL",
		// Eastern Europe
		"PL": "POL", "CZ": "CZE", "SK": "SVK", "HU": "HUN", "RO": "ROU",
		"BG": "BGR", "UA": "UKR", "BY": "BLR", "RU": "RUS", "MD": "MDA",
		"LT": "LTU", "LV": "LVA", "EE": "EST",
		// Southern Europe
		"GR": "GRC", "HR": "HRV", "SI": "SVN", "RS": "SRB", "BA": "BIH",
		"ME": "MNE", "MK": "MKD", "AL": "ALB", "CY": "CYP", "MT": "MLT",
		// Middle East
		"IL": "ISR", "SA": "SAU", "AE": "ARE", "TR": "TUR", "IR": "IRN",
		"IQ": "IRQ", "JO": "JOR", "LB": "LBN", "SY": "SYR", "YE": "YEM",
		"OM": "OMN", "KW": "KWT", "BH": "BHR", "QA": "QAT", "PS": "PSE",
		// East Asia
		"CN": "CHN", "JP": "JPN", "KR": "KOR", "KP": "PRK", "TW": "TWN",
		"HK": "HKG", "MO": "MAC", "MN": "MNG",
		// Southeast Asia
		"TH": "THA", "VN": "VNM", "PH": "PHL", "ID": "IDN", "MY": "MYS",
		"SG": "SGP", "MM": "MMR", "KH": "KHM", "LA": "LAO", "BN": "BRN",
		"TL": "TLS",
		// South Asia
		"IN": "IND", "PK": "PAK", "BD": "BGD", "LK": "LKA", "NP": "NPL",
		"AF": "AFG", "BT": "BTN", "MV": "MDV",
		// Central Asia
		"KZ": "KAZ", "UZ": "UZB", "TM": "TKM", "KG": "KGZ", "TJ": "TJK",
		// Africa - North
		"EG": "EGY", "DZ": "DZA", "MA": "MAR", "TN": "TUN", "LY": "LBY",
		"SD": "SDN", "SS": "SSD",
		// Africa - West
		"NG": "NGA", "GH": "GHA", "CI": "CIV", "SN": "SEN", "ML": "MLI",
		"BF": "BFA", "NE": "NER", "GN": "GIN", "BJ": "BEN", "TG": "TGO",
		"LR": "LBR", "SL": "SLE", "GM": "GMB", "GW": "GNB", "MR": "MRT",
		// Africa - East
		"KE": "KEN", "ET": "ETH", "TZ": "TZA", "UG": "UGA", "SO": "SOM",
		"RW": "RWA", "BI": "BDI", "DJ": "DJI", "ER": "ERI",
		// Africa - Central
		"CD": "COD", "CM": "CMR", "AO": "AGO", "TD": "TCD", "CF": "CAF",
		"CG": "COG", "GA": "GAB", "GQ": "GNQ", "ST": "STP",
		// Africa - South
		"ZA": "ZAF", "ZW": "ZWE", "ZM": "ZMB", "MW": "MWI", "MZ": "MOZ",
		"BW": "BWA", "NA": "NAM", "LS": "LSO", "SZ": "SWZ", "MG": "MDG",
		"MU": "MUS", "SC": "SYC", "KM": "COM", "RE": "REU",
		// Oceania
		"AU": "AUS", "NZ": "NZL", "PG": "PNG", "FJ": "FJI", "NC": "NCL",
		"PF": "PYF", "SB": "SLB", "VU": "VUT", "WS": "WSM", "GU": "GUM",
		"AS": "ASM", "MP": "MNP", "FM": "FSM", "PW": "PLW", "MH": "MHL",
		"KI": "KIR", "TO": "TON", "TV": "TUV", "NR": "NRU",
		// Caribbean
		"CU": "CUB", "DO": "DOM", "HT": "HTI", "JM": "JAM", "TT": "TTO",
		"BB": "BRB", "BS": "BHS", "GD": "GRD", "LC": "LCA", "VC": "VCT",
		"AG": "ATG", "DM": "DMA", "KN": "KNA", "PR": "PRI", "VI": "VIR",
		"TC": "TCA", "KY": "CYM", "BM": "BMU", "AW": "ABW", "CW": "CUW",
		// Central America
		"GT": "GTM", "HN": "HND", "SV": "SLV", "NI": "NIC", "CR": "CRI",
		"PA": "PAN", "BZ": "BLZ",
		// Special
		"Unknown": "",
	}
	if code, ok := mapping[alpha2]; ok {
		return code
	}
	return ""
}

// getCountryName returns human-readable country names
func getCountryName(alpha2 string) string {
	names := map[string]string{
		// North America
		"US": "United States", "CA": "Canada", "MX": "Mexico",
		// South America
		"AR": "Argentina", "BR": "Brazil", "CL": "Chile", "CO": "Colombia",
		"PE": "Peru", "VE": "Venezuela", "EC": "Ecuador", "BO": "Bolivia",
		"PY": "Paraguay", "UY": "Uruguay",
		// Western Europe
		"GB": "United Kingdom", "DE": "Germany", "FR": "France", "ES": "Spain",
		"IT": "Italy", "NL": "Netherlands", "BE": "Belgium", "CH": "Switzerland",
		"AT": "Austria", "PT": "Portugal", "IE": "Ireland", "LU": "Luxembourg",
		// Northern Europe
		"SE": "Sweden", "NO": "Norway", "DK": "Denmark", "FI": "Finland",
		"IS": "Iceland",
		// Eastern Europe
		"PL": "Poland", "CZ": "Czechia", "SK": "Slovakia", "HU": "Hungary",
		"RO": "Romania", "BG": "Bulgaria", "UA": "Ukraine", "BY": "Belarus",
		"RU": "Russia", "MD": "Moldova", "LT": "Lithuania", "LV": "Latvia",
		"EE": "Estonia",
		// Southern Europe
		"GR": "Greece", "HR": "Croatia", "SI": "Slovenia", "RS": "Serbia",
		"BA": "Bosnia and Herzegovina", "ME": "Montenegro", "MK": "North Macedonia",
		"AL": "Albania", "CY": "Cyprus", "MT": "Malta",
		// Middle East
		"IL": "Israel", "SA": "Saudi Arabia", "AE": "United Arab Emirates",
		"TR": "Turkey", "IR": "Iran", "IQ": "Iraq", "JO": "Jordan",
		"LB": "Lebanon", "SY": "Syria", "YE": "Yemen", "OM": "Oman",
		"KW": "Kuwait", "BH": "Bahrain", "QA": "Qatar", "PS": "Palestine",
		// East Asia
		"CN": "China", "JP": "Japan", "KR": "South Korea", "KP": "North Korea",
		"TW": "Taiwan", "HK": "Hong Kong", "MO": "Macau", "MN": "Mongolia",
		// Southeast Asia
		"TH": "Thailand", "VN": "Vietnam", "PH": "Philippines", "ID": "Indonesia",
		"MY": "Malaysia", "SG": "Singapore", "MM": "Myanmar", "KH": "Cambodia",
		"LA": "Laos", "BN": "Brunei", "TL": "Timor-Leste",
		// South Asia
		"IN": "India", "PK": "Pakistan", "BD": "Bangladesh", "LK": "Sri Lanka",
		"NP": "Nepal", "AF": "Afghanistan", "BT": "Bhutan", "MV": "Maldives",
		// Central Asia
		"KZ": "Kazakhstan", "UZ": "Uzbekistan", "TM": "Turkmenistan",
		"KG": "Kyrgyzstan", "TJ": "Tajikistan",
		// Africa - North
		"EG": "Egypt", "DZ": "Algeria", "MA": "Morocco", "TN": "Tunisia",
		"LY": "Libya", "SD": "Sudan", "SS": "South Sudan",
		// Africa - West
		"NG": "Nigeria", "GH": "Ghana", "CI": "Côte d'Ivoire", "SN": "Senegal",
		"ML": "Mali", "BF": "Burkina Faso", "NE": "Niger", "GN": "Guinea",
		"BJ": "Benin", "TG": "Togo", "LR": "Liberia", "SL": "Sierra Leone",
		"GM": "Gambia", "GW": "Guinea-Bissau", "MR": "Mauritania",
		// Africa - East
		"KE": "Kenya", "ET": "Ethiopia", "TZ": "Tanzania", "UG": "Uganda",
		"SO": "Somalia", "RW": "Rwanda", "BI": "Burundi", "DJ": "Djibouti",
		"ER": "Eritrea",
		// Africa - Central
		"CD": "Democratic Republic of the Congo", "CM": "Cameroon", "AO": "Angola",
		"TD": "Chad", "CF": "Central African Republic", "CG": "Republic of the Congo",
		"GA": "Gabon", "GQ": "Equatorial Guinea", "ST": "São Tomé and Príncipe",
		// Africa - South
		"ZA": "South Africa", "ZW": "Zimbabwe", "ZM": "Zambia", "MW": "Malawi",
		"MZ": "Mozambique", "BW": "Botswana", "NA": "Namibia", "LS": "Lesotho",
		"SZ": "Eswatini", "MG": "Madagascar", "MU": "Mauritius", "SC": "Seychelles",
		"KM": "Comoros", "RE": "Réunion",
		// Oceania
		"AU": "Australia", "NZ": "New Zealand", "PG": "Papua New Guinea",
		"FJ": "Fiji", "NC": "New Caledonia", "PF": "French Polynesia",
		"SB": "Solomon Islands", "VU": "Vanuatu", "WS": "Samoa", "GU": "Guam",
		"AS": "American Samoa", "MP": "Northern Mariana Islands", "FM": "Micronesia",
		"PW": "Palau", "MH": "Marshall Islands", "KI": "Kiribati", "TO": "Tonga",
		"TV": "Tuvalu", "NR": "Nauru",
		// Caribbean
		"CU": "Cuba", "DO": "Dominican Republic", "HT": "Haiti", "JM": "Jamaica",
		"TT": "Trinidad and Tobago", "BB": "Barbados", "BS": "Bahamas",
		"GD": "Grenada", "LC": "Saint Lucia", "VC": "Saint Vincent and the Grenadines",
		"AG": "Antigua and Barbuda", "DM": "Dominica", "KN": "Saint Kitts and Nevis",
		"PR": "Puerto Rico", "VI": "U.S. Virgin Islands", "TC": "Turks and Caicos Islands",
		"KY": "Cayman Islands", "BM": "Bermuda", "AW": "Aruba", "CW": "Curaçao",
		// Central America
		"GT": "Guatemala", "HN": "Honduras", "SV": "El Salvador", "NI": "Nicaragua",
		"CR": "Costa Rica", "PA": "Panama", "BZ": "Belize",
		// Special
		"Unknown": "Unknown",
	}
	if name, ok := names[alpha2]; ok {
		return name
	}
	return alpha2
}
