// +build integration

package integration

import (
	"neon-auth/app/domain/model"
	"neon-auth/app/interface/persistence"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"github.com/stretchr/testify/assert"
)

func TestAccountRepositoryIntegrationSaveAccount(t *testing.T) {
	config := persistence.PostgresConfig{Host: "localhost", Port: "5432", User: "test", Password: "test", Name: "account"}

	accountRepoPostgres, err := persistence.NewPsqlAccountRepository(config)
	if err != nil {
		assert.Equal(t, err, nil, "An error occured while connecting to the database in integration tests")
		os.Exit(1)
	}
	account, err := model.NewAccount("test@test.ru", "test")
	assert.Equal(t, err, nil)
	err = accountRepoPostgres.SaveAccount(account)
	if err != nil {
		assert.Equal(t, err, nil, "There should be no errors when adding to a clean base")
		os.Exit(1)
	}
	//When you add the same account again, the code should return an error
	err = accountRepoPostgres.SaveAccount(account)
	assert.Equal(t, err.Error(), "pq: duplicate key value violates unique constraint \"accounts_email_key\"")

}
func TestAccountRepositoryIntegrationFindAccount(t *testing.T) {
	config := persistence.PostgresConfig{Host: "localhost", Port: "5432", User: "test", Password: "test", Name: "account"}

	accountRepoPostgres, err := persistence.NewPsqlAccountRepository(config)
	if err != nil {
		assert.Equal(t, err, nil, "An error occured while connecting to the database in integration tests")
		os.Exit(1)
	}
	account, err := model.NewAccount("find@test.ru", "test")
	assert.Equal(t, err, nil)
	err = accountRepoPostgres.SaveAccount(account)
	if err != nil {
		assert.Equal(t, err, nil, "There should be no errors when adding to a clean base")
		os.Exit(1)
	}

	newAccount, err := accountRepoPostgres.FindByEmail("find@test.ru")
	if err != nil {
		assert.Equal(t, err, nil, "When receiving from the database there should be no errors")
		os.Exit(1)
	}

	assert.Equal(t, newAccount.Email, account.Email, "Emails must match when received from the repository")
	assert.Equal(t, newAccount.Password, account.Password, "Passwords must match when received from the repository")
	assert.Equal(t, newAccount.Status, account.Status, "Status must match when received from the repository")

}
