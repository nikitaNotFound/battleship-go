package matchmaking

import "context"

type MatchmakingService struct{}

func NewGameService() *MatchmakingService {
	return &MatchmakingService{}
}

func (s *MatchmakingService) JoinMatchmakingQueue(ctx context.Context) error {

}

func (s *MatchmakingService) LeaveMatchmakingQueue(ctx context.Context) error {

}
