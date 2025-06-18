package v1

import (
	"encoding/json"
	stdErrors "errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sreekar2307/khata/controller/http/v1/response"
	"github.com/sreekar2307/khata/errors"
	"github.com/sreekar2307/khata/model"
	"github.com/sreekar2307/khata/model/table"
	pkgSlices "github.com/sreekar2307/khata/pkg/slices"
	"github.com/sreekar2307/khata/service"
	"net/http"
	"slices"
)

func (c controller) NewExpense() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req struct {
				SplitType   model.SplitType `json:"split_type" binding:"required"`
				SplitConfig json.RawMessage `json:"split_config"`
				Expense     struct {
					Description string `json:"description" binding:"required,lte=255"`
					Amount      uint64 `json:"amount" binding:"required"`
				}
			}
			splits = make([]model.Split, 0)
		)

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "debug": err.Error()})
			return
		}

		if !slices.Contains(model.SplitTypeValues, req.SplitType) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid split type"})
			return
		}

		authUser, exists := ctx.Get("user")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}
		user := authUser.(table.User)

		splitType := req.SplitType

		if splitType == model.SplitTypes.Percentage {
			percentageSplits := make([]struct {
				Percentage float64 `json:"percentage" biding:"required"`
				UserID     uint64  `json:"user_id" biding:"required"`
			}, 0)
			if err := binding.JSON.BindBody(req.SplitConfig, &percentageSplits); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid split configuration for percentage type"})
				return
			}
			for _, ps := range percentageSplits {
				splits = append(splits, model.Split{
					BorrowerID: ps.UserID,
					Percentage: ps.Percentage,
					LenderID:   user.ID,
				})
			}
		} else if splitType == model.SplitTypes.Amount {
			amountSplits := make([]struct {
				Amount float64 `json:"amount" biding:"required"`
				UserID uint64  `json:"user_id" biding:"required"`
			}, 0)
			if err := binding.JSON.BindBody(req.SplitConfig, &amountSplits); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid split configuration for amount type"})
				return
			}
			for _, as := range amountSplits {
				splits = append(splits, model.Split{
					BorrowerID: as.UserID,
					Amount:     uint64(as.Amount),
					LenderID:   user.ID,
				})
			}
		} else {
			equalSplits := make([]struct {
				UserID uint64 `json:"user_id" biding:"required"`
			}, 0)
			if err := binding.JSON.BindBody(req.SplitConfig, &equalSplits); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid split configuration for equal type"})
				return
			}
			for _, as := range equalSplits {
				splits = append(splits, model.Split{
					BorrowerID: as.UserID,
					LenderID:   user.ID,
				})
			}
		}

		splitsForUser := make(map[uint64]int)
		for _, split := range splits {
			splitsForUser[split.BorrowerID]++
			if splitsForUser[split.BorrowerID] > 1 {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Each user can only have one split in a single expense"})
				return
			}
		}

		expense, ledgers, err := c.ExpenseService.AddExpense(ctx,
			splitType,
			splits,
			table.Expense{
				Description: req.Expense.Description,
				LenderID:    user.ID,
				Amount:      req.Expense.Amount,
			},
		)
		if err != nil {
			if stdErrors.Is(err, errors.ErrInvalidSplitConfiguration) {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid split configuration", "debug": err.Error()})
				return
			} else if stdErrors.Is(err, errors.ErrCheckIfDependencyExists) {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dependency check failed", "debug": err.Error()})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add expense"})
			return
		}
		ledgers = pkgSlices.Filter(ledgers, func(ledger table.Ledger) bool {
			return ledger.BorrowerID != user.ID
		})
		ctx.JSON(http.StatusOK, gin.H{
			"data": response.NewExpense(expense, user, ledgers),
		})
	}
}

func (c controller) Expenses() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			req struct {
				GtId  uint64 `form:"gt_id" binding:"omitempty"`
				Limit int    `form:"limit" binding:"omitempty,min=1,max=100"`
			}
		)
		if err := ctx.BindQuery(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters", "debug": err.Error()})
			return
		}
		authUser, exists := ctx.Get("user")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}
		user := authUser.(table.User)

		views, err := c.LedgerService.GetUserInvolvedExpenses(ctx, user.ID, req.GtId, req.Limit)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add expense"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": pkgSlices.Map(views, func(view service.SimplifiedView) response.SimplifiedView {
				return response.NewSimplifiedView(view)
			}),
		})
	}
}

func (c controller) Balances() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authUser, exists := ctx.Get("user")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}
		user := authUser.(table.User)

		owes, lends, err := c.LedgerService.GetBalanceReport(ctx, user.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add expense"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": response.NewBalance(owes, lends),
		})
	}
}

func (c controller) BalanceConcise() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authUser, exists := ctx.Get("user")
		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}
		user := authUser.(table.User)

		owes, lends, err := c.LedgerService.GetBalanceReportConcise(ctx, user.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add expense"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": response.NewBalanceConcise(owes, lends),
		})
	}
}
