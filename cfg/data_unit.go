package cfg

const (
	B = 1

	// KB ,MB ,GB ,TB ,PB are Decimal
	KB = 1000 * B
	MB = 1000 * KB
	GB = 1000 * MB
//	TB = 1000 * GB // arm32 not support
//	PB = 1000 * TB // arm32 not support

	// KiB ,MiB ,GiB ,TiB ,PiB are Binary
	KiB = 1024 * B
	MiB = 1024 * KiB
	GiB = 1024 * MiB
//	TiB = 1024 * GiB // arm32 not support
//	PiB = 1024 * TiB // arm32 not support
)

// DataUnitAliases is map to add alias.
// Alias must not contain space
// Only lowercase letters should be used in the map, but aliases are case-insensitive
var DataUnitAliases = map[string]int{
	"kb": KB, "kib": KiB,
	"mb": MB, "mib": MiB,
	"gb": GB, "gib": GiB,
//	"tb": TB, "tib": TiB, // arm32 not support
//	"pb": PB, "pib": PiB, // arm32 not support
	"b": B,
}
