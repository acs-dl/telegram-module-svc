# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## 3.0.0
### Fixed
- Specify full name of field with error. (Return error for `page[limit]`, `filter[name]` instead of `limit`, `name`)

## 2.3.0
### Added
- MustEncode function for decoding with panic on error

## 2.2.0

### Added
* DecodeSilently method for decoding without errors

## 2.0.0
### Added
* Decoding and Encoding support for:
    * TextUnmarshaler types (including structs)
    * TextMarshaler types (including structs)
    * slices of TextMarshaler/Unmasrshaler types
    * slices of primitive types and aliases
* Loose constraints on `url` tags: arbitrary types may be used
* Internal support for `required` tag modifier
* Check of field publicity

### Changed
* Encode signature (may return error)

### Fixed
- Empty includes

## 1.1.0
### Added
* Support of a single sorting key
* Sort Key method (trim '-' prefix)

## 1.0.2

### Fixed

Encoding from filter of type `[]string` led to panic.

## 1.0.1

### Fixed

- Cast error when decoding filter to a `[]string`.

## 1.0.0

### Added

* Support for `search` and `sort` query params.
* Support for nested structures
* Support for type aliases
* Default parameters
* Decode now returns typed errors with explanations of what when wrong (can be rendered to client).

### Fixed
* remove uint only restriction for page params
* pointer destinations were not really supported

## 0.1.0
### Added
* proof-of-concept implementation