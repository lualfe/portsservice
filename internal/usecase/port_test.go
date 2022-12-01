package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/lualfe/portsservice/internal/entity"
	"github.com/lualfe/portsservice/internal/usecase/mock"
	"github.com/lualfe/portsservice/pkg/portsstream"
)

func setup(t *testing.T) (*mock.MockportsStreamer, *mock.MockportsRepo) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	return mock.NewMockportsStreamer(ctrl), mock.NewMockportsRepo(ctrl)
}

func TestPorts_Save(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		streamer, repo := setup(t)

		resultChan := make(chan portsstream.Result)
		portResult := entity.Port{Key: "PORT_KEY"}

		successStreamMock(streamer, resultChan)
		successStreamStartMock(streamer, portResult, resultChan)
		successRepoUpsert(repo, portResult)

		p := &Ports{
			streamer: streamer,
			repo:     repo,
		}

		err := p.Save(context.Background(), "./testdata/ports.json")
		if err != nil {
			t.Fatalf("Ports.Save() | %v", err)
		}
	})

	t.Run("Error opening file", func(t *testing.T) {
		p := &Ports{}

		err := p.Save(context.Background(), "invalid")
		if err == nil {
			t.Fatal("Ports.Save() | got error nil, want not nil")
		}
	})

	t.Run("Stream Error", func(t *testing.T) {
		streamer, repo := setup(t)

		resultChan := make(chan portsstream.Result)

		successStreamMock(streamer, resultChan)

		want := errors.New("stream error")
		streamer.EXPECT().
			Start(gomock.Any(), gomock.Any()).
			Do(func(_ any, _ any) {
				resultChan <- portsstream.Result{Err: want}
				close(resultChan)
			}).
			Times(1)

		p := &Ports{
			streamer: streamer,
			repo:     repo,
		}

		err := p.Save(context.Background(), "testdata/ports.json")
		if err == nil && !errors.Is(err, want) {
			t.Fatal("Ports.Save() | got error nil, want not nil")
		}
	})

	t.Run("Error saving ports to database", func(t *testing.T) {
		streamer, repo := setup(t)

		resultChan := make(chan portsstream.Result)
		portResult := entity.Port{Key: "PORT_KEY"}

		successStreamMock(streamer, resultChan)
		successStreamStartMock(streamer, portResult, resultChan)

		want := errors.New("upsert error")
		repo.EXPECT().
			UpsertPort(gomock.Any(), gomock.Any()).
			Return(want)

		p := &Ports{
			streamer: streamer,
			repo:     repo,
		}

		err := p.Save(context.Background(), "testdata/ports.json")
		if err == nil && !errors.Is(err, want) {
			t.Fatal("Ports.Save() | got error nil, want not nil")
		}
	})
}

func successStreamStartMock(streamer *mock.MockportsStreamer, port entity.Port, resultChan chan portsstream.Result) {
	streamer.EXPECT().
		Start(gomock.Any(), gomock.Any()).
		Do(func(_ any, _ any) {
			resultChan <- portsstream.Result{Data: port}
			close(resultChan)
		}).
		Times(1)
}

func successStreamMock(streamer *mock.MockportsStreamer, resultChan chan portsstream.Result) {
	streamer.EXPECT().
		Stream().
		Return(resultChan).
		Times(1)
}

func successRepoUpsert(repo *mock.MockportsRepo, port entity.Port) {
	repo.EXPECT().
		UpsertPort(gomock.Any(), port).
		Return(nil)
}
