# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project's packages adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Migrate to App Build Suite (ABS).

## [2.0.0] - 2025-12-11

### Changed

- Chart: Update to upstream v1.33.7. ([#114](https://github.com/giantswarm/azuredisk-csi-driver-app/pull/114))

## [1.32.9] - 2025-09-09

### Changed

- Switch to semver-compatible release name

## [1.32.9-gs1] - 2025-09-04

### Changed

- Chart: Update to upstream v1.32.9.

## [1.31.11-gs1] - 2025-09-02

### Changed

- Chart: Update to upstream v1.31.11.

## [1.30.12-gs1] - 2025-08-26

### Changed

- Chart: Update to upstream v1.30.12.

## [1.30.2-gs2] - 2024-07-30

### Fixed

- Remove duplication key `securityContext`.

## [1.30.2-gs1] - 2024-07-30

### Changed

- Update Kyverno PolicyExceptions to `v2` and fallback to `v2beta1`.
- Bump upstream version to v1.30.2.

## [1.26.2-gs6] - 2024-01-23

### Added

- Add team label in resources.

### Changed

- Configure gsoci.azurecr.io as the default container image registry.

## [1.26.2-gs5] - 2023-11-28

### Added

- Add `global.podSecurityStandards.enforced` value for PSS migration.

## [1.27.0] - 2023-07-14

### Fixed

- **Breaking:** added policy exceptions. this makes app dependent on kyverno.
- Added required values for pss policies.

## [1.26.2-gs4] - 2023-05-10

### Changed

- Remove `capabitlities.apiversion.has` check for VPA to avoid race condition between this app being installed and the api-version providing app being installed
  - With this change the installation of the chart will fail until the `api-version` is available

## [1.26.2-gs3] - 2023-05-09

### Changed

- Upgrade `Chart.yaml` `apiVersion` from `v1` to `v2`
  - this is required to get `capabilities` conditional to work with `HelmReleases`

## [1.26.2-gs2] - 2023-05-03

### Changed

- Disable PSPs for k8s 1.25 and newer.

## [1.26.2-gs1] - 2023-02-15

### Changed
- updates to match upstream helm 1.26.2
  - Bumped `azuredisk-csi` to upstream version 1.26.2
  - Bumped `csi-provisioner` to upstream version 3.3.0
  - Bumped `csi-attacher` to upstream version 4.0.0
  - Bumped `csi-resizer` to upstream version 1.6.0
  - Bumped `liveness-probe` to upstream version 2.8.0
  - Bumped `nodeDriverRegistrar` to upstream version 2.6.2
  - increase csi-provisioner timeout to 30s

## [1.25.2-gs1] - 2022-12-22

* Remove VolumeSnapshotClass since we need to first install the CRDs as a hook before we can push this manifest

## [1.25.1-gs1] - 2022-12-21

* Add StorageClasses for CSI Disk Driver , only enabled for CAPZ Clusters
* Add VolumeSnapshotClass for CSI Disk Driver , only enabled for CAPZ Clusters

## [1.25.0-gs1] - 2022-12-06

### Changed
- Bumped `azuredisk-csi` to upstream version 1.25.0
- Increased qps limits for csi-provisioner and csi-attacher to match upstream
- Increased WorkerThreads for csi-provisioner to match upstream
- Increased csi-attacher timeout from 600s to 1200s to match upstream
- Add support for http/s proxy settings to azuredisk using `cluster-apps-operator` `cluster.proxy` values
- as per [KEP 2067](https://github.com/kubernetes/enhancements/blob/master/keps/sig-cluster-lifecycle/kubeadm/2067-rename-master-label-taint/README.md#renaming-the-node-rolekubernetesiomaster-node-label) add toleration for `node-role.kubernetes.io/control-plane` label and update nodeSelector
- add rbac to get nodes for azuredisk-node to match upstream
- Make some hardcoded fields in the template a value so that we can set them differently in `Vintage` and in `CAPZ`
  - runOnMaster nodeSelector label
  - AZURE_CREDENTIAL_FILE location

## [1.21.0-gs4] - 2022-08-29

### Fixed

- Remove hostPath mount of `/etc/ssl/certs`.

## [1.21.0-gs3] - 2022-08-23

### Added

- Add NetworkPolicy to make app compatible with workload clusters.

## [1.21.0-gs2] - 2022-08-23

### Changed

- Push app to default catalog.

### Added

- Add NetworkPolicy to make app compatible with workload clusters.

## [1.21.0-gs1] - 2022-08-11

### Changed

- Bumped `azuredisk-csi` to upstream version 1.21.0.
- Bumped `csi-provisioner` to upstream version 3.2.1.

## [1.19.0-gs1] - 2022-07-01

### Changed

- Bumped `azuredisk-csi` to upstream version 1.19.0.
- Bumped `csi-provisioner` to upstream version 3.2.0.
- Bumped `csi-attacher` to upstream version 3.5.0.
- Bumped `csi-resizer` to upstream version 1.5.0.
- Bumped `livenessprobe` to upstream version 2.7.0.
- Bumped `csi-node-driver-registrar` to upstream version 2.5.1.
- Update helm chart templates from upstream.
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

[Unreleased]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v2.0.0...HEAD
[2.0.0]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.32.9...v2.0.0
[1.32.9]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.32.9-gs1...v1.32.9
[1.32.9-gs1]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.31.11-gs1...v1.32.9-gs1
[1.31.11-gs1]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.30.12-gs1...v1.31.11-gs1
[1.30.12-gs1]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.30.2-gs2...v1.30.12-gs1
[1.30.2-gs2]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.30.2-gs1...v1.30.2-gs2
[1.30.2-gs1]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.26.2-gs6...v1.30.2-gs1
[1.26.2-gs6]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.26.2-gs5...v1.26.2-gs6
[1.26.2-gs5]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.27.0...v1.26.2-gs5
[1.27.0]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.26.2-gs4...v1.27.0
[1.26.2-gs4]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.26.2-gs3...v1.26.2-gs4
[1.26.2-gs3]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.26.2-gs2...v1.26.2-gs3
[1.26.2-gs2]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.26.2-gs1...v1.26.2-gs2
[1.26.2-gs1]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.25.2-gs1...v1.26.2-gs1
[1.25.2-gs1]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.25.1-gs1...v1.25.2-gs1
[1.25.1-gs1]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.25.0-gs1...v1.25.1-gs1
[1.25.0-gs1]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.21.0-gs4...v1.25.0-gs1
[1.21.0-gs4]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.21.0-gs3...v1.21.0-gs4
[1.21.0-gs3]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.21.0-gs2...v1.21.0-gs3
[1.21.0-gs2]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.21.0-gs1...v1.21.0-gs2
[1.21.0-gs1]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.19.0-gs1...v1.21.0-gs1
[1.19.0-gs1]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.16.0-gs2...v1.19.0-gs1
[1.16.0-gs2]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.16.0-gs1...v1.16.0-gs2
[1.16.0-gs1]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.13.0-gs2...v1.16.0-gs1
[1.13.0-gs2]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v1.13.0-gs1...v1.13.0-gs2
[1.13.0-gs1]: https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v0.0.0...v1.13.0-gs1
