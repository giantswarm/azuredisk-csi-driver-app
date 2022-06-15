# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.16.1] - 2022-06-15

### Changed

- Remove `imagePullSecrets`

## [1.16.0-gs2] - 2022-04-13

### Changed

- Reverted `csi-provisioner` to upstream version 3.1.0.

## [1.16.0-gs1] - 2022-04-13

### Added

- Bumped `azuredisk-csi` to upstream version 1.16.0.
- Bumped `csi-provisioner` to upstream version 3.4.0.
- Bumped `livenessprobe` to upstream version 2.6.0.
- Bumped `csi-node-driver-registrar` to upstream version 2.5.0.

## [1.13.0-gs2] - 2022-03-21

### Added

- Add VerticalPodAutoscaler CR.

## [1.13.0-gs1] - 2022-03-15

### Added

- Initial release with upstream version 1.13.0.

[Unreleased]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.16.1...HEAD
[1.16.1]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.16.0-gs2...v1.16.1
[1.16.0-gs2]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.16.0-gs1...v1.16.0-gs2
[1.16.0-gs1]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.13.0-gs2...v1.16.0-gs1
[1.13.0-gs2]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.13.0-gs1...v1.13.0-gs2
[1.13.0-gs1]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v0.0.0...v1.13.0-gs1
