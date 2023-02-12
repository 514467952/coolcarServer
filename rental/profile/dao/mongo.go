package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1/rental"
	"coolcar/shared/id"
	mgutil "coolcar/shared/mongo"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	accountIDField      = "accountid"
	profileField        = "profile"
	identityStatusField = profileField + ".identitystatus"
)

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("profile"),
	}
}

//ProfileRecord defines the profile in db
type ProfileRecord struct {
	AccountID string            `bson:"accountid"`
	Prfile    *rentalpb.Profile `bson:"profile"`
}

func (m *Mongo) GetProfile(c context.Context, aid id.AccountID) (*rentalpb.Profile, error) {
	res := m.col.FindOne(c, byAccountID(aid))

	if err := res.Err(); err != nil {
		return nil, err
	}

	var pr ProfileRecord
	err := res.Decode(&pr)
	if err != nil {
		return nil, fmt.Errorf("db cannot decode profile record:%v", err)
	}
	return pr.Prfile, nil
}

func (m *Mongo) UpdateProfile(c context.Context, aid id.AccountID, prevStatus rentalpb.IdentityStatus, p *rentalpb.Profile) error {
	//options.Update().SetUpsert(true)
	//找不到profileField，会创建一条新的记录
	_, err := m.col.UpdateOne(c, bson.M{
		accountIDField:      aid.String(),
		identityStatusField: prevStatus,
	}, mgutil.Set(bson.M{
		accountIDField: aid.String(),
		profileField:   p,
	}), options.Update().SetUpsert(true))

	return err
}

func byAccountID(aid id.AccountID) bson.M {
	return bson.M{
		accountIDField: aid.String(),
	}
}
