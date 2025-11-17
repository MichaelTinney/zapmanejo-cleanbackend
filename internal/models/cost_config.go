package models

type CostConfig struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	UserID    uint    `gorm:"not null;unique" json:"user_id"`
	Nutricao  float64 `json:"nutricao"`  // R$/cabeça/mês
	Sanidade  float64 `json:"sanidade"`
	MaoDeObra float64 `json:"mao_de_obra"`
	Terra     float64 `json:"terra"`
	Outros    float64 `json:"outros"`
}
