package tests

import "cyberpull.com/gosrv"

func startClient(s *GoSRVTestSuite) (err error) {
	s.client = gosrv.NewClient(gosrv.ClientOptions{
		ServerHost: ServerHost,
		ServerPort: ServerPort,
		Info: &gosrv.Info{
			Name:        "GoSRV Client",
			Description: "GoSRV Testing Client",
			Alias:       "GoSRV",
		},
	})

	errChan := make(chan error)

	go s.client.Start(errChan)

	err = <-errChan

	return
}
