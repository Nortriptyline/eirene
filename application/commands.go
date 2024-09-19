package application

import "github.com/Nortriptyline/Eirene/application/commands"

type Commands struct {
	CreateBankCommandHandler commands.CreateBankCommandHandler
}

func InitializeCommands(services *Services) *Commands {
	return &Commands{
		CreateBankCommandHandler: *commands.NewCreateBankCommandHandler(&services.BankService),
	}
}
