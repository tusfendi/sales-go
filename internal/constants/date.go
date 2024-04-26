package constants

const (
	FormatDate      = `2006-01-02`
	DateMonthFormat = `02 January 2006`

	// Month
	January   = `January`
	February  = `February`
	March     = `March`
	April     = `April`
	May       = `May`
	June      = `June`
	July      = `July`
	August    = `August`
	September = `September`
	October   = `October`
	November  = `November`
	December  = `December`
)

// Month to Bulan
var MappingMonthToBulan = map[string]string{
	January:   `Januari`,
	February:  `Februari`,
	March:     `Maret`,
	April:     `April`,
	May:       `Mei`,
	June:      `Juni`,
	July:      `Juli`,
	August:    `Agustus`,
	September: `September`,
	October:   `Oktober`,
	November:  `November`,
	December:  `Desember`,
}
