Android Update
---

Cli tool for installing Android SDKs, written in golang.


# how to

Execute command.
```
$ android-update
```

`android-update.yml` should exist in the directory you execute command.
Or optionally, you can set the path for the yml with `--config` parameter

# android-update.yml

```
sdk: /User/xxx/android-sdk          # android sdk path(optional, must be absolute)
packages:                           # packages to install(array)
  - platform-tools
  - tools
  - build-tools-23.0.1
  - android-23
  - extra-android-m2repository
  - extra-android-support
  - extra-google-google_play_services
  - extra-google-m2repository
  - extra-google-play_apk_expansion
  - extra-google-play_billing
  - extra-google-play_licensing
```

`ANDROID_HOME` environment variable is used to determine sdk directory.
