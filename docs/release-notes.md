Warden 1.1.5 Release Notes

IMPROVEMENTS

- Added new status - mute
- No Windows popup messages when you reboot/restart system
- Nice system tray icon
- Warden now you be exited thougth systen tray icon

FIXES

- Shutdown function turns off pc, not logs it out
- Online status wasn't retained function

---

Warden 1.2.0-beta.1 Release Notes

IMPROVEMENTS

- Added new status - mute
- Added test support for display control
- No Windows popup messages when you reboot/restart system

FIXES

- Shutdown function turns off pc, not logs it out
- Online status wasn't retained function

---

Warden 1.1.4 Release Notes

IMPROVEMENTS

- Added new CLI flags
- Changed log file
- Added nodebug mode

---

Warden 1.1.3 Release Notes

IMPROVEMENTS

- Added CLI
- Added aditional info in log file

---

Warden 1.1.2 Release Notes

IMPROVEMENTS

- Changed way to send online status

---

Warden 1.1.1 Release Notes

IMPROVEMENTS

- Autoreconnection when connection to broker is lost/broker is offline
- Added scalable publisher solution
- No more exit(1) when mqtt lost connection

---

Warden 1.1.0 Release Notes

IMPROVEMENTS

- Custom logger package
- Custom yaml configuration package
- Fatal errors traceback now in warden.log
- Config.yaml validation

---

Warden 1.0.5 Release Notes

IMPROVEMENTS

- Period between publishing of the messages changed to 25s
- README.md is changed
- release-noted moved to new directory

FIXES

- Hostname now is part of MQTT Topic Prefix

---

Warden 1.0.3 Release Notes

IMPROVEMENTS

- Name is changed to Warden

---

Warden 1.0.2 Release Notes

FIXES

- Log formating when new message arrive

---

Warden 1.0.1 Release Notes

Fully working and tested release of project

IMPROVEMENTS

- Power managment of pc. Supported:
  - reboot
  - shutdown
- Sound managment of pc. Supported:
  - mute/unmute
  - volume control
- Publication of sound status of system over mqtt

---

Warden 1.0.0-alpha Pre-Release Notes

PROTOTYPE RELEASE
