package services

import "context"

type MatchmakingService struct{}

func NewMatchmakingService() *MatchmakingService {
	return &MatchmakingService{}
}

func (s *MatchmakingService) JoinMatchmakingQueue(ctx context.Context) error {
	return nil
}

func (s *MatchmakingService) LeaveMatchmakingQueue(ctx context.Context) error {
	return nil
}
