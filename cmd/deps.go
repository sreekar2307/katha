package main

import (
	"fmt"
	"github.com/sreekar2307/katha/config"
	"github.com/sreekar2307/katha/db/sql"
	"github.com/sreekar2307/katha/pkg/jwt"
	"github.com/sreekar2307/katha/pkg/jwt/jwtgo"
	"github.com/sreekar2307/katha/repository"
	"github.com/sreekar2307/katha/service"
	"github.com/sreekar2307/katha/service/expense"
	"github.com/sreekar2307/katha/service/ledger"
	"github.com/sreekar2307/katha/service/user"
	"github.com/sreekar2307/katha/simplifier/onelevel"
	splitterFactory "github.com/sreekar2307/katha/splitter/factory"
	"gorm.io/gorm"
)

type Deps struct {
	ExpenseService service.Expense
	UserService    service.User
	LedgerService  service.Ledger
	Conf           config.Config
	PrimaryDB      *gorm.DB
	JwtImpl        jwt.JWT
	Repository     repository.Repository
}

func GetDeps() (Deps, error) {
	conf, err := config.New()
	if err != nil {
		return Deps{}, fmt.Errorf("failed to load configuration: %w", err)
	}
	primaryDB, err := sql.NewSqlConnection(conf.SQL.PrimaryDatabase)
	if err != nil {
		return Deps{}, fmt.Errorf("failed to connect to primary database: %w", err)
	}
	var (
		repo           = repository.NewRepository()
		expenseService = expense.NewExpenseService(splitterFactory.NewFactory(), primaryDB, repo)
	)
	simplifier := onelevel.NewOneLevelSimplifier(primaryDB, repo)
	ledgerService := ledger.NewLedgerService(simplifier, primaryDB, repo)
	jwtImpl, err := jwtgo.NewGoJWT(conf.Server.AuthTokenSecret)
	if err != nil {
		return Deps{}, fmt.Errorf("failed to create JWT implementation: %w", err)
	}
	userService := user.NewUserService(primaryDB, repo, jwtImpl)
	return Deps{
		ExpenseService: expenseService,
		UserService:    userService,
		LedgerService:  ledgerService,
		Conf:           conf,
		PrimaryDB:      primaryDB,
		JwtImpl:        jwtImpl,
		Repository:     repo,
	}, nil
}
