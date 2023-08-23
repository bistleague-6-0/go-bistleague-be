package team

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/model/entity"
	"bistleague-be/services/repository/profile"
	"bistleague-be/services/repository/storage"
	"bistleague-be/services/repository/team"
	"bistleague-be/services/utils"
	"bistleague-be/services/utils/storageutils"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"math/rand"
	"time"
)

type Usecase struct {
	cfg         *config.Config
	repo        *team.Repository
	profileRepo *profile.Repository
	storageRepo *storage.Repository
}

func New(cfg *config.Config, repo *team.Repository, storageRepo *storage.Repository, profileRepo *profile.Repository) *Usecase {
	return &Usecase{
		cfg:         cfg,
		repo:        repo,
		profileRepo: profileRepo,
		storageRepo: storageRepo,
	}
}

func (u *Usecase) CreateTeam(ctx context.Context, req dto.CreateTeamRequestDTO, teamLeaderID string) (*dto.CreateTeamResponseDTO, error) {
	team := entity.TeamEntity{
		TeamName:        req.TeamName,
		TeamLeaderID:    teamLeaderID,
		TeamMemberMails: req.MemberEmails,
	}

	// create redeem code
	teamToken := make([]byte, 4)
	rand.Read(teamToken)
	strTeamToken := hex.EncodeToString(teamToken)

	teamID, err := u.repo.CreateTeam(ctx, team, strTeamToken)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	claims := entity.CustomClaim{
		TeamID: teamID,
		UserID: teamLeaderID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "rest",
			Subject:   "",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 5)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	jwtToken, err := utils.CreateJWTToken(u.cfg.Secret.JWTSecret, claims)
	if err != nil {
		return nil, err
	}
	return &dto.CreateTeamResponseDTO{
		TeamRedeemToken: strTeamToken,
		JWTToken:        jwtToken,
	}, nil
}

func (u *Usecase) GetTeamInformation(ctx context.Context, teamID string, userID string) (*dto.GetTeamInfoResponseDTO, error) {
	resp, err := u.repo.GetTeamInformation(ctx, teamID)
	if err != nil {
		return nil, err
	}
	result := dto.GetTeamInfoResponseDTO{}
	result.TeamID = teamID
	for _, team := range resp {
		result.TeamName = team.TeamName
		result.TeamRedeemCode = team.RedeemCode
		result.IsActive = team.IsActive
		if team.TeamLeaderID == userID {
			result.Payment = team.PaymentFilename
			result.PaymentURL = team.PaymentURL
			result.PaymentStatusCode = team.PaymentStatus
		}
		result.PaymentStatus = entity.VerificationStatusMap[team.PaymentStatus]
		if team.UserID == userID {
			result.StudentCard = team.StudentCard
			result.StudentCardStatusCode = team.StudentCardStatus
			result.StudentCardStatus = entity.VerificationStatusMap[team.StudentCardStatus]
			result.StudentCardURL = team.StudentCardURL

			result.SelfPortrait = team.SelfPortrait
			result.SelfPortraitStatusCode = team.SelfPortraitStatus
			result.SelfPortraitStatus = entity.VerificationStatusMap[team.SelfPortraitStatus]
			result.SelfPortraitURL = team.SelfPortraitURL

			result.Twibbon = team.Twibbon
			result.TwibbonStatusCode = team.TwibbonStatus
			result.TwibbonStatus = entity.VerificationStatusMap[team.TwibbonStatus]
			result.TwibbonURL = team.TwibbonURL

			result.Enrollment = team.Enrollment
			result.EnrollmentStatusCode = team.EnrollmentStatus
			result.EnrollmentStatus = entity.VerificationStatusMap[team.EnrollmentStatus]
			result.EnrollmentURL = result.EnrollmentURL
		}
		result.Members = append(result.Members, dto.GetTeamMemberInfoResponseDTO{
			UserID:            team.UserID,
			Username:          team.Username,
			Fullname:          team.FullName,
			IsLeader:          team.TeamLeaderID == team.UserID,
			IsDocVerified:     team.IsDocVerified,
			IsProfileVerified: team.IsProfileVerified,
		})
	}
	return &result, nil
}

func (u *Usecase) RedeemTeamCode(ctx context.Context, req dto.RedeemTeamCodeRequestDTO, userID string) (string, error) {
	resp, err := u.repo.RedeemTeamCode(ctx, userID, req.RedeemCode)
	if err != nil {
		return "", err
	}
	claims := entity.CustomClaim{
		TeamID: resp.TeamID,
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "rest",
			Subject:   "",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 5)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	jwtToken, err := utils.CreateJWTToken(u.cfg.Secret.JWTSecret, claims)
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func (u *Usecase) InsertTeamDocument(ctx context.Context, req dto.InsertTeamDocumentRequestDTO, teamID string, userID string) (*dto.InputTeamDocumentResponseDTO, error) {
	randName := utils.GenerateRandomName()
	newName := fmt.Sprintf("%s.%s", req.Type, randName)
	b64file, err := storageutils.NewBase64FromString(req.Document, newName)
	if err != nil {
		return nil, err
	}
	fileurl, err := u.storageRepo.UploadDocument(ctx, b64file)
	if err != nil {
		return nil, err
	}
	if req.Type == "payment" {
		err = u.repo.InsertTeamDocument(ctx, req.DocumentName, fileurl, teamID)
	} else if req.Type == "submission_1" || req.Type == "submission_2" {
		err = u.repo.InsertTeamSubmission(ctx, req.DocumentName, fileurl, teamID, req.Type)
	} else {
		err = u.profileRepo.UpdateUserDocument(ctx, userID, req.DocumentName, fileurl, req.Type)
	}

	if err != nil {
		return nil, err
	}
	return &dto.InputTeamDocumentResponseDTO{
		DocumentType: req.Type,
		DocumentName: req.DocumentName,
		DocumentURL:  fileurl,
	}, nil
}
