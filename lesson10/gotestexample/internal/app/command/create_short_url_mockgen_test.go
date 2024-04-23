package command_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"gotestexample/internal/app/command"
	"gotestexample/internal/common"
	"gotestexample/mocks"
)

func TestCreateShortURLCommandTableWithMockgen(t *testing.T) {
	t.Parallel()

	//nolint:gocritic // it's common pattern to use assert
	assert := assert.New(t)

	type in struct {
		err     error
		longURL string
		id      string
	}

	type out struct {
		err error
		id  string
	}

	tests := []struct {
		setup  func(context.Context, *in) *command.CreateShortURLCmd
		assert func(*out)
		name   string
		in     in
	}{
		{
			name: "successfully save",
			in:   in{longURL: "https://google.com", id: "1"},
			setup: func(ctx context.Context, in *in) *command.CreateShortURLCmd {
				ctrl := gomock.NewController(t)
				idp := mocks.NewStubIDProvider(ctrl)
				repo := mocks.NewMockCmdRepo(ctrl)
				idp.EXPECT().Provide().Return(in.id).Times(1)
				repo.EXPECT().Save(ctx, in.id, in.longURL).Return(in.err).Times(1)

				return command.NewCreateShortURLCmd(idp, repo)
			},
			assert: func(out *out) {
				assert.NoError(out.err)
				assert.Equal("1", out.id)
			},
		},
		{
			name: "invalid url",
			in:   in{longURL: "this is invalid url"},
			setup: func(ctx context.Context, in *in) *command.CreateShortURLCmd {
				ctrl := gomock.NewController(t)
				idp := mocks.NewStubIDProvider(ctrl)
				repo := mocks.NewMockCmdRepo(ctrl)
				idp.EXPECT().Provide().Times(0)
				repo.EXPECT().Save(ctx, in.id, in.longURL).Times(0)

				return command.NewCreateShortURLCmd(idp, repo)
			},
			assert: func(out *out) {
				assert.ErrorIs(out.err, common.ErrValidation)
				assert.Empty(out.id)
			},
		},
		{
			name: "repo error on save",
			in:   in{longURL: "https://google.com", id: "1", err: errors.New("unexpected repository error")},
			setup: func(ctx context.Context, in *in) *command.CreateShortURLCmd {
				ctrl := gomock.NewController(t)
				idp := mocks.NewStubIDProvider(ctrl)
				repo := mocks.NewMockCmdRepo(ctrl)
				idp.EXPECT().Provide().Return(in.id).Times(1)
				repo.EXPECT().Save(ctx, in.id, in.longURL).Return(in.err).Times(1)

				return command.NewCreateShortURLCmd(idp, repo)
			},
			assert: func(out *out) {
				assert.ErrorIs(out.err, common.ErrTechnical)
				assert.Empty(out.id)
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Given
			ctx := context.TODO()
			sut := tt.setup(ctx, &tt.in)

			// When
			res, err := sut.Execute(ctx, tt.in.longURL)

			// Then
			tt.assert(&out{id: res, err: err})
		})
	}
}
