package profile

import (
	"context"
	blobpb "coolcar/blob/api/gen/v1/blob"
	rentalpb "coolcar/rental/api/gen/v1/rental"
	"coolcar/rental/profile/dao"
	"time"

	sharedauth "coolcar/shared/auth"
	"coolcar/shared/id"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	BlobClient        blobpb.BlobServiceClient
	PhotoGetExpire    time.Duration
	PhotoUploadExpire time.Duration
	Mongo             *dao.Mongo
	Logger            *zap.Logger
}

//获取个人信息
func (s *Service) GetProfile(c context.Context, req *rentalpb.GetProfileRequest) (*rentalpb.Profile, error) {
	aid, err := sharedauth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	p, err := s.Mongo.GetProfile(c, aid)
	if err != nil {
		code := s.logAndConvertProfileErr(err)
		if code == codes.NotFound {
			return &rentalpb.Profile{}, nil
		}
		return nil, status.Error(code, "")
	}

	return p.Prfile, nil
}

//提交个人信息
func (s *Service) SubmitProfile(c context.Context, i *rentalpb.Identity) (*rentalpb.Profile, error) {
	aid, err := sharedauth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	p := &rentalpb.Profile{
		Identity:       i,
		IdentityStatus: rentalpb.IdentityStatus_PENDING,
	}

	err = s.Mongo.UpdateProfile(c, aid, rentalpb.IdentityStatus_UNSUBMITTED, p)
	if err != nil {
		s.Logger.Error("SubmitProfile cannot update profile", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	//模拟后台审核
	go func() {
		time.Sleep(3 * time.Second)
		err := s.Mongo.UpdateProfile(context.Background(), aid, rentalpb.IdentityStatus_PENDING, &rentalpb.Profile{
			Identity:       i,
			IdentityStatus: rentalpb.IdentityStatus_VERIFIED,
		})

		if err != nil {
			s.Logger.Error("cannot verify identity", zap.Error(err))
		}
	}()
	return p, nil
}

//清空个人信息
func (s *Service) ClearProfile(c context.Context, req *rentalpb.ClearProfileRequest) (*rentalpb.Profile, error) {
	aid, err := sharedauth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	p := &rentalpb.Profile{}

	err = s.Mongo.UpdateProfile(c, aid, rentalpb.IdentityStatus_VERIFIED, p)
	if err != nil {
		s.Logger.Error("ClearProfile cannot update profile", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return p, nil
}

//获取驾照信息
func (s *Service) GetProfilePhoto(c context.Context, req *rentalpb.GetProfilePhotoRequest) (*rentalpb.GetProfilePhotoResponse, error) {
	aid, err := sharedauth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	pr, err := s.Mongo.GetProfile(c, aid)
	if err != nil {
		return nil, status.Error(s.logAndConvertProfileErr(err), "")
	}

	if pr.PhotoBlobID == "" {
		return nil, status.Error(codes.NotFound, "")
	}

	br, err := s.BlobClient.GetBlobURL(c, &blobpb.GetBlobURLRequest{
		Id:         pr.PhotoBlobID,
		TimeoutSec: int32(s.PhotoGetExpire.Seconds()),
	})
	if err != nil {
		s.Logger.Error("cannot get blob", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	return &rentalpb.GetProfilePhotoResponse{
		UploadUrl: br.Url,
	}, nil
}

//创建驾照信息
func (s *Service) CreateProfilePhoto(c context.Context, req *rentalpb.CreateProfilePhotoRequest) (*rentalpb.CreateProfilePhotoResponse, error) {
	aid, err := sharedauth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}

	br, err := s.BlobClient.CreateBlob(c, &blobpb.CreateBlobRequest{
		AccountId:           aid.String(),
		UploadUrlTimeoutSec: int32(s.PhotoUploadExpire.Seconds()),
	})
	if err != nil {
		s.Logger.Error("cannot create blob", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}

	err = s.Mongo.UpdateProfilePhoto(c, aid, id.BlobID(br.Id))
	if err != nil {
		s.Logger.Error("cannot update profile photo", zap.Error(err))
		return nil, status.Error(codes.Aborted, "")
	}

	return &rentalpb.CreateProfilePhotoResponse{
		UploadUrl: br.UploadUrl,
	}, nil
}

func (s *Service) CompleteProfilePhoto(c context.Context, req *rentalpb.CompleteProfilePhotoRequest) (*rentalpb.Identity, error) {
	aid, err := sharedauth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	pr, err := s.Mongo.GetProfile(c, aid)
	if err != nil {
		return nil, status.Error(s.logAndConvertProfileErr(err), "")
	}

	if pr.PhotoBlobID == "" {
		return nil, status.Error(codes.NotFound, "")
	}

	br, err := s.BlobClient.GetBlob(c, &blobpb.GetBlobRequest{
		Id: pr.PhotoBlobID,
	})
	if err != nil {
		s.Logger.Error("cannot get blob", zap.Error(err))
	}

	//制造假数据
	s.Logger.Info("got profile photo", zap.Int("size", len(br.Data)))

	return &rentalpb.Identity{
		LicNumber:       "322152452",
		Name:            "李四",
		Gender:          rentalpb.Gender_FEMAEL,
		BirthDateMillis: 631152000000,
	}, nil
}

//清除图片信息
func (s *Service) ClearProfilePhoto(c context.Context, req *rentalpb.ClearProfilePhotoRequest) (*rentalpb.ClearProfilePhotoResponse, error) {
	aid, err := sharedauth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	err = s.Mongo.UpdateProfilePhoto(c, aid, id.BlobID(""))
	if err != nil {
		return nil, status.Error(codes.Internal, "")
	}
	return &rentalpb.ClearProfilePhotoResponse{}, nil
}

func (s *Service) logAndConvertProfileErr(err error) codes.Code {
	if err == mongo.ErrNoDocuments {
		return codes.NotFound
	}
	s.Logger.Error("cannot get profile", zap.Error(err))
	return codes.Internal
}
