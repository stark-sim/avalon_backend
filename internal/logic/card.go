package logic

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/stark-sim/avalon_backend/internal/db"
	"github.com/stark-sim/avalon_backend/pkg/ent"
	"github.com/stark-sim/avalon_backend/pkg/ent/card"
	"github.com/stark-sim/avalon_backend/tools"
)

func GetShuffledCardsByNum(ctx context.Context, num uint8, selection []*ent.Card) ([]*ent.Card, error) {
	if num < 5 || num > 12 {
		return nil, errors.New("player's num must between 5 and 12")
	}
	client := db.NewDBClient()
	var (
		nameList []card.Name
		roleList []card.Role
	)
	// 如果预先定好可选卡牌，则从其中选择
	if selection != nil {
		if len(selection) != int(num) {
			return nil, errors.New("selection must match player number")
		}
		nameList = make([]card.Name, num)
		roleList = make([]card.Role, num)
		for _, v := range selection {
			nameList = append(nameList, v.Name)
			roleList = append(roleList, v.Role)
		}
	} else {
		// 不传 selection，则根据人数给出对应卡牌
		switch num {
		case 5:
			nameList = []card.Name{card.NameMerlin, card.NamePercival, card.NameGalahad, card.NameMorgana, card.NameAssassin}
			roleList = []card.Role{card.RoleProphet, card.RoleKnight, card.RoleLoyal, card.RoleEnchantress, card.RoleAssassin}
		case 6:
			nameList = []card.Name{card.NameMerlin, card.NamePercival, card.NameGalahad, card.NameBors, card.NameMorgana, card.NameAssassin}
			roleList = []card.Role{card.RoleProphet, card.RoleKnight, card.RoleLoyal, card.RoleEnchantress, card.RoleAssassin}
		case 7:
			nameList = []card.Name{card.NameMerlin, card.NamePercival, card.NameGalahad, card.NameBors, card.NameMorgana, card.NameAssassin, card.NameOberon}
			roleList = []card.Role{card.RoleProphet, card.RoleKnight, card.RoleLoyal, card.RoleEnchantress, card.RoleAssassin}
		case 8:
			nameList = []card.Name{card.NameMerlin, card.NamePercival, card.NameGalahad, card.NameBors, card.NameGawain, card.NameMorgana, card.NameBedivere, card.NameAssassin, card.NameOberon}
			roleList = []card.Role{card.RoleProphet, card.RoleKnight, card.RoleLoyal, card.RoleEnchantress, card.RoleAssassin, card.RoleErlking}
		case 9:
			nameList = []card.Name{card.NameMerlin, card.NamePercival, card.NameGalahad, card.NameBors, card.NameGawain, card.NameBedivere, card.NameMorgana, card.NameAssassin, card.NameMordred}
			roleList = []card.Role{card.RoleProphet, card.RoleKnight, card.RoleLoyal, card.RoleEnchantress, card.RoleAssassin, card.RoleUsurper}
		case 10:
			nameList = []card.Name{card.NameMerlin, card.NamePercival, card.NameGalahad, card.NameBors, card.NameGawain, card.NameBedivere, card.NameMorgana, card.NameAssassin, card.NameMordred, card.NameOberon}
			roleList = []card.Role{card.RoleProphet, card.RoleKnight, card.RoleLoyal, card.RoleEnchantress, card.RoleAssassin, card.RoleUsurper, card.RoleErlking}
		case 11:
			nameList = []card.Name{card.NameMerlin, card.NamePercival, card.NameGalahad, card.NameBors, card.NameGawain, card.NameBedivere, card.NameMorgana, card.NameAssassin, card.NameMordred, card.NameLancelot}
			roleList = []card.Role{card.RoleProphet, card.RoleKnight, card.RoleLoyal, card.RoleEnchantress, card.RoleAssassin, card.RoleUsurper, card.RoleAce, card.RoleSinner}
		case 12:
			nameList = []card.Name{card.NameMerlin, card.NamePercival, card.NameGalahad, card.NameBors, card.NameGawain, card.NameBedivere, card.NameMorgana, card.NameAssassin, card.NameMordred, card.NameOberon, card.NameLancelot}
			roleList = []card.Role{card.RoleProphet, card.RoleKnight, card.RoleLoyal, card.RoleEnchantress, card.RoleAssassin, card.RoleUsurper, card.RoleErlking, card.RoleAce, card.RoleSinner}
		}
	}
	// 获取 Card 表中的卡片，这些数据最好不要删除
	cards, err := client.Card.
		Query().
		Where(
			card.DeletedAt(tools.ZeroTime),
			card.NameIn(nameList...),
			card.RoleIn(roleList...),
		).
		All(ctx)
	if err != nil {
		logrus.Errorf("error at query cards: %v", err)
		return nil, err
	}
	// 洗牌
	for i, v := range tools.Shuffle(cards) {
		cards[i] = v.(*ent.Card)
	}
	return cards, nil
}
