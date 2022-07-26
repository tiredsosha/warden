**Table of Contents**

- [Warden](#warden)
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
  - [Release notes](#release-notes)
  - [License](#license)

# Warden

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
├── go.mod
├── go.sum
├── main.go
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

- `README.md` is a detailed description of the project.
- `go.mod` is a detailed reference manual for module system in this project.
- `go.sum` is a detailed file of the checksum of direct and indirect dependency required by the module.
- `main.go` is main file of programm.
- `docs` contains project documentations.
- `control` contains main packages for interaction with Win 10 API, each subdirecoty of `control` is a main package.
- `mosquitto` contains mqtt logic package.
- `configs` contains example of config.yaml file.

## Running

Download either GO or EXE file from [Releases page](https://github.com/tiredsosha/warden/releases) and execute it:

    go run main.go
    main.exe

## Configuration

Configuration parameters must be placed in configuration files in the working directory from where you launch Warden.

<table>
<tr><th>Property</th><th>Description</th><th>Example</th>
<tr><td>broker</td><td>URL of the MQTT broker to use</td><td>127.0.0.1</td></tr>
<tr><td>username</td><td>Username used when connecting to MQTT broker</td><td>admin</td></tr>
<tr><td>password</td><td>Password used when connecting to MQTT broker</td><td>password</td></tr>
</table>

### config.yaml

Warden will look for this file in the current working directory (directory from where you launched Warden). Create **config.yaml** file and put desired parameters into it. Or just copy an example of this file from config folder in the repo.

Example file:

    broker: 127.0.0.1
    username: admin
    password: password

## Logging

Warden starts logging immediately after launch. It makes **warden.log** file in the current working directory.

## Supported messages

The payload of all messages is either raw string or a valid JSON element (possibly a primitive, like a single integer).

Example valid message payloads:

- `0`
- `100`
- `true`
- `test string`

### Status

**Topic:** warden/PC_HOSTNAME/status/volume<br>
**Payload:** int in range 0-100<br>
**Persistent:** yes<br>

Send current mastem volume status every 2 minutes.

### Commands

**Topic:** warden/PC_HOSTNAME/commands/shutdown<br>
**Payload:** -

Trigger immediate system shutdown.

---

**Topic:** warden/PC_HOSTNAME/commands/reboot<br>
**Payload:** -

Trigger immediate system reboot.

---

**Topic:** warden/PC_HOSTNAME/commands/volume<br>
**Payload:** int in range 0-100<br>

Trigger changes master volume of system.

---

**Topic:** warden/PC_HOSTNAME/commands/mute<br>
**Payload:** boolean

"true" - trigger mutes system volume. "false" - trigger unmutes system volume.

---

- `PC_HOSTNAME` [is system name of your Windows pc](https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/hostname).

## Building

You can build it by yourself.

    go build -o bin/warden.exe -ldflags "-H windowsgui"

## Future

I will gladly add new staff, if anyone will request!

## Alternativies

- [IOT Link](https://iotlink.gitlab.io/)
- [Winthing](https://github.com/msiedlarek/winthing) **Winthing is no longer actively maintained.**

## Release notes

[Releases note.md](https://github.com/tiredsosha/warden/docs/release-notes.md)

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
