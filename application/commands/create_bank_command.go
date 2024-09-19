package commands

import "github.com/Nortriptyline/Eirene/domain/iservice"

// CreateBankCommand is a struct that contains the fields required to create a bank
type CreateBankCommand struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}

type CreateBankCommandHandler struct {
	bankService iservice.IBankService
}

func NewCreateBankCommandHandler(service *iservice.IBankService) *CreateBankCommandHandler {
	return &CreateBankCommandHandler{
		bankService: *service,
	}
}

func (h *CreateBankCommandHandler) Handle(command *CreateBankCommand) error {
	return h.bankService.CreateBank(command.Name, command.Website)
}
