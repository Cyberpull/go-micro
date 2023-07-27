package tests

import "cyberpull.com/gosrv"

func requestHandlerGetDemo(ctx gosrv.Context) gosrv.Output {
	var question string

	if err := ctx.ParseContent(&question); err != nil {
		return ctx.Error(err)
	}

	if question != "Howdy" {
		return ctx.Error("Expected 'Howdy', get '" + question + "'")
	}

	return ctx.Success("Fine, thanks")
}

func requestHandlerDemoUpdate(ctx gosrv.Context) gosrv.Output {
	var question string

	if err := ctx.ParseContent(&question); err != nil {
		return ctx.Error(err)
	}

	if question != "Hello" {
		return ctx.Error("Expected 'Hello', get '" + question + "'")
	}

	ctx.Update("Demo Update")

	return ctx.Success("Hi")
}

// =====================================

func handlerSubscriber(collection gosrv.RequestHandlerCollection) {
	collection.On("GET", "/demo/name", requestHandlerGetDemo)
	collection.On("GET", "/demo/update", requestHandlerDemoUpdate)
}

// =====================================

func startServer(s *GoSRVTestSuite) (err error) {
	s.server = gosrv.NewServer(gosrv.ServerOptions{
		Host: ServerHost,
		Port: ServerPort,
		Info: &gosrv.Info{
			Name:        "GoSRV Server",
			Description: "GoSRV Testing Server",
			Alias:       "GoSRV",
		},
	})

	s.server.RequestHandlers(handlerSubscriber)

	errChan := make(chan error)

	go s.server.Listen(errChan)

	err = <-errChan

	return
}
