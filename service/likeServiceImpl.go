package service

type LikeServiceImpl struct{}

func (LikeServiceImpl) LikeVideo(userID int64, videoID int64, actionType ActionType) error {
	//TODO implement me
	panic("implement me")
}

func (LikeServiceImpl) LikeList(userID int64) ([]int64, error) {
	//TODO implement me
	panic("implement me")
}

func (LikeServiceImpl) GetFavoriteCount(userID int64) (int64, error) {
	//TODO implement me
	return 0, nil
}

func (LikeServiceImpl) IsFavoriteVideo(userID int64, videoID int64)(bool, error){
	//TODO implement me
	return false, nil
}

func (LikeServiceImpl) GetFavoriteCountByVideoID(videoID int64) (int64, error) {
	//TODO implement me
	return int64(0), nil
}