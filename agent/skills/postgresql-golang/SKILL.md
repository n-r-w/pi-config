---
name: postgresql-golang
description: Postgresql golang guidelines and best practices. Use BEFORE ANY activity related to Postgresql in go code (analysis, review, planning, implementation, etc.).
---

<postgresql_go name="Postgresql Golang Guidelines and Best Practices">
    <libraries>
        - `github.com/jackc/pgx/v5`: PostgreSQL driver
        - `github.com/n-r-w/pgh/v2`: Query builder wrapper
        - `github.com/n-r-w/pgh/v2/px`: Query execution helpers
        - `github.com/n-r-w/squirrel`: SQL query builder
        - `github.com/n-r-w/pgh/v2/px/db`: Database connection management
        - `github.com/n-r-w/pgh/v2/txmgr`: Transaction management
        - `github.com/golang-migrate/migrate/v4`: Database migrations
        - `github.com/n-r-w/testdock/v2`: Unit & Integration testing
    </libraries>
    <basic_query>
        <query_builder>
            Use `pgh.Builder()` to create a squirrel query builder with PostgreSQL `$1, $2...` placeholders:

            ```go
            import (
                sq "github.com/n-r-w/squirrel"
                "github.com/n-r-w/pgh/v2"
            )

            query := pgh.Builder().
                Select("id, name, email").
                From("users").
                Where(sq.Eq{"status": "active"})
            ```
        </query_builder>
        <select_operations>
            ```go
            // Select multiple rows
            var users []User
            if err := px.Select(ctx, db, query, &users); err != nil {
                return fmt.Errorf("get users: %w", err)
            }

            // Select single row
            var user User
            if err := px.SelectOne(ctx, db, query, &user); err != nil {
                return fmt.Errorf("get user: %w", err)
            }

            // Process rows one by one (for large datasets)
            err := px.SelectFunc(ctx, db, query, func(row pgx.Row) error {
                var u User
                if err := row.Scan(&u.ID, &u.Name, &u.Email); err != nil {
                    return err
                }
                // process u...
                return nil
            })
            ```
        </select_operations>
        <exec_operations>
            ```go
            query := pgh.Builder().
                Update("users").
                SetMap(map[string]any{
                    "name":  user.Name,
                    "email": user.Email,
                }).
                Where(sq.Eq{"id": user.ID})

            _, err := px.Exec(ctx, tx, query)
            ```
        </exec_operations>
        <batch_operations>
            ```go
            // Execute multiple queries in one batch
            queries := make([]sq.Sqlizer, 0, len(users))
            for _, user := range users {
                q := pgh.Builder().
                    Insert("users").
                    SetMap(map[string]any{"name": user.Name, "email": user.Email})
                queries = append(queries, q)
            }
            _, err := px.ExecBatch(ctx, queries, db)

            // Split large batches (e.g., 100 per batch)
            rowsAffected, err := px.ExecSplit(ctx, db, queries, 100)

            // Insert with batch splitting
            baseQuery := pgh.Builder().
                Insert("users").
                Columns("name", "email")
            values := make([]pgh.Args, 0, len(users))
            for _, u := range users {
                values = append(values, pgh.Args{u.Name, u.Email})
            }
            rowsAffected, err = px.InsertSplit(ctx, db, baseQuery, values, 100)
            ```
        </batch_operations>
        <error_handling>
            ```go
            if err != nil {
                switch {
                case px.IsNoRows(err):
                    // No rows found
                case px.IsUniqueViolation(err):
                    // Unique constraint violation (PostgreSQL code 23505)
                case px.IsForeignKeyViolation(err):
                    // Foreign key violation (PostgreSQL code 23503)
                default:
                    return err
                }
            }
            ```
        </error_handling>
    </basic_query>
    <transaction_management>
        Full control over transaction options:
        ```go
        err := px.BeginTxFunc(ctx, pool, pgx.TxOptions{
            IsoLevel:   pgx.ReadCommitted,
            AccessMode: pgx.ReadWrite,
        }, func(ctx context.Context, tx pgx.Tx) error {
            // Execute queries using tx directly
            _, err := tx.Exec(ctx, "INSERT INTO users (name) VALUES ($1)", "John")
            if err != nil {
                return err // Transaction will be rolled back
            }
            // Or use px helpers with tx
            query := pgh.Builder().Update("users").Set("active", true).Where("id = ?", 1)
            _, err = px.Exec(ctx, tx, query)
            if err != nil {
                return err // Transaction will be rolled back
            }
            return nil // Transaction will be committed
        })
        ```

        Simplified version with default transaction options:
        ```go
        err := px.BeginFunc(ctx, pool, func(ctx context.Context, tx pgx.Tx) error {
            _, err := tx.Exec(ctx, "UPDATE accounts SET balance = balance - $1 WHERE id = $2", 100, fromID)
            if err != nil {
                return err
            }
            _, err = tx.Exec(ctx, "UPDATE accounts SET balance = balance + $1 WHERE id = $2", 100, toID)
            return err
        })
        ```

        **Key behaviors:**
        - Automatic commit on success
        - Automatic rollback on error or panic
        - Panic is re-thrown after rollback
    </transaction_management>
</postgresql_go>
