# CHANGELOG

This CHANGELOG is a format conforming to [keep-a-changelog](https://github.com/olivierlacan/keep-a-changelog). 
It is generated with git-chglog -o CHANGELOG.md
It assumes the use of [conventional commits](https://www.conventionalcommits.org/)

<a name="unreleased"></a>
## [Unreleased]


<a name="v0.7.2"></a>
## [v0.7.2]
### Chores
- update dependencies


<a name="v0.7.1"></a>
## [v0.7.1]
### Bug Fixes
- ci for tool releases


<a name="v0.7.0"></a>
## [v0.7.0]
### Chores
- remove ioutil
- use fabricator
- move to code.cestus.io/tools/wire
- update go1.23

### Features
- add generics support
- Use //go:generate go tool code.cestus.io/tools/wire/cmd/wire
- add version command


<a name="v0.6.0"></a>
## [v0.6.0]
### All
- update golang.org/x/tools dependency to support Go v1.22.0 ([#401](https://github.com/CestusIO/wire/issues/401))
- update stale refs to Travis ([#312](https://github.com/CestusIO/wire/issues/312))
- fix Github Workflow for tests, drop Travis
- update record/replay files for new version of go ([#310](https://github.com/CestusIO/wire/issues/310))
- try to fix test workflow ([#308](https://github.com/CestusIO/wire/issues/308))
- add a Github Workflow to run tests ([#307](https://github.com/CestusIO/wire/issues/307))

### Bug Fixes
- run wire with -mod=mod ([#353](https://github.com/CestusIO/wire/issues/353))

### Doc
- fix badge in readme ([#315](https://github.com/CestusIO/wire/issues/315))
- fix badge in readme ([#314](https://github.com/CestusIO/wire/issues/314))
- fix badge in readme ([#313](https://github.com/CestusIO/wire/issues/313))

### README
- fix install instructions on tutorial ([#357](https://github.com/CestusIO/wire/issues/357))
- fix install instructions ([#339](https://github.com/CestusIO/wire/issues/339))

### Tests
- disable apidiff check ([#316](https://github.com/CestusIO/wire/issues/316))


<a name="v0.5.0"></a>
## [v0.5.0]
### Doc
- update README to point to issues for questions ([#264](https://github.com/CestusIO/wire/issues/264))

### Docs
- fix incorrect generated code in example ([#277](https://github.com/CestusIO/wire/issues/277))
- add FAQ for explicit binds ([#267](https://github.com/CestusIO/wire/issues/267))
- fix syntax error ([#255](https://github.com/CestusIO/wire/issues/255))

### README
- add link to GitHub Discussions ([#273](https://github.com/CestusIO/wire/issues/273))


<a name="v0.4.0"></a>
## [v0.4.0]
### All
- remove unused dependencies ([#228](https://github.com/CestusIO/wire/issues/228))
- enable go 1.13 ([#212](https://github.com/CestusIO/wire/issues/212))

### Cmd
- Add --output_file_prefix, which sets a prefix for the output file name "wire_gen.go"

### README
- reflect new project status ([#217](https://github.com/CestusIO/wire/issues/217))

### Wire
- wire.FieldsOf should not provide pointer to field type for non-pointer structs ([#210](https://github.com/CestusIO/wire/issues/210))
- FieldsOf now provides a pointer to the field type as well as the actual field type ([#209](https://github.com/CestusIO/wire/issues/209))


<a name="v0.3.0"></a>
## [v0.3.0]
### Admin
- removing proxy instructions, script, and .travis.yml reference ([#133](https://github.com/CestusIO/wire/issues/133))

### All
- add dependency checking to the script Travis runs ([#128](https://github.com/CestusIO/wire/issues/128))
- re-enable travis windows build ([#123](https://github.com/CestusIO/wire/issues/123))
- update version of golang.org/x/tools ([#122](https://github.com/CestusIO/wire/issues/122))

### Bind
- takes a pointer for the second argument ([#152](https://github.com/CestusIO/wire/issues/152))

### Cmd
- support FieldsOf in show ([#143](https://github.com/CestusIO/wire/issues/143))

### Codecov
- also disable patch coverage check ([#137](https://github.com/CestusIO/wire/issues/137))
- disable PR comments ([#135](https://github.com/CestusIO/wire/issues/135))

### Doc
- add guide to use the `wire:"-"` tag ([#195](https://github.com/CestusIO/wire/issues/195))

### Docs
- add FieldsOf to the guide ([#159](https://github.com/CestusIO/wire/issues/159))
- clarify docstring about function parameters being of non-identical types ([#167](https://github.com/CestusIO/wire/issues/167))
- update docs with the bind and struct changes ([#155](https://github.com/CestusIO/wire/issues/155))

### Guide
- clarify interface binding examples ([#141](https://github.com/CestusIO/wire/issues/141))

### Internal
- fix issues found by shellcheck ([#177](https://github.com/CestusIO/wire/issues/177))

### Parse
- use reflect.StructTag to detect wire tags [#179](https://github.com/CestusIO/wire/issues/179) ([#181](https://github.com/CestusIO/wire/issues/181))

### Travis
- refresh runtests.sh and Travis config ([#174](https://github.com/CestusIO/wire/issues/174))
- remove coveralls, only use codecov ([#136](https://github.com/CestusIO/wire/issues/136))
- add codecov.yml with threshold=0 ([#132](https://github.com/CestusIO/wire/issues/132))
- try using codecov ([#131](https://github.com/CestusIO/wire/issues/131))
- remove GOPROXY setting ([#130](https://github.com/CestusIO/wire/issues/130))
- disable windows, update import path config, and only run goveralls on linux ([#115](https://github.com/CestusIO/wire/issues/115))

### Wire
- remove color in wire show ([#190](https://github.com/CestusIO/wire/issues/190))
- improve color output to term by using fatih/color ([#186](https://github.com/CestusIO/wire/issues/186))
- support using '-' tag to prevent filling struct fields ([#163](https://github.com/CestusIO/wire/issues/163))
- use subcommands package, improving help ([#173](https://github.com/CestusIO/wire/issues/173))
- reword the deprecation message of struct ([#158](https://github.com/CestusIO/wire/issues/158))
- improve error for a provider set with a binding but no corresponding provider ([#126](https://github.com/CestusIO/wire/issues/126))


<a name="v0.2.2"></a>
## [v0.2.2]

<a name="v0.2.1"></a>
## [v0.2.1]

<a name="v0.2.0"></a>
## [v0.2.0]
### Doc
- add embeds for Travis, godoc, coveralls to README ([#79](https://github.com/CestusIO/wire/issues/79))

### Docs
- fix typo (missing "forgot") in comment ([#92](https://github.com/CestusIO/wire/issues/92))
- update CONTRIBUTING.md ([#78](https://github.com/CestusIO/wire/issues/78))

### Wire
- output diffs to stdout, and return 1 for diffs and 2 for trouble from diff subcommand ([#90](https://github.com/CestusIO/wire/issues/90))


<a name="v0.1.0"></a>
## v0.1.0
### All
- change copyright notice to The Go Cloud Authors (google/go-cloud[#306](https://github.com/CestusIO/wire/issues/306))
- simplify and clarify some expressions (google/go-cloud[#260](https://github.com/CestusIO/wire/issues/260))
- replace more panic(wire.Build(...)) calls with wire.Build(...); return stuff (google/go-cloud[#248](https://github.com/CestusIO/wire/issues/248))
- move back to github.com/google/go-cloud (google/go-cloud[#162](https://github.com/CestusIO/wire/issues/162))
- fix docs and examples (google/go-cloud[#141](https://github.com/CestusIO/wire/issues/141))
- s/go-cloud/go-x-cloud/ (google/go-cloud[#120](https://github.com/CestusIO/wire/issues/120))

### Docs
- remove references to vgo (google/go-cloud[#206](https://github.com/CestusIO/wire/issues/206))

### Gcp
- set up goose providers

### Goose
- require pointer for first argument to goose.Bind (google/go-cloud[#31](https://github.com/CestusIO/wire/issues/31))
- fix outdated docs (google/go-cloud[#27](https://github.com/CestusIO/wire/issues/27))
- s/Use/Build/
- add test that disambiguate contains the base name
- refactor *gen.inject
- add goose.Value directive
- allow non-injector code to live along with injectors
- use marker functions instead of comments
- clean up comments into full sentences
- remove optional directive
- add show command
- add struct field injection
- support provider cleanup functions
- add interface binding
- allow multiple arguments to use and import
- use same directive parsing code during inject
- add optional provider inputs
- strip vendor from generated import paths
- split into multiple files
- use readable variable names
- add a test for imports across packages
- read tests from testdata
- skip function body type-checking
- use function name as implicit provider set
- rename module to provider set
- dependency injection proof of concept

### Wire
- add a diff command (google/go-cloud[#745](https://github.com/CestusIO/wire/issues/745))
- support multiple packages in Generate (google/go-cloud[#729](https://github.com/CestusIO/wire/issues/729))
- add a test for using a function argument as a provider (google/go-cloud[#724](https://github.com/CestusIO/wire/issues/724))
- give wire.Bind access to the arguments to the injector function (google/go-cloud[#715](https://github.com/CestusIO/wire/issues/715))
- serialize test execution (google/go-cloud[#704](https://github.com/CestusIO/wire/issues/704))
- add FAQ for duplicate providers (google/go-cloud[#660](https://github.com/CestusIO/wire/issues/660))
- remove trailing CR at end of error about cycles (google/go-cloud[#662](https://github.com/CestusIO/wire/issues/662))
- use go/packages for analysis (google/go-cloud[#623](https://github.com/CestusIO/wire/issues/623))
- update README installation instructions  (google/go-cloud[#647](https://github.com/CestusIO/wire/issues/647))
- in tests, don't copy want/wire_gen.go into the temporary GOPATH (google/go-cloud[#646](https://github.com/CestusIO/wire/issues/646))
- allow wire.Value to use values with no parent (i.e., struct fields) (google/go-cloud[#596](https://github.com/CestusIO/wire/issues/596))
- expand on package documentation (google/go-cloud[#587](https://github.com/CestusIO/wire/issues/587))
- add FAQ section (google/go-cloud[#573](https://github.com/CestusIO/wire/issues/573))
- use multiline errors instead of one single line for errors with traces (google/go-cloud[#571](https://github.com/CestusIO/wire/issues/571))
- remove duplicated "a" (google/go-cloud[#561](https://github.com/CestusIO/wire/issues/561))
- fix typo in example code (google/go-cloud[#560](https://github.com/CestusIO/wire/issues/560))
- add info from the dependency graph when a type is not provided
- improve error message for provider set conflicts (google/go-cloud[#500](https://github.com/CestusIO/wire/issues/500))
- fix error messages for Bind and InterfaceValue when the arg doesn't implement the interface (google/go-cloud[#491](https://github.com/CestusIO/wire/issues/491))
- add an example and document how to use wire with mocks (google/go-cloud[#488](https://github.com/CestusIO/wire/issues/488))
- report an error if a func with wire.Build in it is an invalid injector (google/go-cloud[#487](https://github.com/CestusIO/wire/issues/487))
- avoid making variable names that are Go reserved keywords (google/go-cloud[#486](https://github.com/CestusIO/wire/issues/486))
- make Vendor test work with Go module (google/go-cloud[#469](https://github.com/CestusIO/wire/issues/469))
- omit the local package identifier if it matches the package name (google/go-cloud[#385](https://github.com/CestusIO/wire/issues/385))
- use package names to disambiguate variables (google/go-cloud[#386](https://github.com/CestusIO/wire/issues/386))
- add a test that fails due to wire.Value on a value from function scope (google/go-cloud[#405](https://github.com/CestusIO/wire/issues/405))
- add a testcase with multiple similar packages (google/go-cloud[#387](https://github.com/CestusIO/wire/issues/387))
- update Provider.Out to be a slice of provided types, and keep track of the provided concrete type in ProviderSet.providerMap (google/go-cloud[#332](https://github.com/CestusIO/wire/issues/332))
- make tests work with Go modules (google/go-cloud[#331](https://github.com/CestusIO/wire/issues/331))
- Add wire.InterfaceValue, required instead of wire.Value if the value is an interface value (google/go-cloud[#322](https://github.com/CestusIO/wire/issues/322))
- use cmp.Diff instead of shelling out to "diff" to compare against golden strings in tests (google/go-cloud[#287](https://github.com/CestusIO/wire/issues/287))
- respect -record flag for tests (google/go-cloud[#282](https://github.com/CestusIO/wire/issues/282))
- fill in ProviderSet.VarName when the set is a package variable (google/go-cloud[#279](https://github.com/CestusIO/wire/issues/279))
- Build now returns an error if it has any unused arguments (google/go-cloud[#268](https://github.com/CestusIO/wire/issues/268))
- use providerMap.Iterate instead of providerMap.Keys() + At() (google/go-cloud[#265](https://github.com/CestusIO/wire/issues/265))
- show help when requested (google/go-cloud[#238](https://github.com/CestusIO/wire/issues/238))
- fix example in the ProvideBaz section (google/go-cloud[#229](https://github.com/CestusIO/wire/issues/229))
- add check command (google/go-cloud[#207](https://github.com/CestusIO/wire/issues/207))
- don't fail at first error (google/go-cloud[#204](https://github.com/CestusIO/wire/issues/204))
- update internal functions to return []error (google/go-cloud[#197](https://github.com/CestusIO/wire/issues/197))
- add best practices for provider set compatibility (google/go-cloud[#191](https://github.com/CestusIO/wire/issues/191))
- handle build tags with relative paths correctly (google/go-cloud[#188](https://github.com/CestusIO/wire/issues/188))
- remove extra trim in parsing golden test output (google/go-cloud[#185](https://github.com/CestusIO/wire/issues/185))
- only apply wireinject build tag to generated package (google/go-cloud[#176](https://github.com/CestusIO/wire/issues/176))
- use return in tests instead of panic (google/go-cloud[#169](https://github.com/CestusIO/wire/issues/169))
- fix test data race (google/go-cloud[#168](https://github.com/CestusIO/wire/issues/168))
- speed up test execution (google/go-cloud[#163](https://github.com/CestusIO/wire/issues/163))
- add error string assertions in tests (google/go-cloud[#149](https://github.com/CestusIO/wire/issues/149))
- store value expressions in package variables (google/go-cloud[#135](https://github.com/CestusIO/wire/issues/135))
- make solver iterative instead of recursive (google/go-cloud[#137](https://github.com/CestusIO/wire/issues/137))
- change call sites to allow multiple errors (google/go-cloud[#118](https://github.com/CestusIO/wire/issues/118))
- detect cycles incrementally (google/go-cloud[#102](https://github.com/CestusIO/wire/issues/102))
- clean up README (google/go-cloud[#100](https://github.com/CestusIO/wire/issues/100))
- build provider map incrementally (google/go-cloud[#96](https://github.com/CestusIO/wire/issues/96))
- allow non-panic version of injector (google/go-cloud[#91](https://github.com/CestusIO/wire/issues/91))
- rename from Goose (google/go-cloud[#59](https://github.com/CestusIO/wire/issues/59))

### Wire
- fix typo in README: App -> Baz (google/go-cloud[#239](https://github.com/CestusIO/wire/issues/239))


[Unreleased]: https://github.com/CestusIO/wire/compare/v0.7.2...HEAD
[v0.7.2]: https://github.com/CestusIO/wire/compare/v0.7.1...v0.7.2
[v0.7.1]: https://github.com/CestusIO/wire/compare/v0.7.0...v0.7.1
[v0.7.0]: https://github.com/CestusIO/wire/compare/v0.6.0...v0.7.0
[v0.6.0]: https://github.com/CestusIO/wire/compare/v0.5.0...v0.6.0
[v0.5.0]: https://github.com/CestusIO/wire/compare/v0.4.0...v0.5.0
[v0.4.0]: https://github.com/CestusIO/wire/compare/v0.3.0...v0.4.0
[v0.3.0]: https://github.com/CestusIO/wire/compare/v0.2.2...v0.3.0
[v0.2.2]: https://github.com/CestusIO/wire/compare/v0.2.1...v0.2.2
[v0.2.1]: https://github.com/CestusIO/wire/compare/v0.2.0...v0.2.1
[v0.2.0]: https://github.com/CestusIO/wire/compare/v0.1.0...v0.2.0
