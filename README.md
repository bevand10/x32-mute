# x32-mute

A utility to remotely control Behringer X32 audio console input channel mute states.

## Configuration
The app reads its configuration from `/etc/default/x32-mute`. The file is in the classic `.ini` format, as this example shows:

```text
[x32]
ip = 127.0.0.1

[channel]
number = 37
```

## Usage
* `x32-mute yes` - mute the channel
* `x32-mute no`  - unmute the channel
