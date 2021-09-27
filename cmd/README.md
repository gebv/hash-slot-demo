(WIP) Quick solution in the cloud

TODO: finished code of a quick solution

GCP app engine
Set the max number of instances that can be run.

Refer to the instance of interest using the following URL format
```
https://INSTANCE_ID-dot-VERSION_ID-dot-SERVICE_ID-dot-PROJECT_ID.REGION_ID.r.appspot.com

https://INSTANCE_ID.VERSION_ID.SERVICE_ID.CUSTOM_DOMAIN

https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed#gcloud

Note: Targeting an instance is not supported in services that are configured for auto scaling or basic scaling. The instance ID must be an integer in the range from 0, up to the total number of instances running. Regardless of your scaling type or instance class, it is not possible to send a request to a specific instance without targeting a service or version within that instance.
```

An example of a yaml file for GAE
```
service: hashslotdemo1
runtime: go115
instance_class: B1
basic_scaling:
  max_instances: 5
  idle_timeout: 1m

handlers:
  - url: /.*
    script: auto
```
