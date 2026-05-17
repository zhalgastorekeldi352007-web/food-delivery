package logger

import (
    "go.uber.org/zap"
)

func InitLogger(level string) (*zap.Logger, error) {
    var config zap.Config
    if level == "debug" {
        config = zap.NewDevelopmentConfig()
    } else {
        config = zap.NewProductionConfig()
    }
    
    logger, err := config.Build()
    if err != nil {
        return nil, err
    }
    
    return logger, nil
}
