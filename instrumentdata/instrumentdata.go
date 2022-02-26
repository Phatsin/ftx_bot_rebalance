package instrumentdata

type InstrumentData struct {
	Minsize float32
}

func GetInstrumentData(broker, instrument string) InstrumentData {

	_ = broker
	data := make(map[string]InstrumentData)

	data["SRM/USD"] = InstrumentData{
		1,
	}

	data["SOL/USD"] = InstrumentData{
		0.01,
	}

	return data[instrument]
}
