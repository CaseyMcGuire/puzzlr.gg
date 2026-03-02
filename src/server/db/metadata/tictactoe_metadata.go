package metadata

type TicTacToeMetadata struct {
	PlayerMarker string `json:"playerMarker"`
}

func (s *TicTacToeMetadata) gameMetadata() {}
