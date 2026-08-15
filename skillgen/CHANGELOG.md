# Changelog

## [1.1.0](https://github.com/adaptive-enforcement-lab/claude-skills/compare/v1.0.3...v1.1.0) (2026-08-15)


### Features

* replace per-doc skills with lean hub skills and offline reference docs ([082ac8e](https://github.com/adaptive-enforcement-lab/claude-skills/commit/082ac8ec85e00ab42d5e0f9aa5d2ecfaf2d433d7))

## [1.0.3](https://github.com/adaptive-enforcement-lab/claude-skills/compare/v1.0.2...v1.0.3) (2026-07-25)


### Maintenance

* correct CLAUDE.md and published plugin descriptions ([#83](https://github.com/adaptive-enforcement-lab/claude-skills/issues/83)) ([c6cb3c1](https://github.com/adaptive-enforcement-lab/claude-skills/commit/c6cb3c15f9cac556c1bfea4c3c884ed54e3cae7d))

## [1.0.2](https://github.com/adaptive-enforcement-lab/claude-skills/compare/v1.0.1...v1.0.2) (2026-07-21)


### Bug Fixes

* **skillgen:** correct source URLs, add validation, remove dead code ([#61](https://github.com/adaptive-enforcement-lab/claude-skills/issues/61)) ([602c368](https://github.com/adaptive-enforcement-lab/claude-skills/commit/602c368fb6211d583fb00745a32dde3034490aee))


### Maintenance

* consolidate renovate updates and move to Go 1.26 ([#69](https://github.com/adaptive-enforcement-lab/claude-skills/issues/69)) ([172b000](https://github.com/adaptive-enforcement-lab/claude-skills/commit/172b000874a81d2f8fb80b1a3ed6d5aa99dc605e))

## [1.0.1](https://github.com/adaptive-enforcement-lab/claude-skills/compare/v1.0.0...v1.0.1) (2026-01-05)


### Bug Fixes

* correct plugin.json version lookup from plugins/* path ([#51](https://github.com/adaptive-enforcement-lab/claude-skills/issues/51)) ([e3e5a84](https://github.com/adaptive-enforcement-lab/claude-skills/commit/e3e5a843dc51ea7fd0483599490808aadeee72e6))

## [1.0.0](https://github.com/adaptive-enforcement-lab/claude-skills/compare/v0.5.0...v1.0.0) (2026-01-05)


### ⚠ BREAKING CHANGES

* Directory structure changed from skills/{plugin}/ to plugins/{plugin}/skills/

### Bug Fixes

* restructure skills to plugins/*/skills/ hierarchy ([#44](https://github.com/adaptive-enforcement-lab/claude-skills/issues/44)) ([10f98f2](https://github.com/adaptive-enforcement-lab/claude-skills/commit/10f98f2de398e98eb5bf5a8ff936a1f52fce6546))

## [0.5.0](https://github.com/adaptive-enforcement-lab/claude-skills/compare/v0.4.1...v0.5.0) (2026-01-05)


### Features

* config-driven marketplace generation ([#35](https://github.com/adaptive-enforcement-lab/claude-skills/issues/35)) ([9396fed](https://github.com/adaptive-enforcement-lab/claude-skills/commit/9396fedea18cf8767b6ced7f2e75a732b56ab421))


### Bug Fixes

* make marketplace.json generation deterministic ([#43](https://github.com/adaptive-enforcement-lab/claude-skills/issues/43)) ([7d36c22](https://github.com/adaptive-enforcement-lab/claude-skills/commit/7d36c2203d678617e8417ee290a2ead1100707c4))

## [0.4.1](https://github.com/adaptive-enforcement-lab/claude-skills/compare/v0.4.0...v0.4.1) (2026-01-05)


### Code Refactoring

* move Go package to skillgen subdirectory ([#23](https://github.com/adaptive-enforcement-lab/claude-skills/issues/23)) ([88e5d1b](https://github.com/adaptive-enforcement-lab/claude-skills/commit/88e5d1bf2e27ac4683773ea1857632ab45d0e212))
