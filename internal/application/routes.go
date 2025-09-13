package application

import "github.com/teesh3rt/wizardapi/internal/handlers"

func (a *App) loadRoutes() {
	wizardHandler := handlers.WizardHandler{
		Queries: a.Queries,
	}

	a.Router.Get("/wizard/:id", wizardHandler.GetWizard)
	a.Router.Get("/wizard", wizardHandler.GetAllWizards)
	a.Router.Post("/wizard", wizardHandler.CreateWizard)
	a.Router.Delete("/wizard/:id", wizardHandler.DeleteWizard)
}
