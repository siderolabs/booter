## [booter 0.4.1](https://github.com/siderolabs/booter/releases/tag/v0.4.1) (2026-10-29)

Welcome to the v0.4.1 release of booter!



Please try out the release binaries and report any issues at
https://github.com/siderolabs/booter/issues.

### Contributors

* Utku Ozdemir
* Mateusz Urbanek
* Spencer Smith
* Andrey Smirnov
* Edward Sammut Alessi
* Orzelius
* Noel Georgi
* Oguz Kilcan
* Artem Chernyshev
* Justin Garrison
* Andy Longwill
* Maja Bojarska
* Gerard de Leeuw
* Klesev Dimitrij
* Mateusz Urbanek
* Mickaël Canévet
* Mikolaj Pawlikowski
* Rowan Voermans

### Changes
<details><summary>4 commits</summary>
<p>

* [`bf8b1a2`](https://github.com/siderolabs/booter/commit/bf8b1a273f9ef8ec13b2119d0f73e0f56d17cd84) refactor: use the PXE boot packages of the bare-metal provider
* [`22fb9cc`](https://github.com/siderolabs/booter/commit/22fb9ccf22c2d0f0be4ca2b4c143ddffbc167a95) chore: bump deps, rekres
* [`ea6228e`](https://github.com/siderolabs/booter/commit/ea6228e49dfe1e60a5fc1a79aa629d41eb605131) fix: add dhcp option 54 to the pxe offer
* [`b70e6ff`](https://github.com/siderolabs/booter/commit/b70e6ffbe86b424888cf1a7af12417ee5d0bd177) test: bind the advertised PXE boot files to the served files
</p>
</details>

### Changes from siderolabs/image-factory
<details><summary>105 commits</summary>
<p>

* [`f3d984d`](https://github.com/siderolabs/image-factory/commit/f3d984d6c191252c256cd14923797feb6ac4f610) release(v1.7.2): prepare release
* [`79b67d7`](https://github.com/siderolabs/image-factory/commit/79b67d765451270b294ca170fd3e220e308682d6) chore: add cdn hmac verification
* [`a54d48a`](https://github.com/siderolabs/image-factory/commit/a54d48a6bfeb988ee13aff65d52d811ba17ca119) refactor(http): remove shared dependency bag
* [`983ec95`](https://github.com/siderolabs/image-factory/commit/983ec95bd3e69d16cf82de88f349262d0250545b) refactor(http): isolate registry services
* [`015a950`](https://github.com/siderolabs/image-factory/commit/015a9509fb6e247394b2a0e086307be6f159f1e3) refactor(http): isolate OCI routing
* [`58b4086`](https://github.com/siderolabs/image-factory/commit/58b408659489dc11e9889a87256d6ea91f2c6c7b) refactor(http): isolate browser adapters
* [`df928e0`](https://github.com/siderolabs/image-factory/commit/df928e0dea688d58cc73e708bef23ac8aba9be7b) refactor(http): isolate image services
* [`7c5173e`](https://github.com/siderolabs/image-factory/commit/7c5173e95cd99e86c9aeb46804c45ff234efc567) refactor(http): unify route authentication
* [`55a9fff`](https://github.com/siderolabs/image-factory/commit/55a9fff0433c462cbebd793ec5984242c058ae01) refactor(http): isolate endpoint owners
* [`ae77d98`](https://github.com/siderolabs/image-factory/commit/ae77d98659416a1ae6b4371444edd3bf56aa4074) feat(http): declare route ownership
* [`be369f1`](https://github.com/siderolabs/image-factory/commit/be369f1f21b36957fa91102c63989f23f958e25a) refactor(http): declare route metadata
* [`77ecb9e`](https://github.com/siderolabs/image-factory/commit/77ecb9ec4da817df82cf0eeea54069ecc6acf6a1) release(v1.7.1): prepare release
* [`fec432d`](https://github.com/siderolabs/image-factory/commit/fec432d0a129e1665ee7064bc6d7d54a4cc07ca1) feat: update dependencies
* [`c99ce3a`](https://github.com/siderolabs/image-factory/commit/c99ce3a3f0ec5f9f707d39832c61a66dd6974aa0) feat: improve the token UI
* [`73f93b2`](https://github.com/siderolabs/image-factory/commit/73f93b26709faef8c2db968365b5341e1eafa408) test: switch to SeaweedFS instead of minio for testing
* [`c35986a`](https://github.com/siderolabs/image-factory/commit/c35986aa6755191ec707a874fa33d615f7ef1537) feat: add 'On Premises' role to the token UI
* [`42d9b84`](https://github.com/siderolabs/image-factory/commit/42d9b84432f0b18a29da143dd3d92147ee8b1705) fix(scanner): wait for database replacement
* [`feaf2f6`](https://github.com/siderolabs/image-factory/commit/feaf2f66c40f3a8b630f7b503c3560c02117e72c) fix: download 1.14.0
* [`5da8c30`](https://github.com/siderolabs/image-factory/commit/5da8c3002270532c26b30c0fd290f96403544f70) release(v1.7.0): prepare release
* [`513036c`](https://github.com/siderolabs/image-factory/commit/513036cf242d7e86ba5b11c7ad6562f49cd55abc) fix: cache unauthenticated iPXE scripts, dedupe cache signing
* [`5fec48b`](https://github.com/siderolabs/image-factory/commit/5fec48b671fb8126640d958491fab885b5565559) test: characterize HTTP frontend behavior
* [`4520b38`](https://github.com/siderolabs/image-factory/commit/4520b38a89a68ce0be0b0fa7b1deb6fa20d0f8d8) fix(enterprise): evaluate VEX against Talos kernel
* [`a27dfe6`](https://github.com/siderolabs/image-factory/commit/a27dfe64b2bd2837ec4d461ab40c6d1c1dc5a857) feat: enforce public API with OpenAPI
* [`773ba3c`](https://github.com/siderolabs/image-factory/commit/773ba3c277979a239f1c902f903f10953d1498e3) fix(ui): gate token create modal on the in-flight POST
* [`2e149c8`](https://github.com/siderolabs/image-factory/commit/2e149c84e9facb75761cf2016c6b2ccdba1b10f0) fix(auth0): drop Bearer challenge from 401 response
* [`9f82dde`](https://github.com/siderolabs/image-factory/commit/9f82dde3576d4593fc2f49ff311cbefa54143eae) fix: rename registryClientRefreshInterval to refreshInterval
* [`f4f79db`](https://github.com/siderolabs/image-factory/commit/f4f79dbe47b51ad17d59ad52eda7297680d485b1) feat(enterprise): repo-per-org token storage, url-safe stored tokens
* [`f4e4d01`](https://github.com/siderolabs/image-factory/commit/f4e4d013eedeaa2c404d36cb14e631dcaf0b305e) feat(enterprise): expose actor profiles for token creation UI
* [`8a77753`](https://github.com/siderolabs/image-factory/commit/8a77753c5487f1609bcdbbd62248f6d060385362) feat(enterprise): replace scope enum with route-defined capability scopes
* [`4d9710f`](https://github.com/siderolabs/image-factory/commit/4d9710fa70ae1b81ac3cb98e1a3eb362cfdbb2d5) feat(enterprise): add CLI-only admin token scope
* [`4286666`](https://github.com/siderolabs/image-factory/commit/42866667e4b9ec73d6d8b01057214acd6c1dab1c) feat(enterprise): unify download and node tokens as scoped API tokens
* [`4becd5d`](https://github.com/siderolabs/image-factory/commit/4becd5d03bafb316ea38b89c3c062f37b47ea453) release(v1.6.1): prepare release
* [`e5ff748`](https://github.com/siderolabs/image-factory/commit/e5ff748cb65d95fd4522ce692fe29a25a2e81724) fix: pull in Talos for the SBC/GRUB/EFI fix
* [`b54b945`](https://github.com/siderolabs/image-factory/commit/b54b945c294855e4799eb9e27739be0a5c8efddb) release(v1.6.0): prepare release
* [`9244e82`](https://github.com/siderolabs/image-factory/commit/9244e820bd67cb3bfc427033f9344e50ac2ee857) feat: sign installer profile when unified secureboot supported
* [`3b58dfa`](https://github.com/siderolabs/image-factory/commit/3b58dfa0fe48cd7d8366cac92d52d42e7f156a50) feat(enterprise): add a node tokens page
* [`82bd329`](https://github.com/siderolabs/image-factory/commit/82bd3298b90aad729c34a5fa848f93d777961f95) feat: add node token methods to the factory client
* [`9312d82`](https://github.com/siderolabs/image-factory/commit/9312d82ff73feb8fd303dfd8d12ae498f34549f5) feat(enterprise): add self-issued node tokens
* [`4691e01`](https://github.com/siderolabs/image-factory/commit/4691e0156da8dd62d741e9c5ba4d8e4ab534b19b) feat: add WithTokenSource to include bearer tokens
* [`b115d17`](https://github.com/siderolabs/image-factory/commit/b115d17c79e68b568f1f3f526d1bf22bf253b398) chore(enterprise): drop the Auth0 Management API node-token design
* [`6f08ed9`](https://github.com/siderolabs/image-factory/commit/6f08ed978167b88d5f1d384493f755aa07cb5cec) feat(enterprise): extend the download-token issuer to a second audience
* [`c51ac14`](https://github.com/siderolabs/image-factory/commit/c51ac142788af8de5217d8116628ad43ec4f9d70) feat: accept download token on /pxe/ and forward it
* [`7252d87`](https://github.com/siderolabs/image-factory/commit/7252d8728c5a9f83b5ce3c648f00a9b3f522961b) chore: allow proxying IF through IF
* [`f45ec71`](https://github.com/siderolabs/image-factory/commit/f45ec713035d424a43a5686b94515a07c946a303) fix(enterprise): read the org_id claim from if_org_id
* [`43adb19`](https://github.com/siderolabs/image-factory/commit/43adb19dfa025332eae614c4daa6a203da3029b3) release(v1.5.1): prepare release
* [`5f1f197`](https://github.com/siderolabs/image-factory/commit/5f1f197fcc7a37c50fe9b3ac374a721653f8cd81) feat: update Talos to 1.14.0-rc.2
* [`b7908c1`](https://github.com/siderolabs/image-factory/commit/b7908c1f25c0ab589c8d09f30ddf965be847eab3) feat(enterprise): add Auth0 Management API client for node tokens
* [`36fedd7`](https://github.com/siderolabs/image-factory/commit/36fedd761da751d17efa61c8d07ad35bbb55767c) feat: add WithBearerToken to include m2m token
* [`e783a3d`](https://github.com/siderolabs/image-factory/commit/e783a3d7111eccd76eee28ddd2cc5822a3f26ced) feat(frontend): preserve whitespace for vuln descriptions
* [`18f56f7`](https://github.com/siderolabs/image-factory/commit/18f56f773239c09b10262be89a1b5ce7f36b148a) chore: make sure check-dirty also checks docs
* [`196a447`](https://github.com/siderolabs/image-factory/commit/196a447532a25d7da6ea27f7f194d713aa88d104) fix: enforce canonical image references
* [`26b95ca`](https://github.com/siderolabs/image-factory/commit/26b95cacb0846378fa5ed75e397eb732d6500c0a) feat(enterprise): require auth0 clientID and clientSecret always
* [`25561f7`](https://github.com/siderolabs/image-factory/commit/25561f7ee883d742cf7b873309f1b666b1b4ee72) fix: retry put when joining a failed get flight
* [`aab14ff`](https://github.com/siderolabs/image-factory/commit/aab14ffc33e2b276356cc022d2bd8e563eaf4cae) feat: add spdx and vex reports to factory client
* [`dc6a9f9`](https://github.com/siderolabs/image-factory/commit/dc6a9f9a6a0ca51b36354788f7370b3fcd6328ef) feat(enterprise): theme and translate Auth0 logout/login-error pages
* [`bd13149`](https://github.com/siderolabs/image-factory/commit/bd13149c37cbb1dc67f865f143fd1ddeedcfd4b5) release(v1.5.0): prepare release
* [`d3c693a`](https://github.com/siderolabs/image-factory/commit/d3c693aaba0abc2248694c2df7a6560387e75569) feat: update talos to v1.14.0-rc.1
* [`5c9839a`](https://github.com/siderolabs/image-factory/commit/5c9839ae915c9e3755989d13b7091ecec4c2caaa) feat: use the CI cluster registry cache for integration tests
* [`70e0590`](https://github.com/siderolabs/image-factory/commit/70e0590b702a8990ce32a6314a289751caace049) feat(enterprise): add auth0 browser login
* [`f5f3128`](https://github.com/siderolabs/image-factory/commit/f5f3128913dee8a49f183c4bcd31c0feecb35f8a) fix: record the status the client actually received
* [`0d2275c`](https://github.com/siderolabs/image-factory/commit/0d2275cb0adfb6848cd524cea56a6f66e5889e3a) feat: make download token TTL configurable
* [`a5824c7`](https://github.com/siderolabs/image-factory/commit/a5824c7793ff2ab090b6d22fd997e0e282528823) docs: complete user-facing API reference
* [`dd485bd`](https://github.com/siderolabs/image-factory/commit/dd485bd39fb9107b13384ff94841eec09de99b71) docs: document endpoint access control
* [`6016631`](https://github.com/siderolabs/image-factory/commit/60166312b08b592dc2bb779e58047b101b4b255c) test: fix two flaky checks in the enterprise integration job
* [`f984ad8`](https://github.com/siderolabs/image-factory/commit/f984ad83451abc84a65a1f38631f3153d9444743) chore: let the OIDC test server serve extra routes
* [`f297b62`](https://github.com/siderolabs/image-factory/commit/f297b62e4d866dfe7bab20849f17a0eabb58757e) chore: update Talos to the latest version
* [`86b0a22`](https://github.com/siderolabs/image-factory/commit/86b0a22873263a41cc134c46dd4960de858fb095) fix: re-identify cached SBOM bundles per schematic
* [`77990dc`](https://github.com/siderolabs/image-factory/commit/77990dcd2b3f6b2e8f9a6165d7a1107e0f81d0a9) fix: schedule grype db refresh to avoid replica desync
* [`ad14c5f`](https://github.com/siderolabs/image-factory/commit/ad14c5ff1738427f989e59e8370b6466f1c6d36c) feat(enterprise): publish installer build evidence
* [`dbcc3c6`](https://github.com/siderolabs/image-factory/commit/dbcc3c6d70fb864de2d81df00723970ad39ddb1e) feat: add auth0 bearer token authentication
* [`ae757b5`](https://github.com/siderolabs/image-factory/commit/ae757b5b1ca6915a11e463a951cd77577318b263) chore: bump cosign to v3.1.2, sign via Rekor v2 + TSA
* [`615b279`](https://github.com/siderolabs/image-factory/commit/615b279db9b70064b1855b2edc88f35b435b37d3) feat(enterprise): serve detached Sigstore bundles for assets
* [`5296d4a`](https://github.com/siderolabs/image-factory/commit/5296d4a8f8d5a7fbc0de296f3b5d515db478e19d) chore: update grpc library
* [`dee1a57`](https://github.com/siderolabs/image-factory/commit/dee1a57377118d9e9ef9bc9bf7e503a5a03efceb) feat: support JWT download tokens
* [`c2339ee`](https://github.com/siderolabs/image-factory/commit/c2339ee0acf197d9c3fe17a6431a1616eba4f1a7) feat: support public routes in enterprise plugins
* [`a81f6e9`](https://github.com/siderolabs/image-factory/commit/a81f6e9d219d03f6ceb411247676f1a5d5e70237) feat: proxy images through backing registry
* [`32a3b08`](https://github.com/siderolabs/image-factory/commit/32a3b0894882062290e707781872b4ce0320ba6c) fix: reuse registry puller for bundle verification
* [`efab38f`](https://github.com/siderolabs/image-factory/commit/efab38fc84e8884baec12c65081948ce0691b9f7) release(v1.4.0): prepare release
* [`9c64235`](https://github.com/siderolabs/image-factory/commit/9c6423572ac34132753a52b881771312fdd744a2) feat: add schematic owner validation
* [`ca87d23`](https://github.com/siderolabs/image-factory/commit/ca87d2388929bc3f391ac138f5d60e6254abc365) fix: add single-flight around schematic factory
* [`d45b5ac`](https://github.com/siderolabs/image-factory/commit/d45b5ac58987ccc6ea1d75b228937e228903fe54) docs: link to Image Factory Enterprise docs page
* [`490a993`](https://github.com/siderolabs/image-factory/commit/490a993a322bbe292d98cd36f3e1cc40c3f85e2e) chore: bump pkgs revision to match talos v1.14.0-alpha.2
* [`f9ff935`](https://github.com/siderolabs/image-factory/commit/f9ff935e8319c73c5877ee05b32fe556b664fea9) chore: bump go pkgs
* [`12cd647`](https://github.com/siderolabs/image-factory/commit/12cd6471ad76cef7a0be82192a2555f1dbc0d5b7) feat: add llms.txt for better LLM usage
* [`f65960f`](https://github.com/siderolabs/image-factory/commit/f65960f064fd9776725a0990bf90710ed2c2acd1) fix: audit file defaults
* [`f26e5e2`](https://github.com/siderolabs/image-factory/commit/f26e5e25e0de055906268406f36bb3a7f514f735) feat: add audit log for authenticated requests
* [`beff6e2`](https://github.com/siderolabs/image-factory/commit/beff6e26c30feb1f8e45646066e014808180c83a) feat: support registry namespace prefix for core artifacts
* [`8c489d0`](https://github.com/siderolabs/image-factory/commit/8c489d08f1c880bf5de022deba7fea6cc736eca4) chore: update dependencies
* [`026f8a8`](https://github.com/siderolabs/image-factory/commit/026f8a8887508d39b29a7dc89a2455bffed16ff8) feat: extra extensions (enterprise only)
* [`915ef76`](https://github.com/siderolabs/image-factory/commit/915ef76a1b0210728d266e9342b17fe1a3b0d3e3) chore: add insecure flag to dev config
* [`3bccbe1`](https://github.com/siderolabs/image-factory/commit/3bccbe1286dde26f81fa084206a44e633f3c0324) fix: handle single arch images
* [`6b1c855`](https://github.com/siderolabs/image-factory/commit/6b1c8556657df805955f57738872845f920987fd) refactor: prepare for more than one artifact registry
* [`bee4fe3`](https://github.com/siderolabs/image-factory/commit/bee4fe39d89b854eb09aab68c146eb55164f9377) feat: narrow sbom cache key to extension list only
* [`e0e4a44`](https://github.com/siderolabs/image-factory/commit/e0e4a4469448dc0eec828fc9a5f8a0db0a0cf3a6) refactor: abstract versioned cache
* [`3359f6c`](https://github.com/siderolabs/image-factory/commit/3359f6c6b8f1fabda0ed0c8b8487687432ee155d) feat: add secureboot enrollKeys schematic option
* [`805c51c`](https://github.com/siderolabs/image-factory/commit/805c51cb53287923c41070344e25194c1b30a008) feat: add per-request correlation ID to logs
* [`8cee96d`](https://github.com/siderolabs/image-factory/commit/8cee96d4e7b51a233529f0d6c423890ba65c455c) feat: assert pxe cache in tests
* [`4ec0789`](https://github.com/siderolabs/image-factory/commit/4ec07895de9bda8df7c799472f9fb808a7fec920) feat: bump go-conainerregistry
* [`425e59e`](https://github.com/siderolabs/image-factory/commit/425e59ee7350fadee7850ead7e329f889ee737f6) release(v1.3.3): prepare release
* [`b5d3d92`](https://github.com/siderolabs/image-factory/commit/b5d3d9240959706c9b8bf5a0ad93b7ea4ca3c4e6) fix: vulnerability scans with extensions
* [`916bcf6`](https://github.com/siderolabs/image-factory/commit/916bcf69f12ef3bd7cf0991a166edba537300a2d) feat: update go-vex
* [`9920386`](https://github.com/siderolabs/image-factory/commit/9920386d1f64717786f6a4a930b0a719e81a95bc) feat: update Image Factory with Talos 1.14.0-alpha.1
* [`d49e952`](https://github.com/siderolabs/image-factory/commit/d49e9523ba1755059eb8e6a2b08004c3fd6a68ad) feat: allow excluding Talos releases
* [`147a3e8`](https://github.com/siderolabs/image-factory/commit/147a3e8a648554a6810bb3423905d687efad6d31) feat: add scan report to factory client
* [`2887e78`](https://github.com/siderolabs/image-factory/commit/2887e78d95a54e3eb397c7327d30c6754e84a22e) feat: add support for embedding machine configuration
</p>
</details>

### Changes from siderolabs/omni-infra-provider-bare-metal
<details><summary>99 commits</summary>
<p>

* [`47026ce`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/47026ceade31db9cfc4c6d87941e732df0d8550d) refactor: move the PXE boot stack into public packages
* [`e24f327`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/e24f3274ed6cef3149f87baefdfab3fa115270f8) release(v0.14.1): prepare release
* [`b381ac5`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/b381ac5a8c64df094d2de56193c71e405152e7b8) fix: select the Redfish system matching the machine UUID
* [`77f3c51`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/77f3c51b48ed92e5ce959f5a42e6412e8db294d2) chore: bump deps, rekres
* [`51b8c60`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/51b8c609c04f70e733fb3bd54851c0151665d70b) fix: send the server identifier in DHCP proxy replies
* [`d8c2440`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/d8c2440d363a97ec24d7b416d74b600b1ea1f7ab) release(v0.14.0): prepare release
* [`2bc4a88`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/2bc4a88511acccd40bcda2ac7b98c05d0b0a5a4d) refactor: rename the imagefactory package to installationmedia
* [`9c0b0e3`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/9c0b0e3d9907209157548a0934d649fd924691b9) feat: default the agent-mode Talos version to Omni's default
* [`d6cddce`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/d6cddce7b07a620c74faceeaf6e8120bca511186) chore: bump talos-metal-agent to v0.1.7
* [`ce71ada`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/ce71ada4601833095cb1bdc64243deec9c3a20c9) chore: bump deps, rekres
* [`e52678c`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/e52678c7e069d4bccc6bfacc466748dd5bde3265) release(v0.13.0): prepare release
* [`a84fbd5`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/a84fbd5a84f8219c3029982567341d252b99ea9f) test: bump the Talos version of the integration test cluster
* [`98edcab`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/98edcabdad9f8d37df108893d6a19273b8b630f6) test: run the integration tests against the enterprise image factory
* [`dc7a460`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/dc7a4606d8eeede68717800335e379357d2d2a33) feat: exercise the real IPMI path against emulated QEMU machines
* [`ee94d6e`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/ee94d6edc9669ee0257af7d16c48e8c521b5cd57) feat: get boot assets from the image factory through Omni
* [`0e1e7e6`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/0e1e7e68ff74108d68f9d0176350214de20f71c9) test: bind the advertised PXE boot files to the served files
* [`74a2564`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/74a2564a52a5ed5daf3daab9dc843471cd14fe34) release(v0.12.0): prepare release
* [`a9b4266`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/a9b4266e4710375254e6f3114831cfbb291581a1) feat: support running the provider and emulated machines natively
* [`2d243ab`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/2d243abd76edce0f62e66e123748be52683c73cd) release(v0.11.0): prepare release
* [`12025df`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/12025df354e30fd820c361aa702d7779d0c3236d) feat: report version to omni client
* [`3faf960`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/3faf96066ee72f5dce6065e5a4bd179ef9c2aa30) chore: bump deps, rekres, fix boot assets, add agent guide
* [`ba06836`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/ba068364448c820e9a6fb2385b4d626abc5a5c06) release(v0.10.2): prepare release
* [`dc42fb6`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/dc42fb6fc93a2245741bf0cc4c1306e0debf528f) fix: don't stall boot.ipxe requests when many nodes PXE boot at once
* [`37cae94`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/37cae94841a1bcfb400a274924f9d6831d41217a) chore: rekres
* [`16fad8a`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/16fad8a23e954865be9820d1ba1412114546ff27) test: download omnictl from the /api/omnictl/ path
* [`150bf44`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/150bf4483ced2db7c45f581459b73a73d7121696) release(v0.10.1): prepare release
* [`c93c09b`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/c93c09bff74d1c9e93a656aae6420b5c1f73f1da) fix: write boot filename to BOOTP header for U-Boot ProxyDHCP
* [`e6da964`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/e6da96461583e94ba4f194329655e05e00de89f7) chore: bump Go and deps, rekres, fix linters
* [`5b57fe3`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/5b57fe34f306ca5eb9e1dca7f1202336b1e6c26c) release(v0.10.0): prepare release
* [`a3694fb`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/a3694fbe957427a8d8f769ffea03f3c765815b59) chore: bump image-factory and talos-metal-agent
* [`4d13952`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/4d13952f9c6032d174c3744fc2201946cde3d5e3) chore: rekres
* [`e27fd88`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/e27fd8802d08a3f11114e17878af09e2827ba95b) chore: bump deps, rekres, Talos v1.13.0
* [`fed52d9`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/fed52d9444c59c5279c2a943ab0c79bba5a2891b) feat: honor power-off requests from omni
* [`2730cf3`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/2730cf37ba80c93e01d20ce19b0ef6d930fcb65d) chore: accept eula in integration tests
* [`f43cfa1`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/f43cfa182cd2e4d5d50f4114b4e130905489287c) release(v0.9.0): prepare release
* [`84b50bf`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/84b50bfd8fe947ab01865c079547f82b2c0d406e) fix: listen on both DHCP port 67 and 4011 per PXE spec
* [`c9b458e`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/c9b458e808cd16f34e718c747662169241cfe8bd) test: add BMC integration tests and fix hardcoded IPMI username
* [`6dc0c43`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/6dc0c4380480dc147d2734722e6594a82d289267) release(v0.8.1): prepare release
* [`9ed9832`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/9ed98329b94f69138635508425c48f30b884cf56) chore: bump deps including new redfish changes
* [`89de24a`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/89de24a5545c9e5fc21e09e3cca2791993b1f9f5) fix: correct /tftp/ HTTP path, debug build tags, and comment typos
* [`529806a`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/529806a8616bcc78cca7ce3d963b4123e34975ff) fix: use a private IP range in tests
* [`96eabb6`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/96eabb69eabec9addca6bfeef682dff25b38c177) release(v0.8.0): prepare release
* [`3cf79cc`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/3cf79ccd07a740e03d452489b8c81bed82aebe80) test: set required Omni SQLite storage path flag to integration test
* [`d9e43cc`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/d9e43cc415dccc1e1d45732b0e02d55a8c26298d) release(v0.7.1): prepare release
* [`821b331`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/821b331cec9fd69fe2ff848f2bda8af472df2a43) fix: always include the extra config docs in machine config
* [`e81d812`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/e81d81242efc933728e44b75a032e025a3abf923) release(v0.7.0): prepare release
* [`61f2a5d`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/61f2a5d55340ee8268901c98c568e61b1276dc83) chore: rekres, bump deps
* [`f303b3f`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/f303b3ff330d7de951f1e04ef7d4817113d9f578) feat: allow providing additional config documents from config endpoint
* [`0c53f3a`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/0c53f3aa2cc5495003f899f801c5968fbb39d836) release(v0.6.0): prepare release
* [`9c50645`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/9c50645668285df586dcb0e916a02478a3865c03) feat: allow specifying a custom DHCP proxy port
* [`1a622b6`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/1a622b643571c69d8ffcae183ee64c73234cac50) release(v0.5.0): prepare release
* [`4e7f89b`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/4e7f89b560cb7d733053a7d870e00b6efadcb886) chore: bump Talos version to 1.11.3, make integration tests parallel
* [`e53367e`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/e53367edc0e9fe5ea11b1ca8c219aa0bd500bd90) chore: rekres, bump deps
* [`5bca2eb`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/5bca2eb8c19c5626e26edc9593fb3b5e3acd45fd) refactor: adapt to new QTransform controllers
* [`172f6b7`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/172f6b7b0e1a6467eaa71d0e301ddacd9455f195) feat: add AMI FutureState fallback for boot settings configuration
* [`5eeb667`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/5eeb667f20cc71a68cdabbcde9aa92a841b228b0) chore: run integration tests on daily schedule
* [`268ad6f`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/268ad6fd4f74c956bb05d5f33798fecb794dbae9) feat: allow making ipxe trust a custom CA certificate
* [`f6a8cdd`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/f6a8cdd5af0d30002295aff70f3f69866a285cab) feat: allow providing own ca cert and self-issued certs
* [`8a209de`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/8a209de34a9f3085218fdfe5e2494018efdb5692) chore: update Omni client module
* [`3d91bf6`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/3d91bf62be50f9618660e656965cdc54b97eaa61) chore: bump deps and rekres
* [`9efacdb`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/9efacdba43f86d87705736ba32500c0e7765aee2) chore: bump deps and rekres
* [`04ccf6d`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/04ccf6de786b19fbb2897ac8f50f8ec4b9883826) test: bring back `--join-tokens-mode=strict` in tests
* [`c08b791`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/c08b791af210f950c01e0d0880e0ee861dbaa390) fix: connect the ipmi client before using it
* [`1cc7efc`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/1cc7efcfa5ecf735c74fef338cd5da9202078993) chore: bump deps
* [`87001b9`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/87001b9420e5bb0248f31e60a4f7e386db786397) fix: do not reboot machines which do not need to be powered on
* [`d033824`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/d033824bc26af171cc388e5eafe61df5043ce6c3) chore: bump Go, deps and rekres
* [`e19e669`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/e19e669beb7e1aa1b65938664966baf16d05879f) fix: add support for u-boot arm64 client arch in DHCP proxy
* [`b57c272`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/b57c272892d3358d12bcfe07416e2ef2f7da98a6) docs: add link to official docs
* [`a4202aa`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/a4202aa2b244c354c1cf8fb354c1c454474da75e) feat: support node unique token join flow
* [`f61cb0b`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/f61cb0b32855b837085eb926effd33ab7ef94264) fix: do not attempt to reboot if BMC config is not yet set
* [`1735d5d`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/1735d5d4f815cd421c9eada19cf90a6fb6c454ee) fix: fix pxe boot mode change logic
* [`9d4dcd2`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/9d4dcd2c4b62cd49d0796f07bff1c276531711c3) feat: add secure boot support
* [`59aefc3`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/59aefc3c088dce0e5925fc660e80088ad35a0e64) feat: allow configuring agent calls/wipe timeouts
* [`f62eb78`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/f62eb78ff42723ccb5466dd77a8d19f6c9078c69) feat: use native ipmi client library
* [`eb7f567`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/eb7f5675d235fb1dc15c523b5cbcfcc48aebb413) fix: make tftp server bind on api advertise address
* [`22fde54`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/22fde54c7ddac4939144092d54e37d231e0aab70) feat: allow user-configured BMC credentials
* [`a522f9c`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/a522f9c79da61c06205e30cc84b72fc86faae9a4) fix: serve iPXE binaries over http
* [`86def06`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/86def06e2ae14af3dd26110f74c2a98e7657a033) fix: set the pxe boot flag if required before power on
* [`c71ff31`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/c71ff3112a5dd3c193541d8ebecb5c67d28459d2) feat: add TLS support for agent GRPC connection
* [`36d5fbd`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/36d5fbd84cb9c07f3afad3c3086acd3995c216cf) feat: ignore rejected and cordoned machines in pxe boot handler
* [`24f4e21`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/24f4e21135947df30bce4530fe38215a7ab5633a) refactor!: rework resources and controllers
* [`4581ace`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/4581aceb3d3f87cb22221544c35d09fdd9a3a479) test: update integration tests script
* [`fb2e410`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/fb2e41030470e00399373920927c7a2c72da5406) fix: fix racy access to the provider ID variable
* [`c4ad6a9`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/c4ad6a9391716d6cf0b79711588080270c35bb32) fix: use the correct repo for the Omni client
* [`b811596`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/b811596bdb1ffe81f2ec5462455c236fa3565f49) feat: add health checks
* [`535696d`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/535696d5f5f3a70d94175c5f5e90a2666d16770a) feat: probe Redfish API and use it when available
* [`fd0e1c9`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/fd0e1c9f9b76de46e66ff357d54e1917d7c61602) feat: implement cordoning infra machines
* [`5a36cf1`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/5a36cf1be0b1706d0217c0789895e278d8c72855) feat: clear existing machine labels before setting new labels
* [`04954bd`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/04954bdc880ea31b5f2f697ecc5966fd6860ac84) feat: implement machine reboot
* [`d5f562a`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/d5f562acf348b7efa75f8ce9ee1e36b9148f008c) chore: bump agent Talos version and deps
* [`d2369e4`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/d2369e4cb4750d017fbdcd5ca8dc05f69d4139de) feat: add initial redfish support
* [`d2f788d`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/d2f788df17c72bd5877b7e32d782e543078dd2b7) feat: implement powering machines on/off
* [`b423085`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/b4230852744a9014665f1bb5a6d36200a6d3e27d) feat: set machine labels directly on `InfraMachineStatus`
* [`4f52230`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/4f522307cc203924135634d6c954add7cc861952) feat: prevent rebooting too frequently
* [`e6bfaa9`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/e6bfaa902ae333f56828009aa9d0366e160edaf4) feat: change the boot order to be first disk, then network
* [`2343815`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/23438154746ff93885b57158d601773dfb1f1d81) feat: address various points after review
* [`f2ce7d1`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/f2ce7d1c50c5064f246f1dced0fa2342dad0f875) feat: implement bare-metal provider
* [`0498eb1`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/0498eb1188b4df8fe3fa567a37bcc7d7fec91966) chore: add github actions workflow files
* [`c3e892b`](https://github.com/siderolabs/omni-infra-provider-bare-metal/commit/c3e892b87244f390b161f0355d93b6341cd0ebf1) chore: add readme
</p>
</details>

### Dependency Changes

* **github.com/cosi-project/runtime**                       v1.16.1 -> v1.16.4
* **github.com/siderolabs/image-factory**                   v1.3.2 -> v1.7.2
* **github.com/siderolabs/omni-infra-provider-bare-metal**  47026ceade31 **_new_**
* **github.com/siderolabs/omni/client**                     v1.8.0 -> v1.13.0
* **github.com/siderolabs/talos/pkg/machinery**             v1.13.3 -> 3ba3d6a5e784
* **github.com/stretchr/testify**                           v1.11.1 -> v1.12.1
* **golang.org/x/sync**                                     v0.22.0 -> v0.24.0

Previous release can be found at [v0.4.0](https://github.com/siderolabs/booter/releases/tag/v0.4.0)

## [booter 0.4.0](https://github.com/siderolabs/booter/releases/tag/v0.4.0) (2026-07-30)

Welcome to the v0.4.0 release of booter!



Please try out the release binaries and report any issues at
https://github.com/siderolabs/booter/issues.

### Contributors

* Utku Ozdemir

### Changes
<details><summary>2 commits</summary>
<p>

* [`0a6ebf8`](https://github.com/siderolabs/booter/commit/0a6ebf86479023d78eb98a67e8a65c93f5aee98f) feat: use embedded iPXE binaries and serve them from memory
* [`a29d5d2`](https://github.com/siderolabs/booter/commit/a29d5d265ebce3c9e4cae3fac28bae0f58577dcb) chore: rekres to fix lock old issues gh actions job
</p>
</details>

### Changes from siderolabs/go-zbin
<details><summary>2 commits</summary>
<p>

* [`720e539`](https://github.com/siderolabs/go-zbin/commit/720e53961c7d27b9e363865a1609baa62d994354) feat: implement the iPXE zbin image compression format
* [`7b104fd`](https://github.com/siderolabs/go-zbin/commit/7b104fda23a9635a212ab75f2f361e218da246f4) chore: add initial kres skeleton
</p>
</details>

### Dependency Changes

* **github.com/siderolabs/go-zbin**  v0.1.0 **_new_**
* **github.com/stretchr/testify**    v1.11.1 **_new_**
* **golang.org/x/sync**              v0.20.0 -> v0.22.0

Previous release can be found at [v0.3.1](https://github.com/siderolabs/booter/releases/tag/v0.3.1)

