package team

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/model/entity"
	"bistleague-be/services/repository/document"
	"bistleague-be/services/repository/profile"
	"bistleague-be/services/repository/team"
	"bistleague-be/services/utils"
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
	docsRepo    *document.Repository
	profileRepo *profile.Repository
}

func New(cfg *config.Config, repo *team.Repository, docsRepo *document.Repository, profileRepo *profile.Repository) *Usecase {
	return &Usecase{
		cfg:         cfg,
		repo:        repo,
		profileRepo: profileRepo,
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
		result.IsActive = team.IsActive
		result.VerificationStatusCode = team.VerificationStatus
		result.VerificationStatus = entity.VerificationStatusMap[team.VerificationStatus]
		if team.UserID == userID {
			result.StudentCard = team.StudentCard
			result.StudentCardStatusCode = team.StudentCardStatus
			result.StudentCardStatus = entity.VerificationStatusMap[team.StudentCardStatus]

			result.SelfPortrait = team.SelfPortrait
			result.SelfPortraitStatusCode = team.SelfPortraitStatus
			result.SelfPortraitStatus = entity.VerificationStatusMap[team.SelfPortraitStatus]

			result.Twibbon = team.Twibbon
			result.TwibbonStatusCode = team.TwibbonStatus
			result.TwibbonStatus = entity.VerificationStatusMap[team.TwibbonStatus]

			result.Enrollment = team.Enrollment
			result.EnrollmentStatusCode = team.EnrollmentStatus
			result.EnrollmentStatus = entity.VerificationStatusMap[team.EnrollmentStatus]
		}
		result.Members = append(result.Members, dto.GetTeamMemberInfoResponseDTO{
			UserID:   team.UserID,
			Username: team.Username,
			Fullname: team.FullName,
			IsLeader: team.TeamLeaderID == team.UserID,
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

func (u *Usecase) InsertTeamDocument(ctx context.Context, req dto.InsertTeamDocumentRequestDTO, teamID string, userID string) (string, error) {
	currentDate := time.Now()
	formattedDate := currentDate.Format("2006-01-02")
	filename := fmt.Sprintf("%s.%s.%s", req.Type, teamID, formattedDate)

	// upload file
	err := u.docsRepo.UploadTeamDocument(ctx, req.Type, filename, teamID, req.Document)
	if err != nil {
		return "", err
	}
	//update db
	if req.Document == "payment" {
		err = u.repo.InsertTeamDocument(ctx, filename, teamID)
	} else {
		err = u.profileRepo.UpdateUserDocument(ctx, userID, filename, req.Type)
	}
	return filename, err
}
