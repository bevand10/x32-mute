# x32-muter

A utility to remotely control Behringer X32 audio console input channel mute states.

## Configuration
The app reads its configuration from `/etc/default/x32-muter`. The file is in the classic `.ini` format, as this example shows:

```text
[x32]
ip = 127.0.0.1

[channel]
number = 37
```

## Usage
`x32-muter on` - configure the channel on (unmute it)
`x32-muter off` - configure the channel off (mute it)
