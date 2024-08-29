package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

// Estrutura do usuário para representar o modelo na tabela 'users'
type User struct {
	ID   uint
	Name string
}

// Suíte de testes para agrupar os testes relacionados ao banco de dados
type DatabaseTestSuite struct {
	suite.Suite
	db *gorm.DB
}

// SetupSuite executa o código de configuração antes de todos os testes
func (suite *DatabaseTestSuite) SetupSuite() {
	db, err := ConnectDB() // Conecta ao banco de dados
	suite.Require().NoError(err, "Erro ao conectar ao banco de dados")
	suite.db = db

	// Migração automática para criar a tabela 'users'
	err = suite.db.AutoMigrate(&User{})
	suite.Require().NoError(err, "Erro ao migrar o banco de dados")
}

// TestUserInsertion testa a inserção de um registro de usuário
func (suite *DatabaseTestSuite) TestUserInsertion() {
	user := User{Name: "Test User"}     // Cria um usuário de teste
	err := suite.db.Create(&user).Error // Insere o usuário no banco
	suite.Require().NoError(err)

	var retrievedUser User // Recupera o usuário inserido
	err = suite.db.First(&retrievedUser, "name = ?", "Test User").Error
	suite.Require().NoError(err)
	assert.Equal(suite.T(), user.Name, retrievedUser.Name) // Verifica se o nome inserido e recuperado são iguais
}

// TearDownSuite executa o código de limpeza após todos os testes
func (suite *DatabaseTestSuite) TearDownSuite() {
	suite.db.Exec("DROP TABLE users;") // Remove a tabela 'users' após os testes
}

// Executa a suíte de testes
func TestDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseTestSuite))
}
