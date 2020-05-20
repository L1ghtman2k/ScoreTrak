package orm

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/logger"
	"github.com/jinzhu/gorm"
)

type configRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewConfigRepo(db *gorm.DB, log logger.LogInfoFormat) config.Repo {
	return &configRepo{db, log}
}

func (c *configRepo) Get() (*config.DynamicConfig, error) {
	cfg := &config.DynamicConfig{}
	c.db.Take(&cfg)
	return cfg, nil
}

func (c *configRepo) Update(cfg *config.DynamicConfig) error {
	c.log.Debugf("updating the config")
	err := c.db.Model(&cfg).Updates(config.DynamicConfig{RoundDuration: cfg.RoundDuration, Enabled: cfg.Enabled}).Error
	if err != nil {
		c.log.Errorf("error while updating the config, reason : %v", err)
		return err
	}
	return nil
}
