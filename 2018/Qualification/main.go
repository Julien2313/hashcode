package main

import "github.com/sirupsen/logrus"

const (
	//fileName = "a_example"
	//fileName = "b_should_be_easy"
	//fileName = "c_no_hurry"
	fileName = "d_metropolis"
	//fileName = "e_high_bonus"

	logLevel = logrus.WarnLevel // DebugLevel, InfoLevel, WarnLevel
)

func main() {
	startLogger()

	myS := &UberGoogle{
		rides: make(map[int]*ride),
	}
	myS.read()
	log.WithFields(logrus.Fields{
		"function": "main",
	}).Debug("Let's start UberGoogle")

	for step := 0; step < myS.T; step++ {
		log.WithFields(logrus.Fields{
			"step": step,
		}).Info("moving cars at the current step")
		myS.moveAllCarsAtStep(step)
	}

	myS.write()
}
