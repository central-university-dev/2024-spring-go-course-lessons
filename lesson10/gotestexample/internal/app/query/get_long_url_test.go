package query_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"gotestexample/internal/app/query"
	"gotestexample/internal/common"
)

type repositoryStub struct {
	mock.Mock
}

func (s *repositoryStub) Get(_ context.Context, id string) (string, error) {
	args := s.Called(id)

	return args.Get(0).(string), args.Error(1)
}

func TestGetLongUrlQueryExecute(t *testing.T) {
	t.Parallel()

	ctx := context.TODO()

	//nolint:gocritic // it's common pattern to use assert
	assert := assert.New(t)

	type in struct {
		id string
	}

	type out struct {
		err     error
		longURL string
	}

	tests := []struct {
		setup  func(*repositoryStub, *in)
		assert func(*out)
		name   string
		in     in
	}{
		{
			name: "happy path",
			in:   in{id: "existing-key"},
			setup: func(repo *repositoryStub, in *in) {
				repo.On("Get", in.id).Return("https://google.com", nil)
			},
			assert: func(out *out) {
				assert.NoError(out.err)
				assert.Equal(out.longURL, "https://google.com")
			},
		},
		{
			name: "not found",
			in:   in{id: "unknown-id"},
			setup: func(repo *repositoryStub, in *in) {
				repo.On("Get", in.id).Return("", common.ErrNotFound)
			},
			assert: func(out *out) {
				assert.ErrorIs(out.err, common.ErrNotFound)
				assert.Empty(out.longURL)
			},
		},
		{
			name: "error while get from repo",
			in:   in{id: "existing-key"},
			setup: func(repo *repositoryStub, in *in) {
				repo.On("Get", in.id).Return("", errors.New("repository down"))
			},
			assert: func(out *out) {
				assert.ErrorIs(out.err, common.ErrTechnical)
				assert.Empty(out.longURL)
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Given
			repo := &repositoryStub{}
			tt.setup(repo, &tt.in)

			sut := query.NewGetLongURLQuery(repo)

			// When
			res, err := sut.Execute(ctx, tt.in.id)

			// Then
			tt.assert(&out{longURL: res, err: err})
		})
	}
}
