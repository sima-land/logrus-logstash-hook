# Logstash hook for logrus <img src="http://i.imgur.com/hTeVwmJ.png" width="40" height="40" alt=":walrus:" class="emoji" title=":walrus:" />

Based on [this](https://github.com/bshuster-repo/logrus-logstash-hook) hook. But instead of using message context
use global 'type' filed for logging

## Usage

```go
package main

import (
        "github.com/Sirupsen/logrus"
        "github.com/sima-land/logrus-logstash-hook"
)

func main() {
        log := logrus.New()
        hook, err := logrus_logstash.NewHook("tcp", "172.17.0.2:9999", "myApp")
        if err != nil {
                log.Fatal(err)
        }
        log.Hooks.Add(hook)
        log.Info("Hello") // <-- type "myApp" will be added here for logstash
}
```

