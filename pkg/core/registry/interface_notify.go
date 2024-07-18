package registry

import (
	"sync"

	dubboRegistry "dubbo.apache.org/dubbo-go/v3/registry"
	gxset "github.com/dubbogo/gost/container/set"
)

// Endpoint nolint
type Endpoint struct {
	Port     int    `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}

type GeneralInterfaceNotifyListener struct {
	ctx            *ApplicationContext
	notifyListener *NotifyListener
	allUrls        gxset.HashSet
	mutex          sync.Mutex
}

func NewGeneralInterfaceNotifyListener(
	ctx *ApplicationContext,
	notifyListener *NotifyListener,
) *GeneralInterfaceNotifyListener {
	return &GeneralInterfaceNotifyListener{
		ctx:            ctx,
		allUrls:        *gxset.NewSet(),
		notifyListener: notifyListener,
		mutex:          sync.Mutex{},
	}
}

func (gilstn *GeneralInterfaceNotifyListener) Notify(event *dubboRegistry.ServiceEvent) {
	url := event.Service
	urlStr := url.String()

	gilstn.mutex.Lock()
	defer gilstn.mutex.Unlock()

	if !gilstn.allUrls.Contains(urlStr) {
		gilstn.allUrls.Add(urlStr)
	}

	listener := NewInterfaceServiceChangedNotifyListener(gilstn.notifyListener)
}

func (gilstn *GeneralInterfaceNotifyListener) NotifyAll(events []*dubboRegistry.ServiceEvent, f func()) {
	for _, event := range events {
		gilstn.Notify(event)
	}
}
