package admin

import (
	"bistleague-be/model/config"
	"bistleague-be/model/dto"
	"bistleague-be/model/entity"
	adminRepo "bistleague-be/services/repository/admin"
	"bistleague-be/services/repository/challenge"
	emailRepo "bistleague-be/services/repository/email"
	"bistleague-be/services/repository/profile"
	"bistleague-be/services/repository/team"
	teamRepo "bistleague-be/services/repository/team"
	"bistleague-be/services/utils"
	"context"
	"strings"
	"golang.org/x/crypto/bcrypt"
)

const (
	emailTemplate = `
			<!DOCTYPE html>
			<html>
			<head>
				<style>
					body {
						font-family: Arial, sans-serif;
						background-color: #f7f7f7;
						margin: 0;
						padding: 20px;
					}
					.email-container {
						background-color: #ffffff;
						max-width: 600px;
						margin: 0 auto;
						padding: 20px;
						border-radius: 5px;
						box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
					}
					.details-header {
						font-weight: bold;
					}
				</style>
					</head>
					<body>
						<div class="email-container">
						<p>Dear {{.NamaLengkap}},</p>
						<p>We are delighted to inform you that your registration for the Bist League 6 has been successfully received and confirmed. Congratulations, and thank you for choosing to participate in our competition!</p>

						<p><span class="details-header">Here are the details of your registration:</span></p>
						<ul>
							<li><span>Full Name:</span> {{.NamaLengkap}}</li>
							<li><span>Email Address:</span> {{.Email}}</li>
							<li><span>Team Name:</span> {{.NamaTim}}</li>
							<li><span>Contact Number:</span> {{.NomorTelepon}}</li>
						</ul>

						<p>Please keep this email for your records. In case you have any questions or need to make any changes to your registration details, please do not hesitate to contact us at <a href="mailto:bistleague@std.stei.itb.ac.id">bistleague@std.stei.itb.ac.id</a> or +62 81290908333.</p>

						<p  class="details-header">Important Dates:</p>
						<ul>
							<li><span>Preliminary Competition Date:</span> October 22, 2023</li>
							<li><span>Location:</span> Online</li>
						</ul>

						<p>Once again, congratulations on your successful registration. We look forward to seeing you and your team at the competition. Best of luck with your preparations, and may the best team win!</p>

						<p>Sincerely,</p>
						<p>Bist league 6</p>
						</div>
					</body>
					</html>
			`
)

type Usecase struct {
	cfg           *config.Config
	repo          *adminRepo.Repository
	profileRepo   *profile.Repository
	teamRepo      *teamRepo.Repository
	challengeRepo *challenge.Repository
	emailRepo     *emailRepo.Repository
}

func New(cfg *config.Config, repo *adminRepo.Repository, profileRepo *profile.Repository, teamRepo *team.Repository, challengeRepo *challenge.Repository, emailRepo *emailRepo.Repository) *Usecase {
	return &Usecase{
		cfg:           cfg,
		repo:          repo,
		profileRepo:   profileRepo,
		teamRepo:      teamRepo,
		challengeRepo: challengeRepo,
		emailRepo:     emailRepo,
	}
}

func (u *Usecase) InsertNewAdmin(ctx context.Context, req dto.RegisterAdminRequestDTO) (*dto.AuthAdminResponseDTO, error) {
	newpw, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	admin := entity.AdminEntity{
		Password: string(newpw),
		FullName: req.FullName,
		Username: req.Username,
	}
	resp, err := u.repo.RegisterNewAdmin(ctx, admin)
	if err != nil {
		return nil, err
	}
	token, err := utils.GenerateAdminJWTToken(u.cfg.Secret.AdminJWT, resp.UID)
	if err != nil {
		return nil, err
	}
	return &dto.AuthAdminResponseDTO{
		Admin: dto.AuthAdminInfoResponse{
			AdminID:  resp.UID,
			Username: resp.Username,
		},
		Token: token,
	}, nil
}

func (u *Usecase) SignInAdmin(ctx context.Context, req dto.SignInAdminRequestDTO) (*dto.AuthAdminResponseDTO, error) {
	admin, err := u.repo.LoginAdmin(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	token, err := utils.GenerateAdminJWTToken(u.cfg.Secret.AdminJWT, admin.UID)
	if err != nil {
		return nil, err
	}
	return &dto.AuthAdminResponseDTO{
		Admin: dto.AuthAdminInfoResponse{
			AdminID:  admin.UID,
			Username: admin.Username,
		},
		Token: token,
	}, nil
}

func (u *Usecase) GetTeamPayment(ctx context.Context, page int, pageSize int) (*dto.PaginationDTOWrapper, error) {
	resp, err := u.teamRepo.GetPayments(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}
	data := []dto.GetTeamPaymentResponseDTO{}
	for _, payment := range resp {
		member_email_str := strings.Trim(payment.TeamMemberMails, "{}")
		member_email := strings.Split(member_email_str, ",")
		data = append(data, dto.GetTeamPaymentResponseDTO{
			TeamID:          payment.TeamID,
			TeamName:        payment.TeamName,
			TeamMemberMails: member_email,
			PaymentFilename: payment.PaymentFilename,
			PaymentURL:      payment.PaymentURL,
			PaymentStatus:   entity.VerificationStatusMap[payment.PaymentStatus],
			Code:            payment.Code,
		})
	}
	totalTeam, err := u.teamRepo.GetTeamCount(ctx)
	if err != nil {
		return nil, err
	}
	var dtoResp dto.PaginationDTOWrapper

	totalPage := (totalTeam + pageSize - 1) / pageSize

	dtoResp = dto.PaginationDTOWrapper{
		PageSize:  pageSize,
		Page:      page,
		TotalPage: totalPage,
		Data:      data,
	}

	return &dtoResp, nil
}

func (u *Usecase) GetUserList(ctx context.Context, page int, pageSize int) (*dto.PaginationDTOWrapper, error) {
	resp, err := u.profileRepo.GetUserList(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}
	data := []dto.UserDocsResponseDTO{}
	for _, user := range resp {
		data = append(data, dto.UserDocsResponseDTO{
			UID:                  user.UID,
			TeamName:             user.TeamName.String,
			FullName:             user.FullName,
			StudentCardFilename:  user.StudentCardFilename.String,
			StudentCardURL:       user.StudentCardURL.String,
			StudentCardStatus:    entity.VerificationStatusMap[user.StudentCardStatus],
			SelfPortraitFilename: user.SelfPortraitFilename.String,
			SelfPortraitURL:      user.SelfPortraitURL.String,
			SelfPortraitStatus:   entity.VerificationStatusMap[user.SelfPortraitStatus],
			TwibbonFilename:      user.TwibbonFilename.String,
			TwibbonURL:           user.TwibbonURL.String,
			TwibbonStatus:        entity.VerificationStatusMap[user.TwibbonStatus],
			EnrollmentFilename:   user.EnrollmentFilename.String,
			EnrollmentURL:        user.EnrollmentURL.String,
			EnrollmentStatus:     entity.VerificationStatusMap[user.EnrollmentStatus],
			IsProfileVerified:    user.IsProfileVerified,
		})
	}
	totalUser, err := u.profileRepo.GetUserCount(ctx)
	if err != nil {
		return nil, err
	}

	var dtoResp dto.PaginationDTOWrapper

	totalPage := (totalUser + pageSize - 1) / pageSize

	dtoResp = dto.PaginationDTOWrapper{
		PageSize:  pageSize,
		Page:      page,
		TotalPage: totalPage,
		Data:      data,
	}

	return &dtoResp, nil
}

func (u *Usecase) UpdateTeamPaymentStatus(ctx context.Context, teamID string, status int, rejection string) error {
	err := u.teamRepo.UpdatePaymentStatus(ctx, teamID, status, rejection)
	if err != nil {
		return err
	}

	teamVerifications, err := u.teamRepo.GetTeamVerification(ctx, teamID)
	if err != nil {
		return err
	}

	verified := true
	var data struct {
		NamaLengkap  string
		Email        string
		NamaTim      string
		NomorTelepon string
	}
	if len(teamVerifications) < 2 {
		verified = false
	}
	for _, teamVerification := range teamVerifications {
		if teamVerification.PaymentStatus != 2 || teamVerification.StudentCardStatus != 2 || teamVerification.SelfPortraitStatus != 2 || teamVerification.TwibbonStatus != 2 {
			verified = false
		}
		if teamVerification.TeamLeaderID == teamVerification.UserID {
			data.NamaLengkap = teamVerification.FullName
			data.Email = teamVerification.Email
			data.NamaTim = teamVerification.TeamName
			data.NomorTelepon = teamVerification.Phone.String
		}
	}

	if verified {
		return u.emailRepo.SendEmailHTML([]string{data.Email},
			"Confirmation of Registration for Business Case Competition by Bist League 6",
			emailTemplate, data)
	}

	return err
}

func (u *Usecase) UpdateUserDocumentStatus(ctx context.Context, userID string, doctype string, status int, rejection string) error {
	teamID, err := u.profileRepo.UpdateUserDocumentStatus(ctx, userID, doctype, status, rejection)

	if err != nil {
		return err
	}

	teamVerifications, err := u.teamRepo.GetTeamVerification(ctx, teamID)
	if err != nil {
		return err
	}
	verified := true
	var data struct {
		NamaLengkap  string
		Email        string
		NamaTim      string
		NomorTelepon string
	}
	if len(teamVerifications) < 2 {
		verified = false
	}
	for _, teamVerification := range teamVerifications {
		if teamVerification.PaymentStatus != 2 || teamVerification.StudentCardStatus != 2 || teamVerification.SelfPortraitStatus != 2 || teamVerification.TwibbonStatus != 2 {
			verified = false
		}
		if teamVerification.TeamLeaderID == teamVerification.UserID {
			data.NamaLengkap = teamVerification.FullName
			data.Email = teamVerification.Email
			data.NamaTim = teamVerification.TeamName
			data.NomorTelepon = teamVerification.Phone.String
		}
	}

	if verified {
		return u.emailRepo.SendEmailHTML([]string{data.Email},
			"Confirmation of Registration for Business Case Competition by Bist League 6",
			emailTemplate, data)
	}

	return err
}

func (u *Usecase) GetMiniChallengesUsecase(ctx context.Context, page uint64, limit uint64) ([]dto.AdminGetMiniChallengeResponseDTO, error) {
	resp, err := u.challengeRepo.GetUserChallenges(ctx, page, limit)
	if err != nil {
		return nil, err
	}
	result := []dto.AdminGetMiniChallengeResponseDTO{}
	for _, chal := range resp {
		result = append(result, dto.AdminGetMiniChallengeResponseDTO{
			UID:      chal.UID,
			FullName: chal.FullName,
			Username: chal.Username,
			Email:    chal.Email,
			InsertChallengeRequestDTO: dto.InsertChallengeRequestDTO{
				IgUsername:       chal.IgUsername,
				IgContentURl:     chal.IgContentURl,
				TiktokUsername:   chal.TiktokUsername,
				TiktokContentURl: chal.TiktokContentURl,
			},
		})
	}
	return result, err
}

func (u *Usecase) GetMiniChallengeByUIDUsecase(ctx context.Context, uid string) (*dto.AdminGetMiniChallengeResponseDTO, error) {
	resp, err := u.challengeRepo.GetUserChallengeWithUserDetail(ctx, uid)
	if err != nil {
		return nil, err
	}
	return &dto.AdminGetMiniChallengeResponseDTO{
		UID:      resp.UID,
		FullName: resp.FullName,
		Username: resp.FullName,
		Email:    resp.Email,
		InsertChallengeRequestDTO: dto.InsertChallengeRequestDTO{
			IgUsername:       resp.IgUsername,
			IgContentURl:     resp.IgContentURl,
			TiktokUsername:   resp.TiktokUsername,
			TiktokContentURl: resp.TiktokContentURl,
		},
	}, nil
}

func (u *Usecase) GetAllSubmissionUsecase(ctx context.Context, page int, pageSize int) (*dto.PaginationDTOWrapper, error) {
    resp, err := u.teamRepo.GetAllSubmission(ctx, page, pageSize)
    if err != nil {
        return nil, err
    }

	result := []dto.GetAllSubmissionResponseDTO{}

    for _, submission := range resp {

        result = append(result, dto.GetAllSubmissionResponseDTO{
            TeamID:               submission.TeamID,
            Submission1Filename:   submission.Submission1Filename,
            Submission1Url:        submission.Submission1Url,
            Submission1LastUpdate: submission.Submission1LastUpdate,
			Submission2Filename:   submission.Submission2Filename,
			Submission2Url:        submission.Submission2Url,
			Submission2LastUpdate: submission.Submission2LastUpdate,
        })
    }

	totalTeam, err := u.teamRepo.GetTeamCount(ctx)
	if err != nil {
		return nil, err
	}

	var dtoResp dto.PaginationDTOWrapper

	totalPage := (totalTeam + pageSize - 1) / pageSize

	dtoResp = dto.PaginationDTOWrapper{
		PageSize:  pageSize,
		Page:      page,
		TotalPage: totalPage,
		Data:      result,
	}


    return &dtoResp, err
}
