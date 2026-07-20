package formatter

const (
	unitB  = 1
	unitKB = 1000
	unitMB = 1000 * unitKB
	unitGB = 1000 * unitMB
	unitTB = 1000 * unitGB
	unitPB = 1000 * unitTB
	unitEB = 1000 * unitPB

	unitBiB = 1
	unitKiB = 1024
	unitMiB = 1024 * unitKiB
	unitGiB = 1024 * unitMiB
	unitTiB = 1024 * unitGiB
	unitPiB = 1024 * unitTiB
	unitEiB = 1024 * unitPiB
)

var (
	decimalByteMap = map[string]uint64{
		"b":  unitB,
		"kb": unitKB,
		"mb": unitMB,
		"gb": unitGB,
		"tb": unitTB,
		"pb": unitPB,
		"eb": unitEB,

		"":  unitB,
		"k": unitKB,
		"m": unitMB,
		"g": unitGB,
		"t": unitTB,
		"p": unitPB,
		"e": unitEB,
	}

	binaryByteMap = map[string]uint64{
		"bi":  unitBiB,
		"kib": unitKiB,
		"mib": unitMiB,
		"gib": unitGiB,
		"tib": unitTiB,
		"pib": unitPiB,
		"eib": unitEiB,

		"":   unitBiB,
		"ki": unitKiB,
		"mi": unitMiB,
		"gi": unitGiB,
		"ti": unitTiB,
		"pi": unitPiB,
		"ei": unitEiB,
	}
)

var (
	decimalByteUnits = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	binaryByteUnits  = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"}
)

func DecimalBytes(size float64, precision ...int) string { _ = "STUB: not implemented"; return "" }

func BinaryBytes(size float64, precision ...int) string { _ = "STUB: not implemented"; return "" }

func calculateByteSize(size float64, base float64, byteUnits []string) (float64, string) {
	_ = "STUB: not implemented"
	return 0, ""
}

func roundToToString(x float64, max ...int) string { _ = "STUB: not implemented"; return "" }

func ParseDecimalBytes(size string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func ParseBinaryBytes(size string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func parseBytes(s string, kind string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }
