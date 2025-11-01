package handlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// buildFilterClause creates SQL WHERE conditions and args from query params
func buildFilterClause(c *fiber.Ctx, baseArgs []interface{}) (string, []interface{}) {
	var conditions []string
	args := baseArgs
	argNum := len(baseArgs) + 1

	// Country filter
	if country := c.Query("country"); country != "" {
		conditions = append(conditions, fmt.Sprintf("s.country = $%d", argNum))
		args = append(args, country)
		argNum++
	}

	// Browser filter
	if browser := c.Query("browser"); browser != "" {
		conditions = append(conditions, fmt.Sprintf("s.browser = $%d", argNum))
		args = append(args, browser)
		argNum++
	}

	// Device filter
	if device := c.Query("device"); device != "" {
		conditions = append(conditions, fmt.Sprintf("s.device = $%d", argNum))
		args = append(args, device)
		argNum++
	}

	// Page filter
	if page := c.Query("page"); page != "" {
		conditions = append(conditions, fmt.Sprintf("e.url_path = $%d", argNum))
		args = append(args, page)
		argNum++
	}

	clause := ""
	if len(conditions) > 0 {
		clause = " AND " + strings.Join(conditions, " AND ")
	}

	return clause, args
}
