package api_test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/cucumber/godog/colors"
	"github.com/gin-gonic/gin"
	"github.com/jdejesus007/gt-api-project/api"
	"github.com/jdejesus007/gt-api-project/api/provider"

	"github.com/cucumber/godog"
)

var err error

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "pretty", // can define default values
}

type TestImplementation struct {
	api.API
	consumed bool
}

type stateContainer struct {
	api               *gin.Engine
	mockProvider      provider.RepositoryProvider
	resp              *httptest.ResponseRecorder
	inMemoryVariables struct {
		customerEmail string
		customerUUID  string
	}
}

var state stateContainer

func resetResponse() {
	state.resp = httptest.NewRecorder()
}

func cleanupDatabaseAndStateVariables() {
	// This clears up all the changes in the DB after migration
	tables := []string{"customers", "orders", "books"}
	var db *sql.DB
	db, err = state.mockProvider.Database().GetConn().DB()
	if err != nil {
		return
	}
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS=0;")
	if err != nil {
		log.Printf("Failed to unset foreign key checks - %+v\n", err)
	}

	for _, v := range tables {
		_, err = db.Exec(fmt.Sprintf("TRUNCATE TABLE %s;", v))
		if err != nil {
			log.Printf("Failed to truncate table - %+v\n", err)
		}
	}

	_, err = db.Exec("SET FOREIGN_KEY_CHECKS=1;")
	if err != nil {
		log.Printf("Failed to set foreign key checks - %+v\n", err)
	}
}

func initializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		// cleanupDatabaseAndStateVariables()
		return ctx, nil
	})

	ctx.Step(`^I send "(GET)" request to "([^"]*)"$`, iSendGetRequestTo)
	ctx.Step(`^I send "(POST)" request to "([^"]*)"$`, iSendPostRequestTo)
	ctx.Step(`^the response code should be (\d+)$`, theResponseCodeShouldBe)
	ctx.Step(`^the response should match string: "([^"]*)"$`, theResponseShouldMatchString)
	ctx.Step(`^the response should match json:`, theResponseShouldMatchJSON)
}

func iSendGetRequestTo(method, endpoint string) (err error) {
	req, err := initializeRequest(&method, &endpoint, nil)
	if err != nil {
		return
	}

	return sendRequestTo(req, endpoint)
}

func iSendPostRequestTo(method, endpoint string, body *godog.DocString) (err error) {
	req, err := initializeRequest(&method, &endpoint, &body.Content)
	if err != nil {
		return
	}

	return sendRequestTo(req, endpoint)
}

func sendRequestTo(req *http.Request, endpoint string) (err error) {
	// handle panic
	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
		}
	}()

	switch endpoint {
	case "/ping":
		state.api.ServeHTTP(state.resp, req)
	case "/customers/":
		state.api.ServeHTTP(state.resp, req)
		if req.Method == "POST" && state.resp.Code == 201 {
			var data map[string]interface{}
			if err = json.Unmarshal(state.resp.Body.Bytes(), &data); err == nil && data["UUID"] != nil {
				state.inMemoryVariables.customerEmail = data["email"].(string)
				state.inMemoryVariables.customerUUID = data["UUID"].(string)
			}
		}
	default:
		state.api.ServeHTTP(state.resp, req)
	}
	return
}

func initializeRequest(method, endpoint *string, body *string) (*http.Request, error) {
	resetResponse()

	*body = strings.Replace(*body, "{email}", state.inMemoryVariables.customerEmail, -1)

	reader := new(strings.Reader)
	if body != nil {
		var formattedData []byte
		var data map[string]interface{}
		if err = json.Unmarshal([]byte(*body), &data); err != nil {
			return nil, err
		}
		for key, value := range data {
			switch value {
			case "{email}":
				data[key] = state.inMemoryVariables.customerEmail
			}
		}
		if formattedData, err = json.Marshal(data); err != nil {
			return nil, err
		}
		*body = string(formattedData)
		reader = strings.NewReader(*body)
	}
	return http.NewRequest(*method, *endpoint, reader)
}

func theResponseCodeShouldBe(code int) error {
	if code != state.resp.Code {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, state.resp.Code)
	}
	return nil
}

func theResponseShouldMatchString(expected string) error {
	actual := state.resp.Body.String()

	// the matching may be adapted per different requirements.
	if expected != actual {
		return fmt.Errorf("expected String does not match actual, %s vs. %s", expected, actual)
	}

	return nil
}

func theResponseShouldMatchJSON(body *godog.DocString) (err error) {
	// Switch standing expected placeholder within recently created customer uuid
	// Customer uuid gets created on the fly and will match customer email and uuid
	(*body).Content = strings.Replace((*body).Content, "{customerUUID}", state.inMemoryVariables.customerUUID, -1)

	var expected, actual []byte
	var v interface{}
	if err = json.Unmarshal([]byte(state.resp.Body.Bytes()), &v); err != nil {
		return
	}
	switch v := v.(type) {
	case []interface{}:
		for index := range v {
			value := v[index].(map[string]interface{})
			cleanNonComparableFields(&value)
			v[index] = value
		}
	case map[string]interface{}:
		cleanNonComparableFields(&v)
	default:
		err = fmt.Errorf("unsupport JSON type for comparison")
	}
	if actual, err = json.Marshal(v); err != nil {
		return
	}
	if err = json.Unmarshal([]byte(body.Content), &v); err != nil {
		return
	}
	if expected, err = json.Marshal(v); err != nil {
		return
	}
	if !bytes.Equal(actual, expected) {
		err = fmt.Errorf("expected json: %s, does not match actual: %s", string(expected), string(actual))
	}
	return
}

// NewTestServer consumes the underlying implementation and returns a test
// server. A test server should only be used for one test case to prevent state
// bbeing dirtied from one test to another. This will panic if called more than
// once.
func (impl *TestImplementation) NewTestServer() *gin.Engine {
	if impl.consumed {
		panic("Implementation was already used")
	}
	impl.consumed = true
	return impl.API.CreateServer()
}

func TestMain(m *testing.M) {
	flag.Parse()
	opts.Paths = flag.Args()

	status := godog.TestSuite{
		Name:                 "GT Main Suite",
		TestSuiteInitializer: initializeSuite,
		ScenarioInitializer:  initializeScenario,
		Options:              &opts,
	}.Run()

	// Optional: Run `testing` package's logic besides godog.
	if st := m.Run(); st > status {
		status = st
	}

	purgeDBResource()

	os.Exit(status)
}

func initializeSuite(ctx *godog.TestSuiteContext) {
	// -- Populate with mock data
	p := newMockProvider()
	state = stateContainer{
		api: (&TestImplementation{
			API: api.NewBuilder().
				WithRepositoryProvider(p).
				Finalize(),
		}).NewTestServer(),
		mockProvider: p,
	}
}

// cleanNonComparableFields
func cleanNonComparableFields(data *map[string]interface{}) {
	delete(*data, "UUID")
	delete(*data, "createdAt")
}
