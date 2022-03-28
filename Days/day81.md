## Fluentd

Another data collector that I wanted to explore as part of this observability section was [Fluentd](https://docs.fluentd.org/). An Open-Source unified logging layer. 

Fluentd treats logs as JSON

[Installing Fluentd](https://docs.fluentd.org/quickstart#step-1-installing-fluentd)

## Collecting logs from files

We all have applications that write to a `.log` file format. Fluentd has the ability to read logs from a file, with the configuration we have set, we will cover that shortly. Below you can see me bringing up the two containers, `docker-compose up -d file-myapp` and then `docker-compose up -d fluentd` 

![](Images/Day81_Monitoring1.png)

If you watch [Introduction to Fluentd: Collect logs and send almost anywhere](https://www.youtube.com/watch?v=Gp0-7oVOtPw&t=447s)from the "That DevOps Guy" (Some amazing content that has been linked throughout this whole challenge) I am using his example here. 

With the container `file-myapp` there is a script in there to add to an example log file. You can see this happening below. 

![](Images/Day81_Monitoring2.png)

The script that we have running inside of our `file-myapp` container looks like this: 

```
#!/bin/sh
while true
do
	echo "Writing log to a file"
  echo '{"app":"file-myapp"}' >> /app/example-log.log
	sleep 5
done
```

We are using the [tail plugin](https://docs.fluentd.org/input/tail) within fluentd which allows us to read those events from the tail of text files, similar to the `tail -F` command. 

We can also use the [HTTP plugin](https://docs.fluentd.org/input/http) which allows us to send events through HTTP requests. This is what we will see with the `http-myapp` container in the repository. This container also runs a similar shell script to generate logs 

```
#!/bin/sh
while true
do
	echo "Sending logs to FluentD"
  curl -X POST -d 'json={"foo":"bar"}' http://fluentd:9880/http-myapp.log
	sleep 5
done
```






### Container logging 








We are going to be using docker-compose to bring up a small demo environment that can demonstrate fluentd. The docker compose file is going to bring up the following containers. 

container_name: fluentd
container_name: http-myapp 
container_name: file-myapp 
container_name: elasticsearch
container_name: kibana



## Resources 

- [Understanding Logging: Containers & Microservices](https://www.youtube.com/watch?v=MMVdkzeQ848)
- [The Importance of Monitoring in DevOps](https://www.devopsonline.co.uk/the-importance-of-monitoring-in-devops/)
- [Understanding Continuous Monitoring in DevOps?](https://medium.com/devopscurry/understanding-continuous-monitoring-in-devops-f6695b004e3b) 
- [DevOps Monitoring Tools](https://www.youtube.com/watch?v=Zu53QQuYqJ0) 
- [Top 5 - DevOps Monitoring Tools](https://www.youtube.com/watch?v=4t71iv_9t_4)
- [How Prometheus Monitoring works](https://www.youtube.com/watch?v=h4Sl21AKiDg) 
- [Introduction to Prometheus monitoring](https://www.youtube.com/watch?v=5o37CGlNLr8)
- [Promql cheat sheet with examples](https://www.containiq.com/post/promql-cheat-sheet-with-examples)
- [Log Management for DevOps | Manage application, server, and cloud logs with Site24x7](https://www.youtube.com/watch?v=J0csO_Shsj0)
- [Log Management what DevOps need to know](https://devops.com/log-management-what-devops-teams-need-to-know/)
- [What is ELK Stack?](https://www.youtube.com/watch?v=4X0WLg05ASw)
- [Fluentd simply explained](https://www.youtube.com/watch?v=5ofsNyHZwWE&t=14s) 

See you on [Day 82](day82.md)

