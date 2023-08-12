package storage

import (
	"bistleague-be/model/config"
	"bistleague-be/services/repository/storage"
	"bistleague-be/services/utils/storageutils"
	"github.com/gofiber/fiber/v2"
)

type Usecase struct {
	cfg         *config.Config
	StorageRepo *storage.Repository
}

func New(cfg *config.Config, storageRepo *storage.Repository) *Usecase {
	return &Usecase{
		cfg:         cfg,
		StorageRepo: storageRepo,
	}
}

func (u *Usecase) UploadFile(ctx context.Context, req dto.UploadStorageRequest) (*dto.UploadStorageResponse, error) {
	// cek empty strings
	if req.FileName == "" || req.Base64Data == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "file_name and base64_data are required",
		})
		return
	}

	// decode data
	decodedData, ext, err := storageutils.DecodeBase64WithFormat(req.Base64Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// generate random name
	encryptedName := utils.GenerateRandomName() + ext

	

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
