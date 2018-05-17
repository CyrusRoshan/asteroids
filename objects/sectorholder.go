package objects

type SectorHolder struct {
	sectors      [][]Sector
	sectorWidth  int
	sectorHeight int
	rows         int
	columns      int
}

func NewSectorHolder(rows int, columns int, sectorWidth int, sectorHeight int) SectorHolder {
	sectors := make([][]Sector, rows)
	for i := range sectors {
		sectors[i] = make([]Sector, columns)
	}

	return SectorHolder{
		sectors:      sectors,
		sectorWidth:  sectorWidth,
		sectorHeight: sectorHeight,
		rows:         rows,
		columns:      columns,
	}
}

func (s SectorHolder) GetSectorWidth() int {
	return s.sectorWidth
}

func (s SectorHolder) GetSectorHeight() int {
	return s.sectorHeight
}

func (s SectorHolder) GetTotalWidth() int {
	return s.sectorWidth * s.columns
}
func (s SectorHolder) GetTotalHeight() int {
	return s.sectorHeight * s.rows
}

func (s SectorHolder) GetSector(row int, column int) *Sector {
	return &s.sectors[row][column]
}

func (s SectorHolder) GetNearbySectors(row int, column int, radius int) (nearbySectors []*Sector) {
	nearbySectors = make([]*Sector, 0)

	height := len(s.sectors)
	width := len(s.sectors[row])

	for i := (row - radius); i < (row + radius); i++ {
		if i < 0 || i >= height {
			continue
		}

		for j := (column - radius); j < (column + radius); j++ {
			if j < 0 || j >= width {
				continue
			}

			nearbySectors = append(nearbySectors, &s.sectors[i][j])
		}
	}

	return
}
