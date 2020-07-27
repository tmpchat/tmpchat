package usecase

import (
	"github.com/tmpchat/tmpchat/message_broker/domain"
	"github.com/tmpchat/tmpchat/message_broker/repository"
)

type ChatRoomUsecase interface {
	AddMessage(roomID string, message *domain.ChatMessage) (*domain.ChatMessage, error)
}

type chatRoomInteractor struct {
	repo repository.ChatRoomRepository
}

func NewChatRoomUsecase() ChatRoomUsecase {
	return chatRoomInteractor{repo: repository.NewChatRoomRepository()}
}

func (itr chatRoomInteractor) AddMessage(roomID string, message *domain.ChatMessage) (*domain.ChatMessage, error) {
	room, err := itr.repo.Find(roomID)
	if err != nil {
		// TODO: 本来はここで Create せず、エラーとすべき。
		//       ただし、現在は api-gateway と同期をとっていないため、暫定対応としてエラーを握り潰し新規作成する。
		// https://github.com/tmpchat/tmpchat/issues/29
		room, err = itr.repo.Create(roomID)
		if err != nil {
			return nil, err
		}
	}

	err = itr.repo.AddMessage(room.ID, *message)
	if err != nil {
		return nil, err
	}

	return message, nil
} 


