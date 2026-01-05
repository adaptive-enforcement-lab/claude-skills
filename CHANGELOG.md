# Changelog

## [0.4.0](https://github.com/adaptive-enforcement-lab/claude-skills/compare/v0.3.1...v0.4.0) (2026-01-05)


### Features

* display category count in generation summary ([8755a8e](https://github.com/adaptive-enforcement-lab/claude-skills/commit/8755a8e4746f17bf2311bc8a89162ff796f3951d))


### Bug Fixes

* correct marketplace extra-files path for release-please ([2179259](https://github.com/adaptive-enforcement-lab/claude-skills/commit/21792591e656ce4b477e9a7109558c811bbfa176))


### Maintenance

* **main:** release marketplace 0.2.1 ([#12](https://github.com/adaptive-enforcement-lab/claude-skills/issues/12)) ([c8e8c29](https://github.com/adaptive-enforcement-lab/claude-skills/commit/c8e8c298560450cd47cec1aca3ddaeeb44242fe7))

## [0.3.1](https://github.com/adaptive-enforcement-lab/claude-skills/compare/v0.3.0...v0.3.1) (2026-01-05)


### Bug Fixes

* correct release-please outputs and build workflow triggers ([98d47a3](https://github.com/adaptive-enforcement-lab/claude-skills/commit/98d47a3bc3e2a4ca775cbacc9dd9ed4e5eab95dc))

## [0.3.0](https://github.com/adaptive-enforcement-lab/claude-skills/compare/v0.2.0...v0.3.0) (2026-01-05)


### Features

* add completion status logging to generator ([9064dd9](https://github.com/adaptive-enforcement-lab/claude-skills/commit/9064dd9366068fa1f665320b8083d7ae4b61d5c0))
* add essential repository documentation and release automation ([4182537](https://github.com/adaptive-enforcement-lab/claude-skills/commit/4182537a6a49250f538e6c54defa451067dba3ca))
* add marketplace aggregate component for catalog versioning ([a309b8d](https://github.com/adaptive-enforcement-lab/claude-skills/commit/a309b8de6fb82f69e8c87524e31bfc8dad82ee3c))
* add marketplace version sync workflow ([a1babbb](https://github.com/adaptive-enforcement-lab/claude-skills/commit/a1babbb66f79b26ac0998c827b4d3704eef86789))
* add version flag with ldflags injection ([084fab7](https://github.com/adaptive-enforcement-lab/claude-skills/commit/084fab77b4287ed4a6c02479c552c75a8b82f853))
* implement clean architecture skill generator ([f653d1d](https://github.com/adaptive-enforcement-lab/claude-skills/commit/f653d1dd3872286623a6163c53dca9e8701d387b))
* implement modular release pipeline pattern ([50e2e58](https://github.com/adaptive-enforcement-lab/claude-skills/commit/50e2e588444c192ab85b79dbc5c4ec22674f33fa))
* implement workflow-dispatch-coordination pattern ([6aaf673](https://github.com/adaptive-enforcement-lab/claude-skills/commit/6aaf67373f888ef9205fa28cbbdb970de6c7ec43))
* integrate Core App authentication and PR workflow ([992c076](https://github.com/adaptive-enforcement-lab/claude-skills/commit/992c0764487770433e0d63453ef590f37bf43ec5))
* separate generator and marketplace versioning ([d9dfa76](https://github.com/adaptive-enforcement-lab/claude-skills/commit/d9dfa7610a0da75513fd7b9d2852571779b4313f))


### Bug Fixes

* add checkout step to dispatch-build job ([695e2b1](https://github.com/adaptive-enforcement-lab/claude-skills/commit/695e2b110214bed65c127875cf81391f35b00942))
* check only staged paths in git status ([df89164](https://github.com/adaptive-enforcement-lab/claude-skills/commit/df8916434e23dcfa8f956206ad71eb0d8b926970))
* configure each skill category as independent component ([ce00f01](https://github.com/adaptive-enforcement-lab/claude-skills/commit/ce00f01a89e25fcdcf86991e934bb57ad414b2f9))
* configure release-please to update marketplace.json versions ([297710a](https://github.com/adaptive-enforcement-lab/claude-skills/commit/297710a244895b86df21fb4647acd6bd5de3f386))
* correct extra-files paths for marketplace.json updates ([ad2f5cd](https://github.com/adaptive-enforcement-lab/claude-skills/commit/ad2f5cdd001b4f57fc98021ab80fce1d74071191))
* don't exit with error code for non-fatal skill generation errors ([2ec9136](https://github.com/adaptive-enforcement-lab/claude-skills/commit/2ec91363e9e4bb5fcf316c23551a5272ff2b8d69))
* enable v prefix for both component tags ([7d89ff3](https://github.com/adaptive-enforcement-lab/claude-skills/commit/7d89ff3c50152957d72e9f150150d3695df0310a))
* high-priority markdown generation issues ([50ec2e3](https://github.com/adaptive-enforcement-lab/claude-skills/commit/50ec2e392ee5bfd25342e977c06c0954b96b7cc3))
* implement single source of truth for builds ([369fde5](https://github.com/adaptive-enforcement-lab/claude-skills/commit/369fde52c7d8663bf0fb629b02f93c448d8bb4ce))
* make skill generation workflow idempotent ([dbd1a26](https://github.com/adaptive-enforcement-lab/claude-skills/commit/dbd1a26ff09a26e3bd1919161f380e8db350b72f))
* properly extract section content in markdown parser ([981a3c5](https://github.com/adaptive-enforcement-lab/claude-skills/commit/981a3c5b8f2502ab3fc2bb8dd042486341428fb1))
* remove component from generator tag for Go compatibility ([0c28620](https://github.com/adaptive-enforcement-lab/claude-skills/commit/0c28620f065733dcb261874407c2fe56c005fec3))
* remove extra-files from skill components ([af37f8f](https://github.com/adaptive-enforcement-lab/claude-skills/commit/af37f8f091e5f3c7a05808c9b8b8700e3d78c20d))
* remove unused os import ([3b5cd7b](https://github.com/adaptive-enforcement-lab/claude-skills/commit/3b5cd7b2d38f941fd03953207f1aee0d21efa567))


### Maintenance

* initial repository structure with Go-based skill generator ([ce0da8f](https://github.com/adaptive-enforcement-lab/claude-skills/commit/ce0da8feaff7a8945ef2c0db7be682485a97ec19))
* **main:** release build 0.1.1 ([#8](https://github.com/adaptive-enforcement-lab/claude-skills/issues/8)) ([7e3ba60](https://github.com/adaptive-enforcement-lab/claude-skills/commit/7e3ba601a2ef7dff237a4744cd330fa0d75762bf))
* **main:** release enforcement 0.2.0 ([251b9dd](https://github.com/adaptive-enforcement-lab/claude-skills/commit/251b9dd22926177a1a407bfeee48b92d34ce8d47))
* **main:** release enforcement 0.2.0 ([c2a5430](https://github.com/adaptive-enforcement-lab/claude-skills/commit/c2a5430f21192d82f85c776fa1e986b6f8d09c31))
* **main:** release generator 0.2.0 ([61f36a0](https://github.com/adaptive-enforcement-lab/claude-skills/commit/61f36a0e7b4fd8287a6b92c5cce059c6df940473))
* **main:** release generator 0.2.0 ([56e1cc4](https://github.com/adaptive-enforcement-lab/claude-skills/commit/56e1cc425125a94260f62f738f8646163d1a2ef2))
* **main:** release marketplace 0.2.0 ([915f0f4](https://github.com/adaptive-enforcement-lab/claude-skills/commit/915f0f43f24cd385737abc49afe2b95492a3f226))
* **main:** release marketplace 0.2.0 ([a46a7b2](https://github.com/adaptive-enforcement-lab/claude-skills/commit/a46a7b2359c5d8d99eb7548f5472e0d75070c9c8))
* **main:** release patterns 0.2.0 ([#7](https://github.com/adaptive-enforcement-lab/claude-skills/issues/7)) ([1cb338c](https://github.com/adaptive-enforcement-lab/claude-skills/commit/1cb338ccac6ef1ad22c31721a7b59646863b6a89))
* **main:** release secure 0.2.0 ([36d6e07](https://github.com/adaptive-enforcement-lab/claude-skills/commit/36d6e074d9392f59891da4b3943b0b032356a63a))
* **main:** release secure 0.2.0 ([95a21d3](https://github.com/adaptive-enforcement-lab/claude-skills/commit/95a21d32626b277f6fe9b8ebfcffdaad6f318401))

## [0.2.0](https://github.com/adaptive-enforcement-lab/claude-skills/compare/generator-v0.1.0...generator-v0.2.0) (2026-01-04)


### Features

* add essential repository documentation and release automation ([4182537](https://github.com/adaptive-enforcement-lab/claude-skills/commit/4182537a6a49250f538e6c54defa451067dba3ca))
* add marketplace aggregate component for catalog versioning ([a309b8d](https://github.com/adaptive-enforcement-lab/claude-skills/commit/a309b8de6fb82f69e8c87524e31bfc8dad82ee3c))
* implement clean architecture skill generator ([f653d1d](https://github.com/adaptive-enforcement-lab/claude-skills/commit/f653d1dd3872286623a6163c53dca9e8701d387b))
* integrate Core App authentication and PR workflow ([992c076](https://github.com/adaptive-enforcement-lab/claude-skills/commit/992c0764487770433e0d63453ef590f37bf43ec5))
* separate generator and marketplace versioning ([d9dfa76](https://github.com/adaptive-enforcement-lab/claude-skills/commit/d9dfa7610a0da75513fd7b9d2852571779b4313f))


### Bug Fixes

* check only staged paths in git status ([df89164](https://github.com/adaptive-enforcement-lab/claude-skills/commit/df8916434e23dcfa8f956206ad71eb0d8b926970))
* configure each skill category as independent component ([ce00f01](https://github.com/adaptive-enforcement-lab/claude-skills/commit/ce00f01a89e25fcdcf86991e934bb57ad414b2f9))
* configure release-please to update marketplace.json versions ([297710a](https://github.com/adaptive-enforcement-lab/claude-skills/commit/297710a244895b86df21fb4647acd6bd5de3f386))
* don't exit with error code for non-fatal skill generation errors ([2ec9136](https://github.com/adaptive-enforcement-lab/claude-skills/commit/2ec91363e9e4bb5fcf316c23551a5272ff2b8d69))
* enable v prefix for both component tags ([7d89ff3](https://github.com/adaptive-enforcement-lab/claude-skills/commit/7d89ff3c50152957d72e9f150150d3695df0310a))
* high-priority markdown generation issues ([50ec2e3](https://github.com/adaptive-enforcement-lab/claude-skills/commit/50ec2e392ee5bfd25342e977c06c0954b96b7cc3))
* make skill generation workflow idempotent ([dbd1a26](https://github.com/adaptive-enforcement-lab/claude-skills/commit/dbd1a26ff09a26e3bd1919161f380e8db350b72f))
* properly extract section content in markdown parser ([981a3c5](https://github.com/adaptive-enforcement-lab/claude-skills/commit/981a3c5b8f2502ab3fc2bb8dd042486341428fb1))
* remove unused os import ([3b5cd7b](https://github.com/adaptive-enforcement-lab/claude-skills/commit/3b5cd7b2d38f941fd03953207f1aee0d21efa567))


### Maintenance

* initial repository structure with Go-based skill generator ([ce0da8f](https://github.com/adaptive-enforcement-lab/claude-skills/commit/ce0da8feaff7a8945ef2c0db7be682485a97ec19))
* **main:** release enforcement 0.2.0 ([251b9dd](https://github.com/adaptive-enforcement-lab/claude-skills/commit/251b9dd22926177a1a407bfeee48b92d34ce8d47))
* **main:** release enforcement 0.2.0 ([c2a5430](https://github.com/adaptive-enforcement-lab/claude-skills/commit/c2a5430f21192d82f85c776fa1e986b6f8d09c31))
* **main:** release marketplace 0.2.0 ([915f0f4](https://github.com/adaptive-enforcement-lab/claude-skills/commit/915f0f43f24cd385737abc49afe2b95492a3f226))
* **main:** release marketplace 0.2.0 ([a46a7b2](https://github.com/adaptive-enforcement-lab/claude-skills/commit/a46a7b2359c5d8d99eb7548f5472e0d75070c9c8))
* **main:** release secure 0.2.0 ([36d6e07](https://github.com/adaptive-enforcement-lab/claude-skills/commit/36d6e074d9392f59891da4b3943b0b032356a63a))
* **main:** release secure 0.2.0 ([95a21d3](https://github.com/adaptive-enforcement-lab/claude-skills/commit/95a21d32626b277f6fe9b8ebfcffdaad6f318401))
