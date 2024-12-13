//go:build failproddeploy

package main

import (
    "errors"
    "log"
    "time"
)

// Simulate production deployment failure
type ProductionDeployment struct {
    Version     string
    DeployTime  time.Time
    Replicas    int
}

func (d *ProductionDeployment) Deploy() error {
    log.Println("Attempting production deployment")
    
    // Simulate pre-deployment checks
    if d.Version == "" {
        return errors.New("production deployment: invalid version")
    }
    
    // Simulate scaling validation
    if d.Replicas < 1 {
        return errors.New("production deployment: insufficient replicas")
    }
    
    // Simulate time-based deployment restriction
    if time.Since(d.DeployTime) < time.Hour {
        return errors.New("production deployment: too soon after previous deployment")
    }
    
    // Always fail production deployment
    return errors.New("intentional production deployment failure")
}