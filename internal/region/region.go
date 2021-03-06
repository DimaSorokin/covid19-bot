package region

var (
	regions = map[string]string{
		"Вінниця":          "Вінницька",
		"Луцьк":            "Волинська",
		"Дніпро":           "Дніпропетровська",
		"Донецьк":          "Донецька",
		"Житомир":          "Житомирська",
		"Ужгород":          "Закарпатська",
		"Запоріжжя":        "Запорізька",
		"Івано-Франківськ": "Івано-Франківська",
		"Київ":             "Київська",
		"Кіровоград":       "Кіровоградська",
		"Львів":            "Львівська",
		//"м.Київ":            "м.Київ",
		"Миколаїв":   "Миколаївська",
		"Одеса":      "Одеська",
		"Полтава":    "Полтавська",
		"Рівне":      "Рівненська",
		"Суми":       "Сумська",
		"Тернопіль":  "Тернопільська",
		"Харків":     "Харківська",
		"Херсонс":    "Херсонська",
		"Хмельницьк": "Хмельницька",
		"Черкаси":    "Черкаська",
		"Чернівці":   "Чернівецька",
		"Чернігівів": "Чернігівська",
		"Луганськ":   "Луганська",
	}
)

func CheckRegion(searchRegion string) string {
	if val, ok := regions[searchRegion]; ok {
		return val
	} else {
		return "unknown"
	}
}

func GetRegion() map[string]string {
	return regions
}
