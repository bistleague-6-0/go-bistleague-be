package team

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/model/entity"
	"bistleague-be/services/repository/profile"
	"bistleague-be/services/repository/storage"
	"bistleague-be/services/repository/team"
	"bistleague-be/services/utils"
	"bistleague-be/services/utils/randomizer"
	"bistleague-be/services/utils/storageutils"
	"context"
	"fmt"
	"log"
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
	strTeamToken := randomizer.RandStringBytes(6)

	teamID, err := u.repo.CreateTeam(ctx, team, strTeamToken)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	jwtToken, err := utils.GenerateJWTToken(u.cfg.Secret.JWTSecret, teamLeaderID, teamID)
	if err != nil {
		return nil, err
	}
	return &dto.CreateTeamResponseDTO{
		TeamRedeemToken: strTeamToken,
		JWTToken:        jwtToken,
	}, nil
}

func (u *Usecase) GetTeamInformation(ctx context.Context, teamID string, userID string) (*dto.GetTeamInfoResponseDTO, error) {
	// Fetch team information without submission data
	teamInfoResp, err := u.repo.GetTeamInformation(ctx, teamID)
	if err != nil {
		return nil, err
	}

	submission1Resp, err := u.repo.GetSubmission(ctx, teamID)
	if err != nil {
		return nil, err
	}

	result := dto.GetTeamInfoResponseDTO{}
	result.TeamID = teamID

	for _, team := range teamInfoResp {
		result.TeamName = team.TeamName
		result.TeamRedeemCode = team.RedeemCode
		result.IsActive = team.IsActive
		if team.TeamLeaderID == userID {
			result.Payment = team.PaymentFilename
			result.PaymentURL = team.PaymentURL
			result.PaymentStatusCode = team.PaymentStatus
			result.PaymentRejection = team.PaymentRejection
		}
		result.PaymentStatus = entity.VerificationStatusMap[team.PaymentStatus]
		if team.UserID == userID {
			result.StudentCard = team.StudentCard
			result.StudentCardStatusCode = team.StudentCardStatus
			result.StudentCardStatus = entity.VerificationStatusMap[team.StudentCardStatus]
			result.StudentCardURL = team.StudentCardURL
			result.StudentCardRejection = team.StudentCardRejection

			result.SelfPortrait = team.SelfPortrait
			result.SelfPortraitStatusCode = int8(team.SelfPortraitStatus)
			result.SelfPortraitStatus = entity.VerificationStatusMap[int8(team.SelfPortraitStatus)]
			result.SelfPortraitURL = team.SelfPortraitURL
			result.SelfPortraitRejection = team.SelfPortraitRejection

			result.Twibbon = team.Twibbon
			result.TwibbonStatusCode = team.TwibbonStatus
			result.TwibbonStatus = entity.VerificationStatusMap[team.TwibbonStatus]
			result.TwibbonURL = team.TwibbonURL
			result.TwibbonRejection = team.TwibbonRejection

			result.Enrollment = team.Enrollment
			result.EnrollmentStatusCode = team.EnrollmentStatus
			result.EnrollmentStatus = entity.VerificationStatusMap[team.EnrollmentStatus]
			result.EnrollmentURL = team.EnrollmentURL
			result.EnrollmentRejection = team.EnrollmentRejection
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

	result.Submission1Url = submission1Resp.Submission1Url
	result.Submission2Url = submission1Resp.Submission2Url

	return &result, nil
}

func (u *Usecase) RedeemTeamCode(ctx context.Context, req dto.RedeemTeamCodeRequestDTO, userID string) (string, error) {
	resp, err := u.repo.RedeemTeamCode(ctx, userID, req.RedeemCode)
	if err != nil {
		return "", err
	}
	jwtToken, err := utils.GenerateJWTToken(u.cfg.Secret.JWTSecret, userID, resp.TeamID)
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

func (u *Usecase) GetTeamSubmission(ctx context.Context, submissionType int, teamID string) (*dto.GetSubmissionResponseDTO, error) {
	resp, err := u.repo.GetSubmission(ctx, teamID)
	if err != nil {
		return nil, err
	}

	var dtoResp dto.GetSubmissionResponseDTO

	switch submissionType {
	case 1:
		dtoResp = dto.GetSubmissionResponseDTO{
			TeamID:               teamID,
			DocumentType:         "submission_1",
			SubmissionFilename:   resp.Submission1Filename.String,
			SubmissionUrl:        resp.Submission1Url.String,
			SubmissionLastUpdate: resp.Submission1LastUpdate.Time,
		}
	case 2:
		dtoResp = dto.GetSubmissionResponseDTO{
			TeamID:               teamID,
			DocumentType:         "submission_2",
			SubmissionFilename:   resp.Submission2Filename.String,
			SubmissionUrl:        resp.Submission2Url.String,
			SubmissionLastUpdate: resp.Submission2LastUpdate.Time,
		}
	default:
		return nil, fmt.Errorf("invalid submission type")
	}

	return &dtoResp, nil
}
