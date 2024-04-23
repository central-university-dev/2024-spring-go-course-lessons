//go:generate mockgen -destination=../../../mocks/create_short_url_cmd.go -package=mocks -mock_names=Repository=MockCmdRepo,IDProvider=StubIDProvider github.com/kulinsky/gotestexample/internal/app/command Repository,IDProvider

package command

import (
	"context"
	"fmt"
	"net/url"

	"gotestexample/internal/common"
)

var (
	ErrInvalidURL = fmt.Errorf("invalid url: %w", common.ErrValidation)
	ErrRepository = fmt.Errorf("command repo error: %w", common.ErrTechnical)
)

type IDProvider interface {
	Provide() string
}

type Repository interface {
	Save(ctx context.Context, id string, longURL string) error
}

type CreateShortURLCmd struct {
	idProvider IDProvider
	repo       Repository
}

func NewCreateShortURLCmd(idp IDProvider, repo Repository) *CreateShortURLCmd {
	return &CreateShortURLCmd{
		idProvider: idp,
		repo:       repo,
	}
}

func (cmd *CreateShortURLCmd) Execute(ctx context.Context, rawURL string) (string, error) {
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return "", ErrInvalidURL
	}

	id := cmd.idProvider.Provide()

	if err := cmd.repo.Save(ctx, id, parsedURL.String()); err != nil {
		return "", fmt.Errorf("%w: %s", ErrRepository, err.Error())
	}

	return id, nil
}
