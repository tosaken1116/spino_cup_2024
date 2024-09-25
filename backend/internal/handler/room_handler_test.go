// internal/handler/room_handler_test.go
package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/domain/model"
	mu "github.com/tosaken1116/spino_cup_2024/backend/internal/mock/usecase"
	"github.com/tosaken1116/spino_cup_2024/backend/internal/usecase"
	"go.uber.org/mock/gomock"
)

type RoomRequest struct {
	Name        string
	Description string
}

func TestRoomHandler_CreateRoom(t *testing.T) {
	tests := []struct {
		name       string
		arg        any
		fn         func(mockUsecase *mu.MockRoomUsecase)
		wantStatus int
		wantBody   string
	}{
		{
			name: "成功: ルームが作成される",
			arg: &RoomRequest{
				Name:        "Test Room",
				Description: "This is a test room",
			},
			fn: func(mockUsecase *mu.MockRoomUsecase) {
				mockUsecase.
					EXPECT().
					CreateRoom(gomock.Any(), gomock.Any()).
					Return(&usecase.RoomDTO{
						ID:          "01AN4Z07BY79KA1307SR9X4MV3",
						Name:        "Test Room",
						Description: "This is a test room",
					}, nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   "{\"room\":{\"id\":\"01AN4Z07BY79KA1307SR9X4MV3\",\"name\":\"Test Room\",\"description\":\"This is a test room\"}}\n",
		},
		{
			name: "失敗: Bindエラー",
			arg: &struct {
				Name        int
				Description string
			}{
				Name:        1,
				Description: "This is a test room",
			},
			fn: func(mockUsecase *mu.MockRoomUsecase) {
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "{\"message\":\"Unmarshal type error: expected=string, got=number, field=name, offset=9\"}\n",
		},
		{
			name: "失敗: ルーム名が空",
			arg: &RoomRequest{
				Name:        "",
				Description: "This is a test room",
			},
			fn: func(mockUsecase *mu.MockRoomUsecase) {
				mockUsecase.
					EXPECT().
					CreateRoom(gomock.Any(), gomock.Any()).
					Return(nil, model.ErrRoomNameRequired)
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "{\"message\":\"room name is required\"}\n",
		},
		{
			name: "失敗: 内部サーバーエラー",
			arg: &RoomRequest{
				Name:        "Test Room",
				Description: "This is a test room",
			},
			fn: func(mockUsecase *mu.MockRoomUsecase) {
				mockUsecase.
					EXPECT().
					CreateRoom(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("internal server error"))
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mu.NewMockRoomUsecase(ctrl)
			tt.fn(mockUsecase)

			e := echo.New()
			body, err := json.Marshal(tt.arg)
			if err != nil {
				t.Fatalf("Failed to marshal request body: %v", err)
			}
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(body)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/rooms")

			handler := NewRoomHandler(mockUsecase)
			err = handler.CreateRoom(c)
			if err != nil {
				e.HTTPErrorHandler(err, c)
			}

			assert.Equal(t, tt.wantStatus, rec.Code)

			if tt.wantBody != "" {
				assert.Equal(t, tt.wantBody, rec.Body.String())
			}
		})
	}
}

func TestRoomHandler_GetRoom(t *testing.T) {
	tests := []struct {
		name       string
		roomID     string
		fn         func(mockUsecase *mu.MockRoomUsecase)
		wantStatus int
		wantBody   string
	}{
		{
			name:   "成功: ルームが見つかる",
			roomID: "01AN4Z07BY79KA1307SR9X4MV3",
			fn: func(mockUsecase *mu.MockRoomUsecase) {
				mockUsecase.
					EXPECT().
					GetRoom(gomock.Any(), "01AN4Z07BY79KA1307SR9X4MV3").
					Return(&usecase.RoomDTO{
						ID:          "01AN4Z07BY79KA1307SR9X4MV3",
						Name:        "Test Room",
						Description: "This is a test room",
					}, nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   "{\"room\":{\"id\":\"01AN4Z07BY79KA1307SR9X4MV3\",\"name\":\"Test Room\",\"description\":\"This is a test room\"}}\n",
		},
		{
			name:   "失敗: ルームIDが不正",
			roomID: "1",
			fn: func(mockUsecase *mu.MockRoomUsecase) {
				mockUsecase.
					EXPECT().
					GetRoom(gomock.Any(), "1").
					Return(nil, model.ErrRoomIDInvalid)
			},
			wantStatus: http.StatusNotFound,
			wantBody:   "{\"message\":\"room id is invalid\"}\n",
		},
		{
			name:   "失敗: ルームが見つからない",
			roomID: "00000000000000000000000000",
			fn: func(mockUsecase *mu.MockRoomUsecase) {
				mockUsecase.
					EXPECT().
					GetRoom(gomock.Any(), "00000000000000000000000000").
					Return(nil, model.ErrRoomNotFound)
			},
			wantStatus: http.StatusNotFound,
			wantBody:   "{\"message\":\"room not found\"}\n",
		},
		{
			name:   "失敗: 内部サーバーエラー",
			roomID: "00000000000000000000000000",
			fn: func(mockUsecase *mu.MockRoomUsecase) {
				mockUsecase.
					EXPECT().
					GetRoom(gomock.Any(), "00000000000000000000000000").
					Return(nil, errors.New("internal server error"))
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mu.NewMockRoomUsecase(ctrl)
			tt.fn(mockUsecase)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/rooms/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.roomID)

			handler := NewRoomHandler(mockUsecase)
			err := handler.GetRoom(c)
			if err != nil {
				e.HTTPErrorHandler(err, c)
			}

			assert.Equal(t, tt.wantStatus, rec.Code)

			if tt.wantBody != "" {
				assert.Equal(t, tt.wantBody, rec.Body.String())
			}
		})
	}
}

func TestRoomHandler_ListRoom(t *testing.T) {
	tests := []struct {
		name       string
		fn         func(mockUsecase *mu.MockRoomUsecase)
		wantStatus int
		wantBody   string
	}{
		{
			name: "成功: ルームが見つかる",
			fn: func(mockUsecase *mu.MockRoomUsecase) {
				mockUsecase.
					EXPECT().
					ListRoom(gomock.Any()).
					Return([]*usecase.RoomDTO{
						{
							ID:          "01AN4Z07BY79KA1307SR9X4MV3",
							Name:        "Test Room 1",
							Description: "This is a test room 1",
						},
						{
							ID:          "02AN4Z07BY79KA1307SR9X4MV4",
							Name:        "Test Room 2",
							Description: "This is a test room 2",
						},
					}, nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   "{\"rooms\":[{\"id\":\"01AN4Z07BY79KA1307SR9X4MV3\",\"name\":\"Test Room 1\",\"description\":\"This is a test room 1\"},{\"id\":\"02AN4Z07BY79KA1307SR9X4MV4\",\"name\":\"Test Room 2\",\"description\":\"This is a test room 2\"}]}\n",
		},
		{
			name: "失敗: 内部サーバーエラー",
			fn: func(mockUsecase *mu.MockRoomUsecase) {
				mockUsecase.
					EXPECT().
					ListRoom(gomock.Any()).
					Return(nil, errors.New("internal server error"))
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mu.NewMockRoomUsecase(ctrl)
			tt.fn(mockUsecase)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/rooms")

			handler := NewRoomHandler(mockUsecase)
			err := handler.ListRoom(c)
			if err != nil {
				e.HTTPErrorHandler(err, c)
			}

			assert.Equal(t, tt.wantStatus, rec.Code)

			if tt.wantBody != "" {
				assert.Equal(t, tt.wantBody, rec.Body.String())
			}
		})
	}
}

func TestRoomHandler_UpdateRoom(t *testing.T) {
	tests := []struct {
		name       string
		roomID     string
		arg        any
		fn         func(mockUsecase *mu.MockRoomUsecase)
		wantStatus int
		wantBody   string
	}{
		{
			name:   "成功: ルームが更新される",
			roomID: "01AN4Z07BY79KA1307SR9X4MV3",
			arg: &RoomRequest{
				Name:        "Updated Room",
				Description: "This is an updated room",
			},
			fn: func(mockUsecase *mu.MockRoomUsecase) {
				mockUsecase.
					EXPECT().
					UpdateRoom(gomock.Any(), gomock.Any()).
					Return(&usecase.RoomDTO{
						ID:          "01AN4Z07BY79KA1307SR9X4MV3",
						Name:        "Updated Room",
						Description: "This is an updated room",
					}, nil)
			},
			wantStatus: http.StatusOK,
			wantBody:   "{\"room\":{\"id\":\"01AN4Z07BY79KA1307SR9X4MV3\",\"name\":\"Updated Room\",\"description\":\"This is an updated room\"}}\n",
		},
		{
			name: "失敗: Bindエラー",
			arg: &struct {
				Name        int
				Description string
			}{
				Name:        1,
				Description: "This is a test room",
			},
			fn: func(mockUsecase *mu.MockRoomUsecase) {
			},
			wantStatus: http.StatusBadRequest,
			wantBody:   "{\"message\":\"Unmarshal type error: expected=string, got=number, field=name, offset=9\"}\n",
		},
		{
			name:   "失敗: ルームIDが不正",
			roomID: "1",
			arg: &RoomRequest{
				Name:        "",
				Description: "This is an updated room",
			},
			fn: func(mockUsecase *mu.MockRoomUsecase) {
				mockUsecase.
					EXPECT().
					UpdateRoom(gomock.Any(), gomock.Any()).
					Return(nil, model.ErrRoomIDInvalid)
			},
			wantStatus: http.StatusNotFound,
			wantBody:   "{\"message\":\"room id is invalid\"}\n",
		},
		{
			name:   "失敗: 内部エラー",
			roomID: "01AN4Z07BY79KA1307SR9X4MV3",
			arg: &RoomRequest{
				Name:        "Updated Room",
				Description: "This is an updated room",
			},
			fn: func(mockUsecase *mu.MockRoomUsecase) {
				mockUsecase.
					EXPECT().
					UpdateRoom(gomock.Any(), gomock.Any()).
					Return(nil, errors.New("internal server error"))
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mu.NewMockRoomUsecase(ctrl)
			tt.fn(mockUsecase)

			e := echo.New()
			body, err := json.Marshal(tt.arg)
			if err != nil {
				t.Fatalf("Failed to marshal request body: %v", err)
			}
			req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(body)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/rooms/:id")
			c.SetParamNames("id")
			c.SetParamValues(tt.roomID)

			handler := NewRoomHandler(mockUsecase)
			err = handler.UpdateRoom(c)
			if err != nil {
				e.HTTPErrorHandler(err, c)
			}

			assert.Equal(t, tt.wantStatus, rec.Code)

			if tt.wantBody != "" {
				assert.Equal(t, tt.wantBody, rec.Body.String())
			}
		})
	}
}
