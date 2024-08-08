package shared

import "log"

// FailOnError a helper function to check the return value for each amqp call
func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("【%s】: %s", msg, err)
	}
}
