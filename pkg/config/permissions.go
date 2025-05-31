package config

import (
	"github.com/pkg/errors"
)

type PermissionsStruct struct {
	PathConfig     string
	ServiceKeyPath string
	PermissionType string
	PermissionRole string
	EmailAddress   string
}
type permissions struct {
	permissionsParams `json:"permissions"`
}

type permissionsParams struct {
	PathToConfig     string `json:"path_config"`
	PathToServiceKey string `json:"path_service_key"`
	Type             string `json:"type"`
	Role             string `json:"role"`
	Email            string `json:"email"`
}

func (p *permissions) validate() error {
	return errors.Wrap(p.check(), "failed to validate permissions")
}

func (p *permissions) check() error {
	if p.PathToConfig == "" {
		return errors.New("path to config file is empty")
	}
	if p.PathToServiceKey == "" {
		return errors.New("path to service account key is empty")
	}
	if p.Type == "" {
		return errors.New("permission type is empty")
	}
	if p.Role == "" {
		return errors.New("role is empty")
	}
	if p.Email == "" {
		return errors.New("email address is empty")
	}

	return nil
}

// PathConfig returns the path to the configuration file.
func (p *permissions) pathConfig() string {
	return p.permissionsParams.PathToConfig
}

// ServiceKeyPath returns the path to the service account key.
func (p *permissions) serviceKeyPath() string {
	return p.permissionsParams.PathToServiceKey
}

// PermissionType returns the permission type (e.g. "user", "domain").
func (p *permissions) permissionType() string {
	return p.permissionsParams.Type
}

// PermissionRole returns the role (e.g. "reader", "writer").
func (p *permissions) permissionRole() string {

	return p.permissionsParams.Role
}

// EmailAddress returns the email for which permissions will be granted.
func (p *permissions) emailAddress() string {

	return p.permissionsParams.Email
}
func (p *permissions) Permissions() *PermissionsStruct {

	return &PermissionsStruct{
		PathConfig:     p.pathConfig(),
		ServiceKeyPath: p.serviceKeyPath(),
		PermissionType: p.permissionType(),
		PermissionRole: p.permissionRole(),
		EmailAddress:   p.emailAddress(),
	}
}
