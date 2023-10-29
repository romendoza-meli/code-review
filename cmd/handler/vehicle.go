package handler

func NewControllerVehicle(st service.ServiceVehicle, adapter HandlerErrorAdapter, sm mapper.StructMapper) *ControllerVehicle {
	return &ControllerVehicle{st: st,
		adapter: adapter,
		sm:      sm,
	}
}

type ControllerVehicle struct {
	st         service.ServiceVehicle
	adapter    mapper.StructMapper
	errAdapter HandlerErrorAdapter
}
