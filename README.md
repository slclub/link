# link
code link point.   

When you use other source code, try to encapsulate it as much as possible. 

In the future, you can easily replace them with others

The fewer interfaces, the better.

If it has too many interfaces, it may only be used by coupling, That is your choice.


### Wrap viper.

Call directly because there are too many interfaces. 

This goes against the meaning of the project itself.

But viper is so great, So powerful. We don't usually give it up.

Add more source code if it like log system and config etc in the future.


```go
// return *viper.Viper
link.Config() 
```
[viper](https://github.com/spf13/viper)

### Wrap glog.

Here I encapsulate my own log system. 

Maybe there will be something better and more like in the future. I can easily replace it here

```go
link.INFO(args ...interface{})
link.DEBUG(args ...interface{})
link.WARN(args ...interface{})
link.ERROR(args ...interface{})
link.FATAL(args ...interface{})
```

[glog](https://github.com/slclub/glog)


### etc floder just for testing used.

