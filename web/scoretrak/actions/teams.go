package actions

import (
	"fmt"
	"net/http"
	"scoretrak/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/x/responder"
	"scoretrak/constants"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Team)
// DB Table: Plural (teams)
// Resource: Plural (Teams)
// Path: Plural (/teams)
// View Template Folder: Plural (/templates/teams/)

// TeamsResource is the resource for the Team model
type TeamsResource struct {
	buffalo.Resource
}

// List gets all Teams. This function is mapped to the path
// GET /teams
func (v TeamsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	teams := &models.Teams{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Teams from the DB
	if err := q.All(teams); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("teams", teams)
		return c.Render(http.StatusOK, r.HTML("/teams/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(teams))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(teams))
	}).Respond(c)
}

// Show gets the data for one Team. This function is mapped to
// the path GET /teams/{team_id}
func (v TeamsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Team
	team := &models.Team{}

	// To find the Team the parameter team_id is used.
	if err := tx.Find(team, c.Param("team_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("team", team)

		return c.Render(http.StatusOK, r.HTML("/teams/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(team))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(team))
	}).Respond(c)
}

// New renders the form for creating a new Team.
// This function is mapped to the path GET /teams/new
func (v TeamsResource) New(c buffalo.Context) error {
	c.Set("team", &models.Team{})
	c.Set("roles", constants.Roles)
	return c.Render(http.StatusOK, r.HTML("/teams/new.plush.html"))
}

// Create adds a Team to the DB. This function is mapped to the
// path POST /teams
func (v TeamsResource) Create(c buffalo.Context) error {
	// Allocate an empty Team
	team := &models.Team{}

	// Bind team to the html form elements
	if err := c.Bind(team); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(team)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("team", team)
			c.Set("roles", constants.Roles)
			return c.Render(http.StatusUnprocessableEntity, r.HTML("/teams/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "team.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/teams/%v", team.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(team))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(team))
	}).Respond(c)
}

// Edit renders a edit form for a Team. This function is
// mapped to the path GET /teams/{team_id}/edit
func (v TeamsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Team
	team := &models.Team{}
	c.Set("roles", constants.Roles)
	if err := tx.Find(team, c.Param("team_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	
	c.Set("team", team)
	return c.Render(http.StatusOK, r.HTML("/teams/edit.plush.html"))
}

// Update changes a Team in the DB. This function is mapped to
// the path PUT /teams/{team_id}
func (v TeamsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Team
	team := &models.Team{}

	if err := tx.Find(team, c.Param("team_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Team to the html form elements
	if err := c.Bind(team); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(team)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("team", team)
			c.Set("roles", constants.Roles)
			return c.Render(http.StatusUnprocessableEntity, r.HTML("/teams/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "team.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/teams/%v", team.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(team))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(team))
	}).Respond(c)
}

// Destroy deletes a Team from the DB. This function is mapped
// to the path DELETE /teams/{team_id}
func (v TeamsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Team
	team := &models.Team{}

	if err := tx.Eager().Find(team, c.Param("team_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if team.Role == constants.Black{
		teams := []models.Team{}
		//Query for all teams with role constants.Black
		if err := tx.Where("role = (?)", constants.Black).All(&teams); err != nil {
			return c.Error(http.StatusInternalServerError, err)
		}
		
		//if there is 1 or less black teams left
		if len(teams) <= 1{
			//Then deny the deletion
			c.Flash().Add("Failed", fmt.Sprintf("You should have at least one team with role \"%s\"", constants.Black))
			return responder.Wants("html", func(c buffalo.Context) error {
				c.Set("team", team)
				return c.Render(http.StatusForbidden, r.HTML("/teams/show.plush.html"))
			}).Respond(c)
		}

	}

	//Disallow deletion of the last user in last constants.Black team
	if len(team.Users) >= 1{
		c.Flash().Add("Failed", "The team contains user(s) that need to be removed first")
		return responder.Wants("html", func(c buffalo.Context) error {
			c.Set("team", team)
			return c.Render(http.StatusConflict, r.HTML("/teams/show.plush.html"))
		}).Respond(c)
	}

	if err := tx.Destroy(team); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "team.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/teams")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(team))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(team))
	}).Respond(c)
}
