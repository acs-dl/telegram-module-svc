# Changelog

## 1.7.0
### Changed
- Made `SetContextLog` method public to be able to mock

## 1.6.0
### Added
- forbidden error handling for NotAllowed

## 1.5.0
### Added
- Added request ID setter and getter for default middleware

## 1.4.0
### Changed
- lock raven to 0.2.0

## 1.3.0
### Added
- urlval support

## 1.2.0
### Added
- Serve with sane default timeouts & graceful shutdown
## 1.1.0
### Added
- TTL cache

## 1.0.0
### Fixed
* Pass args to all middlewares in DefaultMiddleware

## 0.14.3

### Fixed

* LoganMiddleware entry stacking

## 0.14.2

### Fixed

* Add stack to recover log

## 0.14.1

### Fixed

* Proper type assert on log getter

## 0.14.0

### Added

* Typo fixing alias for `CtxMiddleware`
* `DefaultMiddlewares` helper to init router with safe defaults
* `Log` entry getter from request context

## 0.13.1

### Fixed

* proper dep chi constraint

## 0.13.0

### Changed

* problems.BadRequest determines status code based on interfaces, not types
