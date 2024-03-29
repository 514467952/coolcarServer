package profile

import (
	"context"
	blobpb "coolcar/blob/api/gen/v1/blob"
	rentalpb "coolcar/rental/api/gen/v1/rental"
	"coolcar/rental/profile/dao"
	sharedauth "coolcar/shared/auth"
	"coolcar/shared/id"
	"coolcar/shared/sharedserver"
	mongotesting "coolcar/shared/testing"
	"os"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestProfileLifecyle(t *testing.T) {
	c := context.Background()

	s := newService(c, t)

	aid := id.AccountID("account1")
	c = sharedauth.ContextWithAccountID(c, aid)
	cases := []struct {
		name       string
		op         func() (*rentalpb.Profile, error)
		wantName   string
		wantStatus rentalpb.IdentityStatus
		wantErr    bool
	}{
		{
			name: "get_empty",
			op: func() (*rentalpb.Profile, error) {
				return s.GetProfile(c, &rentalpb.GetProfileRequest{})
			},
			wantStatus: rentalpb.IdentityStatus_UNSUBMITTED,
		},
		{
			name: "submit",
			op: func() (*rentalpb.Profile, error) {
				return s.SubmitProfile(c, &rentalpb.Identity{
					Name: "abc",
				})
			},
			wantName:   "abc",
			wantStatus: rentalpb.IdentityStatus_PENDING,
		},
		{
			name: "submit_again",
			op: func() (*rentalpb.Profile, error) {
				return s.SubmitProfile(c, &rentalpb.Identity{
					Name: "abc",
				})
			},
			wantErr: true,
		},
		{
			name: "todo_force_verify",
			op: func() (*rentalpb.Profile, error) {
				p := &rentalpb.Profile{
					Identity: &rentalpb.Identity{
						Name: "abc",
					},
					IdentityStatus: rentalpb.IdentityStatus_VERIFIED,
				}
				err := s.Mongo.UpdateProfile(c, aid, rentalpb.IdentityStatus_PENDING, p)
				if err != nil {
					return nil, err
				}
				return p, nil
			},
			wantName:   "abc",
			wantStatus: rentalpb.IdentityStatus_VERIFIED,
		},
		{
			name: "clear",
			op: func() (*rentalpb.Profile, error) {
				return s.ClearProfile(c, &rentalpb.ClearProfileRequest{})
			},
			wantStatus: rentalpb.IdentityStatus_UNSUBMITTED,
		},
	}
	for _, cc := range cases {
		p, err := cc.op()
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s: want error; got none", cc.name)
			} else {
				continue
			}
		}
		if err != nil {
			t.Errorf("%s: operation failed: %v", cc.name, err)
		}
		gotName := ""
		if p.Identity != nil {
			gotName = p.Identity.Name
		}
		if gotName != cc.wantName {
			t.Errorf("%s: name field incorrect: want %q, got %q", cc.name, cc.wantName, gotName)
		}
		if p.IdentityStatus != cc.wantStatus {
			t.Errorf("%s: status field incorrect: want %s, got %s", cc.name, cc.wantStatus, p.IdentityStatus)
		}
	}
}

func TestProfilePhotoLifecycle(t *testing.T) {
	c := sharedauth.ContextWithAccountID(context.Background(), id.AccountID("account1"))
	s := newService(c, t)
	s.BlobClient = &blobClient{
		idForCreate: "blob1",
	}

	getPhotoOp := func() (string, error) {
		r, err := s.GetProfilePhoto(c, &rentalpb.GetProfilePhotoRequest{})
		if err != nil {
			return "", err
		}
		return r.UploadUrl, nil
	}

	cases := []struct {
		name        string
		op          func() (string, error)
		wantURL     string
		wantErrCode codes.Code
	}{
		{
			name:        "get_photo_before_upload", //未创建图片，获取图片是拿不到
			op:          getPhotoOp,
			wantErrCode: codes.NotFound,
		},
		{
			name: "create_photo",
			op: func() (string, error) {
				r, err := s.CreateProfilePhoto(c, &rentalpb.CreateProfilePhotoRequest{})
				if err != nil {
					return "", err
				}
				return r.UploadUrl, nil
			},
			wantURL: "upload_url for blob1",
		},
		{
			name: "complete_photo_upload",
			op: func() (string, error) {
				_, err := s.CompleteProfilePhoto(c, &rentalpb.CompleteProfilePhotoRequest{})
				return "", err
			},
		},
		{
			name:    "get_photo_url",
			op:      getPhotoOp,
			wantURL: "get_url for blob1",
		},
		{
			name: "clear_photo",
			op: func() (string, error) {
				_, err := s.ClearProfilePhoto(c, &rentalpb.ClearProfilePhotoRequest{})
				return "", err
			},
		},
		{
			name:        "get_photo_after_clear", //清除图片后，获取图片是拿不到
			op:          getPhotoOp,
			wantErrCode: codes.NotFound,
		},
	}

	for _, cc := range cases {
		got, err := cc.op()
		code := codes.OK
		if err != nil {
			if s, ok := status.FromError(err); ok {
				code = s.Code()
			} else {
				t.Errorf("%s: operation failed:%v", cc.name, err)
			}
		}
		if code != cc.wantErrCode {
			t.Errorf("%s: wrong error code:want:%d,got:%d", cc.name, cc.wantErrCode, code)
		}
		if got != cc.wantURL {
			t.Errorf("%s:wrong url:%q,got:%q", cc.name, cc.wantURL, got)
		}
	}
}

//公共方法--创建客户端
func newService(c context.Context, t *testing.T) *Service {
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot create new mongo client%v", err)
	}

	db := mc.Database("trip")
	mongotesting.SetupIndexes(c, db)
	logger, err := sharedserver.NewZapLogger()
	if err != nil {
		t.Fatalf("cannot create logger : %v", err)
	}

	return &Service{
		Mongo:  dao.NewMongo(db),
		Logger: logger,
	}
}

//自己写一个blob客户端
type blobClient struct {
	idForCreate string //控制blob的id
}

func (b *blobClient) CreateBlob(ctx context.Context, in *blobpb.CreateBlobRequest, opts ...grpc.CallOption) (*blobpb.CreateBlobResponse, error) {
	return &blobpb.CreateBlobResponse{
		Id:        b.idForCreate,
		UploadUrl: "upload_url for " + b.idForCreate,
	}, nil
}

func (b *blobClient) GetBlob(ctx context.Context, in *blobpb.GetBlobRequest, opts ...grpc.CallOption) (*blobpb.GetBlobResponse, error) {
	return &blobpb.GetBlobResponse{}, nil
}

func (b *blobClient) GetBlobURL(ctx context.Context, in *blobpb.GetBlobURLRequest, opts ...grpc.CallOption) (*blobpb.GetBlobURLResponse, error) {
	return &blobpb.GetBlobURLResponse{
		Url: "get_url for " + in.Id,
	}, nil
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
