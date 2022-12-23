package logic

import "github.com/sirupsen/logrus"

func GetMissionCapacityByNumAndSeq(num uint8, seq int) uint8 {
	if num < 5 || num > 12 {
		logrus.Errorf("GetMissionCapacity wrong player num")
		return 0
	}
	if seq < 1 || seq > 5 {
		logrus.Errorf("GetMissionCapacity wrong mission seq")
	}
	switch seq {
	case 1:
		switch num {
		case 5, 6, 7:
			return 2
		case 8, 9, 10:
			return 3
		case 11, 12:
			return 4
		}
	case 2:
		switch num {
		case 5:
			return 2
		case 6, 7:
			return 3
		case 8, 9, 10:
			return 4
		case 11, 12:
			return 4
		}
	case 3:
		switch num {
		case 5:
			return 2
		case 6, 7:
			return 3
		case 8, 9, 10:
			return 4
		case 11, 12:
			return 5
		}
	case 4:
		switch num {
		case 5, 6:
			return 3
		case 7:
			return 4
		case 8, 9, 10:
			return 5
		case 11, 12:
			return 6
		}
	case 5:
		switch num {
		case 5:
			return 3
		case 6, 7:
			return 4
		case 8, 9, 10:
			return 5
		case 11, 12:
			return 6
		}
	}
	return 0
}
