package accountdata

type AccountData struct {
	Key, Secret, Sub string
	Instument        []string
}

func GetAccountData(broker, acc string) AccountData {

	_ = broker
	data := make(map[string]AccountData)

	// Maker Port
	data["maker"] = AccountData{
		"",
		"",
		"MMmorbuck",
		[]string{"SRM/USD", "XRP/USD"},
	}

	// Scalp Port
	data["scalp"] = AccountData{
		"",
		"",
		"slot9",
		[]string{"BTC/USD", "XRP/USD", "XRP-PERP", "BTC-PERP", "SRM/USD"},
	}

	// Mudley Unknow Port
	data["morbucks"] = AccountData{
		"",
		"",
		"Morbucks",
		[]string{"XRP/USD", "DOGE/USD"},
	}

	// Tiger Port
	data["tiger"] = AccountData{
		"",
		"",
		"",
		[]string{"XRP/USD", "BTC-PERP", "BTC/USD"},
	}

	// Mudley Token Port
	data["mltoken"] = AccountData{
		"",
		"",
		"MudleyCompanyB",
		[]string{"SRM/USD", "SOL/USD"},
	}

	// MWSponsor002 Port
	data["sponsor"] = AccountData{
		"",
		"",
		"MWSponsor002",
		[]string{"SOL/USD"},
	}

	data["mlwspot"] = AccountData{
		"",
		"",
		"MW0026",
		[]string{"SOL/USD"},
	}

	return data[acc]
}
