package streaming

import "log"

/*worker listens to metric through channel*/

func StartWorker(id int) {

	/*Testing purpose only*/
	log.Printf("Worker %d started\n", id)

	for metric := range MetricsChannel {
		log.Printf("Worker %d processing: %+v\n", id, metric)

	}
}
