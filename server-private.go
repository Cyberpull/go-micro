package gosrv

import (
	"net"
)

func (p *pServer) handleIncomingConnection(conn net.Conn) {
	var (
		err      error
		instance *serverClientInstance
	)

	defer conn.Close()

	if instance, err = newServerClientInstance(p, conn); err != nil {
		return
	}

	if err = instance.prepare(); err != nil {
		return
	}

	// Execute client boot event =======
	if err = p.execClientBoot(instance); err != nil {
		return
	}

	p.addClientInstance(instance)

	defer p.removeClientInstance(instance)

	// Execute client ready event =======
	if err = p.execClientReady(instance); err != nil {
		return
	}

	instance.Start()
}

func (p *pServer) addClientInstance(i *serverClientInstance) {
	p.mutex.Lock()

	defer p.mutex.Unlock()

	p.instances[i.uuid] = i
}

func (p *pServer) removeClientInstance(i *serverClientInstance) {
	p.mutex.Lock()

	defer p.mutex.Unlock()

	delete(p.instances, i.uuid)
}

// Event Executors ======================

func (p *pServer) execClientBoot(i *serverClientInstance) error {
	for _, handler := range p.clientBoot {
		err := handler(i.client)

		if err != nil {
			return err
		}
	}

	return nil
}

func (p *pServer) execClientReady(i *serverClientInstance) error {
	for _, handler := range p.clientReady {
		err := handler(i.client)

		if err != nil {
			return err
		}
	}

	return nil
}
