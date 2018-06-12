# Change Log

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

<a name="4.0.0"></a>
# [4.0.0](https://github.com/suzuki-shunsuke/go-set/compare/v3.0.0...v4.0.0) (2018-06-12)


### Features

* don't use the pointer ([d8459e5](https://github.com/suzuki-shunsuke/go-set/commit/d8459e5))


### BREAKING CHANGES

* * change API interface: *StrSet -> StrSet
* remove StrSet.Len method: use len(StrSet)



<a name="3.0.0"></a>
# [3.0.0](https://github.com/suzuki-shunsuke/go-set/compare/v2.3.4...v3.0.0) (2018-06-12)


### Code Refactoring

* change StrSet structure ([11dc885](https://github.com/suzuki-shunsuke/go-set/commit/11dc885))


### BREAKING CHANGES

* StrSet's structure changes



<a name="2.3.4"></a>
## [2.3.4](https://github.com/suzuki-shunsuke/go-set/compare/v2.3.3...v2.3.4) (2018-05-11)


### Bug Fixes

* let Add method return error when set is nil ([17d22ef](https://github.com/suzuki-shunsuke/go-set/commit/17d22ef))
* let Adds method return error when set is nil ([d94d6e5](https://github.com/suzuki-shunsuke/go-set/commit/d94d6e5))
* let AddSet method return error when set is nil ([b054856](https://github.com/suzuki-shunsuke/go-set/commit/b054856))
* let AddSets method return error when set is nil ([e741fc0](https://github.com/suzuki-shunsuke/go-set/commit/e741fc0))
* let UnmarshalJSON return error when set is nil ([eb6e986](https://github.com/suzuki-shunsuke/go-set/commit/eb6e986))
* return empty list or map instead of nil ([c41c600](https://github.com/suzuki-shunsuke/go-set/commit/c41c600))



<a name="2.3.3"></a>
## [2.3.3](https://github.com/suzuki-shunsuke/go-set/compare/v2.3.2...v2.3.3) (2018-05-11)


### Bug Fixes

* fix nil check in methods ([#13](https://github.com/suzuki-shunsuke/go-set/issues/13)) ([733cdfc](https://github.com/suzuki-shunsuke/go-set/commit/733cdfc))



<a name="2.3.2"></a>
## [2.3.2](https://github.com/suzuki-shunsuke/go-set/compare/v2.3.1...v2.3.2) (2018-04-16)


### Bug Fixes

* fix ToList when size is 0 ([#9](https://github.com/suzuki-shunsuke/go-set/issues/9)) ([fda2653](https://github.com/suzuki-shunsuke/go-set/commit/fda2653))



<a name="2.3.1"></a>
## [2.3.1](https://github.com/suzuki-shunsuke/go-set/compare/v2.3.0...v2.3.1) (2018-03-30)


### Bug Fixes

* fix comment ([e2ec3c6](https://github.com/suzuki-shunsuke/go-set/commit/e2ec3c6))


### Performance Improvements

* improve MapstructureDecodeHookFromListToStrSet ([9f6a18b](https://github.com/suzuki-shunsuke/go-set/commit/9f6a18b))



<a name="2.3.0"></a>
# [2.3.0](https://github.com/suzuki-shunsuke/go-set/compare/v2.2.0...v2.3.0) (2018-03-29)


### Features

* implements mapstructure.DecodeHookFuncType ([#7](https://github.com/suzuki-shunsuke/go-set/issues/7)) ([c112171](https://github.com/suzuki-shunsuke/go-set/commit/c112171))



<a name="2.2.0"></a>
# [2.2.0](https://github.com/suzuki-shunsuke/go-set/compare/v2.1.1...v2.2.0) (2018-03-29)


### Features

* add some methods to StrSet ([#6](https://github.com/suzuki-shunsuke/go-set/issues/6)) ([a2e077b](https://github.com/suzuki-shunsuke/go-set/commit/a2e077b))



<a name="2.1.1"></a>
## [2.1.1](https://github.com/suzuki-shunsuke/go-set/compare/v2.1.0...v2.1.1) (2018-03-28)


### Bug Fixes

* prevent panic when data is nil ([#5](https://github.com/suzuki-shunsuke/go-set/issues/5)) ([e1a5811](https://github.com/suzuki-shunsuke/go-set/commit/e1a5811))



<a name="2.1.0"></a>
# [2.1.0](https://github.com/suzuki-shunsuke/go-set/compare/v2.0.0...v2.1.0) (2018-03-28)


### Features

* add remove methods to StrSet ([#3](https://github.com/suzuki-shunsuke/go-set/issues/3)) ([ce25d6e](https://github.com/suzuki-shunsuke/go-set/commit/ce25d6e))



<a name="2.0.0"></a>
# [2.0.0](https://github.com/suzuki-shunsuke/go-set/compare/v1.0.0...v2.0.0) (2018-03-28)


### Bug Fixes

* change receiver to pointer ([#2](https://github.com/suzuki-shunsuke/go-set/issues/2)) ([9e7f364](https://github.com/suzuki-shunsuke/go-set/commit/9e7f364))


### BREAKING CHANGES

* change receiver to pointer



<a name="1.0.0"></a>
# 1.0.0 (2018-03-28)


### Features

* implement StrSet ([#1](https://github.com/suzuki-shunsuke/go-set/issues/1)) ([1be010d](https://github.com/suzuki-shunsuke/go-set/commit/1be010d))
