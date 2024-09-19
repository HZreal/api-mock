package constant

/**
 * @Author elastic·H
 * @Date 2024-08-27
 * @File: topic.go
 * @Description:
 */

type Topic struct {
	Exchange string
	Route    string
	Queue    string
}

var (
	DemoTopic = Topic{
		Exchange: "project.demo.exchange",
		Route:    "project.demo.route",
		Queue:    "project.demo.queue",
	}
	FirstTopic = Topic{
		Exchange: "project.first.exchange",
		Route:    "project.first.route",
		Queue:    "project.first.queue",
	}
)
