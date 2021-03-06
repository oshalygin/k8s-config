## 1.2.0 (August 09, 2017)

- Updated deployment service to now properly handle null and zero values with `omitempty` for `fieldRef` environment variable properties.

## 1.1.0 (July 30, 2017)

- Updated deployment service to now properly handle null and zero values with `omitempty`.

## 1.0.2 (July 30, 2017)

- Updated deployment test file which includes a strict `value` instead of pulling from configuration

## 1.0.1 (July 15, 2017)

- Add deployment script error handling
    - If the script fails for any reason it will exit with the appropriate status code(non 0)
    
## 1.0.0 (July 15, 2017)

- Addition of a deployment script.
    - The deployment script is used to create the binary for consumption in a Travis environment.  At this time, only linux amd386 is built.  In the future, this will spit out all the binaries using `gox`
    
## 0.1.2 (July 2, 2017)

- Add code coverage via `gocov`
    - Wire up of CodeClimate

## 0.1.1 (July 2, 2017)

- Add initial layer of tests.

## 0.1.0 (July 2, 2017)

- Add initial code to read arguments.

## 0.0.4 (July 2, 2017)

- Updates to the Travis configuration to only `gox` build after success

## 0.0.4 (July 2, 2017)

- Addition of Travis build job

## 0.0.3 (July 2, 2017)

- Removing required `type` field

## 0.0.2 (July 1, 2017)

- Additional documentation and project layout

## 0.0.1 (July 1, 2017)

- Initial project structure buildout