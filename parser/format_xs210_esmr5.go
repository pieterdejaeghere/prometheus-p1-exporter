package parser

var XS210ESMR5TelegramFormat = TelegramFormat{
	KeyPeak:                       "1-0:1.6.0",
	KeyTimestamp:                  "0-0:1.0.0",
	KeyElectricityUsageLow:        "1-0:1.8.1",
	KeyElectricityUsageHigh:       "1-0:1.8.2",
	KeyElectricityReturnedLow:     "1-0:2.8.1",
	KeyElectricityReturnedHigh:    "1-0:2.8.2",
	KeyActiveTariff:               "0-0:96.14.0",
	KeyActualElectricityDelivered: "1-0:1.7.0",
	KeyActualElectricityRetreived: "1-0:2.7.0",
	KeyPowerFailuresShort:         "0-0:96.7.21",
	KeyPowerFailuresLong:          "0-0:96.7.9",
	KeyGasUsage:                   "0-1:24.2.3",
}
