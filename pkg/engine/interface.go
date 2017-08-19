package engine

import (
	"capsulecd/pkg/config"
	"capsulecd/pkg/pipeline"
	"capsulecd/pkg/scm"
)

// Create mock using:
// mockgen -source=pkg/engine/interface.go -destination=pkg/engine/mock/mock_engine.go
type Interface interface {
	Init(pipelineData *pipeline.Data, config config.Interface, sourceScm scm.Interface) error

	// Validate that required executables are available for the following build/test/package/etc steps
	ValidateTools() error

	// Assemble the package contents
	// Validate that any required files (like dependency management files) exist
	// Create any recommended optional/missing files we can in the structure. (.gitignore, etc)
	// Read & Bump the version in the metadata file(s)
	// CAN NOT override
	// MUST set CurrentMetadata
	// MUST set NextMetadata
	// REQUIRES pipelineData.GitLocalPath
	AssembleStep() error

	// Validate & download dependencies for this package.
	// Generate *.lock files for dependencies (should be deleted in PackageStep if necessary)
	// CAN override
	// REQUIRES pipelineData.GitLocalPath
	// REQUIRES CurrentMetadata
	// REQUIRES NextMetadata
	DependenciesStep() error

	// Compile the source for this package (if required)
	// CAN override
	// USES engine_disable_compile
	// USES engine_cmd_compile
	// REQUIRES pipelineData.GitLocalPath
	CompileStep() error

	// Validate code syntax & execute test runner
	// CAN override
	// Run linter
	// Run unit tests
	// Generate coverage reports
	// USES engine_disable_test
	// USES engine_disable_lint
	// USES engine_disable_security_check
	// USES engine_enable_code_mutation - allows CapsuleCD to modify code using linting tools (only available on some systems)
	// USES engine_cmd_lint
	// USES engine_cmd_test
	// USES engine_cmd_security_check
	TestStep() error

	// Commit any local changes and create a git tag. Nothing should be pushed to remote repository yet.
	// Make sure you remove any unnecessary files from the repo before making the commit
	// CAN NOT override
	// MUST set ReleaseCommit
	// MUST set ReleaseVersion
	// REQUIRES pipelineData.GitLocalPath
	// REQUIRES NextMetadata
	// USES engine_package_keep_lock_file
	PackageStep() error

	// Push the release to the package repository (ie. npm, chef supermarket, rubygems)
	// Should validate any required credentials are specified.
	// CAN override
	// REQUIRES pipelineData.GitLocalPath
	// REQUIRES NextMetadata
	// USES chef_supermarket_username
	// USES chef_supermarket_key
	// USES npm_auth_token
	// USES pypi_repository
	// USES pypi_username
	// USES pypi_password
	// USES rubygems_api_key
	DistStep() error
}
