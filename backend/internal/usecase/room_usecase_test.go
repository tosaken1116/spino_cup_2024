package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/model"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/repository"
	mockrepo "github.com/tosaken1116/spino_cup_2024/backend/internal/mock/domain/repository"
	"go.uber.org/mock/gomock"
)

func TestNewRoomDTOFromModel(t *testing.T) {
	t.Parallel()

	id, err := model.NewRoomID()
	assert.NoError(t, err)

	type args struct {
		m *model.Room
	}
	tests := []struct {
		name string
		args args
		want *RoomDTO
	}{
		{
			name: "success",
			args: args{
				m: &model.Room{
					ID:          id,
					Name:        "room1",
					Description: "room1 description",
				},
			},
			want: &RoomDTO{
				ID:          id.String(),
				Name:        "room1",
				Description: "room1 description",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, NewRoomDTOFromModel(tt.args.m))
		})
	}
}

func TestNewRoomUsecase(t *testing.T) {
	type args struct {
		repo repository.RoomRepository
	}
	tests := []struct {
		name string
		args args
		want RoomUsecase
	}{
		{
			name: "success",
			args: args{
				repo: nil,
			},
			want: &roomUsecase{
				repo: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewRoomUsecase(tt.args.repo))
		})
	}
}

func Test_roomUsecase_CreateRoom(t *testing.T) {
	type fields struct {
	}
	type args struct {
		ctx context.Context
		dto *RoomDTO
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *RoomDTO
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &roomUsecase{}
			got, err := r.CreateRoom(tt.args.ctx, tt.args.dto)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_roomUsecase_GetRoom(t *testing.T) {
	t.Parallel()

	rawid := ulid.MustNew(uint64(time.Now().UnixMilli()), nil).String()
	type args struct {
		ctx   context.Context
		rawid string
	}
	tests := []struct {
		name      string
		args      args
		fn        func(repo *mockrepo.MockRoomRepository)
		want      *RoomDTO
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			args: args{
				ctx:   context.TODO(),
				rawid: rawid,
			},
			fn: func(repo *mockrepo.MockRoomRepository) {
				id, _ := model.ParseRoomID(rawid)
				repo.EXPECT().GetRoom(gomock.Any(), id).Return(&model.Room{
					ID:          id,
					Name:        "room1",
					Description: "room1 description",
				}, nil)
			},
			want: &RoomDTO{
				ID:          rawid,
				Name:        "room1",
				Description: "room1 description",
			},
			assertion: assert.NoError,
		},
		{
			name: "failed (invalid id)",
			args: args{
				ctx:   context.TODO(),
				rawid: "invalid",
			},
			fn:   func(repo *mockrepo.MockRoomRepository) {},
			want: nil,
			assertion: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.ErrorIs(t, err, model.ErrRoomIDInvalid)
			},
		},
		{
			name: "failed (repo error)",
			args: args{
				ctx:   context.TODO(),
				rawid: rawid,
			},
			fn: func(repo *mockrepo.MockRoomRepository) {
				id, _ := model.ParseRoomID(rawid)
				repo.EXPECT().GetRoom(gomock.Any(), id).Return(nil, model.ErrRoomNotFound)
			},
			want: nil,
			assertion: func(t assert.TestingT, err error, _ ...interface{}) bool {
				return assert.ErrorIs(t, err, model.ErrRoomNotFound)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			repo := mockrepo.NewMockRoomRepository(mockCtrl)
			tt.fn(repo)

			r := &roomUsecase{
				repo: repo,
			}
			got, err := r.GetRoom(tt.args.ctx, tt.args.rawid)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_roomUsecase_ListRoom(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name      string
		args      args
		want      []*RoomDTO
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &roomUsecase{}
			got, err := r.ListRoom(tt.args.ctx)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_roomUsecase_UpdateRoom(t *testing.T) {
	type fields struct {
	}
	type args struct {
		ctx context.Context
		dto *RoomDTO
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *RoomDTO
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &roomUsecase{}
			got, err := r.UpdateRoom(tt.args.ctx, tt.args.dto)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
