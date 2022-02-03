# by-email-notification-sdk


## compile protos

Some languajes required extra dependencies and binaries. you can build the protos with docker:

- In a terminal into the project root folder run:
```
docker run --rm -v `pwd`:/temp/buf-gen netsaj/buf-generate
```

- check `gen` folder your protos was be compiled. 


**Option1:** If you have all dependencies you can build running in your terminal 
```
buf generate
```

**Option2:** Run command in your terminal:
```
sh docker.sh
```

**Note1:** docker image repository : https://github.com/netsaj/buf-generate