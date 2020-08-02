package domain

import (
	"encoding/json"
)

type CraeteChatRoomRequest struct {
	ID string `json:"id"`
}

func DecodeCreateChatRoomRequest(buf []byte) (*CraeteChatRoomRequest, error) {
	var req *CraeteChatRoomRequest
	err := json.Unmarshal(buf[:], req)
	if err != nil {
	  return nil, err
	}

	return req, nil
}
