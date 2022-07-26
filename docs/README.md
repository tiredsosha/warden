**Table of Contents**

- [Wardener](#wardener)
  - [Features](#features)
  - [Requirements](#requirements)
    - [Layout](#layout)
  - [Running](#running)
  - [Configuration](#configuration)
    - [config.yaml](#config.yaml)
  - [Logging](#logging)
  - [Supported messages](#supported-messages)
    - [Status](#status)
    - [Commands](#commands)
  - [Building](#building)
  - [Future](#future)
  - [Alternativies](#alternativies)
  - [License](#license)

# Wardener

A simple background service that remotely controls Windows over MQTT.

## Features

- Control
  - Mute/Unmute system
  - Change volume level
  - Reboot system
  - Shutdown system
- Publishing of current volume status
- Works as a background proccess, so no pop-up windows and no need in nircmd

## Requirements

If you use binary file:

- `Windows 10`

If you use source code:

- `Windows 10`
- `Go 1.18 or greater`

### Layout

```tree
├── .github
├── .gitignore
├── README.md
├── docs
│   └── README.md
├── control
│   ├── power
│   │   └── power.go
│   └── sound
│       └── sound.go
├── mosquitto
│   └── mosquitto.go
└── configs
    └── config.yaml
```

A brief description of the layout:

- `.github` has two template files for creating PR and issue. Please see the files for more details.
- `.gitignore` varies per project, but all projects need to ignore `bin` directory.
- `.golangci.yml` is the golangci-lint config file.
- `Makefile` is used to build the project. **You need to tweak the variables based on your project**.
- `CHANGELOG.md` contains auto-generated changelog information.
- `OWNERS` contains owners of the project.
- `README.md` is a detailed description of the project.
- `bin` is to hold build outputs.
- `cmd` contains main packages, each subdirecoty of `cmd` is a main package.
- `build` contains scripts, yaml files, dockerfiles, etc, to build and package the project.
- `docs` for project documentations.
- `hack` contains scripts used to manage this repository, e.g. codegen, installation, verification, etc.
- `pkg` places most of project business logic and locate `api` package.
- `release` [chart](https://github.com/caicloud/charts) for production deployment.
- `test` holds all tests (except unit tests), e.g. integration, e2e tests.
- `third_party` for all third party libraries and tools, e.g. swagger ui, protocol buf, etc.
- `vendor` contains all vendored code.

## Running

Download either GO or EXE file from [Releases page](https://github.com/tiredsosha/wardener/releases) and execute it:

    go run main.go
    main.exe

## Configuration

Configuration parameters must be placed in configuration files in the working directory from where you launch Wardener.

<table>
<tr><th>Property</th><th>Description</th><th>Example</th>
<tr><td>broker</td><td>URL of the MQTT broker to use</td><td>127.0.0.1</td></tr>
<tr><td>username</td><td>Username used when connecting to MQTT broker</td><td>admin</td></tr>
<tr><td>password</td><td>Password used when connecting to MQTT broker</td><td>password</td></tr>
</table>

### config.yaml

Wardener will look for this file in the current working directory (directory from where you launched Wardener). Create **config.yaml** file and put desired parameters into it. Or just copy an exemple of this file from config folder in the repo.

Example file:

    broker: 127.0.0.1
    username: admin
    password: password

## Logging

Wardener starts logging immediately after launch. It makes **wardener.log** file in the current working directory.

## Supported messages

The payload of all messages is either raw string or a valid JSON element (possibly a primitive, like a single integer).

Example valid message payloads:

- `0`
- `100`
- `true`
- `test string`

### Status

**Topic:** wardener/status/volume<br>
**Payload:** int in range 0-100<br>
**Persistent:** yes<br>

Send current mastem volume status every 2 minutes.

### Commands

**Topic:** wardener/commands/shutdown<br>
**Payload:** -

Trigger immediate system shutdown.

---

**Topic:** wardener/commands/reboot<br>
**Payload:** -

Trigger immediate system reboot.

---

**Topic:** wardener/commands/volume<br>
**Payload:** int in range 0-100<br>

Trigger changes master volume of system.

---

**Topic:** wardener/commands/mute<br>
**Payload:** boolean

"true" - trigger mutes system volume. "false" - trigger unmutes system volume.

---

## Building

You can build it by yourself.

    go build -o bin/wardener.exe -ldflags "-H windowsgui"

## Future

I will gladly add new staff, if anyone will request!

## Alternativies

- [IOT Link](https://iotlink.gitlab.io/)
- [Winthing](https://github.com/msiedlarek/winthing) **Winthing is no longer actively maintained.**

## License

Copyright 2022 Alexandra Chichko &lt;tiredsosha@gmail.com&gt;

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this software except in compliance with the License.
You may obtain a copy of the License at

> http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
