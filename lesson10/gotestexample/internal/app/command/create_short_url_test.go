package command_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gotestexample/internal/app/command"
	"gotestexample/internal/common"
)

type repositoryMock struct {
	mock.Mock
}

func (m *repositoryMock) Save(_ context.Context, id, fullURL string) error {
	args := m.Called(id, fullURL)

	return args.Error(0)
}

type idProviderStub struct {
	ID string
}

func (d idProviderStub) Provide() string {
	return d.ID
}

func TestCreateShortURLCommandTable(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()

	//nolint:gocritic // it's common pattern to use assert
	assert := assert.New(t)

	type in struct {
		fullURL string
	}

	type out struct {
		err error
		id  string
	}

	tests := []struct {
		setup  func(*repositoryMock, *idProviderStub, *in)
		assert func(*repositoryMock, *out)
		name   string
		in     in
	}{
		{
			name: "successfully save",
			in:   in{fullURL: "https://google.com"},
			setup: func(repo *repositoryMock, idp *idProviderStub, in *in) {
				idp.ID = "1"
				repo.On("Save", "1", in.fullURL).Return(nil)
			},
			assert: func(repo *repositoryMock, out *out) {
				assert.NoError(out.err)
				assert.Equal("1", out.id)
				repo.AssertExpectations(t)
			},
		},
		{
			name: "invalid url",
			in:   in{fullURL: "this is invalid url"},
			setup: func(_ *repositoryMock, idp *idProviderStub, _ *in) {
				idp.ID = "1"
			},
			assert: func(_ *repositoryMock, out *out) {
				assert.ErrorIs(out.err, common.ErrValidation)
				assert.Empty(out.id)
			},
		},
		{
			name: "repo error on save",
			in:   in{fullURL: "https://google.com"},
			setup: func(repo *repositoryMock, idp *idProviderStub, in *in) {
				idp.ID = "1"
				repo.On("Save", "1", in.fullURL).Return(errors.New("unexpected repository error"))
			},
			assert: func(repo *repositoryMock, out *out) {
				assert.ErrorIs(out.err, common.ErrTechnical)
				assert.Empty(out.id)
				repo.AssertExpectations(t)
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Given
			idp := &idProviderStub{}
			repo := &repositoryMock{}
			tt.setup(repo, idp, &tt.in)

			sut := command.NewCreateShortURLCmd(idp, repo)

			// When
			res, err := sut.Execute(ctx, tt.in.fullURL)

			// Then
			tt.assert(repo, &out{id: res, err: err})
		})
	}
}
